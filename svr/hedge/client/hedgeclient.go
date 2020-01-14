package hedge

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"gqyb/comm"
	"gqyb/comm/errors"
	pb "gqyb/proto/proto"
)

var cli pb.HedgeClient
var conn *grpc.ClientConn

func CreateHedgeClient()(Errors.ErrorCoder){
	const svrName = "Hedge"
	ip,port,err := comm.ReadSvrAddr(svrName)
	comm.Log.Debugf("load HedgeClient config ip:%s,port:%s",ip,port)
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
	cli = pb.NewHedgeClient(conn)
	if cli == nil {
		comm.Log.Fatalf("connect to hedge svr failed.")
		return Errors.E_SYSTEM
	}
	// Contact the server and print out its response.
	return Errors.E_SUCCESS
}

func GetMyslot(uid uint64)([]byte,Errors.ErrorCoder){
	req := &pb.GetMyslotRequest{}
	req.Uid = uid
	if cli == nil{
	comm.Log.Debugf("cli is nil.")}
	rsp,err := cli.GetMyslot(context.Background(),req)
	if!Errors.CheckError(err){
		return nil, Errors.Wrap(err)
	}
	return rsp.Myslots,nil
}

func Predict(symbols []string)([]byte,Errors.ErrorCoder){
	comm.Log.Debugf("send predict() message,symbol=%v",symbols)
	if len(symbols) == 0{
		req := &pb.TimerPredictRequest{}
		rsp,err := cli.TimerPredict(context.Background(),req)
		if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
		return rsp.JsKlines,Errors.E_SUCCESS
	}else {
		req := &pb.PredictRequest{}
		req.Symbols = symbols
		rsp, err := cli.Predict(context.Background(), req)
		if !Errors.CheckError(err) {
			return nil, Errors.Wrap(err)
		}
		return rsp.JsKline, Errors.E_SUCCESS
	}
}

func SetMysymbol(uid uint64, sid uint32, symbol string)(*Slot,Errors.ErrorCoder){
	req := &pb.SetMysymbolRequest{}
	req.Uid = uid
	req.Symbol = symbol
	req.SlotId = sid
	rsp,err := cli.SetMysymbol(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	slt := &Slot{}
	err = json.Unmarshal(rsp.Slot,slt)
	if !Errors.CheckError(err){
		return nil,Errors.E_DATA
	}
	return slt,Errors.E_SUCCESS

}

/////获取默认的大盘指数预测数据
func TimerPredict()([]byte, Errors.ErrorCoder){
	req := &pb.TimerPredictRequest{}
	rsp,err := cli.TimerPredict(context.Background(),req)
	if !Errors.CheckError(err){
		comm.Log.Errorf("TimerPredict() error:%v",err)
		return nil,Errors.Wrap(err)}
	//comm.Log.Debugf("hedge cli rcv response:%v",rsp.JsKlines)
	return rsp.JsKlines,Errors.E_SUCCESS
}

////获取默认的大盘对冲分析数据
func TimeHedge()([]byte, Errors.ErrorCoder){
	req := &pb.TimerHedgeRequest{}
	rsp,err := cli.TimerHedge(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.JsResult,Errors.E_SUCCESS
}

func Hedge(scope string, bsym string)([]byte, Errors.ErrorCoder){
	req :=&pb.HedgeRequest{
		Scope:scope,
		Bsymbol: bsym,
		Ktype: KL_DAY}
	rsp,err := cli.Hedge(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.JsResult,Errors.E_SUCCESS
}


func OpenMyslot(uid uint64,months int32, days int32)([]byte, Errors.ErrorCoder){
	req := &pb.OpenMyslotRequest{}
	req.Uid = uid
	req.Months = months
	req.Days = days
	comm.Log.Debug(req)
	rsp,err := cli.OpenMyslot(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Myslots,Errors.E_SUCCESS
}

func RenewSlot(uid uint64,sid uint32,months int32, days int32)([]byte, Errors.ErrorCoder){
	req := &pb.RenewSlotRequest{}
	req.Uid = uid
	req.SlotId = sid
	req.Months = months
	req.Days = days
	rsp,err := cli.RenewSlot(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Myslots,Errors.E_SUCCESS
}

func Reward(uid uint64,months int32, days int32)([]byte, Errors.ErrorCoder){
	req := &pb.RewardRequest{}
	req.Uin = uid
	req.Months = months
	req.Days = days
	rsp,err := cli.Reward(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Myslots,Errors.E_SUCCESS
}

func SetMyslot(uid uint64,sid uint32, name string, desc string)(*Slot, Errors.ErrorCoder){
	if len(name) > 128 || len(desc) > 250{
		return nil,Errors.E_PARAM
	}
	req :=&pb.SetMyslotRequest{}
	req.Uid = uid
	req.Sid = sid
	req.Name = name
	req.Desc = desc
	rsp,err := cli.SetMyslot(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	slt := &Slot{}
	err = json.Unmarshal(rsp.Slot,slt)
	if !Errors.CheckError(err){
		return nil,Errors.E_DATA
	}
	return slt,Errors.E_SUCCESS
}

func InitSlots(uid uint64)([]byte, Errors.ErrorCoder){
	req := &pb.InitSlotsRequest{
		Uid: uid}
	rsp,err := cli.InitSlots(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Myslots,Errors.E_SUCCESS
}

////for non-slot version
func GetStocks(uid uint64)([]string,  Errors.ErrorCoder){
	req := &pb.GetAllStocksRequest{
		Uid: uid,
	}
	rsp,err := cli.GetAllStocks(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Stocks, Errors.E_SUCCESS
}

func AddStock(uid uint64, stock string)([]string, Errors.ErrorCoder){
	req := &pb.AddStockRequest{
		Uid:uid,
		Stock:stock,
	}
	rsp,err := cli.AddStock(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Stocks, Errors.E_SUCCESS
}

func DelStock(uid uint64, stock string)([]string, Errors.ErrorCoder){
	req := &pb.DelStockRequest{
		Uid:uid,
		Stock:stock,
	}
	rsp,err := cli.DelStock(context.Background(),req)
	if !Errors.CheckError(err){return nil,Errors.Wrap(err)}
	return rsp.Stocks, Errors.E_SUCCESS
}