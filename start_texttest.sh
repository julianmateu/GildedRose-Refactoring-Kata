#!/bin/sh

if [ ! -d "venv" ]; then
    python -m venv venv
fi
venv/bin/pip install texttest

export GOPATH=$(pwd)/go
venv/bin/texttest -d . -con "$@"
