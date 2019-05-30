/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func fetchHosts() ([]byte, error) {
	var httpClient = &http.Client{Timeout: 60 * time.Second}
	r, err := httpClient.Get("https://resilienceblocker.info/data/hosts")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	return body, err
}
