package acc

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"gqyb/comm"
	"gqyb/comm/errors"
	pb "gqyb/proto/proto"
	"strconv"
	"time"
	hg "gqyb/hedge/client"
)

var cli pb.AccountClient
var conn *grpc.ClientConn

func CreateAccountClient()(Errors.ErrorCoder){
	const svrName = "Account"
	ip,port,err := comm.ReadSvrAddr(svrName)
	comm.Log.Debugf("load AccClient config ip:%s,port:%s",ip,port)
	if err != nil{
		//comm.Log.Error("load config failed.")
		return Errors.E_CONFIG
	}
	address := ip+":"+port
	//address := "127.0.0.1:9901"
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		//comm.Log.Fatalf("did not connect: %v", err)
		return Errors.E_CONFIG
	}
	//defer conn.Close()
	cli = pb.NewAccountClient(conn)
	if cli == nil {
		comm.Log.Fatalf("connect to account svr failed.")
		return Errors.E_SYSTEM
	}
	// Contact the server and print out its response.
	return Errors.E_SUCCESS
}

func CloseAccountClient(){
	defer conn.Close()
}

func OrgRegist(regName string,regValue string,vcValue string,info ORG)(*ORG,Errors.ErrorCoder){
	var req pb.OrgRegistRequest
	switch regName {
	case "phone":
		req.Id = &pb.OrgRegistRequest_Phone{regValue}
		req.Verify = &pb.OrgRegistRequest_VerifyCode{vcValue}
	case "email":
		req.Id = &pb.OrgRegistRequest_Email{regValue}
		req.Verify = &pb.OrgRegistRequest_Pswd{regValue}
	default:
		return nil,Errors.E_PARAM
	}
	b,err := json.Marshal(info)
	if !Errors.CheckError(err){
		return nil,Errors.E_DATA
	}
	req.OrgInfo = b
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	comm.Log.Debugf("Request Regist message to server...")
	rsp,err := cli.OrgRegist(ctx,&req)
	if !Errors.CheckError(err) {
		//comm.Log.Fatalf("could not greet: %v", err)
		return nil,Errors.E_SYSTEM
	}
	org := &ORG{}
	json.Unmarshal(rsp.OrgInfo,org)
	return org, Errors.E_SUCCESS
}

func OrgLogin(paramName string,paramValue string,vcValue string)(*ORG, Errors.ErrorCoder){
	comm.Log.Debugf("acc org login,paramName=%s,paramValue=%s",paramName,paramValue)
	var req pb.OrgLoginRequest
	switch paramName {
	case "oid":
		u,_ := strconv.Atoi(paramValue)
		req.Id = &pb.OrgLoginRequest_Oid{uint64(u)}
	case "email":
		req.Id = &pb.OrgLoginRequest_Email{paramValue}
		req.Verify = &pb.OrgLoginRequest_Pswd{vcValue}
	case "phone":
		req.Id = &pb.OrgLoginRequest_Phone{paramValue}
		req.Verify = &pb.OrgLoginRequest_VerifyCode{vcValue}
	default:
		return nil,Errors.E_PARAM
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp,err := cli.OrgLogin(ctx,&req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("login error:%v",err)
		return nil,Errors.E_SYSTEM
	}
	comm.Log.Debugf("org login: oid=%d,err=%v",rsp.Oid,err)

	org := &ORG{}
	json.Unmarshal(rsp.OrgInfo,org)
	return org, Errors.E_SUCCESS
}


func Login(paramName string,paramValue string,vcValue string,chn uint64)(*USER,[]hg.Slot, Errors.ErrorCoder){
	comm.Log.Debugf("acc login,paramName=%s,paramValue=%s",paramName,paramValue)
	var req pb.LoginRequest
	switch paramName {
	case "jscode": //wechat login
		req.Id = &pb.LoginRequest_Wxcode{paramValue}
	case "uid":
		u,_ := strconv.Atoi(paramValue)
		req.Id = &pb.LoginRequest_Uid{uint64(u)}
	case "email":
		req.Id = &pb.LoginRequest_Email{paramValue}
		req.Verify = &pb.LoginRequest_Pswd{vcValue}
	case "phone":
		req.Id = &pb.LoginRequest_Phone{paramValue}
		req.Verify = &pb.LoginRequest_VerifyCode{vcValue}
	default:
		return nil,nil,Errors.E_PARAM
	}
	req.Channel = chn
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rsp,err := cli.Login(ctx,&req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("login error:%v",err)
		return nil,nil,Errors.E_SYSTEM
	}
	comm.Log.Debugf("cli login: uid=%d,err=%v",rsp.Uid,err)

	usr := &USER{}
	json.Unmarshal(rsp.UserInfo,usr)
	slts := []hg.Slot{}
	json.Unmarshal(rsp.Slots,&slts)
	return usr,slts, Errors.E_SUCCESS
}



func Regist(regName string,regValue string,vcValue string,info USER,chn uint64)(*USER, []hg.Slot,Errors.ErrorCoder){
	var req pb.RegistRequest
	switch regName {
	case "wxcode":
		req.Id = &pb.RegistRequest_Wxcode{regValue}
	case "phone":
		req.Id = &pb.RegistRequest_Phone{regValue}
		req.Verify = &pb.RegistRequest_VerifyCode{vcValue}
	case "email":
		req.Id = &pb.RegistRequest_Email{regValue}
		req.Verify = &pb.RegistRequest_Pswd{vcValue}
	default:
		return nil,nil,Errors.E_PARAM
	}
	b,err := json.Marshal(info)
	if !Errors.CheckError(err){
		return nil,nil,Errors.E_DATA
	}
	req.UserInfo = b
	req.Channel = chn
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	comm.Log.Debugf("Request Regist message to server...")
	rsp,err := cli.Regist(ctx,&req)
	if !Errors.CheckError(err) {
		//comm.Log.Fatalf("could not greet: %v", err)
		return nil,nil,Errors.E_SYSTEM
	}
	usr := &USER{}
	json.Unmarshal(rsp.UserInfo,usr)
	slts := []hg.Slot{}
	json.Unmarshal(rsp.Slots,&slts)
	return usr,slts, Errors.E_SUCCESS
}

func BindPhone(uid uint64,phone string, vc string)(Errors.ErrorCoder){
	var req pb.BindPhoneRequest
	req.Uid = uid
	req.Phone = phone
	req.Verify = vc

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	_,err := cli.BindPhone(ctx,&req)
	if err != nil {
		//comm.Log.Fatalf("could not greet: %v", err)
		return Errors.E_SYSTEM
	}
	return Errors.E_SUCCESS
}

func UpdateInfo(uid uint64, jsinfo []byte)(*USER, Errors.ErrorCoder){
	///info: json string
	req := &pb.UpdateInfoRequest{}
	req.Uid = uid;
	req.UserInfo = jsinfo;
	rsp,err := cli.UpdateInfo(context.Background(),req)
	if !Errors.CheckError(err){
		return nil,Errors.Wrap(err)
	}
	usr := &USER{}
	json.Unmarshal(rsp.UserInfo,usr)
	return usr, Errors.E_SUCCESS
}
