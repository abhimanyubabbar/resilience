/* SPDX-License-Identifier: MIT
 * Copyright © 2019-2020 Nadim Kobeissi <nadim@nadim.computer>.
 * All Rights Reserved. */
package main

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

func aboutDialog() {
	var aboutText = strings.Join([]string{
		"Resilience " + versionString + "\n",
		"Resilience is an easy to use content blocker for your computer.",
		"For news and information, please visit:",
		"https://resilienceblocker.info",
	}, "\n")
	dialog.Message(aboutText).Title("About Resilience").Info()
}

func aboutHelpPage() {
	helpPage := "https://resilienceblocker.info/#help"
	switch runtime.GOOS {
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", helpPage).Start()
	case "linux":
		exec.Command("xdg-open", helpPage).Start()
	case "darwin":
		exec.Command("open", helpPage).Start()
	}
}
