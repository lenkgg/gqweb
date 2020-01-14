package mg

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gqyb/comm"
	"io/ioutil"
	"os"
	"gqyb/comm/errors"
)


type configure struct{
	uri string      //"mongodb://localhost:27017"
}

func readConf(svrName string)(*configure){
	file, _ := os.Open("../etc/mongo.json")
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil{
		return nil
	}
	conf := make(map[string]interface{})
	json.Unmarshal([]byte(data), &conf)
	mg := &configure{}
	mg.uri = conf[svrName].(map[string]interface{})["uri"].(string)
	return mg
}

var mg_cli *mgo.Session

func InitMongo(svrName string) (*mgo.Session,error){
	conf:=readConf(svrName)
	comm.Log.Debugf("uri:%s",conf.uri)
	// 连接数据库
	sess,err := mgo.Dial(conf.uri)
	if !Errors.CheckError(err) {return nil,err}
	sess.SetMode(mgo.Monotonic, true)

	//db := client.Database(conf.db)
	mg_cli = sess
	return sess,nil
}

func CloseMongo(){
	defer mg_cli.Close()
}

type T_AUTOINC struct {
	_id string
	col string			`bson:"col"`
	id int				`bson:"id"`
}

const AUTOINC_DB  string = "autoinc"
const AUTOINC_COL string = "col_inc"
const AUTOINC_UID string = "user_uid"
const AUTOINC_GID string = "group_id"

func AutoIncUid()(int64){
	//eg: {"col":"user","id":1000}
	//cli,err := InitMongo("AutoInc")
	//if !comm.CheckError(err) {return -1}
	col := mg_cli.DB(AUTOINC_DB).C(AUTOINC_COL)
	IDInt64 := struct {
		Value int64 `bson:"id"`
	}{Value: 1}
	_, err := col.Find(bson.M{"col": AUTOINC_UID}).Apply(mgo.Change{Update: bson.M{"$inc": IDInt64},
		Upsert: true, ReturnNew: true}, &IDInt64)
	if !Errors.CheckError(err){return -1}
	comm.Log.Debugf("AutoInc id=%d",IDInt64)
	return IDInt64.Value
}

func AutoIncGid()(int64){
	//eg: {"col":"user","id":1000}
	//cli,err := InitMongo("AutoInc")
	//if !comm.CheckError(err) {return -1}
	col := mg_cli.DB(AUTOINC_DB).C(AUTOINC_COL)
	IDInt64 := struct {
		Value int64 `bson:"id"`
	}{Value: 1}
	_, err := col.Find(bson.M{"col": AUTOINC_GID}).Apply(mgo.Change{Update: bson.M{"$inc": IDInt64},
		Upsert: true, ReturnNew: true}, &IDInt64)
	if !Errors.CheckError(err){return -1}
	comm.Log.Debugf("AutoInc id=%d",IDInt64)
	return IDInt64.Value
}
