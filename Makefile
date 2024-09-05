MAKEFLAGS += --no-print-directory
WEBPACK_FLAGS ?= --mode development

.PHONY: default
default: build

.PHONY: all
all:
	@echo "AnonChess build"
	@echo "  `go version`"
	@echo "  nodejs `node --version`"
	@echo "  npm `npm --version`"
	@echo "  npx `npx --version`"
	@$(MAKE) build-deps test release publish

.PHONY: clean
clean:
	@TPL=clean go run ./tpl
	@rm -vrf static/pkg static/ui publish docker/build
	@rm -vf ui/wasm_exec.js

.PHONY: distclean
distclean: clean
	@rm -vrf node_modules

.PHONY: docker
docker:
	@docker/build.sh

#
# build
#
.PHONY: build-deps
build-deps:
	npm install

.PHONY: build-js
build-js:
	@$(MAKE) about
	@$(MAKE) home-js
	@$(MAKE) play-js

.PHONY: build-go
build-go:
	@rm -f ui/wasm_exec.js
	@install -v -m 0640 -t ui "`go env GOROOT`/misc/wasm/wasm_exec.js"
	@rm -rf static/pkg
	@install -m 0750 -d static/pkg
	@$(MAKE) home-go
	@$(MAKE) play-go

.PHONY: html
html:
	TPL=gen go run ./tpl
	@install -v -m 0640 -t static ./package.json

.PHONY: ui
ui: build-js html

.PHONY: build
build: build-go ui

#
# about
#
.PHONY: about
about:
	npx webpack build $(WEBPACK_FLAGS) --entry ./ui/about/index.js -o $(PWD)/static/ui/about

#
# Home
#
.PHONY: home-js
home-js:
	npx webpack build $(WEBPACK_FLAGS) --entry ./ui/home/index.js -o $(PWD)/static/ui/home

.PHONY: home-go
home-go:
	wasm/build.sh -o static/pkg/anonchess-home.wasm ./wasm/home

.PHONY: home
home: home-go home-js

#
# Play
#
.PHONY: play-js
play-js:
	npx webpack build $(WEBPACK_FLAGS) --entry ./ui/play/index.js -o $(PWD)/static/ui/play

.PHONY: play-go
play-go:
	wasm/build.sh -o static/pkg/anonchess-play.wasm ./wasm/play

.PHONY: play
play: play-go play-js

#
# test
#
.PHONY: test
test:
	go test . ./cmd/... ./lib/...
	wasm/test.sh ./wasm/...

#
# release
#
.PHONY: release
release:
	@$(MAKE) build WEBPACK_FLAGS='--mode production'

#
# publish
#
.PHONY: publish
publish:
	@rm -rf publish
	@install -m 0755 -d ./publish
	@install -m 0755 -d ./publish/pkg
	@install -m 0755 -d ./publish/ui
#	vendor
	@./static/vendor-publish.sh
	@install -m 0755 -d ./publish/w3css
	@install -m 0755 -d ./publish/w3css/4
	@install -v -m 0644 -t ./publish/w3css/4 ./static/w3css/4/w3.css
#	css
	@install -m 0755 -d ./publish/css
	@install -v -m 0644 -t ./publish/css ./static/css/*.css
#	root
	@install -v -m 0644 -t ./publish ./static/*.html ./static/*.js
	@install -v -m 0644 -t ./publish ./static/package.json
#	About
	@install -m 0755 -d ./publish/ui/about
	@install -v -m 0644 -t ./publish/ui/about ./static/ui/about/main.js
#	Home
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess-home.wasm
	@install -m 0755 -d ./publish/ui/home
	@install -v -m 0644 -t ./publish/ui/home ./static/ui/home/main.js
#	Play
	@install -v -m 0644 -t ./publish/pkg ./static/pkg/anonchess-play.wasm
	@install -m 0755 -d ./publish/play
	@install -v -m 0644 -t ./publish/play ./static/play/*.html
	@install -m 0755 -d ./publish/ui/play
	@install -v -m 0644 -t ./publish/ui/play ./static/ui/play/main.js
