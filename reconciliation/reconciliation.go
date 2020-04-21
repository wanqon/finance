package reconciliation

import "reflect"

const (
	SourceBase = "/data0/paydata/src_data/db/"
	TargetBase = "/data0/paydata/finance/"
	TIME_LAYIN  = "2006-01-02"
	TIME_LAYOUT = "2006-01-02 15:04:05"
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