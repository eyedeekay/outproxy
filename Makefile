
GO111MODULE=on

build:
	cd outproxy && go build -tags="netgo" \
		-ldflags '-w -extldflags "-static"'

try:
	cd outproxy && ./outproxy

clean:
	find . -name '*.go' -exec gofmt -w -s {} \;
	find . -name '*.i2pkeys' -exec rm {} \;
