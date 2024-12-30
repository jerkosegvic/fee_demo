#!/bin/bash
cd shell
go build
cd ..
fee shell/demoshell --target amd64 > fee_shell.py