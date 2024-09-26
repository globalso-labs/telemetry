SHELL=/bin/bash -e -o pipefail
PWD = $(shell pwd)


all: git-hooks tidy ## Initializes all tools

out:
	@mkdir -p out

git-hooks:
	@git config --local core.hooksPath .githooks/

download: ## Downloads the dependencies
	@go mod download

update: ## Updates the dependencies
	@go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)

tidy: ## Cleans up go.mod and go.sum
	@go mod tidy

imports: bin/goimports
	@bin/goimports -w -l .

agent: imports  ## Run the app
	 CGO_ENABLED=1 go build -ldflags "-s -w" -o ./agent/out/agent agent/cmd/*.go && ./agent/out/agent --config examples/telemetry.yaml

test-build: ## Tests whether the code compiles
	@go build -o /dev/null ./...

build: imports out/bin ## Builds all binaries

GO_BUILD = mkdir -pv "$(@)" && go build -ldflags="-w -s" -o "$(@)" ./...
.PHONY: out/bin
out/bin:
	$(GO_BUILD)

GOLANGCI_LINT = bin/golangci-lint
$(GOLANGCI_LINT):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b bin

lint: imports $(GOLANGCI_LINT) download ## Lints all code with golangci-lint
	@$(GOLANGCI_LINT) run --verbose

lint-fix: imports $(GOLANGCI_LINT) download ## Fixes all code with golangci-lint
	@$(GOLANGCI_LINT) run --fix --verbose

lint-reports: out/lint.xml

.PHONY: out/lint.xml
out/lint.xml: $(GOLANGCI_LINT) out download
	@$(GOLANGCI_LINT) run ./... --out-format checkstyle | tee "$(@)"

test: ## Runs all tests
	@go test $(ARGS) ./...

benchmark: ## Runs all tests
	@go test -bench=. -run=^# $(ARGS) ./...

coverage: out/report.json ## Displays coverage per func on cli
	go tool cover -func=out/cover.out

html-coverage: out/report.json ## Displays the coverage results in the browser
	go tool cover -html=out/cover.out

test-reports: out/report.json

.PHONY: out/report.json
out/report.json: out
	@go test -count 1 ./... -coverprofile=out/cover.out --json | tee "$(@)"

clean: ## Cleans up everything
	@rm -rf bin out protodeps

# Go dependencies versioned through tools.go
GO_DEPENDENCIES = golang.org/x/tools/cmd/goimports

define make-go-dependency
  # target template for go tools, can be referenced e.g. via /bin/<tool>
  bin/$(notdir $1):
	GOBIN=$(PWD)/bin go install $1
endef

# this creates a target for each go dependency to be referenced in other targets
$(foreach dep, $(GO_DEPENDENCIES), $(eval $(call make-go-dependency, $(dep))))

ci: lint-reports test-reports ## Executes lint and test and generates reports

help: ## Shows the help
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
        awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ''
