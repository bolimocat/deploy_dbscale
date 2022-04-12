package main

import (
	"fmt"
	"flag"
	"strings"
//	"database/sql"
	loadProperties "goDeploy/load" //读取配置文件
	remote "goDeploy/remote"
//	mysqlstr "goDeploy/mysqlstr"
)

var command string

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
	var BASEPTH string	//声明全局变量
	var DBSCALEPTH string
//	var DBSCALNM string
	var dbscale_host string
	var DBSCALE_IP []string = make([]string,0)
	
	
	if command == "help"{
		fmt.Println("使用方法：-c 参数")
		fmt.Println("start ： 初始化完成的实例，全部后台启动。")
		fmt.Println("stop ： 停止所有节点上启动的实例。")
	}
	
	if command == "start"{
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
				if strings.Contains(value,"DBSCALEPTH"){
					DBSCALEPTH = strings.Split(value, ":")[1]
					
				}
		//		过滤本次部署的工具包名称
//				if strings.Contains(value, "DBSCALNM"){
//					DBSCALNM = strings.Split(value, ":")[1]
//				}
				//过滤DBSCALE_IP
				if strings.Contains(value, "DBSCALE_HOST"){
					dbscale_host = strings.Split(value,":")[1]
					DBSCALE_IP = append(DBSCALE_IP,dbscale_host)
					
				}
			}
			
			
			for index,_ := range DBSCALE_IP{
				fmt.Println("启动： "+DBSCALE_IP[index]+" 上的dbscale.")
				fmt.Println(BASEPTH+DBSCALEPTH+"/dbscale/daemon/dbscale_daemon.py")

				remote.Nodemission(user, password, DBSCALE_IP[index], 22, "cd "+BASEPTH+DBSCALEPTH+"/dbscale;daemon/dbscale_daemon.py ;sleep 2")
			}
			fmt.Println("完成所有节点dbscale的启动.")
		}
	
	if command == "init"{
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
				if strings.Contains(value,"DBSCALEPTH"){
					DBSCALEPTH = strings.Split(value, ":")[1]
					
				}
		//		过滤本次部署的工具包名称
//				if strings.Contains(value, "DBSCALNM"){
//					DBSCALNM = strings.Split(value, ":")[1]
//				}
				//过滤DBSCALE_IP
				if strings.Contains(value, "DBSCALE_HOST"){
					dbscale_host = strings.Split(value,":")[1]
					DBSCALE_IP = append(DBSCALE_IP,dbscale_host)
					
				}
			}
			
			
			for index,_ := range DBSCALE_IP{
				
				if index == 0 {
					fmt.Println("给集群中的root用户执行授权."+DBSCALE_IP[index])
					fmt.Println("mysql -uroot -p123456 -h127.0.0.1 -P16310 -e 'grant all on *.* to root@'%';'")
					remote.Nodemission(user, password, DBSCALE_IP[index], 22, "mysql -uroot -p123456 -h127.0.0.1 -P16310 -e 'grant all on *.* to root@%;'")
//					remote.Nodemission(user, password, DBSCALE_IP[index], 22, "mysql -uroot -p123456 -h"+DBSCALE_IP[index]+" -P16310 -e 'grant all on *.* to root@'%';'")
//					remote.Nodemission(user, password, DBSCALE_IP[index], 22, "mysql -uroot -p123456 -h"+DBSCALE_IP[index]+" -P16310 -e 'dbscale dynamic add normal_table dataspace test1.tb_0 datasource = \"global_ds\";'")
				}
				
			}
			fmt.Println("完成dbscale上root初始授权的启动.")
		}
	
	
	
	if command == "stop"{
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
				if strings.Contains(value,"DBSCALEPTH"){
					DBSCALEPTH = strings.Split(value, ":")[1]
					
				}
		//		过滤本次部署的工具包名称
//				if strings.Contains(value, "DBSCALNM"){
//					DBSCALNM = strings.Split(value, ":")[1]
//				}
				//过滤DBSCALE_IP
				if strings.Contains(value, "DBSCALE_HOST"){
					dbscale_host = strings.Split(value,":")[1]
					DBSCALE_IP = append(DBSCALE_IP,dbscale_host)
					
				}
			}
			

			for index,_ := range DBSCALE_IP{
				fmt.Println("停止： "+DBSCALE_IP[index]+" 上的dbscale.")
				remote.Nodemission(user, password, DBSCALE_IP[index], 22, "cd "+BASEPTH+DBSCALEPTH+"/dbscale;daemon/stopper.sh ;sleep 5")
			}
			fmt.Println("完成所有节点dbscale的停止.")
		}
	
	if command == "conn"{
		fmt.Println("尝试连接dbscale.")
//		mysqlstr.MysqlQueryRow("123")
	}
	
}