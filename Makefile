KERNEL_VERSION = $(shell uname -r)

default: complie

.PHONY: clean
clean:
	rm -rf headers
	rm -rf linux

.PHONY: test
test: headers
	CGO_ENABLED=1 CGO_CFLAGS="-I$$(pwd)/headers/include" go test ./tracer ./printer ./filter

linux:
	# sudo yum install kernel-devel-$(KERNEL_VERSION) kernel-headers-$(KERNEL_VERSION) -y
	# cp -r /usr/src/kernels/$(uname -r) ./linux
	sudo yum install kernel-devel kernel-headers -y
	cp -r /usr/src/kernels/5.14.0-362.24.2.el9_3.x86_64 ./linux
headers: linux
	cd linux && make headers_install ARCH=x86_64 INSTALL_HDR_PATH=../headers

.PHONY: build
build: headers
	CGO_ENABLED=1 CGO_CFLAGS="-I$$(pwd)/headers/include" go build --ldflags '-linkmode external -extldflags "-static"'

complie:
	CGO_ENABLED=1 CGO_CFLAGS="-I$$(pwd)/headers/include" go build --ldflags '-linkmode external -extldflags "-static"'

.PHONY: install
install: headers
	CGO_ENABLED=1 CGO_CFLAGS="-I$$(pwd)/headers/include"  go install --ldflags '-linkmode external -extldflags "-static"'

.PHONY: demo
demo: build
	./grace -- cat /dev/null
