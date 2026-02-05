// Copyright IBM Corp. 2017, 2026
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResource_basic(t *testing.T) {
	dsn := "null_resource.test"
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "id"),
				),
			},
		},
	})
}

func TestAccResource_basic_FrameworkMigration(t *testing.T) {
	dsn := "null_resource.test"

	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config:            testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "id"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				PlanOnly:                 true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsn, "id"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_EmptyMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_EmptyMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_NullMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_NullMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_NullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_NullValues(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_Value(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Keep_Values(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_EmptyMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_NullMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_NullValueToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_ValueToEmptyMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_ValueToNullMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config:                   testAccResourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_ValueToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_Replace_ValueToNewValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapToMultipleNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapToMultipleValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

func TestAccResource_Triggers_FrameworkMigration_NullMapValueToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion311(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id1),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "null_resource" "test" {
					triggers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("null_resource.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("null_resource.test", "triggers.%", "2"),
				),
			},
		},
	})
}

const testAccResourceConfig_basic = `
resource "null_resource" "test" {
}
`
