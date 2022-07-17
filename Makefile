DOCKER := $(shell which docker)
GOCMD := $(shell which go)
SWAG = $(shell which swag)
MOCKERY = $(shell which mockery)

lint:
	golangci-lint run -c golangci.yml

swagger:
	echo TODO

mocks:
	@${MOCKERY} --all
