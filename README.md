This is a go gRPC server which shows how you can create a wallet via the Lightning Network (LND) and display its balance.

You have to run the Bitcoin Core Daemon (bitcoind) in order for lnd to perform its actions.

More methods and functionalities available from the LND client here: 
https://pkg.go.dev/github.com/lightningnetwork/lnd/lnrpc?utm_source=godoc#LightningClient

## Initial configuration

Install Go:
- `brew install go`

Install Bitcoin Core Daemon:
- `brew install bitcoind`

Add the bitcoin.conf file:
```
mkdir -p "/Users/${USER}/Library/Application Support/Bitcoin"
touch "/Users/${USER}/Library/Application Support/Bitcoin/bitcoin.conf"
chmod 600 "/Users/${USER}/Library/Application Support/Bitcoin/bitcoin.conf
```
with the following content:

```
rpcuser=lnd-user
rpcpassword=lnd-password
mainnet=1
testnet=0
regtest=0
server=1
daemon=1
zmqpubrawblock=tcp://127.0.0.1:28332
zmqpubrawtx=tcp://127.0.0.1:28333
```

Install LND:

- `git clone https://github.com/lightningnetwork/lnd`
- `cd lnd`
- `make install`

Create a LND config file:

```
mkdir -p "/Users/${USER}/Library/Application Support/Lnd"
touch "/Users/${USER}/Library/Application Support/Lnd/Lnd.conf"
chmod 600 "/Users/${USER}/Library/Application Support/Lnd/Lnd.conf
```

with the following content: 
```
 ## LND Settings
# Lets LND know to run on top of Bitcoin (as opposed to Litecoin)
bitcoin.active=true
bitcoin.mainnet=true
# Lets LND know you are running Bitcoin Core (not btcd or Neutrino)
bitcoin.node=bitcoind


## Bitcoind Settings
# Tells LND what User/Pass to use to RPC to the Bitcoin node
bitcoind.rpcuser=lnd-user
bitcoind.rpcpass=lnd-password
# Allows LND & Bitcoin Core to communicate via ZeroMQ
bitcoind.zmqpubrawblock=tcp://127.0.0.1:28332
bitcoind.zmqpubrawtx=tcp://127.0.0.1:28333
```

Add these aliases to you ~/.bash_profile or ~/.zshrc:
```
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin

alias lnd=$HOME/go/bin/lnd
alias lncli=$HOME/go/bin/lncli
```


## How to run

Run lnd for the first time and create a wallet:

- `lnd`
- `lncli create`

To run and bitcoind:

- `bitcoind`

Run lnd for a second time and unlock the previously created wallet with your password:

- `lnd`
- `lncli unlock`


## How to stop

Stop bitcoind

- `bitcoin-cli stop`

Stop lncli:

- `lncli stop`


