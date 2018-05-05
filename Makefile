#!/usr/bin/make -f

all:
	go build
	
clean:
	@[ -f voting-topic ] && rm -r voting-topic || true

.PHONY: all clean