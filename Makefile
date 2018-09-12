#!/usr/bin/make -f

all:
	go build
	
clean:
	go clean

.PHONY: all clean
