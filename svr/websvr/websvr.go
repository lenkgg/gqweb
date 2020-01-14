package main

import (
	"github.com/kataras/iris"
	"gqyb/account/client"
	"gqyb/comm"
	"gqyb/comm/errors"
	"gqyb/comm/session"
	hg "gqyb/hedge/client"
	mkt "gqyb/market/client"
)


func verifySkey(ctx iris.Context){
	//comm.Log.Debugf("verify gqkey.:%v",ctx.Request())
	//pc := ctx.Request().Cookies()
	//comm.Log.Debugf("cookie len=%d,%v",len(pc),pc)
	gqkey := ctx.GetCookie("gqkey")
	uid := ctx.GetCookie("uid")
	comm.Log.Debugf("uid=%s,gqkey=%s",uid,gqkey)
	if gqkey == "" || uid == ""{
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	ok := skey_cache.Verify(uid,gqkey)
	if !ok{
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	ctx.Next()
}

var skey_cache  *session.KeyCache
func main() {
	defer Errors.TryError()
	comm.ConfigLocalFilesystemLogger("websvr.log")
	comm.Log.Info("server started.")
	app := iris.Default()
	skey_cache = session.InitSessionCache()
	ret := acc.CreateAccountClient()
	//defer acc.CloseAccountClient()
	Errors.CheckError(ret)

	ret = hg.CreateHedgeClient()
	Errors.CheckError(ret)

	ret = mkt.CreateMarketClient()
	Errors.CheckError(ret)
	/*
	ret = sns.CreateSnsClient()
	Errors.CheckError(ret)
	*/
	//hg.CreateBridgeClient()
	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/index",index)
	app.Handle("GET", "/notice",Notice)
	app.Handle("POST","/login", Login)
	app.Handle("POST","/regist",Regist)
	app.Handle("GET","/bphone",verifySkey, BindPhone)
	app.Handle("POST","/updateinfo", verifySkey, UpdateInfo)
	///////////////hedge server
	app.Handle("GET","/predict",verifySkey, Predict)
	app.Handle("GET","/hedge",verifySkey, Hedge)
	app.Handle("GET","/getslot",verifySkey, GetSlot)
	app.Handle("GET","/setsymbol",verifySkey, SetSymbol)
	app.Handle("GET","/getstocks",verifySkey, GetStocks)
	app.Handle("GET","/addstock",verifySkey, AddStock)
	app.Handle("GET","/delstock",verifySkey, DelStock)
	//app.Handle("GET","/getslot",verifySkey, GetSlot)
	//app.Handle("GET","/openslot",verifySkey, OpenSlot
	app.Handle("GET","/setslot",verifySkey, SetSlot)
	app.Handle("GET","/loadcncode", GetCnsymbol)

	//app.Handle("GET","/payslot",verifySkey, PaySlot)
	app.Handle("GET","/listshare",verifySkey, ListUsrShare)
	app.Handle("GET","/prebill", verifySkey, PreBill)

	//////////////for operation/////////
	app.Handle("GET","/op/addslot", verifyop, OpAddSlot)
	app.Handle("GET","/vtrade_pos", VtradePosition)
	app.Handle("GET","/vtrade_hedge", VtradeHedge)

	////////////for api //////////////
	app.Handle("POST","/api/login", OrgLogin)
	app.Handle("POST","/api/regist",OrgRegist)
	app.Handle("GET","/api/predict",verifySkey, Predict)

	///////for sns////////////////////////
	/*
	app.Handle("GET","/sns/CreateGroup",verifySkey, CreateGroup)
	app.Handle("GET","/sns/GetGroups",verifySkey, GetGroups)
	app.Handle("GET","/sns/ApplyGroup",verifySkey, ApplyGroup)
	app.Handle("GET","/sns/QuitGroup",verifySkey, QuitGroup)
	app.Handle("GET","/sns/SetGroup",verifySkey, SetGroup)
	app.Handle("GET","/sns/SetManager",verifySkey, SetManager)
	app.Handle("GET","/sns/ListApply",verifySkey, ListApply)
	app.Handle("GET","/sns/AcceptMember",verifySkey, AcceptMember)
	app.Handle("GET","/sns/CreateTopic",verifySkey, CreateTopic)
	app.Handle("GET","/sns/SetTopic",verifySkey, SetTopic)
	app.Handle("POST","/sns/PubPost",verifySkey, PubPost)
	app.Handle("GET","/sns/ListPost",verifySkey, ListPost)
	app.Handle("GET","/sns/GroupsPost",verifySkey, GroupsPost)
	app.Handle("GET","/sns/LikePost",verifySkey, LikePost)
	app.Handle("GET","/sns/DeletePost",verifySkey, DeletePost)
	app.Handle("POST","/sns/ReplyPost",verifySkey, ReplyPost)
	app.Handle("GET","/sns/DeleteReply",verifySkey, DeleteReply)
	app.Handle("POST","/upload",verifySkey,upload)
	*/

	////////////////////for test///////////
	//app.Handle("GET","/test_py_predict", Test_py_predict)

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	//app.Get("/ping", func(ctx iris.Context) {
	// ctx.WriteString("pong")
	//})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	//app.Get("/index", func(ctx iris.Context) {
	// ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	//})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	//app.Run(iris.Addr(":8080"))
	//app.Run(iris.TLS(":8443", "./ssl/2034837_www.lilishare.com.pem", "./ssl/2034837_www.lilishare.com.key"))
	app.Run(iris.TLS(":8443", "./ssl/lilishare.com.pem", "./ssl/lilishare.com.key"))
}
