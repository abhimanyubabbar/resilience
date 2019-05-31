/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */

package main

func togglerEnable() error {
	stateState.enabled = true
	return nil
}

func togglerDisable() error {
	stateState.enabled = false
	return nil
}
