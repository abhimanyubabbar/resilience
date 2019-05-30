/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import "github.com/getlantern/systray"

func main() {
	systray.Run(guiOnReady, guiOnExit)
}
