# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: aether android ios aether-cross swarm evm all test clean
.PHONY: aether-linux aether-linux-386 aether-linux-amd64 aether-linux-mips64 aether-linux-mips64le
.PHONY: aether-linux-arm aether-linux-arm-5 aether-linux-arm-6 aether-linux-arm-7 aether-linux-arm64
.PHONY: aether-darwin aether-darwin-386 aether-darwin-amd64
.PHONY: aether-windows aether-windows-386 aether-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

aether:
	build/env.sh go run build/ci.go install ./cmd/aether
	@echo "Done building."
	@echo "Run \"$(GOBIN)/aether\" to launch aether."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/aether.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/aether.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen

# Cross Compilation Targets (xgo)

aether-cross: aether-linux aether-darwin aether-windows aether-android aether-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/aether-*

aether-linux: aether-linux-386 aether-linux-amd64 aether-linux-arm aether-linux-mips64 aether-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-*

aether-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/aether
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep 386

aether-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/aether
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep amd64

aether-linux-arm: aether-linux-arm-5 aether-linux-arm-6 aether-linux-arm-7 aether-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep arm

aether-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/aether
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep arm-5

aether-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/aether
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep arm-6

aether-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/aether
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep arm-7

aether-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/aether
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep arm64

aether-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/aether
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep mips

aether-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/aether
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep mipsle

aether-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/aether
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep mips64

aether-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/aether
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/aether-linux-* | grep mips64le

aether-darwin: aether-darwin-386 aether-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/aether-darwin-*

aether-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/aether
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/aether-darwin-* | grep 386

aether-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/aether
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/aether-darwin-* | grep amd64

aether-windows: aether-windows-386 aether-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/aether-windows-*

aether-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/aether
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/aether-windows-* | grep 386

aether-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/aether
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/aether-windows-* | grep amd64
