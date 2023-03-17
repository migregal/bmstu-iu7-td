PKG := markup2

TESTFLAGS:=-race

# lint =========================================================================

define lint_project
	golangci-lint run --build-tags="unit" ./$(1)/... && golangci-lint run --build-tags="unit" ./cmd/$(1)/...
endef

define lint_folder
	golangci-lint run --build-tags="testing" ./$(1)/...
endef

.PHONY: lint
lint: lint_markupapi lint_pkg

.PHONY: lint_markupapi
lint_markupapi:
	$(call lint_project,markupapi)

.PHONY: lint_pkg
lint_pkg:
	$(call lint_folder,pkg)

# lint ^========================================================================

# unit =========================================================================

.PHONY: unit-test
unit-test: OUT:=./out
unit-test:
	mkdir -p $(OUT)/unit
	go test -cover -coverprofile=$(OUT)/unit/coverage.out -covermode=atomic \
	-coverpkg=$(PKG)/... $(TESTFLAGS) --tags="unit" $(PKG)/...

.PHONY: coverage
coverage: OUT:=./out
coverage:
	go tool cover -html=$(OUT)/unit/coverage.out -o $(OUT)/unit/index.html

# unit ^========================================================================

# build ========================================================================

.PHONY: build
build: OUT=./out
build:
	go build -o $(OUT)/markupapi.out ./cmd/markupapi/main.go

# build ^=======================================================================
