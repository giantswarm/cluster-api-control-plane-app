# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

- Inject ca into dynamic webhook.

## [0.3.22-gs10-crd] - 2021-11-24

## [0.3.22-gs10] - 2021-11-24

### Fixed

- Prepare to run multiple apps (CRD Webhook, Dynamic Webhook and Controller).

## [0.3.22-gs9] - 2021-10-22

### Fixed

- Use 9443 as service port.

## [0.3.22-gs8] - 2021-10-22

### Changed

- Align CRD webhook name with `cluster-api-app`.

## [0.3.22-gs7] - 2021-10-20

### Fixed

- Change webhooks to use `v1beta1` of `admissionReviewVersions`

## [0.3.22-gs6] - 2021-10-15

## [0.3.22-gs5] - 2021-10-15

### Changed

- Adjust settings for CRD webhook.

## [0.3.22-gs4] - 2021-10-15

### Changed

- Prepare helm values for configuration management.
- Prepare webhook CRD conversion.

## [0.3.22-gs3] - 2021-08-31

## [0.3.22-gs2] - 2021-08-02

## [0.3.22-gs1] - 2021-08-02

### Changed

- Updated cluster-api to v0.3.22.

## [0.3.19-gs1] - 2021-06-21

### Changed

- Updated cluster-api to v0.3.19.

## [0.3.14-gs3] - 2021-06-01

### Removed

- Remove kube-rbac-proxy for the metrics endpoint.

### Fixed

- Fixed label selector for webhook and manager services.

## [0.3.14-gs2] - 2021-05-27

## [0.3.14-gs1] - 2021-05-12

## [0.0.1] - 2021-03-18

[Unreleased]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs10...HEAD
[0.3.22-gs10]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs9...v0.3.22-gs10
[0.3.22-gs9]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs8...v0.3.22-gs9
[0.3.22-gs8]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs7...v0.3.22-gs8
[0.3.22-gs7]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs6...v0.3.22-gs7
[0.3.22-gs6]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs5...v0.3.22-gs6
[0.3.22-gs5]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs4...v0.3.22-gs5
[0.3.22-gs4]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs3...v0.3.22-gs4
[0.3.22-gs3]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs2...v0.3.22-gs3
[0.3.22-gs2]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.22-gs1...v0.3.22-gs2
[0.3.22-gs1]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.19-gs1...v0.3.22-gs1
[0.3.19-gs1]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs3...v0.3.19-gs1
[0.3.14-gs3]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs2...v0.3.14-gs3
[0.3.14-gs2]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.3.14-gs1...v0.3.14-gs2
[0.3.14-gs1]: https://github.com/giantswarm/cluster-api-control-plane-app/compare/v0.0.1...v0.3.14-gs1
[0.0.1]: https://github.com/giantswarm/cluster-api-control-plane-app/releases/tag/v0.0.1
