package provider

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "The null_data_source was historically used to construct intermediate values to re-use elsewhere " +
			"in configuration, the same can now be achieved using locals",

		Description: `The ` + "`null_data_source`" + ` data source implements the standard data source lifecycle but does not
interact with any external APIs.

Historically, the ` + "`null_data_source`" + ` was typically used to construct intermediate values to re-use elsewhere in configuration. The
same can now be achieved using [locals](https://www.terraform.io/docs/language/values/locals.html).
`,

		Read: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"inputs": {
				Description: "A map of arbitrary strings that is copied into the `outputs` attribute, and accessible directly for interpolation.",
				Type:        schema.TypeMap,
				Optional:    true,
			},
			"outputs": {
				Description: "After the data source is \"read\", a copy of the `inputs` map.",
				Type:        schema.TypeMap,
				Computed:    true,
			},
			"random": {
				Description: "A random value. This is primarily for testing and has little practical use; prefer the [hashicorp/random provider](https://registry.terraform.io/providers/hashicorp/random) for more practical random number use-cases.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"has_computed_default": {
				Description: "If set, its literal value will be stored and returned. If not, its value defaults to `\"default\"`. This argument exists primarily for testing and has little practical use.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},

			"id": {
				Description: "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
				Computed:    true,
				Deprecated:  "This attribute is only present for some legacy compatibility issues and should not be used. It will be removed in a future version.",
				Type:        schema.TypeString,
			},
		},
	}
}

func dataSourceRead(d *schema.ResourceData, meta interface{}) error {

	inputs := d.Get("inputs")
	d.Set("outputs", inputs)

	d.Set("random", fmt.Sprintf("%d", rand.Int()))
	if d.Get("has_computed_default") == "" {
		d.Set("has_computed_default", "default")
	}

	d.SetId("static")

	return nil
}
