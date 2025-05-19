# 📊 TRUF.Network Inflation & Index Risk Analyzer

Analyze real-time inflation and index data from TRUF.Network using Go. This project fetches decentralized economic data streams, computes simplified Value-at-Risk (VaR), and raises alerts when risks exceed thresholds.

---

## 🧱 Features

- Fetch real-time inflation and index data from TRUF.Network
- Compute basic risk metrics (average-based VaR)
- Alert on elevated risk levels
- Modular and readable code
- Powered by the official TRUF Go SDK

---

## 📦 Prerequisites

Make sure you have:

- **Go** (v1.20 or later)
- **Git**
- **A text editor** (e.g., VS Code)
- **A `.env` file** with your TRUF Ethereum wallet `PRIVATE_KEY` and `WALLET_ADDRESS`
- **Internet connection** to reach [https://staging.tsn.truflation.com](https://staging.tsn.truflation.com)

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/akshayamadhuri/TRUF.NETWORK-SDK.git
cd truf-inflation-risk-analyzer
```

### 2. Create a `.env` File

```env
PRIVATE_KEY=your_private_key_here
WALLET_ADDRESS=your_wallet_address_here
```

> 🔐 **Important:** Never share your private key publicly!  
> You can generate your key from MetaMask or any Ethereum wallet.

---

## ⚙️ Setup and Installation

### 3. Initialize Go Module (if not already)

```bash
go mod init truf-inflation-risk-analyzer
```

### 4. Install Dependencies

```bash
go get github.com/trufnetwork/sdk-go
go get github.com/joho/godotenv
go get github.com/kwilteam/kwil-db/core/crypto
go get github.com/kwilteam/kwil-db/core/crypto/auth
go get github.com/golang-sql/civil
```

### 5. Tidy Up

This will generate your `go.sum` and finalize `go.mod`:

```bash
go mod tidy
```

---

## 🧠 Project Structure

```
.
├── main.go            # Main application logic
├── go.mod             # Go module file
├── go.sum             # Dependency checksums
└── .env               # Environment variables (not checked into Git)
```

---

## 🔧 Usage

Run the program using:

```bash
go run main.go
```

This will:

- Load your `.env` variables
- Connect to TRUF.Network
- Fetch inflation and index data
- Calculate VaR
- Print alerts if risk is too high

---

## 📈 Sample Output

```bash
Inflation Data:
Date (DateValue): 2023-01-03, Value: 31274.80
...

Inflation Risk Metrics: map[VaR:28502.3363]
ALERT: Portfolio Value-at-Risk exceeds threshold! VaR: 28502.34
```

---

## 📝 License

Distributed under the MIT License. See `LICENSE` for more information.
