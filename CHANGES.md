# Changelog

## [0.2.0](https://github.com/globalso-labs/telemetry/compare/v0.1.0...v0.2.0) (2024-08-01)


### ⚠ BREAKING CHANGES

* golang.org/x/tools updated from v0.21.1-0.20240508182429-e35e4ccd0d2d to v0.23.0, google.golang.org/grpc updated from v1.64.0 to v1.65.0.

### Miscellaneous Chores

* release 0.2.0 ([f6ee5d9](https://github.com/globalso-labs/telemetry/commit/f6ee5d9505b9cf7d15a2c4337ec703de0ec677af))
* update dependencies to latest versions ([4c810c7](https://github.com/globalso-labs/telemetry/commit/4c810c7ff63c8ba1f5bc47eb94087ad9bc441ba6))

## 0.1.0 (2024-08-01)


### ⚠ BREAKING CHANGES

* Change in assertion method may affect test outcomes if there are pointer comparisons.
* The metrics package has been renamed to meter. The configuration structure has been updated to reflect the new meter options and logger integration.

### Features

* add configuration and basic test to allow config initialization ([84075a8](https://github.com/globalso-labs/telemetry/commit/84075a88379bcc1dd80410b77c47c8bfb220e9db))
* add GitHub Actions workflow for automated releases ([cc2fce3](https://github.com/globalso-labs/telemetry/commit/cc2fce33adcde5ac6b44613e636969ac6f0ee820))
* add logging and metrics with configuration support ([02fb9b2](https://github.com/globalso-labs/telemetry/commit/02fb9b28eac9323a95d5eb5b7574fa88c879e986))
* add metrics implementation ([ace9e82](https://github.com/globalso-labs/telemetry/commit/ace9e82426eceeeb2f4f48235ffed2f8edd6ad04))


### Bug Fixes

* update equality assertions in logger tests ([5365a5f](https://github.com/globalso-labs/telemetry/commit/5365a5f3adb191710807409e258808b208f4e321))
