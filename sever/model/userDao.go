package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)


var (
	MyUserDao *UserDao
)
type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao) GetUserById(conn redis.Conn, id int) (user *User,err error) {

	res,err := redis.String(conn.Do("hget","users",id))

	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	err = json.Unmarshal([]byte(res),&user)
	if err != nil{
		fmt.Print("json.Unmarshal err = ",err)
		return
	}


	return
}


func (this *UserDao) Login(userId int, userPwd string) (user *User,err error) {

	conn := this.pool.Get()
	defer conn.Close()

	user,err = this.GetUserById(conn,userId)

	if err != nil{
		return
	}

	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}

	return
}


func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{pool:pool,}
	return
}