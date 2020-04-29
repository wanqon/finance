package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

const (
	//SourceBase = "/data0/paydata/src_data/db/"
	//TargetBase = "/data0/paydata/finance/"
	SourceBase = "/Users/wangqiong1/app/data1/paydata/src_data/db/"
	TargetBase = "/Users/wangqiong1/app/data1/paydata/finance/"
	TIME_LAYIN  = "2006-01-02"
	TIME_LAYOUT = "2006-01-02 15:04:05"
)

type Bill struct {
	biz	string
}

func tryInit() {
	loger := cron.PrintfLogger(log.New(os.Stdout,"cron: ", log.LstdFlags))
	c := cron.New(cron.WithSeconds(),cron.WithChain(cron.SkipIfStillRunning(loger)))

	//job := cron.NewChain(cron.SkipIfStillRunning(loger)).Then(reconciliation.New("WbpayCharge"))
	//c.AddJob("*/5 * * * * *", job)

	p := 0

	c.AddFunc("*/10 * * * * *", func() {
		p++
		fmt.Printf("5 second start:%d\n", p)
		for i:=0; i<8; i++ {
			time.Sleep(time.Second)
			fmt.Printf("5 second sleep:%d\n", i)
		}
		fmt.Printf("5 second end:%d\n", p)

	})
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("10 second once")

	})
	c.Start()
	signalHandle(c)
}

func signalHandle(c *cron.Cron)  {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v", sig)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("graceful shutdown")
			return
		case syscall.SIGUSR2:
			ctx := c.Stop()
			select {
			case <-ctx.Done():
				fmt.Println("context was done too quickly immediately")
				//case <-time.After(750 * time.Millisecond):
				// expected, because the job sleeping for 1 second is still running
			}
			log.Printf("graceful reload")
			return
		}
	}
}

func New(biz string) *Bill {
	return &Bill{
		biz:biz,
	}
}

func (bill *Bill) Run()  {
	reflect.ValueOf(bill).MethodByName(bill.biz).Call(nil)
}
