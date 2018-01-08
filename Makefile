
BINARY = e2-unused-vmcleanup
GOARCH = amd64
TEST_REPORT = e2-unused-vmcleanup_tests.xml
VERSION?=?
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=sidlinux22
#BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}
BUILD_DIR=$(PWD)/cmd/tools/
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: link clean test vet linux darwin windows

link:
	BUILD_DIR=${BUILD_DIR}; \
	BUILD_DIR_LINK=${BUILD_DIR_LINK}; \
	CURRENT_DIR=${CURRENT_DIR}; \
	if [ "$${BUILD_DIR_LINK}" != "$${CURRENT_DIR}" ]; then \
	    echo "Fixing symlinks for build"; \
	    rm -f $${BUILD_DIR}; \
	    ln -s $${CURRENT_DIR} $${BUILD_DIR}; \
	fi

linux: 
	echo ${BUILD_DIR}; cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PWD}/bin/${BINARY}-linux-${GOARCH} *.go ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PWD}/bin/${BINARY}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${PWD}/bin/${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

# test:
# 	cd ${BUILD_DIR}; \
# 	godep go test -v ./... 2>&1 | go2xunit -output ${TEST_REPORT} ; \
# 	cd - >/dev/null

get:
	-cd ${BUILD_DIR}; \
	go get ./... ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f bin/${BINARY}-*

.PHONY: link linux darwin windows test vet fmt clean