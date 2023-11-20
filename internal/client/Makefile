.PHONY: generate-client
generate-client:
	oapi-codegen -config ./config/client.yml ./openapi/openapi.yml

.PHONY: generate-models
generate-models:
	oapi-codegen -config ./config/models.yml ./openapi/openapi.yml

.PHONY: generate-spec
generate-spec:
	oapi-codegen -config ./config/spec.yml ./openapi/openapi.yml

.PHONY: generate
generate:
	make generate-client
	make generate-models
	make generate-spec
