package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSource_basic(t *testing.T) {
	dsn := "data.null_data_source.test"
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "random"),
					resource.TestCheckResourceAttr(dsn, "has_computed_default", "default"),
				),
			},
		},
	})
}

func TestAccDataSource_inputs(t *testing.T) {
	dsn := "data.null_data_source.test"
	resource.UnitTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceConfig_inputs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "random"),
					resource.TestCheckResourceAttr(dsn, "outputs.foo", "bar"),
				),
			},
		},
	})
}

const testAccDataSourceConfig_basic = `
data "null_data_source" "test" {
}
`

const testAccDataSourceConfig_inputs = `
data "null_data_source" "test" {
  inputs = {
    foo = "bar"
  }
}`
