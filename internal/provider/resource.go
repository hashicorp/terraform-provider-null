package provider

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func nullResource() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`null_resource`" + ` resource implements the standard resource lifecycle but takes no further action.

The ` + "`triggers`" + ` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.`,

		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
			},

			"triggers_update": {
				Description: "A map of arbitrary strings that, when changed, will update the null resource, re-running any associated provisioners. On update of a resource, only the create action is executed.",
				Type:        schema.TypeMap,
				Optional:    true,
			},

			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceUpdate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	d.MarkNewResource()
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
