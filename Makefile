VERSION = $(shell head -1 VERSION)

all: build

deps:
	go get -d

build: deps
	go build

test: deps
	go test ./...

install: deps
	go install

tag:
	git commit -m "Release: v$(VERSION)"
	git tag v$(VERSION)

# quick way to check if it really works
integration_test: build
	./gotr en en foo > test.file
	test $$(cat test.file) = 'foo'
	rm test.file

build_all:
	mkdir -p bin
	@for goos in linux windows darwin ; do \
		for goarch in amd64 386; do \
			echo "building bin/gotr_$$goos-$$goarch"; \
			GO_ENABLED=0 GOOS=$$goos GOARCH=$$goarch go build -o "./bin/gotr_$$goos-$$goarch" gotr.go; \
		done; \
	done

homebrew:
	mkdir -p bin/bin
	cp bin/gotr_darwin-amd64 bin/gotr
	tar -czvf "gotr-v$(VERSION).tar.gz" bin/gotr README.md
