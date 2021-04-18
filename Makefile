.PHONY: build run test watch cover watch/cover

build:
	go build -o bin/escargot cmd/escargot/main.go

run:
	go run cmd/escargot/main.go

test:
	 ginkgo --randomizeAllSpecs --randomizeSuites --failOnPending --race --trace --compilers=2 -v ./...

watch:
	 ginkgo watch --randomizeAllSpecs --failOnPending --race --trace --compilers=2 -v ./...

cover:
	 ginkgo --randomizeAllSpecs --randomizeSuites --failOnPending --race --trace --compilers=2 --cover --outputdir='coverage' -v ./...

watch/cover:
	 ginkgo watch --randomizeAllSpecs --failOnPending --race --trace --compilers=2 --cover --outputdir='coverage' -v ./...
