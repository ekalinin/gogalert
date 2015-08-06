NAME=gogalert
EXEC=${NAME}.bin
GOVER=1.5rc1
ENVNAME=${NAME}${GOVER}
GHBASE=github.com/ekalinin
GHNAME=${GHBASE}/${NAME}

deps:
	@go get gopkg.in/alecthomas/kingpin.v2

# with https://github.com/ekalinin/envirius
env-create:
	@bash -c ". ~/.envirius/nv && nv mk ${ENVNAME} --go-prebuilt=${GOVER}"

env-fix:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make env-fix-paths'"

env-fix-paths:
	@if [ -d "${GOPATH}/src/${GHNAME}" ]; then \
		echo "Already fixed. No actions need."; \
	else \
		mkdir -p ${GOPATH}/src/${GHBASE}; \
		ln -s `pwd` ${GOPATH}/src/${GHBASE}; \
	fi

env-deps:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make deps'"

env-build:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make build'"

env-init: env-create env-fix env-deps

env:
	@bash -c ". ~/.envirius/nv && nv use ${ENVNAME}"

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
