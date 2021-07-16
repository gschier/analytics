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

	websiteID := ensureDummyWebsite()
	fmt.Println("[main] Website", websiteID)

	h := SetupRouter()
	fmt.Printf("[schier.co] \033[32;1mStarted server on http://%s:%s\033[0m\n", Config.Host, Config.Port)
	log.Fatal(http.ListenAndServe(Config.Host+":"+Config.Port, h))
}

func ensureDummyWebsite() string {
	account, accountExists := GetAccountByEmail(GetDB(), context.Background(), "greg@schier.co")

	if accountExists {
		websites := FindWebsitesByAccountID(GetDB(), context.Background(), account.ID)
		return websites[0].ID
	}

	a := CreateAccount(GetDB(), context.Background(), "greg@schier.co", "my-pass!")
	w := CreateWebsite(GetDB(), context.Background(), a.ID, "My Blog")

	return w.ID
}
