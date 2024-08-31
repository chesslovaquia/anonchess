.PHONY: default
default: build

.PHONY: all
all:
	@env
	@echo "GOOS=`go env GOOS` GOARCH=`go env GOARCH`"
	@echo "  `go version`"
	@echo "  nodejs `node --version`"
	@$(MAKE) build
	@$(MAKE) test
	@$(MAKE) publish

.PHONY: clean
clean:
	@rm -vrf static/pkg publish

.PHONY: distclean
distclean: clean

.PHONY: docker
docker:
	@docker/build.sh

.PHONY: build
build:
	@rm -rf static/pkg
	@install -m 0750 -d static/pkg
	@install -v -m 0640 -t static/pkg "`go env GOROOT`/misc/wasm/wasm_exec.js"
	@go build -o static/pkg/main.wasm

.PHONY: test
test:
	@go test -exec="`go env GOROOT`/misc/wasm/go_js_wasm_exec"

.PHONY: publish
publish:
	@rm -rf publish
	@install -v -m 0755 -d ./publish

	@./vendor-publish.sh

	@install -v -m 0755 -d ./publish/w3css
	@install -v -m 0755 -d ./publish/w3css/4
	@install -v -m 0644 -t ./publish/w3css/4 ./static/w3css/4/w3.css

	@install -v -m 0644 -t ./publish ./static/*.html

	@install -v -m 0755 -d ./publish/css
	@install -v -m 0644 -t ./publish/css ./static/css/*.css

	@install -v -m 0755 -d ./publish/pkg
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/wasm_exec.js
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess.wasm

	@install -v -m 0755 -d ./publish/ui
	@install -v -m 0644 -t ./publish/ui ./static/ui/bundle.js
