clean:
	rm -rf dist/*

build-darwin: export GOOS=darwin
build-darwin: clean
	go build -o dist/preview