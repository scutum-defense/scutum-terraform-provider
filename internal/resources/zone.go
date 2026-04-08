package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceZone returns the schema.Resource for the scutum_zone resource.
func ResourceZone() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages a protected zone in the Scutum platform",
		CreateContext: resourceZoneCreate,
		ReadContext:   resourceZoneRead,
		UpdateContext: resourceZoneUpdate,
		DeleteContext: resourceZoneDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the protected zone",
			},
			"classification": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone classification: restricted, controlled, monitored, exclusion, buffer, corridor",
			},
			"boundary": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Polygon boundary coordinates",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lat": {Type: schema.TypeFloat, Required: true},
						"lng": {Type: schema.TypeFloat, Required: true},
					},
				},
			},
		},
	}
}

func resourceZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("zone-" + d.Get("name").(string))
	return nil
}

func resourceZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
