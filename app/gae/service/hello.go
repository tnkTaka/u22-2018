package service

import "github.com/u22-2018/app/gae/model"

func GetHello() model.Hello{
	var massage model.Hello
	massage.Massage = "Hello"
	return massage
}
