package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"gqyb/comm"
	"gqyb/comm/errors"
	"strconv"
	mkt "gqyb/market/client"
	hg "gqyb/hedge/client"
)

func PaySlot(ctx iris.Context){
	suid := ctx.GetCookie("uid")
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

func ListUsrShare(ctx iris.Context){
	uid := getUid(ctx)
	if uid == 0{
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	start := ctx.URLParam("start")
	end := ctx.URLParam("end")
	comm.Log.Debugf("list share request: %s~%s",start,end)
	ret,err := mkt.ListUsrShare(uint64(uid),start,end)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	share := mkt.MarketUsr{}
	er := json.Unmarshal(ret,&share)
	if !Errors.CheckError(er){
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,share)
}

func PreBill(ctx iris.Context){
	uid := getUid(ctx)
	if uid == 0{
		Reply(ctx,Errors.E_NO_LOGIN,nil)
		return
	}
	slots,_ := ctx.URLParamInt("slots")
	months,_ := ctx.URLParamInt("months")
	phone := ctx.URLParam("phone")

	_ = mkt.PreBill(uid,uint32(months),uint32(slots),phone)
	Reply(ctx,Errors.E_SUCCESS,nil)

}