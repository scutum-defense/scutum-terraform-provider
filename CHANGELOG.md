# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2024-01-01

### Added

- Initial release
- Provider configuration with endpoint, API key, and region
- Resource: `scutum_zone` -- manage protected zones with polygon boundaries
- Resource: `scutum_corridor` -- manage maritime/aerial/ground corridors
- Resource: `scutum_rule` -- manage detection rules
- Resource: `scutum_policy` -- manage operational policies
- Data source: `scutum_incident` -- read active incidents
- Data source: `scutum_asset` -- read asset information
- Example Terraform configuration
- CI workflow with Go build and vet
