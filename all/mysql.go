package main

import (
	"database/sql"
	"fmt"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

//github.com/go-sql-driver/mysql 包的路径问题
//https://github.com/go-sql-driver/mysql/blob/master/README.md
//GOPATH 下执行： go get -u github.com/go-sql-driver/mysql 即可

func main()  {
	//DSN (Data Source Name)
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	//username:password@protocol(address)/dbname?param=value
	//Except for the databasename, all values are optional. So the minimal DSN is:
	//
	//   /dbname

	//Examples

	//root:pw@unix(/tmp/mysql.sock)/myDatabase?loc=Local
	//user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true/

	//TCP using default port (3306) on localhost:
	//user:password@tcp/dbname?charset=utf8mb4,utf8&sys_var=esc%40ped

	//Use the default protocol (tcp) and host (localhost:3306):
	//user:password@/dbname

	db, err := sql.Open("mysql", "root:pwd@tcp(localhost:3308)/test")
	cheErr(err)

	//fmt.Println(db);

	//add
	//stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?") 会报错

	//stmt, err := db.Prepare("INSERT into userinfo values (?,?,?,?)")
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")

	cheErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2018-12-09")
	cheErr(err)
	id, err := res.LastInsertId()
	cheErr(err)

	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	cheErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	cheErr(err)

	affect, err := res.RowsAffected()
	cheErr(err)

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM userinfo")
	cheErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department  string
		var created string

		err = rows.Scan(&uid, &username, &department, &created)
		cheErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//cheErr(err)
	//
	//res, err = stmt.Exec(id)
	//cheErr(err)
	//
	//affect, err = res.RowsAffected()
	//cheErr(err)
	//
	//fmt.Println(affect)

	db.Close()

}

func cheErr(err error)  {
	if err != nil {
		panic(err.Error())
	}
}
