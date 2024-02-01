all: build install

build:
	go build -o path main.go

clean:
	rm path

install:
	cp path ~/.local/bin


.PHONY: build
