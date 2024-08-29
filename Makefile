WASM_FLAGS ?= --dev
WEBPACK_FLAGS ?= --mode development

.PHONY: default
default: build

.PHONY: all
all:
	@$(MAKE) build
	@$(MAKE) test
	@$(MAKE) release
	@$(MAKE) publish

.PHONY: clean
clean:
	@rm -vrf Cargo.lock pkg publish static/pkg docker/build

.PHONY: distclean
distclean: clean
	@rm -vrf target node_modules

.PHONY: docker
docker: distclean
	@./docker/build.sh

.PHONY: fmt
fmt:
	@rustfmt -l ./src/*.rs

.PHONY: rs-deps
rs-deps:
	@./docker/build-deps.sh

.PHONY: build-rs
build-rs:
	wasm-pack build --target web --out-dir ./static/pkg $(WASM_FLAGS)

.PHONY: js-deps
js-deps:
	@npm update

.PHONY: build-js
build-js:
	npx webpack build $(WEBPACK_FLAGS)

.PHONY: build-deps
build-deps:
	@$(MAKE) rs-deps
	@$(MAKE) js-deps

.PHONY: build
build:
	@$(MAKE) build-rs
	@$(MAKE) build-js

.PHONY: release
release:
	@rm -vrf static/pkg static/ui
	@$(MAKE) build WASM_FLAGS=--release WEBPACK_FLAGS='--mode production'

.PHONY: test
test:
	@wasm-pack test --node --lib

.PHONY: publish
publish:
	@rm -rf ./publish
	@install -v -m 0755 -d ./publish

	@./vendor-publish.sh

	@install -v -m 0755 -d ./publish/w3css
	@install -v -m 0755 -d ./publish/w3css/4
	@install -v -m 0644 -t ./publish/w3css/4 ./static/w3css/4/w3.css

	@install -v -m 0644 -t ./publish ./static/*.html ./static/*.css

	@install -v -m 0755 -d ./publish/pkg
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess*

	@install -v -m 0755 -d ./publish/ui
	@install -v -m 0644 -t ./publish/ui ./static/ui/bundle.js
