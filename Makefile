PROJECT ?= watermark-service
IMAGE ?= watermark-service
VERSION ?= $(shell date +v%Y%m%d)-$(shell git describe --tags --always --dirty)


all: lint test build

# Run lint
lint:
	golangci-lint run ./pkg/... ./cmd/...

# Run tests
test:
	go test -v ./pkg/... ./cmd/... -coverprofile coverage.out

test-coverage:
	./scripts/test-coverage.sh coverage.out

run:
	go run ./cmd/server/main.go

build:
	docker build --cache-from docker.io/$(PROJECT)/$(IMAGE):latest \
		-t docker.io/$(PROJECT)/$(IMAGE):$(VERSION) \
		-t docker.io/$(PROJECT)/$(IMAGE):latest -f Dockerfile .


push:
	docker push docker.io/$(PROJECT)/$(IMAGE):latest
	docker push docker.io/$(PROJECT)/$(IMAGE):$(VERSION)


.PHONY: image push

release: lint test build push
