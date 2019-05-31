/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

import (
	"github.com/elazarl/goproxy"
)

type state struct {
	enabled   bool
	hostsHash string
	proxy     *goproxy.ProxyHttpServer
	rules     *adblockRules
}

func stateInstatiate() state {
	return state{
		enabled:   true,
		hostsHash: "",
		proxy:     nil,
		rules:     nil,
	}
}

var stateState = stateInstatiate()
