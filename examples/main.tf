terraform {
  required_providers {
    scutum = {
      source = "scutum-defense/scutum"
    }
  }
}

provider "scutum" {
  endpoint = "https://api.scutum.local"
  api_key  = var.scutum_api_key
  region   = "abudhabi"
}

variable "scutum_api_key" {
  type      = string
  sensitive = true
}

resource "scutum_zone" "fuel_storage" {
  name           = "Fuel Storage Exclusion Zone"
  classification = "exclusion"

  boundary {
    lat = 24.450
    lng = 54.370
  }
  boundary {
    lat = 24.450
    lng = 54.390
  }
  boundary {
    lat = 24.470
    lng = 54.390
  }
  boundary {
    lat = 24.470
    lng = 54.370
  }
}

resource "scutum_rule" "drone_approach" {
  name     = "Drone approach toward protected perimeter"
  severity = "high"
  category = "perimeter"
  enabled  = true
}

resource "scutum_policy" "no_autonomous_ot" {
  name     = "No autonomous OT write"
  category = "safety"
  verdict  = "deny"
  enabled  = true
}

data "scutum_incident" "current" {}

output "active_incident" {
  value = data.scutum_incident.current.title
}
