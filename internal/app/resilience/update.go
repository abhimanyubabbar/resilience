/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sqweek/dialog"
	"golang.org/x/crypto/blake2b"
)

type updateData struct {
	Latest int
}

func updateHosts(explicit bool) error {
	var httpClient = &http.Client{Timeout: 60 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/hosts")
	if err != nil {
		updateHostsError()
		return err
	}
	defer r.Body.Close()
	hosts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateHostsError()
		return err
	}
	hostsHash := blake2b.Sum256(hosts)
	if bytes.Compare(stateState.hostsHash[:], hostsHash[:]) == 0 {
		if explicit {
			dialog.Message(
				"No updates are available for your Resilience block list.",
			).Title("Resilience Update").Info()
		}
	} else {
		stateState.hostsHash = hostsHash
		return denierUpdate(hosts)
	}
	return nil
}

func updateClient(explicit bool) error {
	var updateResult updateData
	var httpClient = &http.Client{Timeout: 20 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/updateClient.json")
	if err != nil {
		updateClientError()
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateClientError()
		return err
	}
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		updateClientError()
		return err
	}
	if updateResult.Latest > versionBuild {
		dialog.Message(strings.Join([]string{
			"An update is available for your Resilience client.\n",
			"To download the latest version, please visit:",
			"https://resilienceblocker.info",
		}, "\n")).Title("Resilience Update").Info()
	} else {
		if explicit {
			dialog.Message(
				"No updates are available for your Resilience client.",
			).Title("Resilience Update").Info()
		}
	}
	return nil
}

func updateHostsError() {
	dialog.Message(
		"Could not update your Resilience block list.",
	).Title("Resilience Update").Error()
}

func updateClientError() {
	dialog.Message(
		"Could not check for updates for Resilience.",
	).Title("Resilience Update").Error()
}
