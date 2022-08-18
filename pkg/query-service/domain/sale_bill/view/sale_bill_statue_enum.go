package view

type SaleBillStatue int

const (
	SaleBillStatueConfirm  SaleBillStatue = 2
	SaleBillStatueCreating SaleBillStatue = 1
	SaleBillStatueReceive  SaleBillStatue = 3
)

func (e SaleBillStatue) String() string {
	res := "UNKNOWN"
	switch e {
	case SaleBillStatueConfirm:
		res = "Confirm"
	case SaleBillStatueCreating:
		res = "Creating"
	case SaleBillStatueReceive:
		res = "Receive"
	default:
		res = "UNKNOWN"
	}
	return res
}

func (e SaleBillStatue) Title() string {
	res := "UNKNOWN"
	switch e {
	case SaleBillStatueConfirm:
		res = "确认下单"
	case SaleBillStatueCreating:
		res = "创建中"
	case SaleBillStatueReceive:
		res = "已收货"
	default:
		return "UNKNOWN"
	}
	return res
}
