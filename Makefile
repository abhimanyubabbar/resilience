all:
	@/bin/echo -n "[Resilience] Building Resilience... "
	@go build -ldflags="-s -w" -o cmd/resilience/resilience internal/app/resilience/*.go
	@/bin/echo "      OK"

dependencies:
	@/bin/echo -n "[Resilience] Installing dependencies..."
	@/bin/echo " OK"

clean:
	@/bin/echo -n "[Resilience] Cleaning up... "
	@rm -f internal/app/resilience/parser.go
	@rm -f cmd/resilience/resilience
	@/bin/echo "            OK"

.PHONY: dependencies clean all api assets cmd docs internal
