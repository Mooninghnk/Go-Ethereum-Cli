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
)

func print(val ...interface{}) {
	for _, x := range val {
		fmt.Println(x)
	}
}

func setAddress(address string) common.Address {
	return common.HexToAddress(address)
}

func getFixedBalance(client *ethclient.Client, account common.Address) (*big.Int, error) {
	return client.BalanceAt(context.Background(), account, nil)
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
	print("1 ReadBalance - 0 Exit")
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
					print("1 ReadBalance - 0 Exit")
					break
				}else{
					print("not long enough: " + hexAddress)
					print("1 ReadBalance - 0 Exit")
					break
				}
			}
		case "0":
			os.Exit(0)
		default:
			print("1 ReadBalance - 0 Exit")
		}
	}
}
