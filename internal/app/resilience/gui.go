/* SPDX-License-Identifier: GPL-3.0-or-later
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"runtime"

	"github.com/getlantern/systray"
)

func guiOnReady() {
	systray.SetIcon(iconData)
	if runtime.GOOS != "darwin" {
		systray.SetTitle("Resilience")
	}
	systray.SetTooltip("Resilience")
	mStatus := systray.AddMenuItem("Resilience is Enabled", "")
	mStatus.Disable()
	mToggle := systray.AddMenuItem("Disable", "")
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
					err := togglerDisable()
					if err == nil {
						mStatus.SetTitle("Resilience is Disabled")
						mToggle.SetTitle("Enable")
					}
				} else {
					err := togglerEnable()
					if err == nil {
						mStatus.SetTitle("Resilience is Enabled")
						mToggle.SetTitle("Disable")
					}
				}
			case <-mUpdate.ClickedCh:
				go func() {
					updateHosts(false)
					updateClient(true)
				}()
			case <-mAbout.ClickedCh:
				aboutDialog()
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func guiOnExit() {}
