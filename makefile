firestore:
	go run cmd/mockserver/firestore/main.go

bytestream:
	go run cmd/mockserver/bytestream/main.go

bidi:
	go run cmd/bidistream/main.go

bidimock:
	go run cmd/bidistream/main.go -mock=true

client:
	go run cmd/clientstream/main.go

server:
	go run cmd/serverstream/main.go

lint:
	gofumpt -l -w .
	
.PHONY: lint firestore bytestream bidi bidimock client server