clean:
	rm -rf dist/*

build-darwin: export GOOS=darwin
build-darwin: clean
	go build -o dist/create-image-placeholder
	tar -czf dist/create-image-placeholder.tar.gz -C dist create-image-placeholder
	rm dist/create-image-placeholder