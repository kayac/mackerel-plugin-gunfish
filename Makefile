LATEST_TAG := $(shell git describe --abbrev=0 --tags)
export GO111MODULE := on

test:
	go test -v ./...

lint:
	go vet ./...

dist:
	goxc

packages:
	goreleaser build --skip-validate --rm-dist

packages-snapshot:
	goreleaser build --skip-validate --rm-dist --snapshot

clean:
	rm -fr dist/*

release: dist
	ghr -u kayac -r mackerel-plugin-gunfish $(LATEST_TAG) dist/snapshot/

.PHONY: packages packages-snapshot test lint clean dist
