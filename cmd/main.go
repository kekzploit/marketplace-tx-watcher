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

	//hello := viper.Get("VALUE.HELLO").(string)

	txExists, image, title, description, secret, storeType, hash := tx.TxWatch()

	if txExists {
		db.CheckDB()
		store := fmt.Sprintf("\nImage: %s\nTitle: %s\nDescription: %s\nSecret: %s\nType: %s\nHash: %s\n", image, title, description, secret, storeType, hash)
		fmt.Println(store)
	} else {
		fmt.Println("no new store registrations")
	}
}
