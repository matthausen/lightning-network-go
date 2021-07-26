## How to run LND with docker
- `docker pull lightninglabs/lnd:v0.12.0-beta` 
- `docker run lightninglabs/lnd [command-line options]`

### Volumes
A Docker volume will be created with your .lnd directory automatically, and will persist through container restarts.
You can also optionally manually specify a local folder to be used as a volume:
- `docker create --name=mylndcontainer -v /media/lnd-docker/:/root/.lnd myrepository/lnd [command-line options]`

### Run lnd with bitcoind
-`docker run --name lnd-mainnet lightninglabs/lnd --bitcoin.active --bitcoin.mainnet --bitcoin.node=bitcoind`
