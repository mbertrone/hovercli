#!/bin/bash

set -x

#Install hovercli
go install github.com/mbertrone/hovercli/

#Launch hovercli using file parameter
sudo $GOPATH/bin/hovercli -hover http://127.0.0.1:5002
