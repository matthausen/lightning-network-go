## Run bitcoind with docker

### Create a data volume
- `docker volume create --name=bitcoind-data`

### Run the container with your preferred configuration
- `docker run -v bitcoind-data:/bitcoin/.bitcoin --name=bitcoind-node -d \
-p 8333:8333 \
-p 127.0.0.1:8332:8332 \
kylemanna/bitcoind`
  
### Pass configuration as environment variable
- `docker run -v bitcoind-data:/bitcoin/.bitcoin --name=bitcoind-node -d \
  -p 8333:8333 \
  -p 127.0.0.1:8332:8332 \
  -e DISABLEWALLET=1 \
  -e PRINTTOCONSOLE=1 \
  -e RPCUSER=mysecretrpcuser \
  -e RPCPASSWORD=mysecretrpcpassword \
  kylemanna/bitcoind`
  
### or pass configuration directly from bitcoin.conf file
-`docker run -v bitcoind-data:/bitcoin/.bitcoin --name=bitcoind-node -d \
-p 8333:8333 \
-p 127.0.0.1:8332:8332 \
-v /etc/bitcoin.conf:/bitcoin/.bitcoin/bitcoin.conf \
kylemanna/bitcoind`