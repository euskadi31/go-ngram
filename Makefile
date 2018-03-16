.PHONY: all clean fmt vet test cover bench release

PACKAGES = $(shell go list ./... | grep -v vendor)

release:
	@echo "Release v$(version)"
	@git pull
	@git checkout master
	@git pull
	@git checkout develop
	@git flow release start $(version)
	@git flow release finish $(version) -p -m "Release v$(version)"
	@git checkout develop
	@echo "Release v$(version) finished."

all: test

clean:
	@go clean -i ./...

fmt:
	@go fmt $(PACKAGES)

vet:
	@go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

travis:
	@for PKG in $(PACKAGES); do go test -ldflags '-s -w $(LDFLAGS)' -cover -covermode=count -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

cover: test
	@echo ""
	@mkdir -p coverage/
	@echo "mode: set" > ./coverage/test.cov
	@for PKG in $(PACKAGES); do if [ -f $$GOPATH/src/$$PKG/coverage.out ]; then tail -q -n +2 $$GOPATH/src/$$PKG/coverage.out >> ./coverage/test.cov; fi; done;
	@go tool cover -func ./coverage/test.cov
	#@go tool cover -html=./coverage/test.cov

bench:
	@for PKG in $(PACKAGES); do go test -benchmem -bench=. $$PKG || exit 1; done;

