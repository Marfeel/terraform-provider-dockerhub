package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dh "github.com/BarnabyShearer/dockerhub/v2"
)

func resourceToken() *schema.Resource {
	return &schema.Resource{
		Description:   "A hub.docker.io personal access token (for uploading images).",
		CreateContext: resourceTokenCreate,
		DeleteContext: resourceTokenDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "The UUID of the token.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Token label.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Token to use as password",
			},
			"scopes": {
				Type:     schema.TypeList,
				Required: true,
			},
		},
	}
}

func resourceTokenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	token, err := client.CreatePersonalAccessToken(ctx, dh.CreatePersonalAccessToken{
		TokenLabel: d.Get("label").(string),
		Scopes:     d.Get("scopes").([]string),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(token.UUID)
	d.Set("token", token.Token)
	return nil
}

func resourceTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*dh.Client)
	err := client.DeletePersonalAccessToken(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}