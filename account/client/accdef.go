package acc

type USER struct {
	Uid  uint64			`bson:"uid" json:"uid"`
	Name string			`bson:"name" json:"name"`
	Nick string			`bson:"nick" json:"nick"`
	Age int32			`bson:"age" json:"age"`
	Gender int32		`bson:"gender" json:"gender"`
	Country string		`bson:"country" json:"country"`
	Province string		`bson:"province" json:"province"`
	City string			`bson:"city" json:"city"`
	Desc string			`bson:"desc" json:"desc"`
	Other string		`bson:"other" json:"other"`
	Phone string		`bson:"phone" json:"phone"`
	Wxid string			`bson:"wxid" json:"wxid"`    ///openid
	Email string		`bson:"email" json:"email"`
	Avata string		`bson:"avata" json:"avata"`
	Language string		`bson:"lang" json:"lang"`
}
type UserRec struct{
	Regtime string			`bson:"regtime"`
	Info USER        		`bson:"info"`
	Pswd string 			`bson:"pswd"`
	Lastlogin string		`bson:"lastlogin"`
	Channel  uint64			`bson:"channel"`

}

type ORG struct{
	Oid uint64 			`bson:"oid" json:"oid"`
	Name string 		`bson:"name" json:"name"`
	Country string		`bson:"country" json:"country"`
	Province string		`bson:"province" json:"province"`
	City string			`bson:"city" json:"city"`
	Desc string			`bson:"desc" json:"desc"`
	OrgUid	string		`bson:"orguid" json:"orguid"`  /// 工商三证合一的代码
	Other string		`bson:"other" json:"other"`
	Phone string		`bson:"phone" json:"phone"`
	Email string		`bson:"email" json:"email"`
	Logo string			`bson:"logo" json:"logo"`
	Language string		`bson:"lang" json:"lang"`
}

type OrgRec struct{
	Regtime string			`bson:"regtime"`
	Info ORG        		`bson:"info"`
	Pswd string 			`bson:"pswd"`
	Lastlogin string		`bson:"lastlogin"`
}


