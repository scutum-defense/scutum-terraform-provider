```
 ____            _                    _____                     __
/ ___|  ___ _   _| |_ _   _ _ __ ___  |_   _|__ _ __ _ __ __ _ / _| ___  _ __ _ __ ___
\___ \ / __| | | | __| | | | '_ ` _ \   | |/ _ \ '__| '__/ _` | |_ / _ \| '__| '_ ` _ \
 ___) | (__| |_| | |_| |_| | | | | | |  | |  __/ |  | | | (_| |  _| (_) | |  | | | | | |
|____/ \___|\__,_|\__|\__,_|_| |_| |_|  |_|\___|_|  |_|  \__,_|_|  \___/|_|  |_| |_| |_|

 ____                 _     _
|  _ \ _ __ _____   _(_) __| | ___ _ __
| |_) | '__/ _ \ \ / / |/ _` |/ _ \ '__|
|  __/| | | (_) \ V /| | (_| |  __/ |
|_|   |_|  \___/ \_/ |_|\__,_|\___|_|
```

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://go.dev/)
[![Terraform](https://img.shields.io/badge/Terraform-1.0+-7B42BC.svg)](https://www.terraform.io/)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Resources](https://img.shields.io/badge/Resources-4-green.svg)]()
[![Data Sources](https://img.shields.io/badge/Data%20Sources-2-green.svg)]()

**Terraform provider for managing Scutum platform resources as infrastructure-as-code.**

Deploy and manage sovereign defense infrastructure declaratively. Define protected zones, detection rules, corridors, and operational policies in Terraform -- version-controlled, auditable, and repeatable across environments.

---

## Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Resources and Data Sources](#resources-and-data-sources)
- [Quick Start](#quick-start)
- [Provider Configuration](#provider-configuration)
- [Usage Examples](#usage-examples)
- [Development](#development)
- [License](#license)

---

## Overview

The Scutum Terraform Provider brings infrastructure-as-code practices to sovereign defense operations. Instead of manually configuring zones, rules, and policies through a UI, teams can define their entire operational configuration in Terraform -- enabling version control, peer review, and automated deployment pipelines.

**Key capabilities:**

- **Sovereign deployment** -- region-aware configuration for Gulf-first deployments
- **Zone management** -- define protected zones with polygon boundaries and classifications
- **Corridor management** -- manage maritime, aerial, and ground corridors
- **Detection rules** -- deploy and manage detection rules with severity levels
- **Policy enforcement** -- define operational policies with verdicts (allow/deny)
- **Incident visibility** -- read active incidents as data sources
- **Asset discovery** -- query asset information from the platform

---

## Architecture

```
+------------------+       +---------------------+       +------------------+
|                  |       |                     |       |                  |
|  Terraform CLI   +------>+  Scutum Provider    +------>+  Scutum API      |
|                  |       |                     |       |                  |
|  - plan          |       |  Resources:         |       |  - /zones        |
|  - apply         |       |  - scutum_zone      |       |  - /corridors    |
|  - destroy       |       |  - scutum_corridor  |       |  - /rules        |
|                  |       |  - scutum_rule      |       |  - /policies     |
|  State:          |       |  - scutum_policy    |       |  - /incidents    |
|  - terraform.    |       |                     |       |  - /assets       |
|    tfstate       |       |  Data Sources:      |       |                  |
|                  |       |  - scutum_incident  |       |  Auth:           |
|                  |       |  - scutum_asset     |       |  - API Key       |
|                  |       |                     |       |  - Region        |
+------------------+       +---------------------+       +------------------+
```

---

## Resources and Data Sources

### Resources

| Resource           | Description                                       |
| ------------------ | ------------------------------------------------- |
| `scutum_zone`      | Protected zone with polygon boundary and classification (restricted, controlled, monitored, exclusion, buffer, corridor) |
| `scutum_corridor`  | Maritime, aerial, ground, or pipeline corridor with configurable width |
| `scutum_rule`      | Detection rule with severity level and category   |
| `scutum_policy`    | Operational policy with verdict (allow/deny)      |

### Data Sources

| Data Source        | Description                                       |
| ------------------ | ------------------------------------------------- |
| `scutum_incident`  | Read the current active incident                  |
| `scutum_asset`     | Read asset information by ID                      |

---

## Quick Start

### 1. Configure the provider

```hcl
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
```

### 2. Define resources

```hcl
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
```

### 3. Apply

```bash
terraform init
terraform plan
terraform apply
```

---

## Provider Configuration

| Attribute   | Type     | Required | Environment Variable | Description                       |
| ----------- | -------- | -------- | -------------------- | --------------------------------- |
| `endpoint`  | `string` | Yes      | `SCUTUM_ENDPOINT`    | Scutum platform API endpoint      |
| `api_key`   | `string` | Yes      | `SCUTUM_API_KEY`     | API key for authentication        |
| `region`    | `string` | No       | `SCUTUM_REGION`      | Sovereign deployment region       |

All attributes support environment variable fallback for CI/CD pipelines:

```bash
export SCUTUM_ENDPOINT="https://api.scutum.local"
export SCUTUM_API_KEY="sk-..."
export SCUTUM_REGION="abudhabi"
```

---

## Usage Examples

### Zone with exclusion boundary

```hcl
resource "scutum_zone" "refinery_perimeter" {
  name           = "Refinery Perimeter"
  classification = "restricted"

  boundary {
    lat = 24.500
    lng = 54.400
  }
  boundary {
    lat = 24.500
    lng = 54.420
  }
  boundary {
    lat = 24.520
    lng = 54.420
  }
  boundary {
    lat = 24.520
    lng = 54.400
  }
}
```

### Detection rule

```hcl
resource "scutum_rule" "drone_approach" {
  name     = "Drone approach toward protected perimeter"
  severity = "high"
  category = "perimeter"
  enabled  = true
}
```

### Operational policy

```hcl
resource "scutum_policy" "no_autonomous_ot" {
  name     = "No autonomous OT write"
  category = "safety"
  verdict  = "deny"
  enabled  = true
}
```

### Maritime corridor

```hcl
resource "scutum_corridor" "shipping_lane" {
  name  = "Primary Shipping Lane"
  type  = "maritime"
  width = 500.0
}
```

### Reading active incident

```hcl
data "scutum_incident" "current" {}

output "active_incident_title" {
  value = data.scutum_incident.current.title
}
```

### Reading asset info

```hcl
data "scutum_asset" "radar_01" {
  id = "asset-radar-01"
}

output "radar_status" {
  value = data.scutum_asset.radar_01.status
}
```

---

## Development

### Prerequisites

- Go 1.22+
- Terraform 1.0+

### Build

```bash
go build ./...
```

### Lint

```bash
go vet ./...
```

### Project structure

```
scutum-terraform-provider/
  main.go                        # Provider entry point
  go.mod                         # Go module definition
  internal/
    provider/
      provider.go                # Provider schema, resources, and data sources
    resources/
      zone.go                    # Zone resource (full CRUD)
      stubs.go                   # Corridor, rule, policy, incident, asset stubs
  examples/
    main.tf                      # Example Terraform configuration
  docs/                          # Documentation
```

### Adding a new resource

1. Define the resource function in `internal/provider/provider.go`
2. Add it to the `ResourcesMap` in the provider
3. Implement CRUD context functions
4. Add example configuration to `examples/`
5. Update documentation

---

## License

Apache License 2.0 -- see [LICENSE](LICENSE) for details.

Copyright 2024 Scutum Defense.
