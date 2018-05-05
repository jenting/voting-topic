#!/bin/bash

set -o errexit

## make
make clean
make all

## run binary
./voting-topic
