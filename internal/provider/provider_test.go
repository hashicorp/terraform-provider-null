package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func protoV5ProviderFactories() map[string]func() (tfprotov5.ProviderServer, error) {
	return map[string]func() (tfprotov5.ProviderServer, error){
		"null": providerserver.NewProtocol5WithError(New()),
	}
}

func providerVersion311() map[string]resource.ExternalProvider {
	return map[string]resource.ExternalProvider{
		"null": {
			VersionConstraint: "3.1.1",
			Source:            "hashicorp/null",
		},
	}
}
