APP=latest
APP_EXECUTABLE="./out/$(APP)"
SRC_PACKAGES=$(shell go list ./... | grep -v "vendor" | grep -v "latest.cli/latest")
DEP:=$(shell command -v dep 2> /dev/null)
GOLINT:=$(shell command -v golint 2> /dev/null)
RICHGO=$(shell command -v richgo 2> /dev/null)

ifeq ($(RICHGO),)
	GOBIN=go
else
	GOBIN=richgo
endif

ensure-out-dir:
	mkdir -p $(PWD)/out

setup:
ifeq ($(DEP),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif
ifeq ($(GOLINT),)
	$(GOBIN) get -u golang.org/x/lint/golint
endif
ifeq ($(RICHGO),)
	$(GOBIN) get -u github.com/kyoh86/richgo
endif

build-deps:
	dep ensure -v

test: ensure-out-dir
	ENVIRONMENT=test $(GOBIN) test $(SRC_PACKAGES) -coverprofile ./out/coverage -v

lint:
	@for p in $(SRC_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint -set_exit_status=1 $$p; \
	done
	$1

fmt:
	$(GOBIN) fmt $(SRC_PACKAGES)

vet:
	$(GOBIN) vet $(SRC_PACKAGES)

build: ensure-out-dir
	$(GOBIN) build -o $(APP_EXECUTABLE) ./main.go

all: setup build-deps test fmt vet lint build
