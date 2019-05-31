# [Resilience](https://resilienceblocker.info)
<img src="https://raw.githubusercontent.com/kaepora/resilience/master/assets/logo.svg" height="96" />
Resilience Blocker.

- No Acceptable Ads.
- No compromises.

## Goals
### Features List
- [x] Download blocked hosts list from secure server.
- [X] Parse [EasyList format](https://adblockplus.org/filter-cheatsheet) into HTTP proxy rules.
- [X] HTTP proxy for Windows, Linux and macOS.
- [x] Launch at startup and show tray icon.
- [x] Automatically update hosts list on startup and every 24 hours.
- [x] Check for Resilience client updates automatically.
- [x] Store block list and load locally on startup -- check remote hash for updates
- [ ] Interface for managing blocks, adding exceptions or custom blocks.

### Assets Needed
- [x] Logo
- [ ] Website
- [ ] Website: Manifesto
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

## License
Authored by [Nadim Kobeissi](https://nadim.computer) and released under the GNU General Public License, version 3.
