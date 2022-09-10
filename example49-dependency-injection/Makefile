EXECUTABLE := user-service
SOURCES ?= $(shell find . -name "*.go" -type f)
GO ?= go

build: $(EXECUTABLE)

$(EXECUTABLE): $(SOURCES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o bin/$@ .
