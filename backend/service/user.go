package service

import (
	"errors"
	"github.com/YCLiang95/CS160Group1OFS/backend/dao"
	"github.com/YCLiang95/CS160Group1OFS/backend/common/protocal"
)

var(
	RegisterError=errors.New("1000:failed to register user")
	LoginFailedErr=errors.New("1001:email/password is incorrect")
	UserSystemErr=errors.New("2000:network error,please try again")
)


func Register(email,password string)error{

	if err:=dao.GetInstance().Register(email,password);err!=nil {
		return RegisterError
	}
	return nil
}

func Login(email,password string)( *protocal.ProjectUser,error) {

	user, err := dao.GetInstance().GetUser(email)
	if err != nil{
		return nil,UserSystemErr
	}
	if user==nil||user.Password!=password{
		return nil,LoginFailedErr
	}
	return user,nil
}