package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ScutumClient holds the authenticated client configuration for the Scutum API.
type ScutumClient struct {
	Endpoint string
	APIKey   string
	Region   string
}

// New returns a configured Scutum Terraform provider.
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Scutum platform API endpoint",
				DefaultFunc: schema.EnvDefaultFunc("SCUTUM_ENDPOINT", nil),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "API key for authentication",
				DefaultFunc: schema.EnvDefaultFunc("SCUTUM_API_KEY", nil),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sovereign deployment region",
				DefaultFunc: schema.EnvDefaultFunc("SCUTUM_REGION", "default"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scutum_zone":     resourceZone(),
			"scutum_corridor": resourceCorridor(),
			"scutum_rule":     resourceDetectionRule(),
			"scutum_policy":   resourcePolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"scutum_incident": dataSourceIncident(),
			"scutum_asset":    dataSourceAsset(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	endpoint := d.Get("endpoint").(string)
	apiKey := d.Get("api_key").(string)
	region := d.Get("region").(string)

	client := &ScutumClient{
		Endpoint: endpoint,
		APIKey:   apiKey,
		Region:   region,
	}

	return client, nil
}

// ---------------------------------------------------------------------------
// Resource: scutum_zone
// ---------------------------------------------------------------------------

func resourceZone() *schema.Resource {
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

// ---------------------------------------------------------------------------
// Resource: scutum_corridor
// ---------------------------------------------------------------------------

func resourceCorridor() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a maritime or aerial corridor in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name": {Type: schema.TypeString, Required: true},
			"type": {Type: schema.TypeString, Required: true, Description: "maritime, aerial, ground, pipeline"},
			"width": {Type: schema.TypeFloat, Required: true, Description: "Corridor width in meters"},
		},
	}
}

// ---------------------------------------------------------------------------
// Resource: scutum_rule
// ---------------------------------------------------------------------------

func resourceDetectionRule() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a detection rule in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"severity": {Type: schema.TypeString, Required: true},
			"category": {Type: schema.TypeString, Required: true},
			"enabled":  {Type: schema.TypeBool, Optional: true, Default: true},
		},
	}
}

// ---------------------------------------------------------------------------
// Resource: scutum_policy
// ---------------------------------------------------------------------------

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a policy in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"category": {Type: schema.TypeString, Required: true},
			"verdict":  {Type: schema.TypeString, Required: true},
			"enabled":  {Type: schema.TypeBool, Optional: true, Default: true},
		},
	}
}

// ---------------------------------------------------------------------------
// Data Source: scutum_incident
// ---------------------------------------------------------------------------

func dataSourceIncident() *schema.Resource {
	return &schema.Resource{
		Description: "Reads the current active incident from the Scutum platform",
		ReadContext: dataSourceIncidentRead,
		Schema: map[string]*schema.Schema{
			"id":       {Type: schema.TypeString, Computed: true},
			"title":    {Type: schema.TypeString, Computed: true},
			"severity": {Type: schema.TypeString, Computed: true},
			"status":   {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceIncidentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("current")
	return nil
}

// ---------------------------------------------------------------------------
// Data Source: scutum_asset
// ---------------------------------------------------------------------------

func dataSourceAsset() *schema.Resource {
	return &schema.Resource{
		Description: "Reads asset information from the Scutum platform",
		ReadContext: dataSourceAssetRead,
		Schema: map[string]*schema.Schema{
			"id":     {Type: schema.TypeString, Required: true},
			"name":   {Type: schema.TypeString, Computed: true},
			"type":   {Type: schema.TypeString, Computed: true},
			"status": {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceAssetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId(d.Get("id").(string))
	return nil
}
