package main

import (
	"account/biz"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"time"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "cxs20030416"
	dbname   = "account"
)

type User struct {
	Mobile   string
	Password string
	NickName string
	Gender   string
}

func main() {
	// 打开数据库连接
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成一万条随机数据并插入到数据库中
	for i := 0; i < 10000; i++ {
		user := generateUser()
		insertUser(db, user)
	}
}

func generateUser() User {
	// 生成随机mobile
	mobile := "1" + strconv.Itoa(rand.Intn(899999999)+100000000)
	// 生成随机password
	password := randomString(8)
	hashedPwd, err := biz.GetMd5(password)
	if err != nil {
		panic(err)
	}
	// 生成随机nick_name
	nickName := randomString(8)
	// 生成随机gender
	gender := "male"
	if rand.Intn(2) == 1 {
		gender = "female"
	}
	return User{Mobile: mobile, Password: hashedPwd, NickName: nickName, Gender: gender}
}

func insertUser(db *sql.DB, user User) {
	// 插入数据到数据库中 这里可以用values 然后传一个slice 应该会比一条条插快不少
	stmt, err := db.Prepare("INSERT INTO account (mobile, password, nick_name, gender, created_at, updated_at ) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Mobile, user.Password, user.NickName, user.Gender, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
