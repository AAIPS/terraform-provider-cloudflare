package acctest

import (
	"os"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// Integration test account ID.
	testAccCloudflareAccountID string = "f037e56e89293a057740de681ac9abbe"

	// Integration test account zone ID.
	testAccCloudflareZoneID string = "0da42c8d2132a9ddaf714f9e7c920711"
	// Integration test account zone name.
	testAccCloudflareZoneName string = "terraform.cfapi.net"

	// Integration test account alternate zone ID.
	testAccCloudflareAltZoneID string = "b72110c08e3382597095c29ba7e661ea"
	// Integration test account alternate zone name.
	testAccCloudflareAltZoneName string = "terraform2.cfapi.net"
)

var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cloudflare": providerserver.NewProtocol6WithError(provider.New("dev")()),
}

func TestAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

	testAccPreCheckEmail(t)
	testAccPreCheckApiKey(t)
	testAccPreCheckDomain(t)
	testAccPreCheckZoneID(t)
}

func TestAccSkipForDefaultAccount(t *testing.T, reason string) {
	if os.Getenv("CLOUDFLARE_ACCOUNT_ID") == testAccCloudflareAccountID {
		t.Skipf("Skipping acceptance test for default account (%s). %s", testAccCloudflareAccountID, reason)
	}
}

func testAccPreCheckEmail(t *testing.T) {
	if v := os.Getenv(consts.EmailEnvVarKey); v == "" {
		t.Fatalf("%s must be set for acceptance tests", consts.EmailEnvVarKey)
	}
}

func testAccPreCheckApiKey(t *testing.T) {
	if v := os.Getenv(consts.APIKeyEnvVarKey); v == "" {
		t.Fatalf("%s must be set for acceptance tests", consts.APIKeyEnvVarKey)
	}
}

func testAccPreCheckDomain(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_DOMAIN"); v == "" {
		t.Fatal("CLOUDFLARE_DOMAIN must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}

func testAccPreCheckZoneID(t *testing.T) {
	if v := os.Getenv("CLOUDFLARE_ZONE_ID"); v == "" {
		t.Fatal("CLOUDFLARE_ZONE_ID must be set for acceptance tests")
	}
}
