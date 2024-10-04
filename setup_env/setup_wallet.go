package setup_env

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/blocto/solana-go-sdk/types"
	"os"
)

type PrivateKeyFileDataRead struct {
	PrivateKey string
	PublicKey  string
}

func SetupWallet() types.Account {
	data, err := os.ReadFile("wallet.json")
	if err != nil {
		fmt.Printf("Error catched: %s", err)
		return createWallet()
	}
	var privateKeyData PrivateKeyFileDataRead
	err = json.Unmarshal(data, &privateKeyData)
	privateKeyBytes, err := hex.DecodeString(privateKeyData.PrivateKey)
	account, err := types.AccountFromBytes(privateKeyBytes)
	return account
}

func createWallet() types.Account {
	address := types.NewAccount()
	keyPairData := struct {
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"` // Store private key as a hex string
	}{
		PublicKey:  address.PublicKey.String(),
		PrivateKey: hex.EncodeToString(address.PrivateKey), // Encode private key as hex string
	}
	file, _ := os.Create("wallet.json")
	err := json.NewEncoder(file).Encode(keyPairData)
	if err != nil {
		return types.Account{}
	}
	return address
}
