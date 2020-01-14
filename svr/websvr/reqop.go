package main

import(
	"encoding/json"
	"github.com/kataras/iris"
	"gqyb/comm"
	"gqyb/comm/errors"
	"io/ioutil"
	"os"
	"strconv"
	mkt "gqyb/market/client"
	hg "gqyb/hedge/client"
)


func verifyop(ctx iris.Context){
	op := ctx.URLParam("op")
	pswd := ctx.URLParam("pswd")
	if (op != "op_lenk") ||(pswd != "gogoGqybA10F"){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	ctx.Next()
}

func OpAddSlot(ctx iris.Context){
	suid := ctx.URLParam("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	sid,_ := ctx.URLParamInt("sid")
	months,_ := ctx.URLParamInt("months")
	ret,er := mkt.PaySlot(uint64(uid),uint32(sid),int32(months),0)
	if !Errors.CheckError(er){
		Reply(ctx,er,nil)
		return
	}
	slts := []hg.Slot{}
	err = json.Unmarshal(ret,&slts)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,slts)
}

func VtradePosition(ctx iris.Context){
	filePth := "../data/vtrade/quant.json"
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

	ctx.WriteString(string(buf))
}

func VtradeHedge(ctx iris.Context){
	filePth := "../data/vtrade/hedge.json"
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

	ctx.WriteString(string(buf))
}
