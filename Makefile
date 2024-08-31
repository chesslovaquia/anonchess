WEBPACK_FLAGS ?= --mode development

.PHONY: default
default: build

.PHONY: all
all:
	@echo "GOOS=`go env GOOS` GOARCH=`go env GOARCH`"
	@echo "  `go version`"
	@echo "  nodejs `node --version`"
	@echo "  npm `npm --version`"
	@echo "  npx `npx --version`"
	@$(MAKE) build-deps test release publish

.PHONY: clean
clean:
	@rm -vrf static/pkg static/ui publish

.PHONY: distclean
distclean: clean
	@rm -vrf node_modules

.PHONY: docker
docker: distclean
	@docker/build.sh

.PHONY: build-deps
build-deps:
	npm install

.PHONY: build-js
build-js:
	npx webpack build $(WEBPACK_FLAGS)

.PHONY: build-go
build-go:
	@rm -rf static/pkg
	@install -m 0750 -d static/pkg
	@install -v -m 0640 -t static/pkg "`go env GOROOT`/misc/wasm/wasm_exec.js"
	go build -o static/pkg/anonchess.wasm

.PHONY: build
build: build-go build-js

.PHONY: test
test:
	@go test -exec="`go env GOROOT`/misc/wasm/go_js_wasm_exec"

.PHONY: release
release:
	@rm -rf static/pkg static/ui
	@$(MAKE) build WEBPACK_FLAGS='--mode production'

.PHONY: publish
publish:
	@rm -rf publish
	@install -m 0755 -d ./publish

	@./vendor-publish.sh

	@install -m 0755 -d ./publish/w3css
	@install -m 0755 -d ./publish/w3css/4
	@install -v -m 0644 -t ./publish/w3css/4 ./static/w3css/4/w3.css

	@install -v -m 0644 -t ./publish ./static/*.html

	@install -m 0755 -d ./publish/css
	@install -v -m 0644 -t ./publish/css ./static/css/*.css

	@install -m 0755 -d ./publish/pkg
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/wasm_exec.js
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess.wasm

	@install -m 0755 -d ./publish/ui
	@install -v -m 0644 -t ./publish/ui ./static/ui/bundle.js
