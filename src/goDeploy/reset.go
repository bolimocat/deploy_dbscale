package main

import (
	"fmt"
	"strings"
	//加载本地包
	loadProperties "goDeploy/load" //读取配置文件
	remote "goDeploy/remote"	//远程目录控制，创建目录等操作shell
)

func main(){
	fmt.Println("恢复测试环境")
	//开始读取配置文件
	var properties_info = loadProperties.Loadinfo("file/properties_reset")
	
	//定义需要的配置项数组或变量
	//执行用户
	var user string
	var password string
	var HOST_IP []string = make([]string,0)	//声明分片
	var BASEPTH string	//声明全局变量
	
	//从分片中过滤配置信息
	for _,value := range properties_info{
		//过滤用户
		if strings.Contains(value, "user"){
			user = strings.Split(value, ":")[1]
//			fmt.Println("执行用户：",user)
		}
		
		//过滤密码
		if strings.Contains(value, "password"){
			password = strings.Split(value, ":")[1]
//			fmt.Println("用户密码：",password)
		}
		
		//过滤HOST_IP
		if strings.Contains(value, "HOST_IP"){
			host := strings.Split(value, ":")[1] //得到一个host_ip
			//将每个IP追加到HOST_IP分片中
			HOST_IP = append(HOST_IP,host)
		}
		
		
	
		//过滤基础位置
		if strings.Contains(value,"BASTPTH"){
			BASEPTH = strings.Split(value, ":")[1]
//			fmt.Println("BASEPTH ",BASEPTH)
		}
		
//		fmt.Println("开始 创建所有需要的目录 \n")
	for _,value := range HOST_IP{
//			fmt.Println("删除： ",value," 基础目录： ",BASEPTH+"\n")
			remote.Nodemission(user, password, value, 22, "rm -rf "+BASEPTH)

		}
		
	}
}