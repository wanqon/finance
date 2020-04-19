package reconciliation

import "reflect"

const (
	SourceDir = "/data0/paydata/src_data/db/"
	TargetDir = "/data0/paydata/finance/"
)

type recon func()

type Bill struct {
	recon recon
	biz	string
}

func New(biz string) *Bill {
	return &Bill{
		biz:biz,
	}
}

func (bill *Bill) Run()  {
	//bill.recon()
	reflect.ValueOf(bill).MethodByName(bill.biz).Call(nil)
}

func GetZeroTime()  {
	
}