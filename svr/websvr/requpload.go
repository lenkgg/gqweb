package main

import (
	"encoding/base64"
	"github.com/kataras/iris"
	"gqyb/comm"
	"mime/multipart"
	"strings"
)

const FILE_PATH string="../data/file/"

func upload(ctx iris.Context){
	ctx.UploadFormFiles(FILE_PATH, beforeSave)
}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {
	ip := ctx.RemoteAddr()
	ip = strings.Replace(ip, ".", "", -1)
	ip = strings.Replace(ip, ":", "", -1)
	ext := strings.Split(file.Filename,".")[1]
	filename := ip[4:] +comm.Now().String()[13:]+ file.Filename
	ext = strings.ToLower(ext)
	path := FILE_PATH
	if  ext== "jpg" ||ext == "png" ||ext=="gif"||ext=="bmp"{
		path = "image/"
	}else{
		path = "doc/"
	}

	b := []byte(filename[:16])
	//加密
	file.Filename =  base64.URLEncoding.EncodeToString(b) + "."+ext
	ctx.Writef("|%s", path+file.Filename)
}