/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

type state struct {
	enabled   bool
	hostsHash [32]byte
}

func stateInstatiate() state {
	return state{
		enabled:   true,
		hostsHash: [32]byte{},
	}
}

var stateState = stateInstatiate()
