NAME=gogalert
EXEC=${NAME}.bin
GOVER=1.5beta3
ENVNAME=${NAME}${GOVER}
GHNAME=github.com/ekalinin/${NAME}

get-deps:
	@go get gopkg.in/alecthomas/kingpin.v2

# with https://github.com/ekalinin/envirius
env-init:
	@bash -c ". ~/.envirius/nv && nv mk ${ENVNAME} --go-prebuilt=${GOVER}"

env:
	@bash -c ". ~/.envirius/nv && nv use ${ENVNAME}"

env-fix:
	@if [ -d "$GOPATH/src/${GHNAME}" ]; then \
		mkdir -p $GOPATH/src/${GHNAME}; \
		ln -s `pwd` $GOPATH/src/{$GHNAME}; \
	else \
		echo "Already fixed. No actions need."; \
	fi

build:
	@go build -a -tags netgo \
			--ldflags '-s -extldflags "-lm -lstdc++ -static"' \
			-i -o ${EXEC}

run:
	@go run gogalert.go \
		--file=ganglia-responses/ganglia-meta-response.xml \
		--metric=mem_free \
		--list-clusters

runxml:
	@gogalert response/test.xml disk_free

test:
	@go test -cover ./...

test-coverage: env-fix
	@rm coverage.out
	@go test -coverprofile=coverage.out ${GHNAME}/gmeta/response
	@go tool cover -html=coverage.out
