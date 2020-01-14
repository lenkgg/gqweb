package main

import (
	"gqyb/account/client"
	"gqyb/hedge/client"
)

type tResponse struct {
	Code uint32    		`json:"code"`
	Data interface{} 	`json:"data"`
	Msg  string 		`json:"msg"`
}


type LR_RES struct {
	Info acc.USER			`json:"info"`
	Slots []hedge.Slot		`json:"slots"`
}

type LR_ORG struct {
	Info acc.ORG			`json:"info"`
}

type tNotice struct{
	Id int32			`json:"id"`
	Title string 		`json:"title"`
	Start string		`json:"start"`
	End  string		`json:"end"`
}

type rNotice struct{
	Id int32			`json:"id"`
	Title string		`json:"title"`
	Content string		`json:"content"`
}