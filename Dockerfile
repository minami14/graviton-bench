FROM golang:1.17

WORKDIR $GOPATH/src/github.com/minami14/graviton-bench
COPY . .

CMD ["go", "test", "-bench", ".", "-benchmem"]
