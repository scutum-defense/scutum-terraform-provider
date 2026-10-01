package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var corridorTypes = []string{"maritime", "aerial", "ground", "pipeline"}
var ruleSeverities = []string{"info", "low", "medium", "high", "critical"}
var policyVerdicts = []string{"allow", "deny", "warn", "require_approval"}

// ResourceCorridor returns the schema.Resource for the scutum_corridor resource.
func ResourceCorridor() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a maritime or aerial corridor in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name":  {Type: schema.TypeString, Required: true},
			"type":  {Type: schema.TypeString, Required: true, Description: "maritime, aerial, ground, pipeline", ValidateFunc: validation.StringInSlice(corridorTypes, false)},
			"width": {Type: schema.TypeFloat, Required: true, Description: "Corridor width in meters", ValidateFunc: validation.FloatAtLeast(1)},
		},
	}
}

// ResourceDetectionRule returns the schema.Resource for the scutum_rule resource.
func ResourceDetectionRule() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a detection rule in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"severity": {Type: schema.TypeString, Required: true, ValidateFunc: validation.StringInSlice(ruleSeverities, false)},
			"category": {Type: schema.TypeString, Required: true},
			"enabled":  {Type: schema.TypeBool, Optional: true, Default: true},
		},
	}
}

// ResourcePolicy returns the schema.Resource for the scutum_policy resource.
func ResourcePolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a policy in the Scutum platform",
		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"category": {Type: schema.TypeString, Required: true},
			"verdict":  {Type: schema.TypeString, Required: true, ValidateFunc: validation.StringInSlice(policyVerdicts, false)},
			"enabled":  {Type: schema.TypeBool, Optional: true, Default: true},
		},
	}
}

// DataSourceIncident returns the schema.Resource for the scutum_incident data source.
func DataSourceIncident() *schema.Resource {
	return &schema.Resource{
		Description: "Reads the current active incident from the Scutum platform",
		Schema: map[string]*schema.Schema{
			"id":       {Type: schema.TypeString, Computed: true},
			"title":    {Type: schema.TypeString, Computed: true},
			"severity": {Type: schema.TypeString, Computed: true},
			"status":   {Type: schema.TypeString, Computed: true},
		},
	}
}

// DataSourceAsset returns the schema.Resource for the scutum_asset data source.
func DataSourceAsset() *schema.Resource {
	return &schema.Resource{
		Description: "Reads asset information from the Scutum platform",
		Schema: map[string]*schema.Schema{
			"id":     {Type: schema.TypeString, Required: true},
			"name":   {Type: schema.TypeString, Computed: true},
			"type":   {Type: schema.TypeString, Computed: true},
			"status": {Type: schema.TypeString, Computed: true},
		},
	}
}
