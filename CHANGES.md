# Changelog

## [0.6.0](https://github.com/globalso-labs/telemetry/compare/v0.5.0...v0.6.0) (2024-08-07)


### Features

* add version tracking and telemetry middleware ([c638f39](https://github.com/globalso-labs/telemetry/commit/c638f39f63a62fc61555dc89f22929e58aacbc5f))

## [0.5.0](https://github.com/globalso-labs/telemetry/compare/v0.4.0...v0.5.0) (2024-08-06)


### Features

* add With function to create zerolog.Context ([65c81e7](https://github.com/globalso-labs/telemetry/commit/65c81e786606b40d6c57f8ee2cfcfeb29274097a))

## [0.4.0](https://github.com/globalso-labs/telemetry/compare/v0.3.0...v0.4.0) (2024-08-06)


### Features

* remove WithFields function from zerolog logger ([3f8f37c](https://github.com/globalso-labs/telemetry/commit/3f8f37cd608997cac6c801c73df93736f6977453))

## [0.3.0](https://github.com/globalso-labs/telemetry/compare/v0.2.0...v0.3.0) (2024-08-04)


### Features

* add tracer implementation ([c377b5a](https://github.com/globalso-labs/telemetry/commit/c377b5aa9db9b4a93ff37b0beb38684eb1e18dc8))
* update logger to use zerolog levels ([9fc556c](https://github.com/globalso-labs/telemetry/commit/9fc556ced0bf5a8b37d5cc06fd9c4f323538449d))


### Miscellaneous Chores

* release 0.3.0 ([934ddac](https://github.com/globalso-labs/telemetry/commit/934ddaccc3283570336f3a0c2c846eb33146658b))

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
