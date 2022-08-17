package field

type SaleBillStatue int

const (
	SaleBillStatueCreated  SaleBillStatue = 2 // 下单完成
	SaleBillStatueCreating SaleBillStatue = 1 // 创建中
	SaleBillStatueReceive  SaleBillStatue = 3 // 已收货
)

func (e SaleBillStatue) String() string {
	res := "UNKNOWN"
	switch e {
	case SaleBillStatueCreated:
		res = "Created"
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
	case SaleBillStatueCreated:
		res = "下单完成"
	case SaleBillStatueCreating:
		res = "创建中"
	case SaleBillStatueReceive:
		res = "已收货"
	default:
		return "UNKNOWN"
	}
	return res
}
