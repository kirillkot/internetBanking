.PHONE: build
build:
	docker build -t banking .

run: build
	-docker run -ti --rm -p 8080:8080 banking 
