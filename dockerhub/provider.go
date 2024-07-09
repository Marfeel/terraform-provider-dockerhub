package dockerhub

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	rtd "github.com/Marfeel/dockerhub/v2"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_USERNAME", nil),
				Description: "Username for authentication.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DOCKER_PASSWORD", nil),
				Description: "Password for authentication.",
			},
			"sleep_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLEEP_INTERVAL", "750"),
				Description: "Amount of milliseconds to sleep between requests. Defaults to 750",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dockerhub_repository":      resourceRepository(),
			"dockerhub_group":           resourceGroup(),
			"dockerhub_repositorygroup": resourceRepositoryGroup(),
			"dockerhub_token":           resourceToken(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return rtd.NewClient(d.Get("username").(string), d.Get("password").(string), d.Get("sleep_interval").(string)), nil
}
