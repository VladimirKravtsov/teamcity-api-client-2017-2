This is comprehensive Go library to interact with TeamCity 2017.2 API server. 
This client is generated via [go-swagger](https://github.com/go-swagger/go-swagger) tool, using official swagger spec, provided by TeamCity server. 

## Get Started
```go
package main

import (
	"fmt"
	teamcityclient "github.com/VladimirKravtsov/teamcity-api-client-2017-2/client"
	"github.com/VladimirKravtsov/teamcity-api-client-2017-2/client/build"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"net/url"
	"time"
)

func main() {
	tcURL, err := url.Parse("https://teamcity.yourcompany.com/httpAuth")
	if err != nil {
		panic(err)
	}

	runtime := httptransport.New(tcURL.Host, tcURL.Path, []string{tcURL.Scheme})
	teamcity := teamcityclient.New(runtime, strfmt.Default)

	tcUser, tcPassword := "user", "password"
	authentication := httptransport.BasicAuth(tcUser, tcPassword)

	horizonDuration, err := strfmt.ParseDuration("48h")
	sinceTime := time.Now().UTC().Add(-horizonDuration)
	formattedSinceTime := sinceTime.Format("20060102T150405-0700")
	projectId := "Your_ProjId"
	buildLocator := fmt.Sprintf("affectedProject:(id:%s),defaultFilter:false,queuedDate:(date:%s,condition:after)", projectId, formattedSinceTime)
	fields := "build(buildTypeId,status,state,webUrl,id,problemOccurrences(*),queuedDate)"

	buildsResponse, err := teamcity.Build.ServeAllBuilds(
		build.NewServeAllBuildsParams().WithLocator(&buildLocator).WithFields(&fields), authentication)

	if err != nil {
		panic(err)
	}

	bytes, _ := buildsResponse.Payload.MarshalBinary() //serialize to JSON
	fmt.Println(string(bytes))
}
```

How to compose locators and fields params is described [here (TC 2017.2)](https://confluence.jetbrains.com/display/TCD10/REST+API#RESTAPI-BuildLocator)


## How to re-generate TeamCity API client 

This project contains auto-generated code, which is for querying TeamCity API server.  
At the root of ci-dashboard dir you can find `teamcity-swagger-spec.json`(the content is fetched from https://teamcity.mesosphere.io/httpAuth/app/rest/swagger.json). 
Which is essentially comprehensive [swagger](https://swagger.io) specification of TeamCity API. 
By passing the spec to [go-swagger](https://github.com/go-swagger/go-swagger) tool 
we are able to get working client code(API requests and DTO models) automatically. This approach enables to be type-safe 
and maintain consistency the client with actual API. More about benefits you can find on the its site.
  
As for now, the client provides only Build API client.  
  
Normally, you no need to re-generate the client, but in case of TC server was updated or if you will need access to other 
APIs re-generation should take place. There are 2 possible options - 

1) Install [go-swagger](https://github.com/go-swagger/go-swagger) binary locally, for instance, using homebrew. Then run -
```swagger generate client -f teamcity-swagger-spec.json --skip-validation --default-scheme=https```

2) Execute generation via predefined target in makefile `make generate-teamcity-client`. 
It will run generation inside container, using appropriate docker image `quay.io/goswagger/swagger`.