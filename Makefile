SHELL = /bin/bash

install: build 
	chmod 555 ./wi
	mv ./wi /usr/local/bin/wi

build:
	go build . -o wi

run:
	go run .

clean:  
	rm -f ./wi

.PHONY: build install run clean