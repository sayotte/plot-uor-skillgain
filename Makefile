SHELL=/bin/bash
ROOT = $(shell pwd)
ROOTBASENAME = $(shell basename ${ROOT})
PARENTDIR := $(shell dirname `pwd`)
export GOPATH := ${PARENTDIR}/${ROOTBASENAME}-gopath
export PATH := ${PATH}:${GOPATH}/bin
PROJ_ROOT = github.com/sayotte/plot-uor-skillgain
TEST_FLAGS ?= -cover

.PHONY: fmt install

install:
	go fmt ${PROJ_ROOT}/...
	go install ${PROJ_ROOT}/...

test:
	go test ${TEST_FLAGS} -timeout=1m ${PROJ_ROOT}/...

shell:
	bash
