package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func TestProvider_impl(t *testing.T) {
	var _ provider.Provider = New("dev")()
}
