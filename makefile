TARGETDIR=.\deploy
proj=github.com/alexandr-andreyev/rkeeper7-simple-api
sha1ver := $(shell git rev-parse HEAD)
test := $(shell date /t)


all: vet test  buildEXE

vet:
	go vet -all -shadow .\cmd\rk7simpleapi-win
	go vet -all -shadow .\app

test: 
	go.exe test -timeout 30s $(proj)\app

# The sha1 stuff isn't working as of now
buildEXE:
	go build -o "rk7simpleapi.exe" -a -ldflags "-X main.sha1ver=$(sha1ver)" .\cmd\rk7simpleapi-win  
