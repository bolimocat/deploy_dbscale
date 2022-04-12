package mysqlstr

import (
	"strings"
	"database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func MysqlQueryRow(sqlstr string){
	//开始读取配置文件
	var properties_info = loadProperties.Loadinfo("file/properties")
	var user string
	var password string
	var dbscale_host string
	var DBSCALE_IP []string = make([]string,0)
	
	
	for _,value := range properties_info{
		//过滤用户
		if strings.Contains(value, "user"){
				user = strings.Split(value, ":")[1]
				
		}
		//过滤密码
		if strings.Contains(value, "password"){
				password = strings.Split(value, ":")[1]
		}
		
		for index,_ := range DBSCALE_IP{
				if index == 0{
					dbscale_host = DBSCALE_IP[index]
				}
			}
	}
	
//	//数据库连接
//    db,_:=sql.Open("mysql",user+":"+password+"@("+dbscale_host+":16310)/mysql")
//    err :=db.Ping()
//    if err != nil{
//        fmt.Println("数据库链接  失败")
//    } 
//　　defer db.Close()

 //数据库连接
    db,_:=sql.Open("mysql","root:root@(127.0.0.1:3306)/golang")
    
    err :=db.Ping()
    if err != nil{
        fmt.Println("数据库链接失败")
    }
　　　

　　defer db.Close()

}

