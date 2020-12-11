package provider

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSensitive() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`null_resource_sensitive`" + ` resource implements the standard resource lifecycle but takes no further action.

The ` + "`triggers`" + ` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.`,

		Create: resourceSensitiveCreate,
		Read:   resourceSensitiveRead,
		Delete: resourceSensitiveDelete,

		Schema: map[string]*schema.Schema{
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners. Its value is marked as Sensitive and will be masked in the UI.",
				Type:        schema.TypeMap,
				Optional:    true,
				Sensitive:   true,
				ForceNew:    true,
			},

			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceSensitiveCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func resourceSensitiveRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSensitiveDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
