build:
	go build .

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*

setup:
	cp ./config/_app.local.json ./config/app.local.json