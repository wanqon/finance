package main

import (
	"finance/logger"
	"finance/reconciliation"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	if reconciliation.Uri == "charge" {
		reconciliation.WbpayCharge(reconciliation.Date)
		logger.Info("charge done",zap.String("uri",reconciliation.Uri),zap.String("date", reconciliation.Date))
	}

	logger.Info("test", logger.LogField{"a":6, "b":"c"})
	return


	logger.Info("Failed to ferch URL",zap.String("url", "a/b/c"), zap.Int("attempt", 10),zap.Duration("backoff",3))
	logger.Error("Failed to ferch URL",zap.String("url", "a/b/c"), zap.Int("attempt", 10),zap.Duration("backoff",3))
	logger.Info("",zap.String("url", "a/b/c"), zap.Int("attempt", 10),zap.Duration("backoff",3))

	return
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
