package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"os"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	//"golang.org/x/crypto/sha3"
)

func print(val ...interface{}) {
	for _, x := range val {
		fmt.Println(x)
	}
}

func genPrivateKey()(*ecdsa.PrivateKey, error){
	//generate a privateKey
	return crypto.GenerateKey()
}
func getPrivateKeyBites(privkey *ecdsa.PrivateKey) ([]byte) {
	//turn the key into  bytes
	return crypto.FromECDSA(privkey)
}

func hexEncode(privkeybytes []byte) (string) {
	//turn priavate key bytes into string
	return hexutil.Encode(privkeybytes)
}
func genPublicKey(privkey *ecdsa.PrivateKey) (interface{}) {
	return privkey.Public()
}


func publicKeyBytes(publicKeyECDSA *ecdsa.PublicKey) ([]byte) {
	return crypto.FromECDSAPub(publicKeyECDSA)
}

func setAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func getFixedBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), account, nil)
}

func getPublicAddress(publicKeyECDSA *ecdsa.PublicKey) (string) {
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

func toEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func getPendingBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.PendingBalanceAt(context.Background(), account)
}

func main() {
	mainNetClient, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch text {
		case "1":
			print("Input your address")
			for input.Scan() {
				hexAddress := input.Text()
				if len(hexAddress) == 42 {
					set := setAddress(hexAddress)
					fixedBalance, err := getFixedBalance(mainNetClient, set)
					if err != nil {
						log.Fatal(err)
					}
					resultBalnace := toEth(fixedBalance)
					fmt.Printf("%g%s\n", resultBalnace, "  ETH")
					print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
					break
				}else{
					print("not long enough: " + hexAddress)
					print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
					break
				}
			}
		case "0":
			os.Exit(0)
		case "2":
			privKey, err := genPrivateKey()
			if err != nil {
				log.Fatal(err)
			}
			privKeyBytes := getPrivateKeyBites(privKey)
			publickey := genPublicKey(privKey)
			publickeyACD, ok := publickey.(*ecdsa.PublicKey)
			if !ok  {
				log.Fatal(err)
			} 
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit\n")
			fmt.Printf("%s\n%s\n", "Private Key: "+ hexEncode(privKeyBytes), "PublicKey: "+ getPublicAddress(publickeyACD))
		default:
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
		}
	}
}
