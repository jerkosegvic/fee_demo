#!/bin/bash
cd odbrojavanje
gcc odbrojavanje.c -o odbrojavanje
cd ..
fee odbrojavanje/odbrojavanje --target amd64 > fee_odb.py