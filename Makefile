.PHONY: build

build:
	go build -o chronoruler main.go

clean:
	rm -f chronoruler
