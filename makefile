    # Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOGET=$(GOCMD) get
    SWAGGER=docker run --rm -it -e GOPATH=/go -v "$(PWD)":/usr/src/github.com/VladimirKravtsov/teamcity-api-client-2017-2 -w /usr/src/github.com/github.com/VladimirKravtsov/teamcity-api-client-2017-2 quay.io/goswagger/swagger

    all: deps run
    generate-teamcity-client:
		docker pull quay.io/goswagger/swagger
		$(SWAGGER) generate client -f teamcity-swagger-spec.json -t ./teamcity --skip-validation --default-scheme=https
    build:
		$(GOBUILD)
    clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
    run: build
		./$(BINARY_NAME)
    deps:
		$(GOGET)

    # Cross compilation
    build4linux:
		env GOOS=linux GOARCH=amd64 $(GOBUILD)
