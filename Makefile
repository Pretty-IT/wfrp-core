lib:
	go build -buildmode=c-shared -o build/wfrp-core.so export/export.go

exec:
	go build -o build/wfrp-core main.go

test:
	go test -v ./...
