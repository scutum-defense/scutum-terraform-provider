package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var zoneClassifications = []string{
	"restricted", "controlled", "monitored", "exclusion", "buffer", "corridor",
}

// ResourceZone returns the schema.Resource for the scutum_zone resource.
func ResourceZone() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a protected zone in the Scutum platform",
		CreateContext: resourceZoneCreate,
		ReadContext:   resourceZoneRead,
		UpdateContext: resourceZoneUpdate,
		DeleteContext: resourceZoneDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the protected zone",
			},
			"classification": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Zone classification: restricted, controlled, monitored, exclusion, buffer, corridor",
				ValidateFunc: validation.StringInSlice(zoneClassifications, false),
			},
			"boundary": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Polygon boundary coordinates",
				MinItems:    3,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lat": {
							Type:         schema.TypeFloat,
							Required:     true,
							Description:  "Latitude (-90 to 90)",
							ValidateFunc: validateLatitude,
						},
						"lng": {
							Type:         schema.TypeFloat,
							Required:     true,
							Description:  "Longitude (-180 to 180)",
							ValidateFunc: validateLongitude,
						},
					},
				},
			},
		},
	}
}

func validateLatitude(i interface{}, _ string) ([]string, []error) {
	lat, ok := i.(float64)
	if !ok {
		return nil, []error{fmt.Errorf("latitude must be a number")}
	}
	if lat < -90 || lat > 90 {
		return nil, []error{fmt.Errorf("latitude %v out of range [-90, 90]", lat)}
	}
	return nil, nil
}

func validateLongitude(i interface{}, _ string) ([]string, []error) {
	lng, ok := i.(float64)
	if !ok {
		return nil, []error{fmt.Errorf("longitude must be a number")}
	}
	if lng < -180 || lng > 180 {
		return nil, []error{fmt.Errorf("longitude %v out of range [-180, 180]", lng)}
	}
	return nil, nil
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
