/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"github.com/getlantern/systray"
)

func guiOnReady() {
	systray.SetIcon(iconData)
	systray.SetTitle("Resilience")
	systray.SetTooltip("Resilience Blocker")
	mToggle := systray.AddMenuItem("Disable", "Resilience is Enabled.")
	systray.AddSeparator()
	mUpdate := systray.AddMenuItem("Update", "Check for Updates.")
	systray.AddSeparator()
	mAbout := systray.AddMenuItem("About", "About Resilience.")
	mQuit := systray.AddMenuItem("Quit", "Quit Resilience.")
	go func() {
		for {
			select {
			case <-mToggle.ClickedCh:
				if stateState.enabled {
					stateState.enabled = false
					mToggle.SetTitle("Disabled")
				} else {
					stateState.enabled = true
					mToggle.SetTitle("Enabled")
				}
			case <-mUpdate.ClickedCh:
				systray.Quit()
			case <-mAbout.ClickedCh:
				aboutDialog()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func guiOnExit() {}
