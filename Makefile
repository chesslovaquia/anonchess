WASM_FLAGS ?= --dev

.PHONY: default
default: build

.PHONY: all
all:
	@$(MAKE) build
	@$(MAKE) test
	@$(MAKE) publish

.PHONY: clean
clean:
	@rm -vrf Cargo.lock pkg publish static/pkg

.PHONY: distclean
distclean: clean
	@rm -vrf target

.PHONY: docker
docker:
	@./docker/build.sh

.PHONY: fmt
fmt:
	@rustfmt -l ./src/*.rs

.PHONY: build-deps
build-deps:
	@./docker/build-deps.sh

.PHONY: build
build:
	wasm-pack build --target web --out-dir ./static/pkg $(WASM_FLAGS)

.PHONY: release
release:
	@rm -vrf ./static/pkg
	@$(MAKE) build WASM_FLAGS=--release

.PHONY: test
test:
	@wasm-pack test --node --lib

.PHONY: publish
publish:
	@rm -rf ./publish
	@install -v -m 0755 -d ./publish
	@install -v -m 0644 -t ./publish ./static/*.html ./static/*.css
	@install -v -m 0755 -d ./publish/w3css
	@install -v -m 0755 -d ./publish/w3css/4
	@install -v -m 0644 -t ./publish/w3css/4 ./static/w3css/4/w3.css
	@install -v -m 0755 -d ./publish/pkg
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess*
