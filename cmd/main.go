package main

import (
	"fmt"
	"github.com/kekzploit/marketplace-tx-watcher/pkg/db"
	"github.com/kekzploit/marketplace-tx-watcher/pkg/tx"

	"github.com/spf13/viper"
)

func main() {
	// initiate environment variable config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./../configs/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	txExists, image, title, description, secret, storeType, hash := tx.TxWatch()
	//txExists, _, _, _, _, _, hash := tx.TxWatch()

	if txExists {
		vendorExists := db.CheckDB(viper.Get("MONGO.URI").(string), hash)
		if !vendorExists {
			vendorAdded := db.AddVendor(viper.Get("MONGO.URI").(string), image, title, description, secret, storeType, hash)
			if !vendorAdded {
				fmt.Println("error adding vendor")
			} else {
				fmt.Println("added new vendor")
			}
		}
	}
}
