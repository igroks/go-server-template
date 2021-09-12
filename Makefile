PKGS = $$(go list ./... | grep -v /vendor/)

.PHONY: setup
setup:
	go mod download
	go mod tidy
	go mod vendor

.PHONY: lint
lint:
	go fmt $(PKGS)