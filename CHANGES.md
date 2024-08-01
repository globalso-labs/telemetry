# Changelog

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
