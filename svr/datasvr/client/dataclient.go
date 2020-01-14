package dc

import (
	"context"
	"google.golang.org/grpc"
	"gqyb/comm"
	pb "gqyb/proto/proto"
	"io"
	"time"
)

var sym4p_chan = make(chan string,1000)

type DcreqCli interface {
	OnRecive(DFPredict)
}

type DataClient struct{
	cli pb.DataClient
	conn *grpc.ClientConn
	stream pb.Data_PredictSymbolClient
	user *DcreqCli
}



func (p* DataClient)PredRequest(){
	for{
		select {
		case symbol := <- sym4p_chan:
			if p.stream == nil{
				time.Sleep(time.Duration(500 * time.Millisecond))
				break
			}
			req := &pb.PredictSymbolRequest{
				Symbol:symbol}

			err:=p.stream.Send(req)
			if err != nil{
				comm.Log.Errorf("request datacenter failed:%v",err)
			}
		default:
			time.Sleep(time.Duration(500 * time.Millisecond))

		}
	}
}

func (p* DataClient)PredRecive(){
	for {
		if p.stream == nil{
			time.Sleep(time.Millisecond*500)
			continue
		}
		rsp,err := p.stream.Recv()
		if err == io.EOF {
			break
		}else if err != nil{
			comm.Log.Errorf("stream recive data error:%v",err)
			continue
		}
		kline := []string{}
		for _,v := range rsp.LastOhlc{
			kline = append(kline, v)
		}
		kline = append(kline,rsp.PredictOhlc)
		prd := DFPredict{}
		prd.Symbol = rsp.Symbol
		prd.Ptime = rsp.Ptime
		prd.Kline = kline

		if p.user != nil{
			(*p.user).OnRecive(prd)
		}
		time.Sleep(time.Millisecond*500)
	}
}

func NewDataClient()*DataClient{
	p := &DataClient{}
	const svrName = "Data"
	ip,port,err := comm.ReadSvrAddr(svrName)
	comm.Log.Debugf("load DataClient config ip:%s,port:%s",ip,port)
	if err != nil{
		//comm.Log.Error("load config failed.")
		return nil
	}
	address := ip+":"+port
	//address := "127.0.0.1:9901"
	p.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		//comm.Log.Fatalf("did not connect: %v", err)
		return nil
	}
	//defer conn.Close()
	p.cli = pb.NewDataClient(p.conn)
	if p.cli == nil {
		comm.Log.Fatalf("connect to account svr failed.")
		return nil
	}
	p.stream = nil
	p.CheckPredictStream()
	// Contact the server and print out its response.
	return p
}


func(p* DataClient)RegistRcvcli(prd DcreqCli )bool{
	if prd == nil{return false}
	p.user = &prd
	return true
}

func (p*DataClient)PredictSymbols(symbols []string){
	if !p.CheckPredictStream(){return}
	for _,v := range symbols{
		sym4p_chan <- v
	}
}

func (p*DataClient)PredictSymbol(symbol string){
	if p.CheckPredictStream(){
		sym4p_chan <- symbol
	}

}

func (p*DataClient)CheckPredictStream()bool{
	if p.stream == nil{
		stream, err := p.cli.PredictSymbol(context.Background())
		if err != nil {
			comm.Log.Errorf("创建数据流失败: [%v] ", err)
			p.stream = nil
			return false
		}else{
			p.stream = stream
			go p.PredRecive()
			go p.PredRequest()
		}
	}
	return true
}