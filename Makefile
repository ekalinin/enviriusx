.PHONY: env build

NAME=enviriusx
EXEC=nvx
GOVER=1.5
ENVNAME=${NAME}${GOVER}
GHBASE=github.com/ekalinin
GHNAME=${GHBASE}/${NAME}

#
# For virtual environment create with
# https://github.com/ekalinin/envirius
#
env-create:
	@bash -c ". ~/.envirius/nv && nv mk ${ENVNAME} --go-prebuilt=${GOVER}"

env-fix:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make fix-paths'"

env-deps:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make deps'"

env-build:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make build'"

env-init: env-create env-fix env-deps

env:
	@bash -c ". ~/.envirius/nv && nv use ${ENVNAME}"

#
# Other targets
#

deps:
	@go get gopkg.in/alecthomas/kingpin.v2

fix-paths:
	@if [ -d "${GOPATH}/src/${GHNAME}" ]; then \
		echo "Already fixed. No actions need."; \
	else \
		mkdir -p ${GOPATH}/src/${GHBASE}; \
		ln -s `pwd` ${GOPATH}/src/${GHBASE}; \
	fi

build:
	@go build -a -tags netgo \
		--ldflags '-s -extldflags "-lm -lstdc++ -static"' \
		-o ${EXEC} ./cmd/nvx/main.go

