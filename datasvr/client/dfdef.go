package dc

type DFitem struct{
	Dtime   string   	`bson:"dtime" json:"dtime"`
	Open    float64		`bson:"open" json:"open"`
	High    float64		`bson:"high" json:"high"`
	Low     float64		`bson:"low" json:"low"`
	Close   float64		`bson:"close" json:"close"`
	Volume  float64     `bson:"volume" json:"volume"`
	Date    int64		`bson:"date" json:"date"`
	DateWeek  int32		`bson:"date_week" json:"date_week"`
}

type DFPredict struct{
	Symbol string  		`bson:"symbol" json:"symbol"`
	Kline  []string		`bson:"kline" json:"kline"`
	Ptime  string		`bson:"ptime" json:"ptime"`
}

type JQSHAREitem struct{
	Symbol string		`bson:"symbol" json:"symbol"`
	ShareRmb float32	`bson:"share_rmb" json:"share_rmb"`
	ChangeDate string   `bson:"change_date" json:"change_date"`
}


/*
python define:
E_KLINE_DAY = 'DAY'
    E_KLINE_WEEK = 'WEK'
    E_KLINE_MONTH = 'MON'

    E_KLINE_MINUTE = 'MIN'
    E_KLINE_5MIN = '5MIN'
    E_KLINE_15MIN = '15MIN'
    E_KLINE_HOUR = 'HOR'
    E_KLINE_2HOUR = '2HOR'

    E_KLINE_TICK = 'TIK'
 */
const(
	BAR_TYPE_WEEK = "WEK"
	BAR_TYPE_MONNTH = "MON"
  	BAR_TYPE_DAY = "DAY"
	BAR_TYPE_HOUR = "HOR"
	BAR_TYPE_MIN = "MIN"
	BAR_TYPE_5MIN = "5MIN"
	BAR_TYPE_15MIN = "15MIN"
	BAR_TYPE_2HOUR = "2HOR"
	BAR_TYPE_TICK = "TIK"
)

var HS_INDEX_SYMBOL = []string{"sh000001","sz399300","sh000016","sz399001","sz399006"}

const(
	MIN_DATA_LENGTH = 130
	PREDICT_STEP = 10
	MODEL_PATH_STOCK = "../model/stockmodel_B1024_E260/stockmodel_B1024_E260.pb"
	MODEL_PATH_INDEX = "../model/indexmodel_B500_E260/indexmodel_B500_E260.pb"
)

const(
	TASKID_CN_STOCK_FEATURE = 1000
)