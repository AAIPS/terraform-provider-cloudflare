package main

import (
	"context"
	"flag"
	"log"

	framework "github.com/cloudflare/terraform-provider-cloudflare/internal/framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
	commit  string = ""
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers")
	flag.Parse()

	err := providerserver.Serve(
		context.Background(),
		framework.New(version),
		providerserver.ServeOpts{
			Address: "registry.terraform.io/cloudflare/cloudflare",
			Debug:   debug,
		})

	if err != nil {
		log.Fatal(err)
	}
}
