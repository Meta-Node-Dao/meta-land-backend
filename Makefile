#.PHONY: build

VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`


gofmt:
		echo "formating with gofmt..."
		gofmt -s -w ${GOFILES}
		echo "complete with gofmt"
govet:
		echo "doing statics check..."
		go vet $(VETPACKAGES)