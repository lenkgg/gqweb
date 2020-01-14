package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"gqyb/comm"
	"gqyb/comm/errors"
	"strconv"
	sns "gqyb/sns/client"
	"strings"
)

func CreateGroup(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	pms,_ := ctx.URLParamInt("permission")
	jsgroup,err := sns.CreateGroup(uint64(uid),int32(pms))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	comm.Log.Debug(jsgroup)
	group := sns.Group{}
	err = json.Unmarshal(jsgroup,&group)
	if !Errors.CheckError(err){
		comm.Log.Errorf("CreateGroup():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,group)
}

func GetGroups(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	jsgroups,err := sns.GetGroups(uint64(uid))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	groups := []sns.Group{}
	err = json.Unmarshal(jsgroups,&groups)
	if !Errors.CheckError(err){
		comm.Log.Errorf("GetGroups():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,groups)
}

func ApplyGroup(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	gid,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	rgid,bOK,err := sns.ApplyGroup(uint64(uid),uint64(gid))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Gid  uint64			`json:"gid"`
		Apply bool			`json:"aplly"`
	}
	rsp := &RSP{
		Gid:rgid,
		Apply:bOK}

	Reply(ctx,Errors.E_SUCCESS,rsp)
}

func QuitGroup(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	gid,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	rgid,bOK,err := sns.QuitGroup(uint64(uid),uint64(gid))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Gid  uint64			`json:"gid"`
		Quit bool			`json:"quit"`
	}
	rsp := &RSP{
		Gid:rgid,
		Quit:bOK}

	Reply(ctx,Errors.E_SUCCESS,rsp)
}

func SetGroup(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	gid,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	name := ctx.URLParam("name")
	desc := ctx.URLParam("desc")
	image := ctx.URLParam("image")
	permission,err := ctx.URLParamInt("permission")
	if !Errors.CheckError(err){
		permission = 999
	}
	stat,err := ctx.URLParamInt("status")
	if !Errors.CheckError(err){
		stat = 999
	}
	jsgroup,err := sns.SetGroup(uint64(uid),uint64(gid),name,desc,int32(permission),int32(stat),image)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	group := sns.Group{}
	err = json.Unmarshal(jsgroup,&group)
	if !Errors.CheckError(err){
		comm.Log.Errorf("CreateGroup():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,group)
}

func SetManager(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	uid,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	gid,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uin,err := ctx.URLParamInt64("uin")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	op,err := ctx.URLParamInt("op")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	managers,err := sns.SetManager(uint64(uid),uint64(gid),uint64(uin),int32(op))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Gid uint64   	`json:"gid"`
		Manager []uint64	`json:"manager"`
	}
	rsp := &RSP{
		Gid:uint64(gid),
		Manager:managers}
	Reply(ctx,Errors.E_SUCCESS,rsp)
}

func ListApply(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(num)

	applys,err := sns.ListApply(uid,gid)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Gid uint64   		`json:"gid"`
		Applys []uint64		`json:"applys"`
	}
	rsp := &RSP{
		Gid:uint64(gid),
		Applys:applys}
	Reply(ctx,Errors.E_SUCCESS,rsp)
}

func AcceptMember(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(num)

	num,err = ctx.URLParamInt64("uin")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uin := uint64(num)

	mem,err := sns.AcceptMember(uid,gid,uin)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Gid uint64   		`json:"gid"`
		Member []uint64		`json:"member"`
	}
	rsp := &RSP{
		Gid:gid,
		Member:mem}
	Reply(ctx,Errors.E_SUCCESS,rsp)

}


func CreateTopic(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(num)
	tname := ctx.URLParam("name")
	desc := ctx.URLParam("desc")
	topics,err := sns.CreateTopic(uid,gid,tname,desc)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	tps := []sns.Topic{}
	err = json.Unmarshal(topics,&tps)
	if !Errors.CheckError(err){
		comm.Log.Errorf("CreateTopic():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	type RSP struct{
		Gid uint64   		`json:"gid"`
		Topics []sns.Topic		`json:"topics"`
	}
	rsp := &RSP{
		Gid:gid,
		Topics:tps}
	Reply(ctx,Errors.E_SUCCESS,rsp)

}

func SetTopic(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(num)
	tname := ctx.URLParam("name")
	desc := ctx.URLParam("desc")
	sym := ctx.URLParam("symbol")
	image := ctx.URLParam("image")
	stat,err := ctx.URLParamInt("status")
	if !Errors.CheckError(err){
		stat = 999
	}

	topics,err := sns.SetTopic(uid,gid,tname,desc,sym,int32(stat),image)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	tps := []sns.Topic{}
	err = json.Unmarshal(topics,&tps)
	if !Errors.CheckError(err){
		comm.Log.Errorf("SetTopic():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	type RSP struct{
		Gid uint64   		`json:"gid"`
		Topics []sns.Topic		`json:"topics"`
	}
	rsp := &RSP{
		Gid:gid,
		Topics:tps}
	Reply(ctx,Errors.E_SUCCESS,rsp)
}

func PubPost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num := ctx.FormValue("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	tmp,err = strconv.Atoi(num)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(tmp)

	tname := ctx.FormValue("tname")
	s_pid := ctx.FormValue("pid")
	v,_ := strconv.Atoi(s_pid)
	pid := uint64(v)
	title := ctx.FormValue("title")
	text := ctx.FormValue("text")
	attch := ctx.FormValue("attacher")
	attacher := strings.Split(attch,",")
	nick := ctx.FormValue("nick")
	avatar := ctx.FormValue("avatar")

	jspost,err := sns.PubPost(uid,gid,tname,pid,title,text,attacher,nick,avatar)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}

	ReplyString(ctx,Errors.E_SUCCESS,string(jspost))

}


func ListPost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	num,err := ctx.URLParamInt64("gid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gid := uint64(num)
	dir,err := ctx.URLParamInt("dir")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	lastpid,err := ctx.URLParamInt64("lastpid")
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}

	jspost,err := sns.ListPost(uid,gid,int32(dir),uint64(lastpid))
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}

	post := sns.PostsRsp{}
	err = json.Unmarshal(jspost,&post)
	if !Errors.CheckError(err){
		comm.Log.Errorf("ListPost():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}
	Reply(ctx,Errors.E_SUCCESS,post)

}

func GroupsPost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	str := ctx.URLParam("gids")
	list := strings.Split(str,",")
	if len(list) <1{
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	gids := []uint64{}
	for _,v := range(list){
		num,_ := strconv.Atoi(v)
		gids = append(gids,uint64(num))
	}

	jsposts,err := sns.GroupsPost(uid,gids)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}

	posts := []sns.PostsRsp{}
	err = json.Unmarshal(jsposts,&posts)
	if !Errors.CheckError(err){
		comm.Log.Errorf("ListPost():unmarshal failed.")
		Reply(ctx,Errors.E_DATA,nil)
		return
	}

	Reply(ctx,Errors.E_SUCCESS,posts)
}

func LikePost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	s_pid := ctx.URLParam("pid")
	v,_ := strconv.Atoi(s_pid)
	pid := uint64(v)

	count,err := sns.LikePost(uid,pid)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Pid uint64   		`json:"pid"`
		Likes int32			`json:"likes"`
	}
	rsp := &RSP{
		Pid:pid,
		Likes:count}
	Reply(ctx,Errors.E_SUCCESS,rsp)


}

func DeletePost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	s_pid := ctx.URLParam("pid")
	v,_ := strconv.Atoi(s_pid)
	pid := uint64(v)

	bDel,err := sns.DeletePost(uid,pid)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Pid uint64   		`json:"pid"`
		Del bool			`json:"del"`
	}
	rsp := &RSP{
		Pid:pid,
		Del:bDel}
	Reply(ctx,Errors.E_SUCCESS,rsp)

}

func ReplyPost(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)

	attch	:= ctx.FormValue("attacher")
	attacher := strings.Split(attch,",")
	if len(attacher) >3{
		Reply(ctx,Errors.E_POST_MANYPIC,nil)
		return
	}

	s_pid := ctx.FormValue("pid")
	v,_ := strconv.Atoi(s_pid)
	pid := uint64(v)
	s_rid := ctx.FormValue("rid")
	r,_ := strconv.Atoi(s_rid)
	rid := uint64(r)
	text := ctx.FormValue("text")
	mynick := ctx.FormValue("nick")
	str := ctx.FormValue("touin")
	num,err := strconv.Atoi(str)
	touin := uint64(num)
	tonick := ctx.FormValue("tonick")

	js,err := sns.ReplyPost(uid,pid,rid,text,mynick,uint64(touin),tonick,attacher)
	rpls := []sns.Reply{}
	_ = json.Unmarshal(js,&rpls)

	type RSP struct{
		Pid uint64				`json:"pid"`
		Replys []sns.Reply		`json:"reply"`
	}
	rsp := &RSP{
		Pid:pid,
		Replys:rpls}
	Reply(ctx,Errors.E_SUCCESS,rsp)

}

func DeleteReply(ctx iris.Context){
	suid := ctx.GetCookie("uid")
	tmp,err := strconv.Atoi(suid)
	if !Errors.CheckError(err){
		Reply(ctx,Errors.E_PARAM,nil)
		return
	}
	uid := uint64(tmp)
	s_pid := ctx.URLParam("pid")
	v,_ := strconv.Atoi(s_pid)
	pid := uint64(v)
	s_rid := ctx.URLParam("rid")
	r,_ := strconv.Atoi(s_rid)
	rid := uint64(r)

	bDel,err := sns.DeleteReply(uid,pid,rid)
	if !Errors.CheckError(err){
		Reply(ctx,err,nil)
		return
	}
	type RSP struct{
		Pid uint64   		`json:"pid"`
		Rid uint64			`json:"rid"`
		Del bool			`json:"del"`
	}
	rsp := &RSP{
		Pid:pid,
		Rid:rid,
		Del:bDel}
	Reply(ctx,Errors.E_SUCCESS,rsp)

}