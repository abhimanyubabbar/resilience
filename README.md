# [Resilience](https://resilienceblocker.info)
Resilience Blocker.

## Goals
### Features List
- [ ] Download blocked hosts list from secure server.
- [ ] Block URIs on Windows, Linux and macOS.
- [ ] GUI for managing blocks, adding exceptions or custom blocks.
- [ ] Launch at startup and show tray icon.
- [ ] Automatically update hosts list on startup and every 24 hours.

### Assets Needed
- [ ] Logo
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