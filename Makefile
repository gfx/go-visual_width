
test:
	go test --test.v=true

data:
	author/east-asian-width > data.go


.PHONY: data test
