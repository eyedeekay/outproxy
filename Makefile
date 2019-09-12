
PACKAGE=outproxy
USER_GH=eyedeekay
GO111MODULE=on
VERSION := 0.32.081

GO111MODULE=on

echo:
	@echo "gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(PACKAGE) -t v$(VERSION) -d Standalone Go+SAM based outproxies"

tag:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(PACKAGE) -t v$(VERSION) -d "Standalone Go+SAM based Outproxies"


build:
	cd outproxy && go build -tags="netgo" \
		-ldflags '-w -extldflags "-static"'

try:
	cd outproxy && ./outproxy

clean:
	find . -name '*.go' -exec gofmt -w -s {} \;
	find . -name '*.i2pkeys' -exec rm {} \;

