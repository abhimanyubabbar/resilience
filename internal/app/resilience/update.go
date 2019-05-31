/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sqweek/dialog"
)

type updateData struct {
	Latest int
}

func updateHosts(explicit bool) error {
	var httpClient = &http.Client{Timeout: 60 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/blockList.b2sum")
	if err != nil {
		updateHostsError()
		return err
	}
	defer r.Body.Close()
	b2sum, err := ioutil.ReadAll(r.Body)
	if err != nil {
		updateHostsError()
		return err
	}
	if stateState.hostsHash == strings.Trim(string(b2sum), "\r\n ") {
		if explicit {
			dialog.Message(
				"No updates are available for your Resilience block list.",
			).Title("Resilience Update").Info()
		}
		return nil
	}
	r, err = httpClient.Get("https://resilienceblocker.info/data/blockList")
	if err != nil {
		updateHostsError()
		return err
	}
	defer r.Body.Close()
	hosts, err := ioutil.ReadAll(r.Body)
	return denierUpdate(hosts, true)
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
