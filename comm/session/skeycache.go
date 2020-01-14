package session

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"gqyb/comm/errors"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var skey_lock sync.Mutex

const SKEY_EXPIRED = 7    ////7 days
type keyParam struct {
	Value string			`json:"v"`
	Expire time.Time		`json:"e"`
}
type KeyCache struct{
	Cache map[string]*keyParam		`json:"cn"`
}

var g_skey KeyCache

func(p *keyParam)isExpired(v string)bool{
	if p.Expire.After(time.Now()) && (v == p.Value){
		p.Expire = time.Now().AddDate(0,0,SKEY_EXPIRED)
		return true
	}else{
		return false
	}
}

///only for regist & login cmd
func(p *KeyCache)NewUp(k string)(string,Errors.ErrorCoder){
	_,ok := p.Cache[k]
	if !ok {
		b := make([]byte, 16)
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			return "",Errors.Wrap(err)
		}
		//加密
		gqkey :=  base64.URLEncoding.EncodeToString(b)
		pm := keyParam{}
		pm.Expire = time.Now().AddDate(0, 0, SKEY_EXPIRED)
		pm.Value = gqkey

		skey_lock.Lock()
		defer skey_lock.Unlock()
		p.Cache[k] = &pm
		return gqkey,Errors.E_SUCCESS
	}else{
		skey_lock.Lock()
		p.Cache[k].Expire = time.Now().AddDate(0,0,SKEY_EXPIRED)
		skey_lock.Unlock()
		return p.Cache[k].Value,Errors.E_SUCCESS
	}
}

func(p *KeyCache)ClearAllExpired(){
	tm := time.Now()
	for k,v := range p.Cache{
		if v.Expire.Before(tm){
			skey_lock.Lock()
			delete(p.Cache, k)
			skey_lock.Unlock()
		}
	}
	time.AfterFunc(time.Duration(6*time.Hour),p.ClearAllExpired)
}

func (p *KeyCache)Save()Errors.ErrorCoder{
	if len(p.Cache) < 1 {return Errors.E_SUCCESS}
	filePth := "./sessions/gqkey"
	fp, err := os.OpenFile(filePth, os.O_RDWR|os.O_CREATE, 0755)
	if !Errors.CheckError(err){return Errors.Wrap(err)}
	defer fp.Close()
	bt,err := json.Marshal(p.Cache)
	if !Errors.CheckError(err){return Errors.Wrap(err)}
	fp.Write(bt)
	return Errors.E_SUCCESS
}

func(p *KeyCache)Load()Errors.ErrorCoder{
	filePth := "./sessions/gqkey"
	fp, err := os.OpenFile(filePth, os.O_RDWR|os.O_CREATE, 0755)
	if !Errors.CheckError(err){return Errors.Wrap(err)}
	defer fp.Close()
	buf, err := ioutil.ReadAll(fp)
	if !Errors.CheckError(err){return Errors.Wrap(err)}
	err = json.Unmarshal(buf,p.Cache)
	if !Errors.CheckError(err){return Errors.E_DATA}
	return Errors.E_SUCCESS
}

func(p *KeyCache)Verify(k string, v string)(bool){
	kk,ok := p.Cache[k]
	if !ok {
		return false
	}else{
		if kk.isExpired(v){return true}
		return false
	}
}

func InitSessionCache()(*KeyCache){
	g_skey.Cache = make(map[string]*keyParam)
	g_skey.Load()
	go timerSave()
	return &g_skey
}

func timerSave(){
	for true{
		g_skey.ClearAllExpired()
		g_skey.Save()
		//time.AfterFunc(time.Hour,timerSave)
		time.Sleep(time.Hour)
	}
}
