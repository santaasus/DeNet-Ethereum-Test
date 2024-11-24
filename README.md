# Go Test Task

## gRPC server:

Implement GetAccount() method:

Request: { ethereum_address, crypto_signature }

Response: { gastoken_balance, wallet-nonce }

It is necessary to validate the signature (crypto_signature) of the client address (ethereum_address) and return a response only if it is valid, otherwise, return an error

Implement GetAccounts() bidirectional stream method:

Request stream: { [ ] ethereum_addresses, erc20_token_address }

Response stream: { [ ] { ethereum_address, erc20_balance } }

For each address from the list of addresses (ethereum_address), the method must return its token balance (erc20_token_address).

## gRPC client:

Call GetAccount():

You can generate a wallet in the client application or import an existing one and sign the message with it

In order to check that the server method returns the correct data, the wallet balance must contain tokens. You can use test Ethereum tokens

Run the client method for GetAccounts() for 100 addresses and 3-5 tokens (signing the message is not required)

Additionally measure the processing time

Resources:
- https://goethereumbook.org/en
- https://geth.ethereum.org/docs/tools/abigen
- https://grpc.io/docs/languages/go/quickstart/


