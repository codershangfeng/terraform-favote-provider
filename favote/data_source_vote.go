package favote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVote() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVoteRead,
		Schema: map[string]*schema.Schema{
			"vid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "UUID for vote item",
			},
			"topic": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Vote topic",
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Vote options against its topic",
			},
		},
	}
}

func dataSourceVoteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/votes", "http://localhost:8080"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	votes := make([]VoteDataSource, 0)
	err = json.NewDecoder(r.Body).Decode(&votes)
	if err != nil {
		return diag.FromErr(err)
	}

	topic := d.Get("topic").(string)

	for _, v := range votes {
		if v.Topic == topic {
			if err := d.Set("topic", v.Topic); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("options", v.Options); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("vid", v.ID); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(fmt.Sprintf("%s/vote/%d", "http://localhost:8080", v.ID))
			break
		}
	}

	return diags
}

type VoteDataSource struct {
	ID      int      `json:"id"`
	Topic   string   `json:"topic"`
	Options []string `json:"options"`
}
