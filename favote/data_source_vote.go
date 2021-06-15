package favote

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

func dataSourceVote() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVoteRead,
		Schema: map[string]*schema.Schema{
			"vid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"topic": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceVoteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vote/%s", "http://localhost:8080", d.Get("vid")), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusNotFound {
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		return diags
	}

	v := new(VoteDataSource)
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("vid", strconv.Itoa(v.Id)); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("topic", v.Topic); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("options", v.Options); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

type VoteDataSource struct {
	Id      int      `json:"id"`
	Topic   string   `json:"topic"`
	Options []string `json:"options"`
}
