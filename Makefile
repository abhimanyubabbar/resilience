all:
	@/bin/echo -n "[Resilience] Building Resilience... "
	@go build -ldflags="-s -w" -o cmd/resilience/resilience internal/app/resilience/*.go
	@/bin/echo "      OK"

dependencies:
	@/bin/echo -n "[Resilience] Installing dependencies..."
	@go get -u github.com/sqweek/dialog
	@go get -u github.com/getlantern/systray
	@go get -u golang.org/x/crypto/blake2b
	@go get -u github.com/elazarl/goproxy
	@/bin/echo " OK"

clean:
	@/bin/echo -n "[Resilience] Cleaning up... "
	@rm -f cmd/resilience/resilience
	@/bin/echo "            OK"

.PHONY: dependencies clean all api assets cmd docs internal
