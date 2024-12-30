#!/bin/bash
cd shell
go build
cd ..
fee shell/demoshell --target amd64 > output.py