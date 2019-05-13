package ethgo

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//GenPrivateKey function that generates a private key and returns it
func GenPrivateKey() (*ecdsa.PrivateKey, error) {
	//generate a privateKey
	return crypto.GenerateKey()
}

//GetPrivateKeyBites takes in a generated private key and returns that key in bytes
func GetPrivateKeyBites(privkey *ecdsa.PrivateKey) []byte {
	//turn the key into  bytes
	return crypto.FromECDSA(privkey)
}

//HexEncode takes in the key in bytes and returns a string with that key
func HexEncode(privkeybytes []byte) string {
	//turn priavate key bytes into string
	return hexutil.Encode(privkeybytes)
}

//GenPublicKey generates the public key from the privateket
func GenPublicKey(privkey *ecdsa.PrivateKey) interface{} {
	return privkey.Public()
}

//PublicKeyBytes turn the public key into bytes and returns bytes so we can work with it later
func PublicKeyBytes(publicKeyECDSA *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(publicKeyECDSA)
}

//SetAddress sets a hex address and returns it
func SetAddress(address string) common.Address {
	return common.HexToAddress(address)
}

//GetFixedBalance returns a fixed address in wei
func GetFixedBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), account, nil)
}

//GetPublicAddress returns the public address for your account
func GetPublicAddress(publicKeyECDSA *ecdsa.PublicKey) string {
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

//ToEth from wei to eth converter
func ToEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

//GetPendingBalance get pending balance of an account
func GetPendingBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.PendingBalanceAt(context.Background(), account)
}

// Towei converter
func Towei(num float64) int64 {
	return int64(num * 1e18)
}

//SuggestGas returns the best gas price in typof bigInt
func SuggestGas(client *ethclient.Client) (*big.Int, error) {
	return client.SuggestGasPrice(context.Background())
}

//GetNonce returns the nonce number
func GetNonce(client *ethclient.Client, fromAddress common.Address) (uint64, error) {
	return client.PendingNonceAt(context.Background(), fromAddress)
}

//BigInt returns a big int in wei
func BigInt(wei int64) *big.Int {
	return big.NewInt(wei)
}

//GenTransaction retuns a new transaction
func GenTransaction(nonce uint64, toAddress common.Address, val *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *types.Transaction {
	return types.NewTransaction(nonce, toAddress, val, gasLimit, gasPrice, data)
}

//GetNetworkID returns the network id of the current client
func GetNetworkID(client *ethclient.Client) (*big.Int, error) {
	return client.NetworkID(context.Background())
}

//SingTx returns a singed transaction
func SingTx(transaction *types.Transaction, chainID *big.Int, privkey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return types.SignTx(transaction, types.NewEIP155Signer(chainID), privkey)
}

//SendTransaction return an error if process was not seccesfull
func SendTransaction(client *ethclient.Client, singTX *types.Transaction) error {
	return client.SendTransaction(context.Background(), singTX)
}

//HexToEcdsa encodes a private key string to a ECDSA
func HexToEcdsa(key string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(key)
}
