package favote

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"favote_vote": resourceVote(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"favote_vote": dataSourceVote(),
		},
	}
}
