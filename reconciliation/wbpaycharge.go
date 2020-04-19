package reconciliation

import (
	"fmt"
	"time"
)

type ChargeBill struct {
	Bill
}

func (bill *ChargeBill) Run(date time.Time) {

	year,month,day := time.Now().In(time.FixedZone("CST", 8*3600)).AddDate(0,0,-1).Date()
	dir := fmt.Sprintf(SourceDir+"charge/%d/%02d/%02d/", year, int(month), day)
	fmt.Println(dir)
}

func (bill *Bill) WbpayCharge()  {
	year,month,day := time.Now().AddDate(0,0,-1).Date()
	dir := fmt.Sprintf(SourceDir+"charge/%d/%02d/%02d/", year, int(month), day)

	fmt.Println(dir)
	//today := time.Now().Format("2006-01-02")
}
