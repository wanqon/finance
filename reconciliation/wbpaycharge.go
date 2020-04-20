package reconciliation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	//充值状态_关闭
	CHARGE_STATUS_CLOSED = 0
	//充值状态_新创建
	CHARGE_STATUS_CREATE = 1
	//充值状态_充值成功
	CHARGE_STATUS_CHARGED = 2
)

type ChargeBill struct {
	date	string
	StartTime	time.Time
	EndTime		time.Time

}

type HandleInfo func(charge []string, c chan string)

func (bill *ChargeBill) Run() {
	l := time.FixedZone("CST", 8*3600)
	var t time.Time
	if len(bill.date) == 0 {
		t = time.Now().AddDate(0,0,-1).In(l)
	} else {
		t, _ = time.Parse(TIME_LAYIN,bill.date)
	}
	year,month,day := t.Date()
	targetDir := fmt.Sprintf(TargetBase+"charge/%d%02d/%02d/", year, int(month), day)

	if _, err := os.Stat(targetDir);os.IsNotExist(err) {
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	bill.StartTime = time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,l)
	bill.EndTime = time.Date(t.Year(),t.Month(),t.Day(),23,59,59,0,l)

	//startTime := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,l)
	//endTime := time.Date(t.Year(),t.Month(),t.Day(),23,59,59,0,l)

	dbDate := time.Now().In(l).Format("20060102")
	sourceDir := fmt.Sprintf(SourceBase+"charge/%s/",dbDate)

	c := make(chan string)
	var swg *sync.WaitGroup
	swg.Add(1)
	wFile := targetDir+"charge.txt"
	WriteFile(swg,wFile,c)
	for i:=0; i<128; i++ {
		swg.Add(1)
		fileName:=fmt.Sprintf(sourceDir+"snap_%d.txt", i)
		go ReadFile(swg, fileName, c, bill.handleInfo)
	}
	swg.Wait()
}

func (bill *ChargeBill) handleInfo(charge []string, c chan string) {
	chargeTime, _ := time.Parse(TIME_LAYOUT, charge[9])
	status,_ := strconv.Atoi(charge[7])
	if TimeCompare(chargeTime,bill.StartTime,bill.EndTime) && status == CHARGE_STATUS_CHARGED {
		info := strings.Join(charge[:10],"\t")
		info += "\r\n";
		c <- info
	}
}

func TimeCompare(check time.Time, start time.Time, end time.Time) bool {
	return check.Second() >= start.Second() && check.Second() <= end.Second()
}


func ReadFile(swg *sync.WaitGroup, path string, c chan string, handlefuc HandleInfo)  {
	defer swg.Done()
	fr, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer fr.Close()
	reader := bufio.NewReader(fr)
	for {
		line,_,err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		chargeInfo := strings.Split("\t", string(line))
		handlefuc(chargeInfo, c)
	}
}

func WriteFile(swg *sync.WaitGroup, path string, c chan string) {
	defer swg.Done()
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND,os.ModePerm)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for info := range c {
		_, err := f.Write([]byte(info))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}
