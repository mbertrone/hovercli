#!/bin/bash

#set -x

echo "FILE"
echo $1
echo ""

set -x

#Install hovercli
go install github.com/mbertrone/hovercli/

#Launch iovisorovndn using file parameter
sudo $GOPATH/bin/hovercli -hover http://127.0.0.1:5002
