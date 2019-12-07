#############################################################################
# License: MIT
#
# (C) 2019 Juhapekka Piiroinen <juhapekka.piiroinen@1337.fi>
# All Rights Reserved.
#############################################################################
GOPATH=$(PWD)
export GOPATH

all:
	go get -d ./...
	go build  -o bin/helloworld ./...

clean:
	rm -rf bin
