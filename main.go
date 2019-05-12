package main

import (
	"Go-Ethereum-Cli/lib/ethgo"
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
	"strconv"
	//"golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
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
		case "3":
			for input.Scan() {
				str := strings.Split(input.Text(), " ")
				prvk, amounnt, to := str[0], str[1], str[2]
				privateKeym, err := crypto.HexToECDSA(prvk)
				if err != nil {
					log.Fatal(err)
				}
				pubkey := privateKeym.Public()
				pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("Cannot accept type publickey is njot of the type *ecdsa .PublicKey")
				}
				fromAdds := crypto.PubkeyToAddress(*pubkeyECDSA)
				nonce, err := ethgo.GetNonce(mainNetClient, fromAdds)
				parsedPrice, err := strconv.ParseFloat(amounnt, 64)
				val := ethgo.BigInt(ethgo.Towei(parsedPrice))
				gasLimit := uint64(21000)
				gasPrice, err := ethgo.SuggestGas(mainNetClient)
				toAddress := ethgo.SetAddress(to)
				var data []byte
				tx := ethgo.GenTransaction(nonce, toAddress, val, gasLimit, gasPrice, data)
				chainID := ethgo.Gen
				signedTx, err := ethgo.SignTx(tx, chainID, privateKeym)
			}
		default:
			print("1-ReadBalance || 2-GenerateEthereumWallet || 3-Transfer || 0-Exit")
		}
	}
}
