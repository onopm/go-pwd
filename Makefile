
.PHONY: clean test

gpwd: pwd.go
	go build -o gpwd pwd.go

test:
	go test -v .

clean:
	rm gpwd

