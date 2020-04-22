package reconciliation

import (
	"reflect"
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

func NewCharge(date string) *ChargeBill {
	return &ChargeBill{
		date:      date,
		StartTime: time.Time{},
		EndTime:   time.Time{},
		Fw:        nil,
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