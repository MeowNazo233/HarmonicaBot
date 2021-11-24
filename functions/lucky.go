package functions

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Lucky(user_id_HMC uint64) string {
	rand_zb := 0
	year := time.Now().Year()
	month := time.Now().Format("01")
	day := time.Now().Day()
	key := false
	exist := false

	date_now := fmt.Sprintf("%d-%s-%d", year, month, day)
	db, err := sql.Open("sqlite3", "./lucky.db")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var userid int
		var date string
		var rand_a int
		err = rows.Scan(&userid, &date, &rand_a)
		checkErr(err)
		if user_id_HMC == uint64(userid) {
			exist = true
			if date == date_now {
				rand_zb = rand_a
			}
			if date != date_now {
				rand.Seed(time.Now().Unix())
				rand_zb = rand.Intn(100) + 1
				key = true
			}
		}
	}
	if exist == false {

		// //插入数据
		fmt.Print("插入数据, ID=")
		stmt, err := db.Prepare("INSERT INTO userinfo(userid, date, rand)  values(?, ?, ?)")
		checkErr(err)
		rand.Seed(time.Now().Unix())
		rand_zb = rand.Intn(100) + 1
		res, err := stmt.Exec(user_id_HMC, date_now, rand_zb)
		checkErr(err)
		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
	}
	if key { //更新数据

		stmt, err := db.Prepare("update userinfo set date=? where userid=?")
		checkErr(err)
		res, err := stmt.Exec(date_now, user_id_HMC)
		checkErr(err)
		res.RowsAffected()

		stmt, err = db.Prepare("update userinfo set rand=? where userid=?")
		checkErr(err)
		res, err = stmt.Exec(rand_zb, user_id_HMC)
		checkErr(err)
		res.RowsAffected()
	}
	db.Close()
	url1 := fmt.Sprintf("https://codechina.csdn.net/u011570312/senso-ji-omikuji/-/raw/main/%d_0.jpg", rand_zb)
	url2 := fmt.Sprintf("https://codechina.csdn.net/u011570312/senso-ji-omikuji/-/raw/main/%d_1.jpg", rand_zb)
	send := fmt.Sprintf("[CQ:at,qq=%d] [CQ:image,file=%s][CQ:image,file=%s]", user_id_HMC, url1, url2)
	//fmt.Println(send)

	return send
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
