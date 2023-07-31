terraform {
  required_providers {
    cloudflare = {
      source  = "localhost/cloudflare/cloudflare"
    }
  }
}

provider "cloudflare" {
#  api_token = "<API_TOKEN>"
  api_token = "GUPdRtsTRcIejTnmD3m9xw5wimnN0SGK0U1sZzyY"
}

resource "cloudflare_ruleset" "example" {}