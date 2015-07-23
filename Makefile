EXEC=gogalert
ENVNAME=gogalert1.5beta2

get-deps:
	@go get gopkg.in/alecthomas/kingpin.v2

# with https://github.com/ekalinin/envirius
env-init:
	@nv mk ${ENVNAME} --go-prebuilt=1.5beta2

env:
	@bash -c ". ~/.envirius/nv && nv use ${ENVNAME}"

build:
	@go build -a -tags netgo --ldflags '-s -extldflags "-lm -lstdc++ -static"' -i -o ${EXEC}

run:
	@go run gogalert.go

runxml:
	@gogalert response/test.xml disk_free
