test:
	go test ./... -cover

# No of lines written in Go
nolines:
	find . -name '*.go' | xargs wc -l | tail -1

