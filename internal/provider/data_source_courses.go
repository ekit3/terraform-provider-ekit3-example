package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCourses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCoursesRead,
		Schema: map[string]*schema.Schema{
			"courses": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"summary": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"__v": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCoursesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/cours", "http://localhost:4000"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	courses := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&courses)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("courses", courses); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
