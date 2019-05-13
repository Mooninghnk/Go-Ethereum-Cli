package main

import (
	"Go-Ethereum-Cli/lib/ethgo"
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"strconv"

	//"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func print(val ...interface{}) {
	for _, x := range val {
		fmt.Println(x)
	}
}

func main() {
	mainNetClient, err := ethclient.Dial("https://ropsten.infura.io")
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
					set := ethgo.SetAddress(hexAddress)
					fixedBalance, err := ethgo.GetFixedBalance(mainNetClient, set)
					if err != nil {
						log.Fatal(err)
					}
					resultBalnace := ethgo.ToEth(fixedBalance)
					fmt.Printf("%g%s\n", resultBalnace, "  ETH")
					print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
					break
				} else {
					print("not long enough: " + hexAddress)
					print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
					break
				}
			}
		case "0":
			os.Exit(0)
		case "2":
			privKey, err := ethgo.GenPrivateKey()
			if err != nil {
				log.Fatal(err)
			}
			privKeyBytes := ethgo.GetPrivateKeyBites(privKey)
			publickey := ethgo.GenPublicKey(privKey)
			publickeyACD, ok := publickey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal(err)
			}
			fmt.Printf("\n%s\n%s\n\n", "Private Key: "+ethgo.HexEncode(privKeyBytes), "PublicKey: "+ethgo.GetPublicAddress(publickeyACD))
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit\n")
		case "3":
			for input.Scan() {
				client := mainNetClient
				str := strings.Split(input.Text(), " ")
				prvk, amount, to := str[0], str[1], str[2]
				privateKey, err := crypto.HexToECDSA(prvk[2:])
				if err != nil {
					log.Fatal(err)
				}

				publicKey := privateKey.Public()
				publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
				nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
				if err != nil {
					log.Fatal(err)
				}
				preValue, err := strconv.ParseFloat(amount, 64)
				if err != nil {
					log.Fatal(err)
				}
				value := big.NewInt(toWei(preValue)) // in wei (1 eth)
				gasLimit := uint64(21000)            // in units
				gasPrice, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				toAddress := common.HexToAddress(to)
				var data []byte
				tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

				chainID, err := client.NetworkID(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
				if err != nil {
					log.Fatal(err)
				}

				err = client.SendTransaction(context.Background(), signedTx)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

				break
			}
		default:
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
		}
	}
}

func hexToAddress(hex string) common.Address {
	return common.HexToAddress(hex)
}
func suggestGas(client *ethclient.Client) (*big.Int, error) {
	return client.SuggestGasPrice(context.Background())
}
func toWei(num float64) int64 {
	return int64(num * 1e18)
}
func hexToECDSA(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKey)
}

func getNonce(client *ethclient.Client, fromAddress common.Address) (uint64, error) {
	return client.PendingNonceAt(context.Background(), fromAddress)

}
