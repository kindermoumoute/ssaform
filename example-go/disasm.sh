#!/bin/bash -e
#
# All material is licensed under the Apache License Version 2.0, January 2016
# http://www.apache.org/licenses/LICENSE-2.0
#
# You need to install Go before running this script
# https://golang.org/doc/install
for i in `seq 1 2`;
do
    cd "same-binary-$i"
    go build
    go tool objdump -s main.Test "same-binary-$i"
    rm "same-binary-$i"
    cd ..
done