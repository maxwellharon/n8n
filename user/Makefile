.PHONY: dev deps proto

proto:
	$(MAKE) -C proto proto

deps: proto
	go mod download

dev:
	air -c air.conf

