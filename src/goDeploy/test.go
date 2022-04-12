package main

import (
	"fmt"

	transmit "goDeploy/transmit"	//传输文件到远程目录

)

	


func main(){

	fmt.Println("传递dbscale的配置文件\n")
			transmit.Transmit("greatdb", "abc123", "192.168.2.17", 22,"file/cfg/dbscale_192.168.2.2111111.conf","/home/greatdb/goDeploy","dbscale.conf.bk111111","向 192.168.2.17 发送dbscale的配置文件")        
}
	
		

