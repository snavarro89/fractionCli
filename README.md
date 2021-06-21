# fractionCli
Golang CLI Tool to perfom simple operations on fractions

# Requirements

Latest version of golang

# Setup

git clone https://github.com/snavarro89/fractionCli

Build on desired platform

env GOOS=darwin GOARCH=amd64 go build -o bin/cli   //OS X

bin/cli [operations]

# TODO

1) Manage precision, on rounding up to two decimals fraction result might be different
2) Convert the last fractoin to mixed number
3) Optimize reading the operands to use a single function to compile all 4 allowed operands
