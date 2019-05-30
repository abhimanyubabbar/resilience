@go get github.com/sqweek/dialog
@go get github.com/getlantern/systray
@go get golang.org/x/crypto/blake2b
@go build -ldflags="-s -w -H=windowsgui" -o cmd\resilience\resilience.exe^
    internal\app\resilience\main.go^
    internal\app\resilience\state.go^
    internal\app\resilience\gui.go^
    internal\app\resilience\update.go^
    internal\app\resilience\about.go^
    internal\app\resilience\version.go^
    internal\app\resilience\denier.go^
    internal\app\resilience\toggler.go^
    internal\app\resilience\icon.go
::@cmd\resilience\resilience.exe