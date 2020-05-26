package null

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func resource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Delete: resourceDelete,

		Schema: map[string]*schema.Schema{
			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
			},
			"delay" : {
				Type:     schema.TypeInt,
				Optional: true,
				Default: 0,
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	time.Sleep(d.Get("delay").(time.Duration) * time.Second)
	return nil
}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
