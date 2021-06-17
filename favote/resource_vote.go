package favote

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type VoteResource struct {
	VID     *int     `json:"vid,omitempty"`
	Topic   string   `json:"topic"`
	Options []string `json:"options"`
}

func resourceVote() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVoteCreate,
		ReadContext:   resourceVoteRead,
		UpdateContext: resourceVoteUpdate,
		DeleteContext: resourceVoteDelete,
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
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Vote options against its topic",
			},
		},
	}
}

func resourceVoteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics
	ops := d.Get("options").([]interface{})
	options := make([]string, len(ops))
	for i, v := range ops {
		options[i] = v.(string)
	}
	vote := VoteResource{
		Topic:   d.Get("topic").(string),
		Options: options,
	}
	voteJSON, _ := json.Marshal(vote)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/votes", "http://localhost:8080"), bytes.NewBufferString(string(voteJSON)))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusCreated {
		return diag.FromErr(errors.New("resource vote create failed"))
	}

	voteResponse := new(VoteResource)
	err = json.NewDecoder(r.Body).Decode(voteResponse)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("vid", voteResponse.VID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/vote/%d", "http://localhost:8080", vote.VID))

	return diags
}

func resourceVoteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/vote/%s", "http://localhost:8080", d.Id()), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusNotFound {
		return diags
	}

	var vote VoteResource
	err = json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("vid", vote.VID); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceVoteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceVoteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
