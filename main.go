package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	c "omni/terraform-provider-kafka-connect/connect"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: c.Provider})
}
