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
* refactor: add custom error struct in elcontracts writer methods by @maximopalopoli in <https://github.com/Layr-Labs/eigensdk-go/pull/482>
  * The errors now follow this format: `error name (error code) - error description: underlying error`.

    For example, in this code section:
    ```go
	receipt, err := w.txMgr.Send(ctx, tx, waitForReceipt)
	if err != nil {
		wrappedError := SendError(err)
		return nil, wrappedError
	}
    ```
    The returned error if err is not nil would be `Send error (5) - Failed to send tx with err *underlying error*`

### Removed
