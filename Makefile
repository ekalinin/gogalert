NAME=gogalert
EXEC=${NAME}.bin
GOVER=1.5beta2
ENVNAME=${NAME}${GOVER}

get-deps:
	@go get gopkg.in/alecthomas/kingpin.v2

# with https://github.com/ekalinin/envirius
env-init:
	@nv mk ${ENVNAME} --go-prebuilt=${GOVER}

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

