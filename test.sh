#!/usr/bin/env bash
go test ./f3client/... -run=^Test_Unit_ -v -coverprofile=./results/unitcover.out
go test ./f3client/... -p 1 -run=^Test_Integration_ -v -coverprofile=./results/itcover.out