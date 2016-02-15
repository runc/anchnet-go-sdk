# Makefile
#
# Targets:
#   build
#   test
#   clean

# Build local binaries.
build:
	cd anchnet && godep go build . && cd -
.PHONY: build

# Unit test.
test:
	godep go test .
.PHONY: test

# Clean up.
clean:
	rm -rf anchnet-go-sdk/anchnet/anchnet
	rm -rf _output
.PHONY: clean
