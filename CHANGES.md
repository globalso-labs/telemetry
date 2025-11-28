# Changelog

## [1.9.0](https://github.com/globalso-labs/telemetry/compare/v2.0.0...v1.9.0) (2025-11-28)


### ⚠ BREAKING CHANGES

* Removes Sentry logging hook and related dependency.
* OpenTelemetry collector agent is no longer included or supported in this repository. All related binaries, configuration, and integration points are removed. Downstream consumers must migrate to external collector solutions.
* Dependency updates may introduce breaking changes from upstream libraries. Review release notes of updated dependencies for possible API or behavior changes.
* testing method renamed
* The removal of the internal package from the logger initialization may impact other tests relying on it.
* The telemetry configuration now requires the "x-scope-orgid" header.
* golang.org/x/tools updated from v0.21.1-0.20240508182429-e35e4ccd0d2d to v0.23.0, google.golang.org/grpc updated from v1.64.0 to v1.65.0.
* Change in assertion method may affect test outcomes if there are pointer comparisons.
* The metrics package has been renamed to meter. The configuration structure has been updated to reflect the new meter options and logger integration.

### Features

* add configuration and basic test to allow config initialization ([84075a8](https://github.com/globalso-labs/telemetry/commit/84075a88379bcc1dd80410b77c47c8bfb220e9db))
* add constants for service names and update resource initialization ([8a8e012](https://github.com/globalso-labs/telemetry/commit/8a8e0129acf201416799b7961ae7961c0a2e858e))
* add default headers for telemetry configuration ([ce8d08f](https://github.com/globalso-labs/telemetry/commit/ce8d08f2637f5a4d1f43ceb7b053d2eb73307fb7))
* add GitHub Actions workflow for automated releases ([cc2fce3](https://github.com/globalso-labs/telemetry/commit/cc2fce33adcde5ac6b44613e636969ac6f0ee820))
* add goreleaser configuration for multiple platforms ([770f254](https://github.com/globalso-labs/telemetry/commit/770f254e74c8e9cf69f4c19fdd2659a96beef084))
* add initial implementation of telemetry agent ([d143bf4](https://github.com/globalso-labs/telemetry/commit/d143bf48cb43dfb48861360da030b3e06e5bb360))
* add logging and metrics with configuration support ([02fb9b2](https://github.com/globalso-labs/telemetry/commit/02fb9b28eac9323a95d5eb5b7574fa88c879e986))
* add metrics implementation ([ace9e82](https://github.com/globalso-labs/telemetry/commit/ace9e82426eceeeb2f4f48235ffed2f8edd6ad04))
* add Prometheus Remote Write exporter and new receivers ([6783c4d](https://github.com/globalso-labs/telemetry/commit/6783c4dac5eba3e89fc50b05a2914568548cfe8e))
* add resource management to Telemetry ([a2789e9](https://github.com/globalso-labs/telemetry/commit/a2789e94fb302abb5c4a24da78e0bd092673bf0f))
* add tracer implementation ([c377b5a](https://github.com/globalso-labs/telemetry/commit/c377b5aa9db9b4a93ff37b0beb38684eb1e18dc8))
* add update command to manage dependencies ([ee11eaa](https://github.com/globalso-labs/telemetry/commit/ee11eaa6a03e25fb281768d8a63c6dcd71a73ae6))
* add version tracking and telemetry middleware ([c638f39](https://github.com/globalso-labs/telemetry/commit/c638f39f63a62fc61555dc89f22929e58aacbc5f))
* add With function to create zerolog.Context ([65c81e7](https://github.com/globalso-labs/telemetry/commit/65c81e786606b40d6c57f8ee2cfcfeb29274097a))
* add WithKey function and update testing context handling ([2df39ca](https://github.com/globalso-labs/telemetry/commit/2df39cae4dc7181c3eddb770514ab6a69ed25e86))
* add WithTestingCtx method ([bfa449c](https://github.com/globalso-labs/telemetry/commit/bfa449ce0cebcabcfff2308b7d232110a488b0d5))
* enhance Resource functionality ([b315539](https://github.com/globalso-labs/telemetry/commit/b315539c3ad3792bad496f53198d7d14599b21b3))
* enhance telemetry configuration ([3c89f58](https://github.com/globalso-labs/telemetry/commit/3c89f589fe50158102812b6c4798fc1bb948303f))
* export zerolog.Logger as Instance type alias ([9b5f611](https://github.com/globalso-labs/telemetry/commit/9b5f6116cae5e037b63965a7a1133fc51154a080))
* implement merge function for key/value map ([7d922d1](https://github.com/globalso-labs/telemetry/commit/7d922d1d77aa6da7cbbf791e20c24ce3f16fe63b))
* implement parse verbosity function ([9878372](https://github.com/globalso-labs/telemetry/commit/9878372639d43e85e8ed29dcd14531b32af86b94))
* release 1.8.0 ([a586909](https://github.com/globalso-labs/telemetry/commit/a586909f35e8654addf377b25734d943050669e6))
* remove WithFields function from zerolog logger ([3f8f37c](https://github.com/globalso-labs/telemetry/commit/3f8f37cd608997cac6c801c73df93736f6977453))
* update definitions ([fb954db](https://github.com/globalso-labs/telemetry/commit/fb954db6ffcded18f553b01d904e8f5f9f92b454))
* update dependecies ([31232bb](https://github.com/globalso-labs/telemetry/commit/31232bba87e76d2b49a88bbda67261c0c8916439))
* update local dependencies ([7312513](https://github.com/globalso-labs/telemetry/commit/731251357da84192ee518cfac818936a3e0e82ae))
* update local go version ([e157c4c](https://github.com/globalso-labs/telemetry/commit/e157c4cfd301d8525bffdf256a35da7df9c12361))
* update logger to use zerolog levels ([9fc556c](https://github.com/globalso-labs/telemetry/commit/9fc556ced0bf5a8b37d5cc06fd9c4f323538449d))
* update sdk logger ([2611ffb](https://github.com/globalso-labs/telemetry/commit/2611ffb79028c90e1442a9859c8fc15da8557e7b))
* update structure ([f2735b8](https://github.com/globalso-labs/telemetry/commit/f2735b8a8eaba637ae23616d458053ade60f708a))
* update telemetry configuration and remove deprecated constants ([790d57a](https://github.com/globalso-labs/telemetry/commit/790d57a8a62522a9f4bdbba9979dde60d8727f99))
* update telemetry initialization and lint configuration ([753d032](https://github.com/globalso-labs/telemetry/commit/753d032ea4a8e66dac5ba698c29dd9df1b032e17))


### Bug Fixes

* rename module to lower case ([9ef6f1f](https://github.com/globalso-labs/telemetry/commit/9ef6f1f3991c14573a7c03e7858f0ca558b9d83d))
* update equality assertions in logger tests ([5365a5f](https://github.com/globalso-labs/telemetry/commit/5365a5f3adb191710807409e258808b208f4e321))
* update golangci-lint installation command for overwrite safety ([5803477](https://github.com/globalso-labs/telemetry/commit/580347783acaca70fd5a821b7c251a65563527da))
* update golangci-lint installation process ([e50f36c](https://github.com/globalso-labs/telemetry/commit/e50f36c4125f6ef39f43badd606a034e6ac2792a))
* update golangci-lint path to ensure consistency ([7ed292f](https://github.com/globalso-labs/telemetry/commit/7ed292fd19971e3ea2a0043639ac7fa188434ec0))
* update resource initialization to use global variable ([1275f85](https://github.com/globalso-labs/telemetry/commit/1275f85029b8f9f1e4b70f3404b1ea2bca9a47ff))


### Miscellaneous Chores

* release 0.2.0 ([f6ee5d9](https://github.com/globalso-labs/telemetry/commit/f6ee5d9505b9cf7d15a2c4337ec703de0ec677af))
* release 0.3.0 ([934ddac](https://github.com/globalso-labs/telemetry/commit/934ddaccc3283570336f3a0c2c846eb33146658b))
* release 1.9.0 ([7015b5b](https://github.com/globalso-labs/telemetry/commit/7015b5bfea37be9c8ba410d74c075e2ec8dbccab))
* update dependencies to latest versions ([74069f7](https://github.com/globalso-labs/telemetry/commit/74069f74e08901c0c2d1d1a2d8bbb058f2dc181e))
* update dependencies to latest versions ([4c810c7](https://github.com/globalso-labs/telemetry/commit/4c810c7ff63c8ba1f5bc47eb94087ad9bc441ba6))


### Code Refactoring

* remove OpenTelemetry collector agent implementation ([868f941](https://github.com/globalso-labs/telemetry/commit/868f9415b6674948ed95020f35d8676e470eb34a))
* update logger_test to use assert package ([0c8b70a](https://github.com/globalso-labs/telemetry/commit/0c8b70a0a9815e98be7adb25efdbd2698bef709b))
* update workflows and dependencies to latest versions ([ac38dec](https://github.com/globalso-labs/telemetry/commit/ac38dece8e7da795d500c9d51aee36bb8ab82478))

## [2.0.0](https://github.com/globalso-labs/telemetry/compare/v1.8.0...v2.0.0) (2025-11-28)


### ⚠ BREAKING CHANGES

* Removes Sentry logging hook and related dependency.

### Bug Fixes

* rename module to lower case ([9ef6f1f](https://github.com/globalso-labs/telemetry/commit/9ef6f1f3991c14573a7c03e7858f0ca558b9d83d))


### Code Refactoring

* update workflows and dependencies to latest versions ([ac38dec](https://github.com/globalso-labs/telemetry/commit/ac38dece8e7da795d500c9d51aee36bb8ab82478))

## [1.8.0](https://github.com/globalso-labs/telemetry/compare/v2.0.0...v1.8.0) (2025-06-19)


### ⚠ BREAKING CHANGES

* OpenTelemetry collector agent is no longer included or supported in this repository. All related binaries, configuration, and integration points are removed. Downstream consumers must migrate to external collector solutions.
* Dependency updates may introduce breaking changes from upstream libraries. Review release notes of updated dependencies for possible API or behavior changes.
* testing method renamed
* The removal of the internal package from the logger initialization may impact other tests relying on it.
* The telemetry configuration now requires the "x-scope-orgid" header.
* golang.org/x/tools updated from v0.21.1-0.20240508182429-e35e4ccd0d2d to v0.23.0, google.golang.org/grpc updated from v1.64.0 to v1.65.0.
* Change in assertion method may affect test outcomes if there are pointer comparisons.
* The metrics package has been renamed to meter. The configuration structure has been updated to reflect the new meter options and logger integration.

### Features

* add configuration and basic test to allow config initialization ([84075a8](https://github.com/globalso-labs/telemetry/commit/84075a88379bcc1dd80410b77c47c8bfb220e9db))
* add constants for service names and update resource initialization ([8a8e012](https://github.com/globalso-labs/telemetry/commit/8a8e0129acf201416799b7961ae7961c0a2e858e))
* add default headers for telemetry configuration ([ce8d08f](https://github.com/globalso-labs/telemetry/commit/ce8d08f2637f5a4d1f43ceb7b053d2eb73307fb7))
* add GitHub Actions workflow for automated releases ([cc2fce3](https://github.com/globalso-labs/telemetry/commit/cc2fce33adcde5ac6b44613e636969ac6f0ee820))
* add goreleaser configuration for multiple platforms ([770f254](https://github.com/globalso-labs/telemetry/commit/770f254e74c8e9cf69f4c19fdd2659a96beef084))
* add initial implementation of telemetry agent ([d143bf4](https://github.com/globalso-labs/telemetry/commit/d143bf48cb43dfb48861360da030b3e06e5bb360))
* add logging and metrics with configuration support ([02fb9b2](https://github.com/globalso-labs/telemetry/commit/02fb9b28eac9323a95d5eb5b7574fa88c879e986))
* add metrics implementation ([ace9e82](https://github.com/globalso-labs/telemetry/commit/ace9e82426eceeeb2f4f48235ffed2f8edd6ad04))
* add Prometheus Remote Write exporter and new receivers ([6783c4d](https://github.com/globalso-labs/telemetry/commit/6783c4dac5eba3e89fc50b05a2914568548cfe8e))
* add resource management to Telemetry ([a2789e9](https://github.com/globalso-labs/telemetry/commit/a2789e94fb302abb5c4a24da78e0bd092673bf0f))
* add tracer implementation ([c377b5a](https://github.com/globalso-labs/telemetry/commit/c377b5aa9db9b4a93ff37b0beb38684eb1e18dc8))
* add update command to manage dependencies ([ee11eaa](https://github.com/globalso-labs/telemetry/commit/ee11eaa6a03e25fb281768d8a63c6dcd71a73ae6))
* add version tracking and telemetry middleware ([c638f39](https://github.com/globalso-labs/telemetry/commit/c638f39f63a62fc61555dc89f22929e58aacbc5f))
* add With function to create zerolog.Context ([65c81e7](https://github.com/globalso-labs/telemetry/commit/65c81e786606b40d6c57f8ee2cfcfeb29274097a))
* add WithKey function and update testing context handling ([2df39ca](https://github.com/globalso-labs/telemetry/commit/2df39cae4dc7181c3eddb770514ab6a69ed25e86))
* add WithTestingCtx method ([bfa449c](https://github.com/globalso-labs/telemetry/commit/bfa449ce0cebcabcfff2308b7d232110a488b0d5))
* enhance Resource functionality ([b315539](https://github.com/globalso-labs/telemetry/commit/b315539c3ad3792bad496f53198d7d14599b21b3))
* enhance telemetry configuration ([3c89f58](https://github.com/globalso-labs/telemetry/commit/3c89f589fe50158102812b6c4798fc1bb948303f))
* export zerolog.Logger as Instance type alias ([9b5f611](https://github.com/globalso-labs/telemetry/commit/9b5f6116cae5e037b63965a7a1133fc51154a080))
* implement merge function for key/value map ([7d922d1](https://github.com/globalso-labs/telemetry/commit/7d922d1d77aa6da7cbbf791e20c24ce3f16fe63b))
* implement parse verbosity function ([9878372](https://github.com/globalso-labs/telemetry/commit/9878372639d43e85e8ed29dcd14531b32af86b94))
* release 1.8.0 ([a586909](https://github.com/globalso-labs/telemetry/commit/a586909f35e8654addf377b25734d943050669e6))
* remove WithFields function from zerolog logger ([3f8f37c](https://github.com/globalso-labs/telemetry/commit/3f8f37cd608997cac6c801c73df93736f6977453))
* update definitions ([fb954db](https://github.com/globalso-labs/telemetry/commit/fb954db6ffcded18f553b01d904e8f5f9f92b454))
* update dependecies ([31232bb](https://github.com/globalso-labs/telemetry/commit/31232bba87e76d2b49a88bbda67261c0c8916439))
* update local dependencies ([7312513](https://github.com/globalso-labs/telemetry/commit/731251357da84192ee518cfac818936a3e0e82ae))
* update local go version ([e157c4c](https://github.com/globalso-labs/telemetry/commit/e157c4cfd301d8525bffdf256a35da7df9c12361))
* update logger to use zerolog levels ([9fc556c](https://github.com/globalso-labs/telemetry/commit/9fc556ced0bf5a8b37d5cc06fd9c4f323538449d))
* update sdk logger ([2611ffb](https://github.com/globalso-labs/telemetry/commit/2611ffb79028c90e1442a9859c8fc15da8557e7b))
* update structure ([f2735b8](https://github.com/globalso-labs/telemetry/commit/f2735b8a8eaba637ae23616d458053ade60f708a))
* update telemetry configuration and remove deprecated constants ([790d57a](https://github.com/globalso-labs/telemetry/commit/790d57a8a62522a9f4bdbba9979dde60d8727f99))
* update telemetry initialization and lint configuration ([753d032](https://github.com/globalso-labs/telemetry/commit/753d032ea4a8e66dac5ba698c29dd9df1b032e17))


### Bug Fixes

* update equality assertions in logger tests ([5365a5f](https://github.com/globalso-labs/telemetry/commit/5365a5f3adb191710807409e258808b208f4e321))
* update golangci-lint installation command for overwrite safety ([5803477](https://github.com/globalso-labs/telemetry/commit/580347783acaca70fd5a821b7c251a65563527da))
* update golangci-lint installation process ([e50f36c](https://github.com/globalso-labs/telemetry/commit/e50f36c4125f6ef39f43badd606a034e6ac2792a))
* update golangci-lint path to ensure consistency ([7ed292f](https://github.com/globalso-labs/telemetry/commit/7ed292fd19971e3ea2a0043639ac7fa188434ec0))
* update resource initialization to use global variable ([1275f85](https://github.com/globalso-labs/telemetry/commit/1275f85029b8f9f1e4b70f3404b1ea2bca9a47ff))


### Miscellaneous Chores

* release 0.2.0 ([f6ee5d9](https://github.com/globalso-labs/telemetry/commit/f6ee5d9505b9cf7d15a2c4337ec703de0ec677af))
* release 0.3.0 ([934ddac](https://github.com/globalso-labs/telemetry/commit/934ddaccc3283570336f3a0c2c846eb33146658b))
* update dependencies to latest versions ([74069f7](https://github.com/globalso-labs/telemetry/commit/74069f74e08901c0c2d1d1a2d8bbb058f2dc181e))
* update dependencies to latest versions ([4c810c7](https://github.com/globalso-labs/telemetry/commit/4c810c7ff63c8ba1f5bc47eb94087ad9bc441ba6))


### Code Refactoring

* remove OpenTelemetry collector agent implementation ([868f941](https://github.com/globalso-labs/telemetry/commit/868f9415b6674948ed95020f35d8676e470eb34a))
* update logger_test to use assert package ([0c8b70a](https://github.com/globalso-labs/telemetry/commit/0c8b70a0a9815e98be7adb25efdbd2698bef709b))

## [2.0.0](https://github.com/globalso-labs/telemetry/compare/v1.7.0...v2.0.0) (2025-06-19)


### ⚠ BREAKING CHANGES

* OpenTelemetry collector agent is no longer included or supported in this repository. All related binaries, configuration, and integration points are removed. Downstream consumers must migrate to external collector solutions.
* Dependency updates may introduce breaking changes from upstream libraries. Review release notes of updated dependencies for possible API or behavior changes.
* testing method renamed

### Features

* add WithKey function and update testing context handling ([2df39ca](https://github.com/globalso-labs/telemetry/commit/2df39cae4dc7181c3eddb770514ab6a69ed25e86))


### Bug Fixes

* update golangci-lint path to ensure consistency ([7ed292f](https://github.com/globalso-labs/telemetry/commit/7ed292fd19971e3ea2a0043639ac7fa188434ec0))


### Miscellaneous Chores

* update dependencies to latest versions ([74069f7](https://github.com/globalso-labs/telemetry/commit/74069f74e08901c0c2d1d1a2d8bbb058f2dc181e))


### Code Refactoring

* remove OpenTelemetry collector agent implementation ([868f941](https://github.com/globalso-labs/telemetry/commit/868f9415b6674948ed95020f35d8676e470eb34a))

## [1.7.0](https://github.com/globalso-labs/telemetry/compare/v1.6.1...v1.7.0) (2025-06-16)


### Features

* export zerolog.Logger as Instance type alias ([9b5f611](https://github.com/globalso-labs/telemetry/commit/9b5f6116cae5e037b63965a7a1133fc51154a080))

## [1.6.1](https://github.com/globalso-labs/telemetry/compare/v1.6.0...v1.6.1) (2025-05-14)


### Bug Fixes

* update resource initialization to use global variable ([1275f85](https://github.com/globalso-labs/telemetry/commit/1275f85029b8f9f1e4b70f3404b1ea2bca9a47ff))

## [1.6.0](https://github.com/globalso-labs/telemetry/compare/v1.5.0...v1.6.0) (2025-05-14)


### Features

* update definitions ([fb954db](https://github.com/globalso-labs/telemetry/commit/fb954db6ffcded18f553b01d904e8f5f9f92b454))
* update local dependencies ([7312513](https://github.com/globalso-labs/telemetry/commit/731251357da84192ee518cfac818936a3e0e82ae))
* update local go version ([e157c4c](https://github.com/globalso-labs/telemetry/commit/e157c4cfd301d8525bffdf256a35da7df9c12361))
* update sdk logger ([2611ffb](https://github.com/globalso-labs/telemetry/commit/2611ffb79028c90e1442a9859c8fc15da8557e7b))

## [1.5.0](https://github.com/globalso-labs/telemetry/compare/v1.4.0...v1.5.0) (2025-05-13)


### Features

* add WithTestingCtx method ([bfa449c](https://github.com/globalso-labs/telemetry/commit/bfa449ce0cebcabcfff2308b7d232110a488b0d5))
* update dependecies ([31232bb](https://github.com/globalso-labs/telemetry/commit/31232bba87e76d2b49a88bbda67261c0c8916439))

## [1.4.0](https://github.com/globalso-labs/telemetry/compare/v1.3.0...v1.4.0) (2025-03-19)


### Features

* implement parse verbosity function ([9878372](https://github.com/globalso-labs/telemetry/commit/9878372639d43e85e8ed29dcd14531b32af86b94))

## [1.3.0](https://github.com/globalso-labs/telemetry/compare/v1.2.0...v1.3.0) (2025-03-19)


### Features

* update structure ([f2735b8](https://github.com/globalso-labs/telemetry/commit/f2735b8a8eaba637ae23616d458053ade60f708a))

## [1.2.0](https://github.com/globalso-labs/telemetry/compare/v1.1.1...v1.2.0) (2024-09-26)


### Features

* add constants for service names and update resource initialization ([8a8e012](https://github.com/globalso-labs/telemetry/commit/8a8e0129acf201416799b7961ae7961c0a2e858e))
* add Prometheus Remote Write exporter and new receivers ([6783c4d](https://github.com/globalso-labs/telemetry/commit/6783c4dac5eba3e89fc50b05a2914568548cfe8e))
* add update command to manage dependencies ([ee11eaa](https://github.com/globalso-labs/telemetry/commit/ee11eaa6a03e25fb281768d8a63c6dcd71a73ae6))
* enhance telemetry configuration ([3c89f58](https://github.com/globalso-labs/telemetry/commit/3c89f589fe50158102812b6c4798fc1bb948303f))

## [1.1.1](https://github.com/globalso-labs/telemetry/compare/v1.1.0...v1.1.1) (2024-09-26)


### Bug Fixes

* update golangci-lint installation process ([e50f36c](https://github.com/globalso-labs/telemetry/commit/e50f36c4125f6ef39f43badd606a034e6ac2792a))

## [1.1.0](https://github.com/globalso-labs/telemetry/compare/v1.0.0...v1.1.0) (2024-09-26)


### Features

* add goreleaser configuration for multiple platforms ([770f254](https://github.com/globalso-labs/telemetry/commit/770f254e74c8e9cf69f4c19fdd2659a96beef084))


### Bug Fixes

* update golangci-lint installation command for overwrite safety ([5803477](https://github.com/globalso-labs/telemetry/commit/580347783acaca70fd5a821b7c251a65563527da))

## [1.0.0](https://github.com/globalso-labs/telemetry/compare/v0.6.0...v1.0.0) (2024-09-26)


### ⚠ BREAKING CHANGES

* The removal of the internal package from the logger initialization may impact other tests relying on it.
* The telemetry configuration now requires the "x-scope-orgid" header.

### Features

* add default headers for telemetry configuration ([ce8d08f](https://github.com/globalso-labs/telemetry/commit/ce8d08f2637f5a4d1f43ceb7b053d2eb73307fb7))
* add initial implementation of telemetry agent ([d143bf4](https://github.com/globalso-labs/telemetry/commit/d143bf48cb43dfb48861360da030b3e06e5bb360))
* add resource management to Telemetry ([a2789e9](https://github.com/globalso-labs/telemetry/commit/a2789e94fb302abb5c4a24da78e0bd092673bf0f))
* enhance Resource functionality ([b315539](https://github.com/globalso-labs/telemetry/commit/b315539c3ad3792bad496f53198d7d14599b21b3))
* implement merge function for key/value map ([7d922d1](https://github.com/globalso-labs/telemetry/commit/7d922d1d77aa6da7cbbf791e20c24ce3f16fe63b))
* update telemetry configuration and remove deprecated constants ([790d57a](https://github.com/globalso-labs/telemetry/commit/790d57a8a62522a9f4bdbba9979dde60d8727f99))
* update telemetry initialization and lint configuration ([753d032](https://github.com/globalso-labs/telemetry/commit/753d032ea4a8e66dac5ba698c29dd9df1b032e17))


### Code Refactoring

* update logger_test to use assert package ([0c8b70a](https://github.com/globalso-labs/telemetry/commit/0c8b70a0a9815e98be7adb25efdbd2698bef709b))

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
