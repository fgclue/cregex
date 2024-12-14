make:
	mkdir -p build
	go build -o build/cregex main.go

clean:
	rm -fr build/cregex