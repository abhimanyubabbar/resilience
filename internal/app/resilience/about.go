package main

import (
	"strings"

	"github.com/sqweek/dialog"
)

func aboutDialog() {
	var aboutText = strings.Join([]string{
		"Resilience " + version + "\n",
		"Resilience is an easy to use content blocker for your computer.",
		"For news and information, please visit:",
		"https://resilienceblocker.info",
	}, "\n")
	dialog.Message(aboutText).Title("About Resilience").Info()
}
