package courses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_TOKEN", nil),
			},
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ekite_cour": dataSourceCour(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ekite_courses": dataSourceCourses(),
		},
		ConfigureContextFunc: configure,
	}
}

type Config struct {
	api_token string
	api_url   string
}

func configure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics
	config := Config{}

	if v, ok := d.GetOk("api_token"); ok {
		config.api_token = v.(string)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "api_token is required",
		})

		return nil, diags
	}

	if v, ok := d.GetOk("api_url"); ok {
		config.api_url = v.(string)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "api_url is required",
		})

		return nil, diags
	}

	return &config, diags
}
