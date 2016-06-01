.PHONY: env build

NAME=enviriusx
EXEC=nvx
GOVER=1.6.2
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

env-tools:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make tools'"


env-build:
	@bash -c ". ~/.envirius/nv && nv do ${ENVNAME} 'make build'"

env-init: env-create env-fix env-deps env-tools

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

tools:
	@go get -u -v github.com/nsf/gocode
	@go get -u -v github.com/rogpeppe/godef
	@go get -u -v github.com/golang/lint/golint
	#@go get -u -v github.com/lukehoban/go-find-references
	@go get -u -v github.com/lukehoban/go-outline
	@go get -u -v sourcegraph.com/sqs/goreturns
	@go get -u -v golang.org/x/tools/cmd/gorename
	@go get -u -v github.com/tpng/gopkgs
	@go get -u -v github.com/newhook/go-symbols

build:
	@go build -a -tags netgo \
		--ldflags '-s -extldflags "-lm -lstdc++ -static"' \
		-o ${EXEC} ./cmd/nvx/main.go
