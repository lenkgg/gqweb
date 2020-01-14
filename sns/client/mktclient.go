package market

import (
	"context"
	"google.golang.org/grpc"
	"gqyb/comm"
	"gqyb/comm/errors"
	pb "gqyb/proto/proto"
)

var cli pb.MarketClient
var conn *grpc.ClientConn

func CreateMarketClient()(Errors.ErrorCoder){
	const svrName = "Market"
	ip,port,err := comm.ReadSvrAddr(svrName)
	comm.Log.Debugf("load MarketClient config ip:%s,port:%s",ip,port)
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
	cli = pb.NewMarketClient(conn)
	if cli == nil {
		comm.Log.Fatalf("connect to market svr failed.")
		return Errors.E_SYSTEM
	}
	// Contact the server and print out its response.
	return Errors.E_SUCCESS
}

func NotifyAction(uid uint64, uin uint64, act string)(Errors.ErrorCoder){
	req := &pb.NotifyActionRequest{
		Uid: uid,
		Uin: uin,
		Act: act}
	_,err := cli.NotifyAction(context.Background(),req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("Notifyaction failed:%v",err)
		return Errors.Wrap(err)
	}
	return Errors.E_SUCCESS
}

func PaySlot(uid uint64, sid uint32, months int32, days int32)([]byte, Errors.ErrorCoder){
	req := &pb.PaySlotRequest{
		Uid: uid,
		Sid: sid,
		Months: months}
	comm.Log.Debug(req)
	rsp,err := cli.PaySlot(context.Background(),req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("pay slot failed:%v",err)
		return nil,Errors.Wrap(err)
	}
	return rsp.Myslots,Errors.E_SUCCESS
}

func ListUsrShare(uid uint64, start string, end string)([]byte,Errors.ErrorCoder){
	req := &pb.ListUsrShareRequest{}
	req.Uid = uid;
	req.Start= start
	req.End = end;
	rsp,err :=cli.ListUsrShare(context.Background(),req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("ListUsrShare() error,uid=%d,error:%v",uid,err)
		return nil,Errors.Wrap(err)
	}
	return rsp.Shares,Errors.E_SUCCESS
}

func PreBill(uid uint64, months uint32, slots uint32, phone string)(Errors.ErrorCoder){
	req := &pb.PreBillRequest{
		Uid: uid,
		Months: months,
		Slots: slots,
		Phone: phone,
	}
	_,err := cli.PreBill(context.Background(),req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("prebill failed, error=%v",err)
		return Errors.Wrap(err)
	}
	return Errors.E_SUCCESS
}