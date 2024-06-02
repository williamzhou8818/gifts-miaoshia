package database

type Order struct {
	Id     int
	GiftId int
	UserId int
}

// 清除全部订单记录
func ClearOrders() error {
	db := GetGiftDBConnection()
	return db.Where("id>0").Delete(Order{}).Error
}
