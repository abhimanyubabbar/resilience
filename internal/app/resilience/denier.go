/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/elazarl/goproxy"
	"github.com/sqweek/dialog"
	"golang.org/x/crypto/blake2b"
)

func denierHostsInit() error {
	err := denierVerifyConfig()
	if err != nil {
		denierHostsError()
		return err
	}
	hosts, err := denierHostsRead()
	if err != nil {
		return err
	}
	denierUpdate(hosts, false)
	return err
}

func denierProxyInit() {
	stateState.proxy = goproxy.NewProxyHttpServer()
	stateState.proxy.Verbose = false
	stateState.proxy.OnRequest().HandleConnectFunc(
		func(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
			if !stateState.enabled {
				return goproxy.OkConnect, host
			}
			if stateState.rules.ShouldBlock(ctx.Req.URL.String(), map[string]interface{}{
				"domain": host,
			}) {
				return goproxy.RejectConnect, host
			}
			return goproxy.OkConnect, host
		},
	)
	http.ListenAndServe(":7341", stateState.proxy)
}

func denierUpdate(hosts []byte, write bool) error {
	var newRules []string
	var err error
	for _, rule := range strings.Split(string(hosts), "\n") {
		rule = strings.Trim(rule, "\r\n ")
		if len(rule) > 0 {
			newRules = append(newRules, rule)
		}
	}
	tempRules, err := adblockNewRules(newRules, nil)
	if err != nil {
		denierUpdateError()
		return err
	}
	stateState.rules = tempRules
	newHash := blake2b.Sum256(hosts)
	stateState.hostsHash = strings.Join([]string{
		hex.EncodeToString(newHash[:]),
		"blockList",
	}, "  ")
	if write {
		err = denierVerifyConfig()
		if err != nil {
			denierHostsError()
			return err
		}
		err = denierHostsWrite(hosts)
		if err != nil {
			denierHostsError()
			return err
		}
	}
	return err
}

func denierVerifyConfig() error {
	currentUser, _ := user.Current()
	hostsFilePath := path.Join(path.Join(path.Join(
		currentUser.HomeDir, ".config"), "resilience"), "blockList",
	)
	configFolderInfo, err := os.Stat(
		path.Join(currentUser.HomeDir, ".config"),
	)
	if err != nil || !configFolderInfo.IsDir() {
		err = os.Mkdir(path.Join(currentUser.HomeDir, ".config"), 0700)
		if err != nil {
			return err
		}
	}
	configFolderInfo, err = os.Stat(
		path.Join(currentUser.HomeDir, path.Join(".config", "resilience")),
	)
	if err != nil || !configFolderInfo.IsDir() {
		err = os.Mkdir(path.Join(currentUser.HomeDir, path.Join(".config", "resilience")), 0700)
		if err != nil {
			return err
		}
	}
	_, err = os.OpenFile(hostsFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	return nil
}

func denierHostsRead() ([]byte, error) {
	currentUser, _ := user.Current()
	hostsFilePath := path.Join(path.Join(path.Join(
		currentUser.HomeDir, ".config"), "resilience"), "blockList",
	)
	hosts, err := ioutil.ReadFile(hostsFilePath)
	return hosts, err
}

func denierHostsWrite(hosts []byte) error {
	currentUser, _ := user.Current()
	hostsFilePath := path.Join(path.Join(path.Join(
		currentUser.HomeDir, ".config"), "resilience"), "blockList",
	)
	err := ioutil.WriteFile(hostsFilePath, hosts, 0600)
	return err
}

func denierUpdateError() {
	dialog.Message(strings.Join([]string{
		"Could not update your Resilience block list.",
	}, "\n")).Title("Resilience Error").Error()
}

func denierHostsError() {
	dialog.Message(strings.Join([]string{
		"Could not read or write to your local Resilience block list.",
	}, "\n")).Title("Resilience Error").Error()
}
