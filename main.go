package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println()
	fmt.Printf("\u001B[32;1m┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┃                  analytics                  ┃\u001B[0m\n")
	fmt.Printf("\u001B[32;1m┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\u001B[0m\n")

	InitConfig()

	if Config.MigrateOnStart {
		mustMigrate(context.Background(), GetDB())
	}

	h := SetupRouter(ensureDummyWebsite())
	fmt.Printf("[schier.co] \033[32;1mStarted server on http://%s:%s\033[0m", Config.Host, Config.Port)
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, h))
}

func ensureDummyWebsite() string {
	account, accountExists := GetAccountByEmail(GetDB(), context.Background(), "greg@schier.co")

	globalWebsiteID := ""
	if accountExists {
		websites := FindWebsitesByAccountID(GetDB(), context.Background(), account.ID)
		globalWebsiteID = websites[0].ID
	} else {
		a := CreateAccount(GetDB(), context.Background(), "greg@schier.co", "my-pass!")
		w := CreateWebsite(GetDB(), context.Background(), a.ID, "My Blog")
		globalWebsiteID = w.ID
	}

	fmt.Println("[main] Website", globalWebsiteID)
	return globalWebsiteID
}
