SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

RELEASE_DIR=release
BINARY=loop
LDFLAGS='-extldflags "-static"'
TARGEOS = darwin linux

.DEFAULT_GOAL: build
.PHONY: build clean release

build: $(SOURCES)
	CGO_ENABLED=0 go build --ldflags ${LDFLAGS} -o ${BINARY} $(SOURCES)

define build
	mkdir -p ${RELEASE_DIR}/$(1);
	GOOS=$(1) CGO_ENABLED=0 go build --ldflags ${LDFLAGS} -o ${RELEASE_DIR}/$(1)/${BINARY} $(SOURCES)
endef

release: $(SOURCES)
	$(call build,darwin)
	$(call build,linux)

clean:
	rm -rf ${BINARY}
	rm -rf ${RELEASE_DIR}
