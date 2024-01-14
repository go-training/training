EXECUTABLE := web
LDFLAGS ?= -X main.version=$(VERSION) -X main.commit=$(COMMIT)
SOURCES ?= $(shell find . -name "*.go" -type f)
GO ?= go

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

all: build

build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO) build -v -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o bin/$@ main.go

# install air command
.PHONY: air
air:
	@hash air > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi

# run air
.PHONY: dev
dev: air
	air --build.cmd "make build" --build.bin "bin/web"
