package catalog

import (
	"testing"

	"github.com/databrickslabs/terraform-provider-databricks/qa"
)

func TestSchemasData(t *testing.T) {
	qa.ResourceFixture{
		Fixtures: []qa.HTTPFixture{
			{
				Method:   "GET",
				Resource: "/api/2.0/unity-catalog/schemas?catalog_name=a",
				Response: Schemas{
					Schemas: []SchemaInfo{
						{
							Name: "a.c",
						},
						{
							Name: "a.d",
						},
					},
				},
			},
		},
		Resource:    DataSourceSchemas(),
		HCL:         `catalog_name = "a"`,
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ApplyNoError(t)
}

func TestSchemasData_Error(t *testing.T) {
	qa.ResourceFixture{
		Fixtures:    qa.HTTPFailures,
		Resource:    DataSourceSchemas(),
		Read:        true,
		NonWritable: true,
		ID:          "_",
	}.ExpectError(t, "I'm a teapot")
}
