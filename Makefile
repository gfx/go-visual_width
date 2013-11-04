
test:
	go test --test.v=true

bench:
	go test --test.bench=Bench

data:
	author/east-asian-width > data.go


.PHONY: data test
