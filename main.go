package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"time"

	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"gopkg.in/macaroon.v2"
)

//TODO: add docker image for bitcoind client
//TODO: add docker image for lightning network cli (lncli)

type rpcCreds map[string]string

func (m rpcCreds) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return m, nil
}
func (m rpcCreds) RequireTransportSecurity() bool { return true }

// Macaroons global var
var MACAROONOPTION grpc.CallOption

func main() {
	grpcConn := grpcSetup()
	defer grpcConn.Close()

	lncli := lnrpc.NewLightningClient(grpcConn)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	walletBalanceReq := lnrpc.WalletBalanceRequest{}
	walletRes, err := lncli.WalletBalance(ctx, &walletBalanceReq, MACAROONOPTION)
	if err != nil {
		fmt.Printf("Error getting wallet balance: %s", err)
	}
	fmt.Println(walletRes.Balance)

	/*newAddressRequest := lnrpc.NewAddressRequest{Type: 0}
	newAddrRes, err := lncli.NewAddress(ctx, &newAddressRequest, MACAROONOPTION)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newAddrRes.Address)*/

}

func newCreds(bytes []byte) rpcCreds {
	creds := make(map[string]string)
	creds["macaroon"] = hex.EncodeToString(bytes)
	return creds
}

func grpcSetup() *grpc.ClientConn {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	homeDir := usr.HomeDir
	lndDir := fmt.Sprintf("%s/Library/Application Support/Lnd", homeDir)

	// SSL credentials setup
	var serverName string
	certFileLocation := fmt.Sprintf("%s/tls.cert", lndDir)
	creds, err := credentials.NewClientTLSFromFile(certFileLocation, serverName)
	if err != nil {
		fmt.Println(err)
	}

	macaroonFileLocation := fmt.Sprintf("%s/data/chain/bitcoin/mainnet/admin.macaroon", lndDir)
	macaroonMap := map[string]string{"macaroon": macaroonFileLocation}
	macaroonMetadata := metadata.New(macaroonMap)
	MACAROONOPTION = grpc.Header(&macaroonMetadata)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	macaroonBytes, err := ioutil.ReadFile(macaroonFileLocation)
	if err != nil {
		panic(fmt.Sprintln("Cannot read macaroon file", err))
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macaroonBytes); err != nil {
		panic(fmt.Sprintln("Cannot unmarshal macaroon", err))
	}

	conn, err := grpc.DialContext(ctx, "localhost:10009", []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newCreds(macaroonBytes)),
	}...)

	if err != nil {
		panic(fmt.Errorf("unable to connect to localhost:10009: %w", err))
	}

	return conn
}
