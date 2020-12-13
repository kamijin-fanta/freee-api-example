FORCE:
.PHONY: FORCE


oapi: FORCE
	sudo docker run --rm -v `pwd`:/app openapitools/openapi-generator-cli generate -i https://raw.githubusercontent.com/freee/freee-api-schema/master/v2020_06_15/open-api-3/api-schema.json -g go -o /app/pkg/freee --additional-properties=packageName=freee,isGoSubmodule=false
	sudo chown ${USER}:${USER} -R ./pkg/freee/
	rm ./pkg/freee/go.mod ./pkg/freee/go.sum
