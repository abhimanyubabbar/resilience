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

func updateHosts(explicit bool) {
	var httpClient = &http.Client{Timeout: 60 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/hosts")
	if err != nil {
		updateHostsError()
		return
	}
	defer r.Body.Close()
	hosts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateHostsError()
		return
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
		denierUpdate(hosts)
	}
}

func updateClient(explicit bool) {
	var updateResult updateData
	var httpClient = &http.Client{Timeout: 20 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/updateClient.json")
	if err != nil {
		updateClientError()
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateClientError()
		return
	}
	err = json.Unmarshal(body, &updateResult)
	if err != nil {
		updateClientError()
		return
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
