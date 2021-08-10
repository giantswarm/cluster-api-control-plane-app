# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Use constant name across releases for app manifests.
- Bumped `cluster-api` to version `v0.4.0`.
- Removed dedicated container for the webhook.

## [0.3.14-gs3] - 2021-06-01

### Removed

- Remove kube-rbac-proxy for the metrics endpoint.

### Fixed

- Fixed label selector for webhook and manager services.

## [0.3.14-gs2] - 2021-05-27

## [0.3.14-gs1] - 2021-05-12

## [0.0.1] - 2021-03-18

[Unreleased]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs3...HEAD
[0.3.14-gs3]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs2...v0.3.14-gs3
[0.3.14-gs2]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs1...v0.3.14-gs2
[0.3.14-gs1]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.0.1...v0.3.14-gs1
[0.0.1]: https://github.com/giantswarm/cluster-api-control-plane-app/releases/tag/v0.0.1
