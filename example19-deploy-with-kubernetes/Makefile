ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= latest
endif

prepare:
	sed -ie "s/VERSION/$(VERSION)/g" deployment.yml
	sed -ie "s/THIS_STRING_IS_REPLACED_DURING_BUILD/$(shell date)/g" deployment.yml
	cat deployment.yml

build:
ifneq ($(DRONE_TAG),)
	go build -v -ldflags "-X main.version=$(DRONE_TAG)" -a -o release/linux/amd64/hello
else
	go build -v -ldflags "-X main.version=$(DRONE_COMMIT)" -a -o release/linux/amd64/hello
endif
