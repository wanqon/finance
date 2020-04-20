package main

import (
	"finance/reconciliation"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)


const TIME_LAYIN  = "2006-01-02"
const TIME_LAYOUT = "2006-01-02 15:04:05"

func main() {

	time.FixedZone("CST", 8*3600)
	l,_ := time.LoadLocation("America/New_York")
	//time.LoadLocation("China/Beijing")
	t, _ := time.Parse(TIME_LAYIN,"2020-04-05")
	t1 := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,l)
	fmt.Println(t1.Format(TIME_LAYOUT))
	lcst := time.FixedZone("CST", 8*3600)
	t2 := time.Date(t.Year(),t.Month(),t.Day(),23,59,59,0,lcst)
	fmt.Println(t2.Format(TIME_LAYOUT))
	//return

	fmt.Println(time.Now().In(l))
	fmt.Println(time.Now().In(time.FixedZone("CST", 8*3600)))

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
	//c.Entries()
	select {

	}
}
