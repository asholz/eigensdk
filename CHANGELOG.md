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

* refactor: encapsulate parameters into `TaskSignature` in [#487](https://github.com/Layr-Labs/eigensdk-go/pull/487)

  * Introduced `TaskSignature` struct to encapsulate parameters related to task signatures:
  * Updated `ProcessNewSignature` to accept a `TaskSignature` struct instead of multiple parameters.

    ```rust
    // BEFORE
    blsAggServ.ProcessNewSignature(
			context.Background(),
			taskIndex,
			taskResponse,
			blsSigOp1,
			testOperator1.OperatorId,
		)
    
    // AFTER
    taskSignature := NewTaskSiganture(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)

    blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
    ```
  
### Removed
