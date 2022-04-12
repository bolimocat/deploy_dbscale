//rebuild goDeploy by 1.16
package main

import (
	"fmt"
	"strings"
	"strconv"
	
	//加载本地包
	loadProperties "goDeploy/load" //读取配置文件
	remote "goDeploy/remote"	//远程目录控制，创建目录等操作shell
	transmit "goDeploy/transmit"	//传输文件到远程目录
	configure "goDeploy/configure"
	
)

func main(){
	fmt.Println("rebuild goDeploy by 1.16")
	//开始读取配置文件
	var properties_info = loadProperties.Loadinfo("file/properties")
	fmt.Println("main中获得的配置文件信息，properties_info : \n",properties_info) //从配置文件读取所有的配置内容

	
	//定义需要的配置项数组或变量
	var user string
	var password string
	var HOST_IP []string = make([]string,0)	//声明分片
//	var GLOBAL_IP []string = make([]string,3) 
	var ZK_IP []string = make([]string,0)
	var DBSCALE_IP []string = make([]string,0)
	var NORMAL_IP []string = make([]string,0)
	var SHARD_IP []string = make([]string,0)
	
//	var mysql_num int //定义mysql实例数量 
	var BASEPTH string	//声明全局变量
	var ZKPTH string
	var DBSCALEPTH string
	var DBPTH string
	var GLOBAL_HOST string
	var SLAVE_NUM string
	var ZKNM string
	var DBSCALNM string
	var DBNAME string
		
	
	//从分片中过滤配置信息
	for _,value := range properties_info{
		//过滤用户
		if strings.Contains(value, "user"){
			user = strings.Split(value, ":")[1]
			fmt.Println("执行用户：",user)
		}
		
		//过滤密码
		if strings.Contains(value, "password"){
			password = strings.Split(value, ":")[1]
			fmt.Println("用户密码：",password)
		}
		
		//过滤HOST_IP
		if strings.Contains(value, "HOST_IP"){
			host := strings.Split(value, ":")[1] //得到一个host_ip
			//将每个IP追加到HOST_IP分片中
			HOST_IP = append(HOST_IP,host)
		}
		
		//过滤GLOBAL角色IP
		if strings.Contains(value, "GLOBAL_HOST"){
			GLOBAL_HOST = strings.Split(value, ":")[1] 
//			global := strings.Split(value, ":")[1]
//			GLOBAL_IP = append(GLOBAL_IP,global)
			fmt.Println("GLOBAL_IP : ",GLOBAL_HOST)
		}
	
	//过滤zookeeper节点IP
		if strings.Contains(value, "ZK_HOST"){
			zk_host := strings.Split(value,":")[1]
			ZK_IP = append(ZK_IP,zk_host)
		}
		//过滤dbscale节点的IP
		if strings.Contains(value, "DBSCALE_HOST"){
			dbscale_host := strings.Split(value, ":")[1]
			DBSCALE_IP = append(DBSCALE_IP,dbscale_host)
		}
		
		//过滤normal节点ip
		if strings.Contains(value, "NORMAL_HOST"){
			normal_host := strings.Split(value, ":")[1]
			NORMAL_IP = append(NORMAL_IP,normal_host)
		}
		
		//过滤shard节点ip
		if strings.Contains(value, "SHARD_HOST"){
			shard_host := strings.Split(value, ":")[1]
			SHARD_IP = append(SHARD_IP,shard_host)
		}
		
		//过滤slave个数信息
		if strings.Contains(value, "SLAVENUM"){
			SLAVE_NUM = strings.Split(value, ":")[1]
			fmt.Println("SLAVE_NUM ",SLAVE_NUM)
	}
	
		//过滤基础位置
		if strings.Contains(value,"BASTPTH"){
			BASEPTH = strings.Split(value, ":")[1]
			fmt.Println("BASEPTH ",BASEPTH)
		}
		//过滤zk位置
		if strings.Contains(value,"ZKPTH"){
			ZKPTH = strings.Split(value, ":")[1]
			fmt.Println("ZKPTH ",ZKPTH)
		}
		//过滤dbscale位置
		if strings.Contains(value,"DBSCALEPTH"){
			DBSCALEPTH = strings.Split(value, ":")[1]
			fmt.Println("DBSCALEPTH ",DBSCALEPTH)
		}
		//过滤MYSQL的位置
		if strings.Contains(value,"DBPTH"){
			DBPTH = strings.Split(value, ":")[1]
			fmt.Println("DBPTH ",DBPTH)
		}
		
		//过滤本次部署的工具包名称
		if strings.Contains(value, "ZKNM"){
			ZKNM = strings.Split(value, ":")[1]
		}
		if strings.Contains(value, "DBSCALNM"){
			DBSCALNM = strings.Split(value, ":")[1]
		}
		if strings.Contains(value, "DBNAME"){
			DBNAME = strings.Split(value, ":")[1]
		}
		
	}
	

	fmt.Println("开始 创建所有需要的目录 \n")
	//在目标节点创建所有基础目录
	for _,value := range HOST_IP{
			fmt.Println("基础目录主机： ",value," 基础目录： ",BASEPTH+"\n")
			fmt.Println("user : "+user+"  ,password : "+password)
			remote.Efolder(user,password,value,22,BASEPTH)

		}
	
	//在所有zk节点上创建需要的zk目录,上传软件包，本地生成配置文件
	for index,value := range ZK_IP{
		fmt.Println("zookeeper主机： ",value," zookeeper目录： ",ZKPTH+"\n")
		remote.Zkfolder(user,password,value,22, BASEPTH, ZKPTH,strconv.Itoa(index+1))
		fmt.Println("发送zookeeper文件包到目标位置\n")
		//文件传输
			//参数：用户名，密码，host，端口，本地文件，远程目录
			transmit.Transmit("root", "abc123", value, 22,"file/zookeeper/"+ZKNM+".tar.gz",BASEPTH+"/"+ZKPTH,ZKNM+".tar.gz","向 "+value+" 发送zookeeper")
		fmt.Println("远程解压zookeeper文件包\n")
		remote.DecoZK(user,password,value,22, BASEPTH, ZKPTH,ZKNM+".tar.gz")
		fmt.Println("\n生成zookeeper的配置文件\n")
		configure.GenZKCfg("file/cfg/zookeeper", value,BASEPTH,ZKPTH,ZK_IP)
		fmt.Println("\n传送zookeeper的配置文件\n")
			transmit.Transmit("root", "abc123", value, 22,"file/cfg/zookeeper_"+value+".cfg",BASEPTH+""+ZKPTH+"/"+ZKNM+"/conf","zoo.cfg"," 向 "+value+" 发送zookeeper.conf")
		fmt.Println("\n修改zookeeper配置文件的属主 , chown -R "+user+"."+user+" "+BASEPTH+ZKPTH+"/"+ZKNM+"/conf/zoo.cfg")
			remote.Nodemission("root", password, value, 22, "chown -R "+user+"."+user+" "+BASEPTH+ZKPTH+"/"+ZKNM+"/conf/zoo.cfg")
			remote.Nodemission(user, password, value, 22, "sed -i \"s/\\x0//g\" "+BASEPTH+ZKPTH+"/"+ZKNM+"/conf/zoo.cfg")
	}
	
	//在所有的dbscale节点创建需要的dbscale目录
	for _,value := range DBSCALE_IP{
		
		fmt.Println("dbscale主机： ",value," dbscale目录： ",DBSCALEPTH+"\n")
		remote.DBSCALEfolder(user,password,value,22, BASEPTH, DBSCALEPTH)
		fmt.Println("发送dbscale文件包到目标位置\n")
		//文件传输
			//参数：用户名，密码，host，端口，本地文件，远程目录
			transmit.Transmit("root", "abc123", value, 22,"file/dbscale/"+DBSCALNM+".tar.gz",BASEPTH+"/"+DBSCALEPTH,DBSCALNM+".tar.gz","向 "+value+" 发送dbscale")
		fmt.Println("远程解压dbscale文件包\n")
		remote.DecoDBSCALE(user,password,value,22, BASEPTH, DBSCALEPTH,DBSCALNM+".tar.gz")
		fmt.Println("修改 ",value," 节点上的dbscale的dbscale-service.sh\n")
		remote.Nodemission(user, password, value, 22, "sed -i \"s/ulimit -n 102400/ulimit -n 1024000/g\" "+BASEPTH+"/"+DBSCALEPTH+"/dbscale/dbscale-service.sh")
		fmt.Println("生成 ",value," 节点上的dbscale配置文件\n")
		configure.GenDBSCALEfg("file/cfg/dbscale", value, BASEPTH, DBSCALEPTH,HOST_IP,ZK_IP,NORMAL_IP,SHARD_IP,SLAVE_NUM,"Y",BASEPTH,DBPTH,DBNAME)

		fmt.Println("传递dbscale的配置文件\n")
			transmit.Transmit(user, password, value, 22,"file/cfg/dbscale_"+value+".conf",BASEPTH+"/"+DBSCALEPTH+"/dbscale","dbscale.conf","向 "+value+" 发送dbscale的配置文件")
			remote.Nodemission(user, password, value, 22, "sed -i \"s/\\x0//g\" "+BASEPTH+"/"+DBSCALEPTH+"/dbscale/dbscale.conf")
			remote.Nodemission(user, password, value, 22, "sed -i '286,288d' "+BASEPTH+"/"+DBSCALEPTH+"/dbscale/dbscale.conf")
	}
	


	//在所有的mysql节点创建需要的mysql目录
	for _,value := range HOST_IP{
		fmt.Println("mysql主机： ",value," mysql目录： ",DBPTH+"\n")
		remote.MYSQLfolder(user,password,value,22, BASEPTH, DBPTH)
		fmt.Println("发送mysql文件包到目标位置\n")
		//文件传输
			//参数：用户名，密码，host，端口，本地文件，远程目录
			transmit.Transmit("root", "abc123", value, 22,"file/mysql/"+DBNAME+".tar.gz",BASEPTH+"/"+DBPTH,DBNAME+".tar.gz","向 "+value+" 发送mysql")
		fmt.Println("远程解压mysql文件包\n")
		remote.DecoMYSQL(user,password,value,22, BASEPTH, DBPTH,DBNAME+".tar.gz")
		fmt.Println("将mysql二进制文件放入/usr/bin")
		remote.Nodemission("root", password, value, 22, "cp -rf "+BASEPTH+DBPTH+"/"+DBNAME+"/bin/mysql /usr/bin")
	}


	//生成Mysql的配置文件，在cfgDBSCALE中完成		
		//修改文件格式
		remote.FormatCfg("file/distribute_mysql")
		
		
		//向目标节点分发mysql配置文件
		var distribute_plan = loadProperties.Distributemysql()
		for _,value := range distribute_plan{
			if value != ""{
				ip := strings.Split(value, ":")[1]
				port :=  strings.Split(value, ":")[2]
				transmit.Transmit(user, password, ip, 22,"file/cfg/mysql_"+ip+"_"+port+".conf",BASEPTH+"/"+DBPTH,"mysql_"+ip+"_"+port+".conf","向 "+ip+" 发送mysql配置文件")
				//remote.Nodemission(user, password, value, 22, "sed -i \"s/\\x0//g\" "+BASEPTH+"/"+DBPTH+"/mysql_"+value+"_"+port+".conf")
				//remote.Nodemission(user, password, value, 22, "sed -i \"s/\\x0//g\" "+BASEPTH+"/"+DBSCALEPTH+"/dbscale/dbscale.conf")
			}
			
		}
		
		fmt.Println(" —————— 完成全部文件部署 —————— ")

		
}

