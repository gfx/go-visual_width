
test: deps
	go test --test.v=true

bench:
	go test --test.bench=Bench

data:
	author/east-asian-width > data.go

deps:
	go get "github.com/stretchr/testify/assert"


.PHONY: data test deps
