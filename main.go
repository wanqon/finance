package main

import (
	"finance/reconciliation"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

func main() {

	a := reconciliation.NewCharge("")
	a.Run()

	return

	loger := cron.PrintfLogger(log.New(os.Stdout,"cron: ", log.LstdFlags))
	c := cron.New(cron.WithSeconds())

	job := cron.NewChain(cron.SkipIfStillRunning(loger)).Then(reconciliation.New("WbpayCharge"))
	c.AddJob("*/5 * * * * *", job)

	//c.AddFunc("*/5 * * * * *", func() {
	//	p++
	//	fmt.Printf("5 seconde start:%d\n", p)
	//	time.Sleep(7*time.Second)
	//	fmt.Printf("5 seconde end:%d\n", p)
	//})
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("10 seconde once")
	})
	c.Start()

	select {

	}
}
