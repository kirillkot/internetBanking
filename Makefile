vendor: Gopkg.lock Gopkg.toml
	rm -rf vendor/
	dep ensure -v

.PHONE: build
build: vendor
	docker build -t banking .

.PHONE: run
run: build
	-docker run -ti --rm -p 8080:8080 banking
