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
* refactor: update interface on `bls aggregation` in [#485](https://github.com/Layr-Labs/eigensdk-go/pull/485).
  * Introduces a new struct `TaskMetadata` with a constructor `NewTaskMetadata` to initialize a new task and a method `WithWindowDuration` to set the window duration.
  * Refactors `InitializeNewTask` and `singleTaskAggregatorGoroutineFunc` to accept a `TaskMetadata` struct instead of multiple parameters.

    ```rust
    // BEFORE
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

    blsAggServ.InitializeNewTask(
			taskIndex,
			blockNum,
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		)
    
    // AFTER
    blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

    metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
    blsAggServ.InitializeNewTask(metadata)
    ```
### Removed
