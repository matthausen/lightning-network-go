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
rpcuser=bitcoinrpc
rpcpassword=some_password_goes_here
testnet=0
regtest=1
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

wi the following content: 
```
bitcoin.regtest=1
bitcoin.active=1
bitcoin.node=bitcoind
debuglevel=debug
```

Add these aliases to you ~/.bash_profile or ~/.zshrc:
```
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin

alias lnd=$HOME/go/bin/lnd
alias lncli=$HOME/go/bin/lncli
```


## How to run

Run lnd fior the first time and create a wallet:

- `lnd`
- `lncli create`

To run and stop bitcoind:

- `bitcoind`
- `bitcoin-cli stop`

Run lnd for a second time and unlock the previously created wallet with your password:

- `lnd`
- `lncli unlock`

To stop lncli:

- `lncli stop`

