# Signer v2

Signerv2 is a module for signing and verifying messages. It provides a simple and secure way to handle cryptographic signatures.

## Features

- Sign messages using various algorithms using Web3Signer
- Verify message signatures
- Support for multiple key types

### Comparison to Old Signer

In comparison to the old signer, Signerv2 offers:
- New signing mechaninsms
- Simplified API for easier integration

### Using SignerFromConfig

SignerV2 introduces `SignerFromConfig`

The `SignerFromConfig` function allows you to create a signer function based on a configuration. This configuration can specify whether to use a private key signer, a local keystore signer, or a remote web3 signer.

```go
package main

import (
    "github.com/Layr-Labs/eigensdk-go/signerv2"
)

func main() {
    config := signerv2.Config{
        // ...initialize your configuration...
    }
    chainID := // Set your chain ID
    signerFn, signerAddr, err := signerv2.SignerFromConfig(config, chainID)
    if err != nil {
        // Handle error
        return
    }
    // Use signerFn and signerAddr as needed
}
```
### Web3Signer
`Web3SignerFn` creates a signer function that uses a remote signer
It exposes `eth_SignTransaction` endpoint which return rlp encoded signed tx

## Upgrade from Signer (v1)

`NewPrivateKeySigner` and `NewPrivateKeyFromKeystoreSigner` functions should be upgraded to use the new `SignerFromConfig` with its correct Config
