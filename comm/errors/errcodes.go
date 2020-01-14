package Errors

var E_SUCCESS = NewCoder(0,"success.")
var E_SYSTEM 	=	NewCoder(1,"system error.")
var E_PARAM		=	NewCoder(2,"request parameter error.")
var E_CONFIG		=	NewCoder(3,"load config file failed.")
var E_VERIFY		=	NewCoder(4,"verify data failed.")
var E_DATABASE		= 	NewCoder(5,"database error.")
var E_DATA			= 	NewCoder(6, "data error.")
var E_NO_LOGIN		=  	NewCoder(7,"user not login.")
var E_NO_DATA		=   NewCoder(8, "no record in database.")
//////////////for account: 1000~1999//////////
var E_ACC_WX_LOGIN_FAILED		=	NewCoder(1001,"wechat login failed.")
var E_ACC_USER_NOT_EXIST		=	NewCoder(1002,"user not exist.")
var E_ACC_DATABASE		=	NewCoder(1003,"database connect failed.")
var E_ACC_VC		=	NewCoder(1004,"verify code error.")
var E_ACC_PASSWORD		=	NewCoder(1005,"password error.")
var E_ACC_USER_EXIST		=	NewCoder(1006,"user existed.")

///////////for hedge: 2000~2999///////////////
var E_HDG_DECODE			= NewCoder(2001, "json decode failed.")
var E_HDG_DB				= NewCoder(2002, "hedge db error.")
var E_HDG_NO_DATA			= NewCoder(2003, "no predict data.")
var E_HDG_NO_SLOT			= NewCoder(2004, "no more slot.")
var E_MORE_SLOT_LIMIT		= NewCoder(2005, "slots more than limit.")
var E_WAIT_PREDICT			= NewCoder(2006, "wait for a moment,ai is working hard.")
var E_MAX_STOCK				= NewCoder(2007, "max stocks count limited.")

////////for market:3000~3999//////////
var E_MKT_NO_ACT				= NewCoder(3001,"action id not exist.")


///////for sns:4000~4999////////////
var E_SNS_DB				= NewCoder(4001,"sns db error.")
var E_GROUP_LIMITED			= NewCoder(4002,"maximum group number limit.")
var E_GROUP_NOTEXIST		= NewCoder(4003,"group not exist.")
var E_SNS_PERMISSION		= NewCoder(4004,"permission limited.")
var E_GROUP_MAXMANAGER		= NewCoder(4005,"maxinum manager number is 5.")
var E_TOPIC_EXIST			= NewCoder(4006,"the topic has existed.")
var E_POST_NOTOPC			= NewCoder(4007,"not exist this topic.")
var E_POST_NODATA			= NewCoder(4008,"no data.")
var E_POST_MANYPIC			= NewCoder(4009,"too many attachers.")

////////////////////for data center /////////////
var E_HTTP_READY			= NewCoder(5001,"create http request error.")
var E_HTTP_CONNECT			= NewCoder(5002,"http connect failed.")
var E_HTTP_RESPONSE			= NewCoder(5003,"http response error.")
var E_SYMBOL_ERROR			= NewCoder(5004,"symbol error.")
var E_DC_REMOVE				= NewCoder(5005,"remove records failed.")
var E_DC_UPSERT				= NewCoder(5006,"upsert data failed.")
var E_DC_FIND				= NewCoder(5007,"mongodb find failed.")
var E_DC_INSERT				= NewCoder(5008,"mongodb insert failed.")
var E_DC_LESS_DATA			= NewCoder(5009,"less kline data.")
var E_DC_NO_DATA			= NewCoder(5010,"no kline data.")