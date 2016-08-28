check: vet lint test

init:
	go get -t -v ./...
	go get github.com/golang/lint/golint
	go get -u github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover

vet:
	go vet ./...

test:
	go test -v -race ./...

lint:
	golint -set_exit_status ./...

goveralls:
	$HOME/gopath/bin/goveralls -service=travis-ci