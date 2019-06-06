package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
	horizon "github.com/stellar/go/clients/horizonclient"
)

var (
	sellingAssetCode   = flag.String("selling-code", "ABDT", "selling asset code")
	sellingAssetIssuer = flag.String("selling-issuer", "GDZURZR6RZKIQVOWZFWPVAUBMLLBQGXP2K5E5G7PEOV75IYPDFA36WK4", "selling asset issuer")
	buyingAssetCode    = flag.String("buying-code", "native", "buying asset code")
	buyingAssetIssuer  = flag.String("buying-issuer", "", "buying asset issuer")
)

func main() {
	flag.Parse()
	obRequest := horizon.OrderBookRequest{
		Limit:              200,
		SellingAssetCode:   *sellingAssetCode,
		SellingAssetIssuer: *sellingAssetIssuer,
		SellingAssetType:   horizon.AssetType4,
	}

	if *buyingAssetCode == "native" {
		obRequest.BuyingAssetType = horizon.AssetTypeNative
	} else {
		obRequest.BuyingAssetType = horizon.AssetType4
		obRequest.BuyingAssetCode = *buyingAssetCode
		obRequest.BuyingAssetIssuer = *buyingAssetIssuer
	}

	resp, err := horizonclient.DefaultPublicNetClient.OrderBook(obRequest)
	if err != nil {
		logrus.Fatal(err.(*horizon.Error).Problem)
	}

	fmt.Printf("Orderbook %v/%v\n\n", *sellingAssetCode, *buyingAssetCode)

	fmt.Printf("BUY %v with %v:\n", *sellingAssetCode, *buyingAssetCode)

	sum := 0.0
	msum := 0.0
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	w.Write([]byte(fmt.Sprintf("PRICE (%v/%v)\tAMOUNT\tSUM\tMULTIPLIED SUM\n", *buyingAssetCode, *sellingAssetCode)))
	for _, bid := range resp.Bids {
		amount, err := strconv.ParseFloat(bid.Amount, 64)
		if err != nil {
			logrus.Fatal(err)
		}
		price, err := strconv.ParseFloat(bid.Price, 64)
		if err != nil {
			logrus.Fatal(err)
		}
		sum += amount
		msum += amount * price
		w.Write([]byte(fmt.Sprintf("%v\t%v\t%v\t%v\n", bid.Price, amount, sum, msum)))
	}
	w.Flush()

	fmt.Printf("\nSELL %v for %v:\n", *sellingAssetCode, *buyingAssetCode)

	sum = 0.0
	msum = 0.0
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	w.Write([]byte(fmt.Sprintf("PRICE (%v/%v)\tAMOUNT\tSUM\tMULTIPLIED SUM\n", *buyingAssetCode, *sellingAssetCode)))
	for _, bid := range resp.Asks {
		amount, err := strconv.ParseFloat(bid.Amount, 64)
		if err != nil {
			logrus.Fatal(err)
		}
		price, err := strconv.ParseFloat(bid.Price, 64)
		if err != nil {
			logrus.Fatal(err)
		}
		sum += amount
		msum += amount * price
		w.Write([]byte(fmt.Sprintf("%v\t%v\t%v\t%v\n", bid.Price, amount, sum, msum)))
	}
	w.Flush()

}
