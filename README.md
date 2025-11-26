# Local Crypto Provider Plugin  

This plugin provides a local implementation of crypto services for testing. E.g. Signing and verification with **temporary** created keys.

Note: If you reboot the crypto provider with this plugin, all keys are gone:) 

## Usage

Start the Plugin under the desired adress and connect any service to it. 

Environment Variables: 

|Variable|Example Value|
|--------|-----|
|CRYPTO_PROVIDER_HASHICORP_VAULT_ADDRESS|0.0.0.0:50051|