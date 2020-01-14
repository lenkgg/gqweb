package market

//////////营销推广记录表 Market，存储最近一年的数据////////////
/////////渠道商统一分配渠道号 9000000001 开始
const CHANNEL_ORG  = 9000000001
const PA_REG = "user regist."
const PA_OPEN_SLOT = "buy slot."

type PayBill struct{
	Billno string			`bson:"billno" json:"billno"`
	Dtime string			`bson:"dtime" json:"dtime"`
	Amount float32			`bson:"amount" json:"amount"`
	Paychn string			`bson:"paychn" json:"paychn"`
	Paybill string			`bson:"paybill" json:"paybill"`
	Mark string				`bson:"mark" json:"mark"`
	Operator uint64			`bson:"op" json:"op"`
	Checker uint64			`bson:"checker" json:"checker"`
	Items []string			`bson:"items" json:"items"`
}

type MktActItem struct{
	Id    string			`bson:"id" json:"id"`
	Dtime string			`bson:"dtime" json:"dtime"`
	Uin uint64				`bson:"uin" json:"uin"`
	Action string			`bson:"action" json:"action"`
	Paybill  string			`bson:"paybill" json:"paybill"`
}


type MarketUsr struct{
	Uid uint64					`bson:"uid" json:"uid"`   			///uid
	Follow		[]uint64 		`bson:"follow" json:"follow"`  		///所有下线用户列表
	History		[]MktActItem	`bson:"history" json:"history"`			///普通用户赠送槽位时效，即时赠送
}

type MarketOrg struct{
	Orgid 		uint64			`bson:"orgid" json:"orgid"`   			///channelID
	Follow		[]uint64 		`bson:"follow" json:"follow"`  		///所有下线用户列表
	Detail		[]MktActItem	`bson:"detail" json:"detail"`			///未结算，机构涉及现金，月结
	History  	[]MktActItem	`bson:"history" json:"history"` 	///已结算
}

type PreBillRec struct{
	Uid uint64 					`bson:"uid" json:"uid"`
	Months uint32				`bson:"months" json:"months"`
	Slots uint32 				`bson:"slots" json:"slots"`
	Phone string				`bson:"phone" json:"phone"`
	PreTime string				`bson:"pretime" json:"pretime"`
	Status  int32				`bson:"status" json:"status"`
	Operator string				`bson:"oper" json:"oper"`
	OpTime  string				`bson:"optime" json:"optime"`
	Price float32				`bson:"price" json:"price"`
}