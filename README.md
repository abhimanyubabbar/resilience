# [Resilience](https://resilienceblocker.info)
<img src="https://raw.githubusercontent.com/kaepora/resilience/master/assets/logo.svg" height="96" />
Resilience Blocker.

- No Acceptable Ads.
- No compromises.

## Goals
### Features List
- [x] Download blocked hosts list from secure server.
- [x] Block URIs on Windows, Linux and macOS.
- [x] Launch at startup and show tray icon.
- [x] Automatically update hosts list on startup and every 24 hours.
- [x] Check for Resilience client updates automatically.
- [ ] GUI for managing blocks, adding exceptions or custom blocks.

### Assets Needed
- [x] Logo
- [ ] Website

## Build
You must have [Go](https://golang.org) installed in order to build Resilience.

```
make dependencies
make all
```

Resilience will be located at `cmd/resilience/resilience`.

## License
Authored by [Nadim Kobeissi](https://nadim.computer) and released under the GNU General Public License, version 3.
