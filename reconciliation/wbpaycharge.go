package reconciliation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type ChargeBill struct {
	date	string
}

type HandleInfo()

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

	//startTime := time.Date(t.Year(),t.Month(),t.Day(),0,0,0,0,l)
	//endTime := time.Date(t.Year(),t.Month(),t.Day(),23,59,59,0,l)

	dbDate := time.Now().In(l).Format("20060102")
	sourceDir := fmt.Sprintf(SourceBase+"charge/%s/",dbDate)

	for i:=0; i<128; i++ {
		fileName:=fmt.Sprintf(sourceDir+"snap_%d.txt", i)
		go ReadFile(fileName)
		go func(fileName string) {
			fmt.Println("aaa")
		}()

	}

	fmt.Println(sourceDir)

}

func (bill *ChargeBill) handleInfo(charge []string) {
	if charge[9]> {

	}
}

func compareDate(date1 string, date2 string) bool {
	d1, err := time.Parse(TIME_LAYOUT, date1)
	d2, err := time.Parse(TIME_LAYOUT, date2)
	return d1.Second() >= d2.Second()
}


func ReadFile(file string)  {
	fr, err := os.OpenFile(file, os.O_RDONLY, 0600)
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

		fmt.Println(string(line))
	}

}
