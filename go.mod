module github.com/congcongke/restapi

go 1.16

require (
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/runtime v0.19.5
	github.com/go-openapi/spec v0.19.3
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.5
	github.com/goharbor/harbor/src/server/v2.0/models v0.0.0-00010101000000-000000000000
)

replace github.com/goharbor/harbor/src/server/v2.0/models => github.com/congcongke/models v0.0.0-20220317074912-88e04967720f
