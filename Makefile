# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

PKG=github.com/siongui/responsive-embed-generator
GHPAHE_DIR=docs

devserver: test local js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run gopherjs/devserver/server.go -dir=gopherjs

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build gopherjs/*.go -o gopherjs/app.js

test: fmt
	@echo "\033[92mTest ...\033[0m"
	@go test -v

reset_github_repo:
	@echo "\033[92mWarning: Reset https://${PKG}.git ...\033[0m"
	git remote add origin https://${PKG}.git
	git push --force --set-upstream origin master

fmt:
	@echo "\033[92mGo fmt source code ...\033[0m"
	@go fmt *.go
	@go fmt gopherjs/*.go
	@go fmt gopherjs/devserver/*.go

publish:
	@echo "\033[92mPublish to GitHub Pages ...\033[0m"
	[ -d $(GHPAHE_DIR) ] || mkdir -p $(GHPAHE_DIR)
	cp gopherjs/index.html $(GHPAHE_DIR)/index.html
	cp gopherjs/app.js $(GHPAHE_DIR)/app.js
	git add $(GHPAHE_DIR)/
	git commit -m "publish to GitHub Pages"
	git push

install:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs
	@echo "\033[92mInstalling godom ...\033[0m"
	go get -u github.com/siongui/godom

local:
	@[ -d src/${PKG}/ ] || mkdir -p src/${PKG}/
	@cp *.go src/${PKG}/
