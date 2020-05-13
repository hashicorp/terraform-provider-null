package provider

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func dataSource() *schema.Resource {
	return &schema.Resource{
		Description: "The `null_data_source` data source implements the standard data source lifecycle but does not interact with any external APIs.",

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
				Description: "A random value. This is primarily for testing and has little practical use; prefer the [random provider](https://www.terraform.io/docs/providers/random/) for more practical random number use-cases.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"has_computed_default": {
				Description: "If set, its literal value will be stored and returned. If not, its value defaults to `\"default\"`. This argument exists primarily for testing and has little practical use.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
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
