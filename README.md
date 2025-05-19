# 📊 TRUF.Network Inflation & Risk Analyzer

Analyze real-time inflation and index data from TRUF.Network. This repository contains modular components and a combined script to track economic data and compute portfolio risk.


## ⚙ Prerequisites

- Go (v1.20+)
- Git

You will also need a `.env` file in the root with your credentials:

```env
PRIVATE_KEY=your_private_key
WALLET_ADDRESS=your_wallet_address
```

> **Warning**
> Never share your private key publicly!

---

##  Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/akshayamadhuri/TRUF.NETWORK-SDK.git
cd TRUF.NETWORK-SDK
```

### 2. Set Up Your Environment

Create a `.env` file in the root directory:

```env
PRIVATE_KEY=your_private_key
WALLET_ADDRESS=your_wallet_address
```

### 3. Initialize Go Module

If running for the first time:

```bash
go mod init truf-inflation-risk-analyzer
```

Then fetch dependencies:

```bash
go get github.com/trufnetwork/sdk-go
go get github.com/joho/godotenv
go get github.com/kwilteam/kwil-db/core/crypto
go get github.com/kwilteam/kwil-db/core/crypto/auth
go mod tidy
```

---

##  Run Individual Modules

Each folder contains its own `main.go` file—you can run each part independently.

###  Inflation Tracker

```bash
cd Inflation-Tracker
go run main.go
```

###  Index Stream

```bash
cd 'Index value'
go run mai.go
```

###  Risk Calculator

```bash
cd 'Calculating Portfolio Risk'
go run main.go
```

---

##  Run Combined Script

To test the complete end-to-end functionality, use:

```bash
go run combined\ main.go
```

This will:

- Load environment variables
- Connect to TRUF.Network
- Fetch both inflation & index values
- Calculate simplified Value-at-Risk (VaR)
- Print alerts if needed

---

## 🔎 Example Output

```yaml
Inflation Data:
Date: 2023-01-01, Value: 30245.67
...

Risk Metrics: map[VaR:28502.3363]
ALERT: Portfolio Value-at-Risk exceeds threshold! VaR: 28502.34
```


## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---
