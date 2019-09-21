package null

import (
	"os"
	"testing"

	tftest "github.com/apparentlymart/terraform-plugin-test"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testHelper *tftest.Helper

// initProviderTesting is a helper that would ideally be provided by the SDK
// itself if the SDK were to adopt tftest as its testing mechanism, but it
// lives here for now for ease of prototyping.
func initProviderTesting(name string, providerFunc plugin.ProviderFunc) *tftest.Helper {
	if tftest.RunningAsPlugin() {
		// The test program is being re-launched as a provider plugin via our
		// stub program.
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: providerFunc,
		})
		os.Exit(0)
	}

	return tftest.AutoInitProviderHelper(name)
}

func TestMain(m *testing.M) {
	testHelper = initProviderTesting("null", Provider)
	status := m.Run()
	testHelper.Close()
	os.Exit(status)
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}
