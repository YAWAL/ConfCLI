.PHONY: dependencies
dependencies:
		echo "Installing dependencies"
		dep ensure

.PHONY: code-quality
code-quality:
	gometalinter --vendor --tests --skip=mock --exclude='_gen.go' --disable=gotype --disable=errcheck --disable=gas --disable=dupl --deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml

install-helpers:
	echo "Installing GoMetaLinter"
	go get -u github.com/alecthomas/gometalinter
	echo "Installing linters"
	gometalinter --install
	echo "Installing Dep"
	curl    https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

tests: dependencies
	echo "Tests"
	go test ./cmd

build: dependencies
	echo "Build"
	go build

sonar-scanner:
	sonar-scanner   -Dsonar.projectKey=erp1   -Dsonar.sources=.   -Dsonar.host.url=http://localhost:9000   -Dsonar.login=6ee5e1c431f83cfa03cc7e6eb8a3ce9374c3a2b4
