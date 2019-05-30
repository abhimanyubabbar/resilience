package main

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
)

func denierUpdate(hosts []byte) {
	var hostsPath string
	if runtime.GOOS == "windows" {
		hostsPath = path.Join(
			path.Join(
				path.Join(
					path.Join(
						os.Getenv("windir"), "system32",
					), "drivers",
				), "etc",
			), "hosts",
		)
	} else if runtime.GOOS == "linux" {
		hostsPath = path.Join("/etc", "hosts")
	} else if runtime.GOOS == "darwin" {
		hostsPath = path.Join("/etc", "hosts")
	}
	denierUpdateWrite(hosts, hostsPath)
}

func denierUpdateWrite(hosts []byte, hostsPath string) {
	resBegin := "# BEGIN RESILIENCE BLOCK LIST"
	resEnd := "# END RESILIENCE BLOCK LIST"
	resMatch := regexp.MustCompile(
		strings.Join([]string{"(?s)", resBegin, ".+", resEnd}, ""),
	)
	hostsFile, err := ioutil.ReadFile(hostsPath)
	stringHostsFile := string(hostsFile)
	var hostsUpdated string
	if err != nil {
		denierUpdateError()
		return
	}
	isResMatch := resMatch.MatchString(stringHostsFile)
	if isResMatch {
		hostsUpdated = resMatch.ReplaceAllString(stringHostsFile, strings.Join([]string{
			resBegin,
			string(hosts),
			resEnd,
		}, "\n"))
	} else {
		hostsUpdated = strings.Join([]string{
			stringHostsFile,
			resBegin,
			string(hosts),
			resEnd,
		}, "\n")
	}
	err = ioutil.WriteFile(hostsPath, []byte(hostsUpdated), 755)
	if err != nil {
		denierUpdateError()
		return
	}
}

func denierUpdateError() {
	dialog.Message(strings.Join([]string{
		"Could not write to your Resilience block list.",
		"Please run Resilience as an administrator.",
	}, "\n")).Title("Resilience Error").Error()
}
