#!/bin/bash

ENV_NAME="(devenv)"
if ! [[ "$PS1" =~ "$ENV_NAME" ]]; then
    export PS1=$ENV_NAME' '$PS1;
fi

PROJECT_GOPATH="$(pwd)"
if ! [[ "$GOPATH" =~ "$PROJECT_GOPATH" ]]; then
    export GOPATH=$PROJECT_GOPATH:$GOPATH;
fi
