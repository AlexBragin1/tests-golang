package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	infuraURL       = "https://mainnet.infura.io/v3/id"
	chainlinkETHUSD = "0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419"
	wethAddress     = "0xC02aaA39b223FE8D0A0e6F39d5c161e490fed0e9"
	erc20ABI        = `[{"constant":true,"inputs":[],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"}]`
)

func main() {

	address := common.HexToAddress(os.Args[1])

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	ethPrice, err := getPriceInUSD(client, chainlinkETHUSD)
	if err != nil {
		log.Fatalf("Failed to get ETH price: %v", err)
	}

	ethBalanceDecimal := decimal.NewFromBigInt(balance, 0).Div(decimal.NewFromFloat(1e18))
	ethInUSD := ethBalanceDecimal.Mul(ethPrice)

	//	wethAddressHex := common.HexToAddress(wethAddress)
	//	wethBalance, err := getERC20Balance(client, wethAddressHex, address)
	//	if err != nil {
	//		log.Fatalf("Failed to get WETH balance: %v", err)
	//	}

	//	wethBalanceDecimal := decimal.NewFromBigInt(wethBalance, 0).Div(decimal.NewFromFloat(1e18))
	//	wethInUSD := wethBalanceDecimal.Mul(ethPrice)

	fmt.Printf("ETH Balance: %s\n", ethBalanceDecimal.String())
	fmt.Printf("ETH in USD: %s\n", ethInUSD.String())
	// fmt.Printf("WETH Balance: %s\n", wethBalanceDecimal.String())
	// fmt.Printf("WETH in USD: $%.2f\n", wethInUSD)
}

func getPriceInUSD(client *ethclient.Client, priceFeedAddress string) (decimal.Decimal, error) {
	contractAddress := common.HexToAddress(priceFeedAddress)
	parsedABI, err := abi.JSON(strings.NewReader(`[
        {"constant":true,"inputs":[],"name":"latestAnswer","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"}
    ]`))
	if err != nil {
		return decimal.Zero, err
	}

	callData, err := parsedABI.Pack("latestAnswer")
	if err != nil {
		return decimal.Zero, err
	}

	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return decimal.Zero, err
	}

	var price *big.Int
	err = parsedABI.UnpackIntoInterface(&price, "latestAnswer", result)
	if err != nil {
		return decimal.Zero, err
	}

	priceDecimal := decimal.NewFromBigInt(price, 0).Div(decimal.NewFromFloat(1e8))
	fmt.Println(priceDecimal)
	return priceDecimal, nil
}

func getERC20Balance(client *ethclient.Client, tokenAddress common.Address, ownerAddress common.Address) (*big.Int, error) {
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return nil, err
	}

	// Вызываем функцию balanceOf с адресом владельца
	callData, err := parsedABI.Pack("balanceOf", ownerAddress)
	if err != nil {
		return nil, err
	}

	contractAddress := tokenAddress
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddress,
		Data: callData,
	}, nil)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "balanceOf", result)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
