package main

import (
	"fmt"
	"stockpredictionscronjob/service"

	"github.com/robfig/cron/v3"
)

func main() {
	fmt.Println("Automation data handler is running...")
	c := cron.New(cron.WithSeconds())

	service.BootstrapFirstHistories()
	fmt.Println("Bootstrap finished")

	c.AddFunc("0 20 10 * * *", service.SaveLastStockPrices)
	c.AddFunc("0 * 11-18 * * *", service.UpdateLastStockPrices)

	c.AddFunc("0 20 18 * * *", service.UpdatePrdictionLog)
	c.AddFunc("0 30 18 * * *", service.MakePredictions)

	fmt.Println(c.Entries())

	c.Run()
}
