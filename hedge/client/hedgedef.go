package hedge

const(
	M_UNKNOWN     = "UNK";
	M_STOCK_CN    = "STK_CN";       //include sh & sz
	M_STOCK_HK    = "STK_HK";
	M_STOCK_US    = "STK_US";
	M_FUTURE_SGP  = "FT_SGP";
	////////////future///////////
	M_FUTURE_CN_SH   = "FT_CN_SH";
	M_FUTURE_US_CMX  = "FT_US_CMX";

	/////////sub marcket///////////
	M_STOCK_CN_SH = "STK_CN_SH";
	M_STOCK_CN_SZ = "STK_CN_SZ";
)

////hedge scope
const(
	SCP_STK_HS300     = "CN_HS300";
	SCP_STK_SZ50      = "CN_SZ50";
	SCP_STK_CN_ALL    = "CN_ALL";
)

const(
	KL_DAY          = "DAY";
	KL_WEEK         = "WEK";
	KL_MONTH        = "MON";
	KL_YEAR         = "YER";
	KL_5M           = "5M";
	KL_15M          = "15M";
	KL_30M          = "30M";
	KL_HOUR         = "HOR";
	KL_2HOUR        = "2HR";
	KL_CUSTOM       = "UNK";
)

///    `bson:"" json:""`
/*
type F_OBJECT struct{
	Symbol string     		`bson:"symbol" json:"symbol"`
	Mkt string      		`bson:"mkt" json:"mkt"`
	Sub_mkt string     		`bson:"sub_mkt" json:"sub_mkt"`
	Name string        		`bson:"name" json:"name"`
	Industry string     	`bson:"indu" json:"indu"`
	Expire comm.Time           `bson:"expire" json:"expire"`    //future 交割日
	Desc string         	`bson:"desc" json:"desc"`
}
type KL_BAR struct{
	Dtime comm.Time     		`bson:"dtime" json:"dtime"`
	Open float32          	`bson:"o" json:"o"`
	High float32          	`bson:"h" json:"h"`
	Low float32            	`bson:"l" json:"l"`
	Close float32         	`bson:"c" json:"c"`
	Volume uint32         	`bson:"v" json:"v"`
	Rchange float32         `bson:"r" json:"r"`
}

type KLINE struct{
	Symbol string  			`bson:"symbol" json:"symbol"`
	Dtime  comm.Time			`bson:"dtime" json:"dtime"`
	K_type string         	`bson:"ktype" json:"ktype"`
	Bars []KL_BAR    		`bson:"kline" json:"kline"`
}
*/
type Slot struct{
	Sid uint32				`bson:"sid" json:"sid"`
	Name string				`bson:"name" json:"name"`
	Symbol string			`bson:"symbol" json:"symbol"`
	Opentime string		`bson:"opentime" json:"opentime"`
	Expire string		`bson:"expire" json:"expire"`
	Desc string				`bson:"desc" json:"desc"`
}


type Myslots struct {
	Uid uint64					`bson:"uid" json:"uid"`
	Slots []Slot				`bson:"slots" json:"slots"`
}

///for non-slot version
type Mystocks struct {
	Uid uint64					`bson:"uid" json:"uid"`
	Stocks []string				`bson:"stocks" json:"stocks"`
}

type PredictResult struct {
	Symbol string 			`bson:"symbol" json:"symbol"`
	Kline []string			`bson:"kline" json:"kline"`
	Ptime string				`bson:"ptime json:"ptime"`
}

////for front decode
type HedgeResult struct{
	Bsymbol  string				`bson:"bsymbol" json:"bsymbol"`
	Scope   string 				`bson:"scope" json:"scope"`
	Rank  []string				`bson:"rank" json:"rank"`
	Htime string				`bson:"htime" json:"htime"`
}

type BRG_CONF struct {
	Aisvr string            `json:"aisvr"`
	Pre_cache string        `json:"predict"`
	Hdg_cache string        `json:"hedge"`
	Timer_path string		`json:"tm_path"`
	Cn_index string         `json:"default_cn_index"`
	Cn_hedge string 		`json:"default_cn_hedge"`
	Scope string			`json:"scope"`
}

var DEFAULT_CN_INDEX = []string{"sh000001", "sh000016", "sz399001", "sz399006", "sz399300"}
var PREDICT_PATH = "../data/predict/"

//所有的stocks列表存磁盘文件，启动时加载
//所有的predict和hedge结果数据存磁盘文件，一天一个文件，svr周期性刷新加载

