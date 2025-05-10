package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-sql/civil"
	"github.com/joho/godotenv"
	"github.com/kwilteam/kwil-db/core/crypto"
	"github.com/kwilteam/kwil-db/core/crypto/auth"
	"github.com/trufnetwork/sdk-go/core/tnclient"
	"github.com/trufnetwork/sdk-go/core/types"
	"github.com/trufnetwork/sdk-go/core/util"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve PRIVATE_KEY from .env
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatalf("PRIVATE_KEY not found in environment")
	}

	ctx := context.Background()

	// Create a signer using the private key
	pk, err := crypto.Secp256k1PrivateKeyFromHex(privateKey)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	signer := &auth.EthPersonalSigner{Key: *pk}

	// Initialize the TRUF network client
	tnClient, err := tnclient.NewClient(ctx, "https://staging.tsn.truflation.com", tnclient.WithSigner(signer))
	if err != nil {
		log.Fatalf("Failed to initialize TRUF network client: %v", err)
	}

	// Generate a Stream ID or use an existing one
	streamId := util.GenerateStreamId("stf37ad83c0b92c7419925b7633c0e62") // Generates a util.StreamId
	streamLocator := types.StreamLocator{
		StreamId: streamId,
		DataProvider: func() util.EthereumAddress {
			address, err := util.NewEthereumAddressFromString("0x4710a8d8f0d845da110086812a32de6d90d7ff5c")
			if err != nil {
				log.Fatalf("Failed to get address: %v", err)
			}
			return address
		}(),
	}

	// Load Primitive Actions
	primitiveActions, err := tnClient.LoadComposedStream(streamLocator)
	if err != nil {
		log.Fatalf("Failed to load composed stream: %v", err)
	}

	// Fetch and display inflation data using timestamps
	dateFrom := civil.Date{Year: 2023, Month: 1, Day: 1}
	dateTo := civil.Date{Year: 2023, Month: 1, Day: 31}

	// Fetch records
	records, err := primitiveActions.GetRecord(ctx, types.GetRecordInput{
		DateFrom: &dateFrom, // Use civil.Date
		DateTo:   &dateTo,   // Use civil.Date
	})
	if err != nil {
		log.Fatalf("Failed to fetch inflation data: %v", err)
	}

	// Print the fetched inflation data
	fmt.Println("Inflation Data:")
	for _, record := range records {
		fmt.Printf("Date (DateValue): %d, Value: %s\n", record.DateValue, record.Value.String())
	}

	// Fetch index data
	indexes, err := primitiveActions.GetIndex(ctx, types.GetIndexInput{
		DateFrom: &dateFrom, // Use civil.Date
		DateTo:   &dateTo,   // Use civil.Date
	})
	if err != nil {
		log.Fatalf("Failed to fetch index data: %v", err)
	}

	// Print the fetched index data
	fmt.Println("Index Data:")
	for _, index := range indexes {
		fmt.Printf("Date (DateValue): %d, Index Value: %s\n", index.DateValue, index.Value.String())
	}
}
