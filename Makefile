test:
	go test -cover ./... -v

coverage:
	rm -fr coverage
	mkdir coverage
	go test -coverprofile=coverage/cover.out ./...
	go tool cover -html=coverage/cover.out -o=coverage/cover.html

.PHONY: test coverage
