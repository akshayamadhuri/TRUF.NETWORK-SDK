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
	// Step 1: Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		log.Fatalf("PRIVATE_KEY not found in environment")
	}

	// Step 2: Initialize TRUF Network Client
	ctx := context.Background()
	pk, err := crypto.Secp256k1PrivateKeyFromHex(privateKey)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	signer := &auth.EthPersonalSigner{Key: *pk}
	tnClient, err := tnclient.NewClient(ctx, "https://staging.tsn.truflation.com", tnclient.WithSigner(signer))
	if err != nil {
		log.Fatalf("Failed to initialize TRUF network client: %v", err)
	}

	// Define the Truflation Stream ID and Provider Address
	streamID := util.GenerateStreamId("stf37ad83c0b92c7419925b7633c0e62") // Correct Stream ID
	providerAddress := "0x4710a8d8f0d845da110086812a32de6d90d7ff5c"       // Correct Provider Address
	log.Printf("Attempting to load stream with ID: %s and provider: %s", streamID, providerAddress)

	// Fetch data for the stream
	data := fetchData(ctx, tnClient, streamID, providerAddress)
	fmt.Println("Fetched Data:", data)

	// Calculate risk metrics based on the data
	riskMetrics := calculateRiskMetrics(data)
	fmt.Println("Risk Metrics:", riskMetrics)

	// Generate alerts based on the calculated metrics
	generateAlerts(riskMetrics)
}

// Helper Function: Fetch Data for a Single Stream
func fetchData(ctx context.Context, tnClient *tnclient.Client, streamID util.StreamId, providerAddress string) []types.StreamRecord {
	// Create EthereumAddress using NewEthereumAddressFromString
	dataProvider, err := util.NewEthereumAddressFromString(providerAddress)
	if err != nil {
		log.Fatalf("Invalid provider address: %v", err)
	}

	// Define date range directly with civil.Date
	dateFrom := civil.Date{Year: 2023, Month: 1, Day: 1} // 2023-01-01
	dateTo := civil.Date{Year: 2023, Month: 1, Day: 31}  // 2023-01-31

	log.Printf("Fetching stream: %s with provider: %s", streamID, providerAddress)
	streamLocator := types.StreamLocator{
		StreamId:     streamID,     // Use the updated stream ID
		DataProvider: dataProvider, // Use the updated EthereumAddress
	}
	stream, err := tnClient.LoadPrimitiveStream(streamLocator)
	if err != nil {
		log.Fatalf("Failed to load stream %s: %v", streamID, err)
	}

	// Fetch records for the specified date range
	records, err := stream.GetRecord(ctx, types.GetRecordInput{
		DateFrom: &dateFrom,
		DateTo:   &dateTo,
	})
	if err != nil {
		log.Fatalf("Failed to fetch records for stream %s: %v", streamID, err)
	}

	return records
}

// Helper Function: Calculate Risk Metrics
func calculateRiskMetrics(data []types.StreamRecord) map[string]float64 {
	riskMetrics := make(map[string]float64)

	// Example: Calculate Value-at-Risk (VaR) based on inflation data
	var sum float64
	for _, record := range data {
		value, _ := record.Value.Float64()
		sum += value
	}
	if len(data) > 0 {
		riskMetrics["VaR"] = sum / float64(len(data)) // Simple average (example)
	}

	return riskMetrics
}

// Helper Function: Generate Alerts
func generateAlerts(riskMetrics map[string]float64) {
	if VaR, exists := riskMetrics["VaR"]; exists && VaR > 1000 { // Example threshold
		fmt.Printf("ALERT: Portfolio Value-at-Risk exceeds threshold! VaR: %.2f\n", VaR)
	}
}
