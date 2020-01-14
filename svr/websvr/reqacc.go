package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"gqyb/account/client"
	"gqyb/comm"
	"gqyb/comm/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	hg "gqyb/hedge/client"
)


func Reply(ctx iris.Context, err error, data interface{}){
	rsp := &tResponse{}
	rsp.Code = 0;
	if!Errors.CheckError(err) {
		rsp.Code = Errors.Wrap(err).Code()
		rsp.Msg = Errors.Wrap(err).Msg()
		}else{
			rsp.Data = data;
			rsp.Msg = "success."
	}
	b, e := json.Marshal(rsp)
	if e != nil {
		comm.Log.Errorf("error:%v", e)
		return
	}
	ctx.WriteString(string(b))
	comm.Log.Debugf("reply:%s",string(b))
	return
}

func ReplyString(ctx iris.Context,err Errors.ErrorCoder, data string) {
	type tt struct{
		Code uint32    		`json:"code"`
		Data string 		`json:"data"`
		Msg  string 		`json:"msg"`
	}
	rsp := &tt{}
	if !Errors.CheckError(err){
		rsp.Code = err.Code()
		rsp.Msg = err.Msg()
	}else{
		rsp.Code = 0
		rsp.Data = data
		rsp.Msg = "success."
	}
	b, e := json.Marshal(&rsp)
	if e != nil {
		comm.Log.Errorf("error:%v", e)
		return
	}
	ctx.WriteString(string(b))
	comm.Log.Debugf(string(b))
	return
}

func Login(ctx iris.Context){
	jscode := ctx.Request().FormValue("wxcode")
	chnstr := ctx.Request().FormValue("channel")
	chn,err := strconv.Atoi(chnstr)
	if !Errors.CheckError(err){ chn = 0}
	comm.Log.Debugf("login request:jscode=%s,channel=%d",jscode,chn)
	var ret Errors.ErrorCoder
	var info *acc.USER
	var slts []hg.Slot
	if jscode != ""{  //wechat login
		info,slts,ret = acc.Login("jscode",jscode,"",uint64(chn))
	}else if suid := ctx.Request().FormValue("uid"); suid!=""{
		pswd := ctx.Request().FormValue("pswd")
		info,slts,ret = acc.Login("uid",suid, pswd,uint64(chn))
	}else if email :=ctx.Request().FormValue("email");email != ""{
		pswd := ctx.Request().FormValue("pswd")
		info,slts,ret = acc.Login("email",email,pswd,uint64(chn))
	}else if phone :=ctx.Request().FormValue("phone");phone != ""{
		vc := ctx.Request().FormValue("vc")
		info,slts,ret = acc.Login("phone",phone,vc,uint64(chn))
	}else{
		comm.Log.Warning("data error.")
	}
	//comm.Log.Debugf("uid=%d,ret=%v",info.Uid,ret)
	if ret != Errors.E_SUCCESS {
		Reply(ctx,ret,nil)
	}else{
		comm.Log.Debugf("set cookie: uid = %d",info.Uid)
		skey2Cookie(ctx,info.Uid)
		res := &LR_RES{
			Info: *info,
			Slots: slts}
		Reply(ctx,Errors.E_SUCCESS,res)
	}
}

func OrgRegist(ctx iris.Context){
	var regName,regValue,vcValue string
	if email :=ctx.Request().FormValue("email");email != ""{
		regName = "email"
		regValue = email
		vcValue = ctx.Request().FormValue("pswd")
	}else if phone :=ctx.Request().FormValue("phone");phone != ""{
		regName = "phone"
		regValue = phone
		vcValue = ctx.Request().FormValue("vc")
	}else{
		comm.Log.Warning("data error.")
		ctx.EndRequest()
	}

	info := acc.ORG{
		Name: ctx.Request().FormValue("name"),
		Province: ctx.Request().FormValue("province"),
		City: ctx.Request().FormValue("city"),
		Country:ctx.Request().FormValue("country"),
		Desc: ctx.Request().FormValue("desc"),
		Other: ctx.Request().FormValue("other")}
	org,err := acc.OrgRegist(regName,regValue,vcValue,info)
	if !Errors.CheckError(err) {
		Reply(ctx,err,nil)
		return}
	orgkey2Cookie(ctx,org.Oid)
	res := &LR_ORG{
		Info: *org}
	Reply(ctx,Errors.E_SUCCESS,res)
}

func OrgLogin(ctx iris.Context){
	var ret Errors.ErrorCoder
	var info *acc.ORG
	if soid := ctx.GetCookie("oid"); soid!=""{
		pswd := ctx.URLParam("pswd")
		info,ret = acc.OrgLogin("oid",soid, pswd)
	}else if email :=ctx.URLParam("email");email != ""{
		pswd := ctx.URLParam("pswd")
		info,ret = acc.OrgLogin("email",email,pswd)
	}else if phone :=ctx.URLParam("phone");phone != ""{
		vc := ctx.URLParam("vc")
		info,ret = acc.OrgLogin("phone",phone,vc)
	}else{
		comm.Log.Warning("data error.")
	}
	//comm.Log.Debugf("uid=%d,ret=%v",info.Uid,ret)
	if ret != Errors.E_SUCCESS {
		Reply(ctx,ret,nil)
	}else{
		comm.Log.Debugf("set cookie: uid = %d",info.Oid)
		orgkey2Cookie(ctx,info.Oid)
		res := &LR_ORG{
			Info: *info}
		Reply(ctx,Errors.E_SUCCESS,res)
	}
}

//Regist(regName string,regValue string,vcValue string,info pb.UserInfo)(*pb.UserInfo,pb.ERRCODE)
func Regist(ctx iris.Context){
	var regName,regValue,vcValue string
	if wxcode := ctx.Request().FormValue("wxcode");wxcode != ""{  //wechat login
		regName = "wxcode"
		regValue = wxcode
	}else if email :=ctx.Request().FormValue("email");email != ""{
		regName = "email"
		regValue = email
		vcValue = ctx.Request().FormValue("pswd")
	}else if phone :=ctx.Request().FormValue("phone");phone != ""{
		regName = "phone"
		regValue = phone
		vcValue = ctx.Request().FormValue("vc")
	}else{
		comm.Log.Warning(ctx.Request().GetBody())
		ctx.EndRequest()
	}
	chn,err := ctx.URLParamInt64("channel")
	if !Errors.CheckError(err){ chn = 0}
	ag,_ := strconv.Atoi(ctx.Request().FormValue("age"))
	gend,_ := strconv.Atoi(ctx.Request().FormValue("gender"))
	info := acc.USER{
		Name: ctx.Request().FormValue("name"),
		Nick: ctx.Request().FormValue("nick"),
		Age: int32(ag),
		Gender: int32(gend),
		Province: ctx.Request().FormValue("province"),
		City: ctx.Request().FormValue("city"),
		Country:ctx.Request().FormValue("country"),
		Desc: ctx.Request().FormValue("desc"),
		Other: ctx.Request().FormValue("other")}
	user,slts,err := acc.Regist(regName,regValue,vcValue,info,uint64(chn))
	if !Errors.CheckError(err) {
		Reply(ctx,err,nil)
		return}
	skey2Cookie(ctx,user.Uid)
	res := &LR_RES{
		Info: *user,
		Slots: slts}
	Reply(ctx,Errors.E_SUCCESS,res)
}

/////when regist or login cmd done, write uid & gqkey to cookie
func skey2Cookie(ctx iris.Context,uid uint64){
	comm.Log.Debugf("write gqkey.")
	suid := strconv.Itoa(int(uid))
	gqkey,err := skey_cache.NewUp(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_SYSTEM,nil)
		return
	}
	comm.Log.Debugf("gqkey: uid=%d, gqkey=%s",uid,gqkey)
	http.SetCookie(ctx.ResponseWriter(),&http.Cookie{
		Name: "uid",
		Value: suid})
	http.SetCookie(ctx.ResponseWriter(),&http.Cookie{
		Name: "gqkey",
		Value: gqkey})

}
func orgkey2Cookie(ctx iris.Context,oid uint64){
	comm.Log.Debugf("write gqkey.")
	soid := strconv.Itoa(int(oid))
	gqkey,err := skey_cache.NewUp(soid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_SYSTEM,nil)
		return
	}
	comm.Log.Debugf("gqkey: oid=%d, gqkey=%s",oid,gqkey)
	http.SetCookie(ctx.ResponseWriter(),&http.Cookie{
		Name: "oid",
		Value: soid})
	http.SetCookie(ctx.ResponseWriter(),&http.Cookie{
		Name: "gqkey",
		Value: gqkey})

}


func BindPhone(ctx iris.Context){
	u,err := strconv.Atoi(ctx.GetCookie("uid"))
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(u)
	phone := ctx.URLParam("phone")
	vc := ctx.URLParam("vc")
	e := acc.BindPhone(uid,phone,vc)
	if !Errors.CheckError(e){
		Reply(ctx,e,nil)
	}else{
		Reply(ctx,Errors.E_SUCCESS,nil)
	}
}

func UpdateInfo(ctx iris.Context){
	uid,err := strconv.Atoi(ctx.GetCookie("uid"))
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	usr := acc.USER{}
	usr.Uid = uint64(uid)
	usr.Name = ctx.Request().FormValue("nickName")
	sg :=ctx.Request().FormValue("gender");
	gnd,err := strconv.Atoi(sg);
	if !Errors.CheckError(err){
		comm.Log.Warnf("gender error,uid=%d",uid)
		gnd = 0;
	}
	usr.Gender = int32(gnd)
	usr.Language = ctx.Request().FormValue("language")
	usr.City = ctx.Request().FormValue("city")
	usr.Province = ctx.Request().FormValue("province")
	usr.Country = ctx.Request().FormValue("country")
	usr.Avata = ctx.Request().FormValue("avatarUrl")

	comm.Log.Debug(usr)
	bs,err := json.Marshal(usr);
	if !Errors.CheckError(err){
		comm.Log.Errorf("encode user error:%v",err)
		Reply(ctx,err,nil)
		return
	}
	rsp_usr,err := acc.UpdateInfo(uint64(uid),bs)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}else{
		Reply(ctx,Errors.E_SUCCESS,*rsp_usr)
	}

}

func getUid(ctx iris.Context)(uint64){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		return 0
	}
	return uint64(uid)
}

func Notice(ctx iris.Context){
	filePth := "../etc/notice.json"
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
	nt := &tNotice{}
	err = json.Unmarshal(buf,nt)
	if !Errors.CheckError(err){
		comm.Log.Errorf("unmarshal json file failed,file=%s,err=%v",filePth,err)
		Reply(ctx,Errors.E_SUCCESS,nil)
		return}
	now := comm.Now()
	if comm.NewTime(nt.Start).Before(now) && comm.NewTime(nt.End).After(now){
		fn := "../etc/notice/" + nt.Title +".txt"
		fl,err :=os.Open(fn)
		defer fl.Close()
		//comm.Log.Printf("open file:%s", filePth)
		if !Errors.CheckError(err){
			comm.Log.Errorf("open file failed,file=%s,err=%v",filePth,err)
			Reply(ctx,Errors.E_SUCCESS,nil)
			return}

		cnt, err := ioutil.ReadAll(fl)
		if !Errors.CheckError(err){
			comm.Log.Errorf("read file failed,file=%s,err=%v",filePth,err)
			Reply(ctx,Errors.E_SUCCESS,nil)
			return}
		rsp := rNotice{}
		rsp.Id = nt.Id;
		rsp.Title = nt.Title;
		rsp.Content = string(cnt)
		Reply(ctx,Errors.E_SUCCESS,rsp)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,nil)
}