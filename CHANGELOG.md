# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Each version will have a separate `Breaking Changes` section as well. To describe in how to upgrade from one version to another if needed

## [Unreleased]
### Added
### Changed
* fix: change requested pr url in changelog's workflow by @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/575>

### Breaking changes
* refactor: add custom error struct in elcontracts reader methods by @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/477>
  * The errors now follow this format: `error name (error code) - error description: underlying error`.

    For example, in this code section:
    ```go
	distributionRootsLength, err := r.rewardsCoordinator.GetDistributionRootsLength(&bind.CallOpts{Context: ctx})
	if err != nil {
		wrappedError := BindingError("RewardsCoordinator.getDistributionRootsLength", err)
		return &big.Int{}, wrappedError
	}
    ```
    The returned error if err is not nil would be `Binding error (0) - Error happened while calling RewardsCoordinator.getDistributionRootsLength: *underlying error*`

### Removed
