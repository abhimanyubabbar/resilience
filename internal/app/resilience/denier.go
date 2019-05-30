package main

import "runtime"

func denierUpdate(hosts []byte) {
	if runtime.GOOS == "windows" {
		denierUpdateWindows(hosts)
	} else if runtime.GOOS == "linux" {
		denierUpdateLinux(hosts)
	} else if runtime.GOOS == "darwin" {
		denierUpdateMac(hosts)
	}
}

func denierUpdateWindows(hosts []byte) {}

func denierUpdateLinux(hosts []byte) {}

func denierUpdateMac(hosts []byte) {}
