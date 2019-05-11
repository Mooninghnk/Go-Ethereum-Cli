package main

import (
	"Go-Ethereum-Cli/lib/ethgo"
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	//"golang.org/x/crypto/sha3"
)

func print(val ...interface{}) {
	for _, x := range val {
		fmt.Println(x)
	}
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
		default:
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
		}
	}
}
