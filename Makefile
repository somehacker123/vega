APPS := dummyriskmodel vega vegaccount vegastream

ifeq ($(CI),)
	# Not in CI
	VERSION := dev-$(USER)
	VERSION_HASH := $(shell git rev-parse HEAD | cut -b1-8)
else
	# In CI
	ifneq ($(DRONE),)
		# In Drone: https://docker-runner.docs.drone.io/configuration/environment/variables/

		ifneq ($(DRONE_TAG),)
			VERSION := $(DRONE_TAG)
		else
			# No tag, so make one
			VERSION := $(shell git describe --tags 2>/dev/null)
		endif
		VERSION_HASH := $(shell echo "$(CI_COMMIT_SHA)" | cut -b1-8)

	else
		# In an unknown CI
		VERSION := unknown-CI
		VERSION_HASH := unknown-CI
	endif
endif

.PHONY: all
all: build

.PHONY: lint
lint: ## Lint the files
	@t="$$(mktemp)" ; \
	go list ./... | xargs golint | grep -vE '(and that stutters|blank import should be|should have comment|which can be annoying to use)' | tee "$$t" ; \
	code=0 ; test "$$(wc -l <"$$t" | awk '{print $$1}')" -gt 0 && code=1 ; \
	rm -f "$$t" ; \
	exit "$$code"

.PHONY: retest
retest: ## Re-run all unit tests
	@go test -count=1 ./...

.PHONY: test
test: ## Run unit tests
	@go test ./...

.PHONY: integrationtest
integrationtest: ## run integration tests, showing ledger movements and full scenario output
	@go test -v ./integration/... -godog.format=pretty

.PHONY: race
race: ## Run data race detector
	@env CGO_ENABLED=1 go test -race ./...

.PHONY: mocks
mocks: ## Make mocks
	@go generate ./...

.PHONY: msan
msan: ## Run memory sanitizer
	@if ! which clang 1>/dev/null ; then echo "Need clang" ; exit 1 ; fi
	@env CC=clang CGO_ENABLED=1 go test -msan ./...

.PHONY: vet
vet: ## Run go vet
	@go vet -all ./...

.PHONY: vetshadow
vetshadow: # Run go vet with shadow detection
	@go vet -shadow ./... 2>&1 | grep -vE '^(#|gateway/graphql/generated.go|proto/.*\.pb\.(gw\.)?go)' ; \
	code="$$?" ; test "$$code" -ne 0

.PHONY: .testCoverage.txt
.testCoverage.txt:
	@go list ./... |grep -v '/gateway' | xargs go test -covermode=count -coverprofile="$@"
	@go tool cover -func="$@"

.PHONY: coverage
coverage: .testCoverage.txt ## Generate global code coverage report

.PHONY: .testCoverage.html
.testCoverage.html: .testCoverage.txt
	@go tool cover -html="$^" -o "$@"

.PHONY: coveragehtml
coveragehtml: .testCoverage.html ## Generate global code coverage report in HTML

.PHONY: deps
deps: ## Get the dependencies
	@go mod download
	@go mod vendor
	@grep 'google/protobuf' go.mod | awk '{print "# " $$1 " " $$2 "\n"$$1"/src";}' >> vendor/modules.txt
	@modvendor -copy="**/*.proto"

.PHONY: build
build: build-linux-amd64 ## install the binaries (for linux/amd64) in cmd/{progname}/

.PHONY: build-%
build-%: ## Build the binaries for a particular OS+ARCH combination in cmd/{progname}/
	@v="${VERSION}" vh="${VERSION_HASH}" gcflags="" suffix="" ; \
	if test -n "$$DEBUGVEGA" ; then \
		gcflags="all=-N -l" ; \
		suffix="-dbg" ; \
		v="debug-$$v" ; \
	fi ; \
	goos="$$(echo "$*" | cut -f1 -d-)" ; \
	goarch="$$(echo "$*" | cut -f2 -d-)" ; \
	if echo "$$goarch" | grep -q '^arm[5-7]$$' ; then \
		goarch=arm ; goarm="$$(echo "$$goarch" | cut -b4)" ; \
	fi ; \
	ldflags="-X main.Version=$$v -X main.VersionHash=$$vh" ; \
	echo "Version: $$v ($$vh)" ; \
	for app in $(APPS) ; do \
		o="./cmd/$$app/$$app-$$goos-$$goarch$$goarm$$suffix" ; \
		echo "Building $$o" ; \
		env CGO_ENABLED=1 "GOARCH=$$goarch" "GOARM=$$goarm" "GOOS=$$goos" \
			go build -v -ldflags "$$ldflags" -gcflags "$$gcflags" -o "$$o" "./cmd/$$app" \
			|| exit 1 ; \
	done

.PHONY: buildall
buildall: ## Build the binaries for all OS+ARCH combinations
	@go tool dist list | tr / " " | while read -r os arch ; do \
		if test "$$arch" = arm ; then \
			for v in 5 6 7 ; do \
				target="build-$${os}-$${arch}$${v}" ; \
				echo -n "$$target " ; \
				make "$$target" 1>"$$target.log" 2>&1 || echo -n "failed" ; \
				echo ; \
			done ; \
		else \
			target="build-$${os}-$${arch}" ; \
			echo -n "$$target " ; \
			make "build-$${os}-$${arch}" 1>"$$target.log" 2>&1 || echo -n "failed" ; \
			echo ; \
		fi ; \
	done

.PHONY: gofmtsimplify
gofmtsimplify:
	@find . -path vendor -prune -o \( -name '*.go' -and -not -name '*_test.go' -and -not -name '*_mock.go' \) -print0 | xargs -0r gofmt -s -w

.PHONY: install
install: ## install the binaries in GOPATH/bin
	@cat .asciiart.txt
	@echo "Version: ${VERSION} (${VERSION_HASH})"
	@for app in $(APPS) ; do \
		env CGO_ENABLED=1 go install -v -ldflags "-X main.Version=${VERSION} -X main.VersionHash=${VERSION_HASH}" "./cmd/$$app" || exit 1 ; \
	done

.PHONY: gqlgen
gqlgen: ## run gqlgen
	@cd ./gateway/graphql/ && go run github.com/99designs/gqlgen -c gqlgen.yml

.PHONY: gqlgen_check
gqlgen_check: ## GraphQL: Check committed files match just-generated files
	@find gateway/graphql -name '*.graphql' -o -name '*.yml' -exec touch '{}' ';' ; \
	make gqlgen 1>/dev/null || exit 1 ; \
	files="$$(git diff --name-only gateway/graphql/)" ; \
	if test -n "$$files" ; then \
		echo "Committed files do not match just-generated files:" $$files ; \
		test -n "$(CI)" && git diff gateway/graphql/ ; \
		exit 1 ; \
	fi

.PHONY: ineffectassign
ineffectassign: ## Check for ineffectual assignments
	@ia="$$(env GO111MODULE=auto ineffassign . | grep -v '_test\.go:')" ; \
	if test "$$(echo -n "$$ia" | wc -l | awk '{print $$1}')" -gt 0 ; then echo "$$ia" ; exit 1 ; fi

.PHONY: proto
proto: deps ## build proto definitions
	@./proto/generate.sh

.PHONY: proto_check
proto_check: ## proto: Check committed files match just-generated files
	@make proto_clean 1>/dev/null
	@make proto 1>/dev/null
	@files="$$(git diff --name-only proto/)" ; \
	if test -n "$$files" ; then \
		echo "Committed files do not match just-generated files:" $$files ; \
		test -n "$(CI)" && git diff proto/ ; \
		exit 1 ; \
	fi

.PHONY: proto_clean
proto_clean:
	@find proto -name '*.pb.go' -o -name '*.pb.gw.go' -o -name '*.validator.pb.go' -o -name '*.swagger.json' \
		| xargs -r rm
	@find proto/doc -name index.html -o -name index.md \
		| xargs -r rm

# Misc Targets

codeowners_check:
	@if grep -v '^#' CODEOWNERS | grep "," ; then \
		echo "CODEOWNERS cannot have entries with commas" ; \
		exit 1 ; \
	fi

.PHONY: print_check
print_check: ## Check for fmt.Print functions in Go code
	@f="$$(mktemp)" && \
	find -name vendor -prune -o \
		-name cmd -prune -o \
		-name '*_test.go' -prune -o \
		-name '*.go' -print0 | \
		xargs -0 grep -E '^([^/]|/[^/])*fmt.Print' | \
		tee "$$f" && \
	count="$$(wc -l <"$$f")" && \
	rm -f "$$f" && \
	if test "$$count" -gt 0 ; then exit 1 ; fi

.PHONY: docker
docker: ## Make docker container image from scratch
	@test -f "$(HOME)/.ssh/id_rsa" || exit 1
	@docker build \
		--build-arg SSH_KEY="$$(cat ~/.ssh/id_rsa)" \
		-t "docker.pkg.github.com/vegaprotocol/vega/vega:$(VERSION)" \
		.

.PHONY: docker_quick
docker_quick: build ## Make docker container image using pre-existing binaries
	@for app in $(APPS) ; do \
		f="cmd/$$app/$$app" ; \
		if ! test -f "$$f" ; then \
			echo "Failed to find: $$f" ; \
			exit 1 ; \
		fi ; \
		cp -a "$$f" . || exit 1 ; \
	done
	@docker build \
		-t "docker.pkg.github.com/vegaprotocol/vega/vega:$(VERSION)" \
		-f Dockerfile.quick \
		.
	@for app in $(APPS) ; do \
		rm -rf "./$$app" ; \
	done

.PHONY: gettools_build
gettools_build:
	@./script/gettools.sh build

.PHONY: gettools_develop
gettools_develop:
	@./script/gettools.sh develop

# Make sure the mdspell command matches the one in .drone.yml.
.PHONY: spellcheck
spellcheck: ## Run markdown spellcheck container
	@docker run --rm -ti \
		--entrypoint mdspell \
		-v "$(PWD):/src" \
		docker.pkg.github.com/vegaprotocol/devops-infra/markdownspellcheck:latest \
			--en-gb \
			--ignore-acronyms \
			--ignore-numbers \
			--no-suggestions \
			--report \
			'*.md' \
			'docs/**/*.md'

# The integration directory is special, and contains a package called core_test.
.PHONY: staticcheck
staticcheck: ## Run statick analysis checks
	@go list ./... | grep -v /integration | xargs staticcheck
	@f="$$(mktemp)" && find integration -name '*.go' | xargs staticcheck | grep -v 'could not load export data' | tee "$$f" && \
	count="$$(wc -l <"$$f")" && rm -f "$$f" && if test "$$count" -gt 0 ; then exit 1 ; fi

.PHONY: clean
clean: ## Remove previous build
	@for app in $(APPS) ; do rm -f "$$app" build-*.log "cmd/$$app/$$app" "cmd/$$app/$$app"-* ; done

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
