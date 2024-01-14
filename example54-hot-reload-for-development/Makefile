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
