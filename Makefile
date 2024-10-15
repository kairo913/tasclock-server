MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/scripts

lint:
	$(ROOT_DIR)/lint.sh
