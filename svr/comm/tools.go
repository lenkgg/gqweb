package comm

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func Create_random_with_num_low(num int)(string){
	seed := "abcdefghijklmnopqrstuvwxyz0123456789"
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := []byte{}
	for i:=0;i<num;i++{
		ind :=rand1.Int()%(len(seed))
		ret = append(ret,seed[ind])
	}
	return string(ret)
}

func Create_random_with_num(num int)(int){
	seed := "0123456789"
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := []byte{}
	for i:=0;i<num;i++{
		ind :=rand1.Int()%(len(seed))
		ret = append(ret,seed[ind])
	}
	out,_ := strconv.Atoi(string(ret))
	return out
}

func Md5_from_binary(str string)(string){
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Random_from_list(ll []string)(string){
	rand1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := rand1.Int()%(len(ll))
	return ll[id]
}
