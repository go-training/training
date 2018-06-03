ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= latest
endif

prepare:
	sed -ie "s/VERSION/$(VERSION)/g" deployment.yml
	sed -ie "s/THIS_STRING_IS_REPLACED_DURING_BUILD/$(shell date)/g" deployment.yml
	cat deployment.yml
