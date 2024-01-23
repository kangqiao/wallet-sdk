tag := $(shell git describe --tags --exact-match 2>/dev/null || echo "")
commit := $(shell git rev-parse HEAD)
build_time := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LD_FLAGS := -ldflags "-w -s  \
	-X zp.com/wallet-sdk/common.Tag=${tag} \
	-X zp.com/wallet-sdk/common.Commit=${commit} \
	-X zp.com/wallet-sdk/common.BuildTime=${build_time}"

clean:
	rm -fr sdk/Sdk.xcframework
	rm -fr sdk/sdk.aar
	rm -fr sdk/sdk-sources.jar
	rm -fr ./build

ios:
	go mod tidy && \
	cd sdk && \
    rm -fr *.xcframework && \
    go get golang.org/x/mobile/cmd/gomobile@v0.0.0-20231127183840-76ac6878050a && \
    gomobile init &&\
    gomobile bind -target=ios ${LD_FLAGS}

android:
	go mod tidy && \
	cd sdk && \
	go get golang.org/x/mobile/cmd/gomobile@v0.0.0-20231127183840-76ac6878050a && \
	gomobile init &&\
	gomobile bind -target=android/arm64 ${LD_FLAGS}


mobile: clean ios android