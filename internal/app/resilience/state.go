/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

type state struct {
	enabled bool
}

var stateState = state{}

func stateInstatiate() {
	stateState = state{
		enabled: true,
	}
}
