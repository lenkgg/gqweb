package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"gqyb/comm"
	"gqyb/comm/errors"
	hg "gqyb/hedge/client"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Predict(ctx iris.Context){
	///symbol = sh000001,sz399001,...
	str := ctx.URLParam("symbol")
	str = strings.ToLower(str) ///后台symbol全部小写
	symbols := strings.Split(str,",")
	for _,symbol := range symbols{
		symbol = CheckSymbol(symbol) ///后台symbol全部小写
		if symbol == ""{
			Reply(ctx,Errors.E_PARAM,nil)
			return
		}
	}
	ret,err := hg.Predict(symbols)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	rlt := &[]hg.PredictResult{}
	json.Unmarshal(ret,rlt)
	Reply(ctx,Errors.E_SUCCESS,rlt)
}

/*
////hedge scope
const(
	SCP_STK_HS300     = "CN_HS300";
	SCP_STK_SZ50      = "CN_SZ50";
	SCP_STK_CN_ALL    = "CN_ALL";
)
 */
func Hedge(ctx iris.Context){
	////this version only surport const scope+bsymbol
	scope:= ctx.URLParam("scope");
	bsymbol := ctx.URLParam("bsymbol")
	bsymbol = CheckSymbol(bsymbol) ///后台symbol全部小写
	if bsymbol == ""{
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	comm.Log.Debugf("hedge request: scope=%s, bsymbol=%s",scope,bsymbol)
	ret,err := hg.Hedge(scope,bsymbol)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	//comm.Log.Debug(string(ret))
	rlt := hg.HedgeResult{}
	er := json.Unmarshal(ret,&rlt)
	if !Errors.CheckError(er){
		Reply(ctx,er,nil)
		return
	}
	Reply(ctx,nil,rlt)
}

func index(ctx iris.Context){
	rsp,err := hg.TimerPredict()
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	rlt := []hg.PredictResult{}
	//comm.Log.Debugf("rsp:%v",rsp)
	er := json.Unmarshal(rsp,&rlt)
	//comm.Log.Debugf("unmarshal to:%v",rlt)
	if !Errors.CheckError(er){
		comm.Log.Errorf("[index] unmarshal json failed,error:%v", er)
		Reply(ctx,er,nil)
		return;
	}
	Reply(ctx,Errors.E_SUCCESS,rlt)
}

func GetSlot(ctx iris.Context){
	comm.Log.Debug("rcv getslot request.")
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	slt,er := hg.GetMyslot(uint64(uid))
	if !Errors.CheckError(er){Reply(ctx,er,nil);return}

	ReplyString(ctx,Errors.E_SUCCESS,string(slt))
}


func SetSymbol(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	ssid := ctx.URLParam("sid")
	symbol := ctx.URLParam("symbol")

	///symbol == "" when clear slot
	if symbol != ""{
		symbol = CheckSymbol(symbol) ///后台symbol全部小写
		if symbol == ""{
			Reply(ctx,Errors.E_PARAM,nil)
			return
		}
	}

	comm.Log.Debugf("setsymbol request:uid=%s,sid=%s, symbol=%s",suid,ssid,symbol)
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	sid,err := strconv.Atoi(ssid)
	if !Errors.CheckError(err) || sid > 100 || sid < 0{
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	ret,er := hg.SetMysymbol(uint64(uid),uint32(sid),symbol)
	Reply(ctx, er, ret)
}


func SetSlot(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	ssid := ctx.URLParam("sid")
	name := ctx.URLParam("name")
	desc := ctx.URLParam("desc")

	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	sid,err := strconv.Atoi(ssid)
	if !Errors.CheckError(err) || sid > 100 || sid < 0{
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	ret,er := hg.SetMyslot(uint64(uid),uint32(sid),name,desc)
	Reply(ctx,er,ret)
}

func GetCnsymbol(ctx iris.Context){
	filePth := "../etc/stock_code_cn.json"
	f, err := os.Open(filePth)
	defer f.Close()
	//comm.Log.Printf("open file:%s", filePth)
	if !Errors.CheckError(err){
		comm.Log.Errorf("open file failed,file=%s,err=%v",filePth,err)
		Reply(ctx,Errors.E_SUCCESS,nil)
		return}

	buf, err := ioutil.ReadAll(f)

	if !Errors.CheckError(err){
		comm.Log.Errorf("read file failed,file=%s,err=%v",filePth,err)
		Reply(ctx,Errors.E_SUCCESS,nil)
		return}
	comm.Log.Debug("load cn_code.")
	ctx.WriteString(string(buf))
}

func CheckSymbol(symbol string)(string){
	sym := strings.ToLower(symbol)
	if len(sym)<8{
		return ""
	}
	if sym[0:2] == "sz" ||sym[0:2]=="sh"{
		return sym
	}
	return ""
}

//////for  non-slots version begin //////
func GetStocks(ctx iris.Context){
	comm.Log.Debug("rcv getstocks request.")
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	syms,er := hg.GetStocks(uint64(uid))
	if !Errors.CheckError(er){Reply(ctx,er,nil);return}

	Reply(ctx,Errors.E_SUCCESS,syms)
}

func AddStock(ctx iris.Context){
	comm.Log.Debug("rcv addstock request.")
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	stk := ctx.URLParam("stk")
	syms,er := hg.AddStock(uint64(uid),stk)
	if !Errors.CheckError(er){Reply(ctx,er,nil);return}

	Reply(ctx,Errors.E_SUCCESS,syms)
}

func DelStock(ctx iris.Context){
	comm.Log.Debug("rcv delstock request.")
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	stk := ctx.URLParam("stk")
	syms,er := hg.DelStock(uint64(uid),stk)
	if !Errors.CheckError(er){Reply(ctx,er,nil);return}

	Reply(ctx,Errors.E_SUCCESS,syms)
}

//////for  non-slots version end //////