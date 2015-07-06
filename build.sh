#!/bin/bash

docker build -t noop .
docker run --rm noop bash -c 'cd /go/bin && tar -c noop*' | tar -xv
