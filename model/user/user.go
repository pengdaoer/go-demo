package user

import (
	"web-demo-interface-test/dao/userDao"
)

type User struct {
	UserId 		string `json:"userId" db:"userId"`	//用户唯一标识符，由日期（精确到秒）＋ username组成
	Username 	string `json:"username" db:"username"`//用户名，登录用
	Password 	string `json:"password" db:"password"`//密码，登录用，需要加密
	Status 		int	   `json:"status" db:"status"`//用户状态，-1:注销，0:正常，1:封禁
}

type Info struct {
	UserId 		string `json:"userId" db:"userId"`
	Password 	string `json:"password" db:"password"`
	NickName 	string `json:"nickname" db:"nickname"`
}

func SelectUser(username string) string {
	var user User
	err := userDao.QueryRowDB(&user,"select userId from user_table where username=?",username)
	if err != nil{
		return ""
	}
	return user.UserId
}

func LoginUser(username,password string) User {
	var user User
	err := userDao.QueryRowDB(&user,"select userId,status from user_table where username=? and password=?",username,password)
	if err != nil{
		return User{}
	}
	return user
}

func InsertUser(user *User) error  {
	_,err := userDao.ModifyDB("insert into user_table(userid,username,password,status) values (?,?,?,?)",
		user.UserId, user.Username, user.Password, user.Status)
	_,err = userDao.ModifyDB("insert into user_info_table(userid,password,nickname) values (?,?,?)",
		user.UserId, user.Password, user.Username)
	if err != nil{
		return err
	}
	return nil
}

func DeleteUser(userId string) (int64,error) {
	rows,err := userDao.ModifyDB("update user_table set status=-1 where userId=?",userId)
	if err != nil{
		return 0,err
	}
	return rows,nil
}

func UpdateInfo(info *Info) (int64,error) {
	rows,err := userDao.ModifyDB("update user_info_table set nickname=?,password=? where userId=?",info.NickName,info.Password,info.UserId)
	if err != nil{
		return 0,err
	}
	return rows,nil
}
