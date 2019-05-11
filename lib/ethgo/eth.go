package ethgo

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenPrivateKey() (*ecdsa.PrivateKey, error) {
	//generate a privateKey
	return crypto.GenerateKey()
}
func GetPrivateKeyBites(privkey *ecdsa.PrivateKey) []byte {
	//turn the key into  bytes
	return crypto.FromECDSA(privkey)
}

func HexEncode(privkeybytes []byte) string {
	//turn priavate key bytes into string
	return hexutil.Encode(privkeybytes)
}
func GenPublicKey(privkey *ecdsa.PrivateKey) interface{} {
	return privkey.Public()
}

func PublicKeyBytes(publicKeyECDSA *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(publicKeyECDSA)
}

func SetAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func GetFixedBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), account, nil)
}

func GetPublicAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

func ToEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func GetPendingBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.PendingBalanceAt(context.Background(), account)
}
