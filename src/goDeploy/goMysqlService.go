package main

import (
	"fmt"
	"flag"
	"strings"
	"time"
	loadProperties "goDeploy/load" //读取配置文件
	remote "goDeploy/remote"	//远程目录控制
//	tool "goDeploy/tool"
)

	var command string
//定义需要的配置项数组或变量
	var user string
	var password string
	var HOST_IP []string = make([]string,0)	//声明分片
	var BASEPTH string	//声明全局变量
	var DBPTH string
	var DBNAME string
//	var CLUSTER_UPDATE_USER string

func init() {
    flag.StringVar(&command,"c","请选择操作内容：init , start ,stop","log in user")
}

func main() {
	
	flag.Parse()
	//开始读取配置文件
	var properties_info = loadProperties.Loadinfo("file/properties")
	
	//定义需要的配置项数组或变量
	//执行用户
	var user string
	var password string
//	var ZK_IP []string = make([]string,0)	//声明分片

	var BASEPTH string	//声明全局变量
	
	if command == "help"{
		fmt.Println("使用方法：-c 参数")
		fmt.Println("init ： 执行数据库的初始化，新创建的实例只执行一次即可。")
		fmt.Println("startall ： 初始化完成的实例，全部后台启动。")
		fmt.Println("updateuser ： 根据dbscale使用的需要，更新内部用户及相关权限。")
		fmt.Println("stopall ： 停止所有节点上启动的实例。")
	}
	
	if command == "init"{
		fmt.Println("初始化所有的mysql实例")
		//从分片中过滤配置信息
			for _,value := range properties_info{
				//过滤用户
				if strings.Contains(value, "user"){
					user = strings.Split(value, ":")[1]
				}
				//过滤密码
				if strings.Contains(value, "password"){
					password = strings.Split(value, ":")[1]
				}
				//过滤基础位置
				if strings.Contains(value,"BASTPTH"){
					BASEPTH = strings.Split(value, ":")[1]
				}
				//过滤MYSQL的位置
				if strings.Contains(value,"DBPTH"){
					DBPTH = strings.Split(value, ":")[1]
					fmt.Println("DBPTH ",DBPTH)
				}
				if strings.Contains(value, "DBNAME"){
					DBNAME = strings.Split(value, ":")[1]
				}
				
			}
			
			//在所有mysql实例上执行初始化
		var distribute_plan = loadProperties.Distributemysql()
		for _,value := range distribute_plan{
			if value != ""{
				ip := strings.Split(value, ":")[1]
				port :=  strings.Split(value, ":")[2]
				fmt.Println("初始化"+ip+"上，端口是"+port+"的实例")
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysqld  --defaults-file="+BASEPTH+DBPTH+"/mysql_"+ip+"_"+port+".conf  --explicit_defaults_for_timestamp --initialize-insecure --user="+user)
			}
			
		}
	}
	if command == "startall"{
		fmt.Println("启动每个实例")
			//从分片中过滤配置信息
			for _,value := range properties_info{
				//过滤用户
				if strings.Contains(value, "user"){
					user = strings.Split(value, ":")[1]
				}
				//过滤密码
				if strings.Contains(value, "password"){
					password = strings.Split(value, ":")[1]
				}
				//过滤基础位置
				if strings.Contains(value,"BASTPTH"){
					BASEPTH = strings.Split(value, ":")[1]
				}
				//过滤MYSQL的位置
				if strings.Contains(value,"DBPTH"){
					DBPTH = strings.Split(value, ":")[1]
					fmt.Println("DBPTH ",DBPTH)
				}
				if strings.Contains(value, "DBNAME"){
					DBNAME = strings.Split(value, ":")[1]
				}
				
			}
			
			//在所有mysql实例上执行后台启动
		var distribute_plan = loadProperties.Distributemysql()
		for _,value := range distribute_plan{
			if value != ""{
				ip := strings.Split(value, ":")[1]
				port :=  strings.Split(value, ":")[2]
				fmt.Println("启动"+ip+"上，端口为 "+port+" 的mysql实例。")
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysqld_safe  --defaults-file="+BASEPTH+DBPTH+"/mysql_"+ip+"_"+port+".conf   --user="+user+" > /dev/null 2>&1 &")
			}
			
		}
	}
	if command == "updateuser"{
		fmt.Println("初始化用户")
		//从分片中过滤配置信息
			for _,value := range properties_info{
				//过滤用户
				if strings.Contains(value, "user"){
					user = strings.Split(value, ":")[1]
				}
				//过滤密码
				if strings.Contains(value, "password"){
					password = strings.Split(value, ":")[1]
				}
				//过滤基础位置
				if strings.Contains(value,"BASTPTH"){
					BASEPTH = strings.Split(value, ":")[1]
				}
				//过滤MYSQL的位置
				if strings.Contains(value,"DBPTH"){
					DBPTH = strings.Split(value, ":")[1]
					fmt.Println("DBPTH ",DBPTH)
				}
				if strings.Contains(value, "DBNAME"){
					DBNAME = strings.Split(value, ":")[1]
				}
			}
				//在所有mysql实例上执行host用户创建，修改密码等。
		var distribute_plan = loadProperties.Distributemysql()
		for _,value := range distribute_plan{
			if value != ""{
				ip := strings.Split(value, ":")[1]
				port :=  strings.Split(value, ":")[2]
				fmt.Println("在"+ip+"上，端口为 "+port+" 的mysql实例上修改root@localhost，root@%。\r\n")
//				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysql  -uroot -S "+BASEPTH+"/mysqldata"+port+"/mysql.sock  -e  \"grant all on *.* to dbscale_internal@'%' identified by '123456' with grant option;grant all on *.* to root@'%' identified by '123456' with grant option;reset master;reset slave\" " )
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysql  -uroot -S "+BASEPTH+"/mysqldata"+port+"/mysql.sock  -e  \"grant all on *.* to dbscale_internal@'%' identified by '123456' with grant option;\" " )
				time.Sleep(time.Duration(5)*time.Second)	
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysql  -uroot -S "+BASEPTH+"/mysqldata"+port+"/mysql.sock  -e  \"grant all on *.* to root@'%' identified by '123456' with grant option;\" " )
				time.Sleep(time.Duration(5)*time.Second)
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysql  -uroot -S "+BASEPTH+"/mysqldata"+port+"/mysql.sock  -e  \reset master;reset slave;\" " )
				time.Sleep(time.Duration(5)*time.Second)
			}
			
		}
			
	}
	if command == "stopall"{
		fmt.Println("停止所有实例")
		//从分片中过滤配置信息
			for _,value := range properties_info{
				//过滤用户
				if strings.Contains(value, "user"){
					user = strings.Split(value, ":")[1]
				}
				//过滤密码
				if strings.Contains(value, "password"){
					password = strings.Split(value, ":")[1]
				}
				//过滤基础位置
				if strings.Contains(value,"BASTPTH"){
					BASEPTH = strings.Split(value, ":")[1]
				}
				//过滤MYSQL的位置
				if strings.Contains(value,"DBPTH"){
					DBPTH = strings.Split(value, ":")[1]
					fmt.Println("DBPTH ",DBPTH)
				}
				if strings.Contains(value, "DBNAME"){
					DBNAME = strings.Split(value, ":")[1]
				}
			}
						//在所有mysql实例上执行host用户创建，修改密码等。
		var distribute_plan = loadProperties.Distributemysql()
		for _,value := range distribute_plan{
			if value != ""{
				ip := strings.Split(value, ":")[1]
				port :=  strings.Split(value, ":")[2]
				fmt.Println("停止"+ip+"上，端口为 "+port+" 的mysql实例")
				remote.Nodemission(user, password, ip, 22, BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysqladmin  -uroot -S "+BASEPTH+"/mysqldata"+port+"/mysql.sock  shutdown" )
			}
			
		}
			
	}
}