# [Resilience](https://resilienceblocker.info)
<img src="https://raw.githubusercontent.com/kaepora/resilience/master/assets/icon/icon.png" height="96" />

Resilience is an ad blocker for your computer that works with **any browser** on **any operating system**.

Resilience doesn't sell out your privacy with *"acceptable ads"*. Resilience won't be blocked by your web browser's developers. Resilience won't ever stop defending your privacy and your right to block ads on your goddamn computer.

**Resilience is currently in early alpha!** test, contribute, suggest fixes and improvements! The goal is to get it to a state where it's just as easy to use as uBlock.

## Goals 
### Features List`
- [x] Download blocked hosts list from secure server.
- [X] Parse [EasyList format](https://adblockplus.org/filter-cheatsheet) into HTTP proxy rules.
- [X] HTTP proxy for Windows, Linux and macOS.
- [x] Launch at startup and show tray icon.
- [x] Automatically update hosts list on startup and every 24 hours.
- [x] Check for Resilience client updates automatically.
- [x] Store block list and load locally on startup -- check remote hash for updates
- [ ] Interface for managing blocks, adding exceptions or custom blocks.

### Assets Needed
- [ ] Website: Help and Docs

## Build
You must have [Go](https://golang.org) installed in order to build Resilience.

### Windows
```
Build.cmd
```

### Linux and macOS
```
make dependencies
make all
```

Resilience will be located at `cmd/resilience/resilience`.

## Getting Started
In order to use Resilience, you must configure it as your system's HTTP and HTTPS proxy by setting them both to `localhost:7341`. People testing an early alpha already know how to do this, but better help instructions will come in due course.

## License
Authored by [Nadim Kobeissi](https://nadim.computer) and released under the MIT license.
