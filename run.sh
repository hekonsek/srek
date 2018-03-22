#!/usr/bin/env bash

go-bindata ansible/
go run srek.go bindata.go "$@"