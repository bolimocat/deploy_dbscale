package main

import (
	"fmt"
	"strings"
	"flag"	//交互包
	"github.com/samuel/go-zookeeper/zk"	//zk操作包
	"time"
	"strconv"
	loadProperties "goDeploy/load" //读取配置文件
	remote "goDeploy/remote"	//远程目录控制，创建目录等操作shell
	
)

	var command string
	
func init() {
    flag.StringVar(&command,"c","请选择操作内容：init , start ,stop","log in user")
}

func main() {
//	fmt.Println("启动Zookeeper的服务：")
	
	flag.Parse()

	//开始读取配置文件
	var properties_info = loadProperties.Loadinfo("file/properties")
	
	//定义需要的配置项数组或变量
	//执行用户
	var user string
	var password string
//	var ZK_IP []string = make([]string,0)	//声明分片
	var ZK_IP []string	//声明分片
	var BASEPTH string	//声明全局变量
	var ZKPTH string
	var ZKNM string
//	var IP string
	var zk_host string
	
	
		if command == "help"{
		fmt.Println("使用方法：-c 参数")
		fmt.Println("init ： 集群配置信息的初始化，执行一次即可。")
		fmt.Println("deletePath ： 删除所有集群的配置信息。")
		fmt.Println("start ： 启动每个节点上的zk服务。")
		fmt.Println("stop ： 停止每个节点上的zk服务。")
	}
		
		
	
		//初始化zk的基础信息
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
				//过滤HOST_IP
				if strings.Contains(value, "ZK_HOST"){
					zk_host = strings.Split(value,":")[1]
					ZK_IP = append(ZK_IP,zk_host)
				}
			
				//过滤基础位置
				if strings.Contains(value,"BASTPTH"){
					BASEPTH = strings.Split(value, ":")[1]
				}
				if strings.Contains(value,"ZKPTH"){
					ZKPTH = strings.Split(value, ":")[1]
				}
				if strings.Contains(value, "ZKNM"){
					ZKNM = strings.Split(value, ":")[1]
				}
			}
			
			hosts := []string{ZK_IP[0]+":2181"}
					// 连接zk
					conn, _, err := zk.Connect(hosts, time.Second*5)
					defer conn.Close()
					if err != nil {
						fmt.Println(err)
						return
					}
					
			fmt.Println("创建zk信息")
				add(conn,"/dbscale")
				add(conn,"/dbscale/cluster")
				add(conn,"/dbscale/cluster/nodes")
				add(conn,"/dbscale/cluster/online_nodes")
				add(conn,"/dbscale/keepalive")
				add(conn,"/dbscale/keepalive/keepalive_init_info")
				add(conn,"/dbscale/keepalive/keepalive_update_info")
				add(conn,"/dbscale/keepalive/mul_sync_info")
				add(conn,"/dbscale/dynamic_operation")
				add(conn,"/dbscale/dynamic_operation/dynamic_operation_info")
				add(conn,"/dbscale/dynamic_operation/dynamic_cluster_management_info")
				add(conn,"/dbscale/dynamic_operation/dynamic_cluster_management_init_info")
				add(conn,"/dbscale/dynamic_operation/dynamic_config_info")
				add(conn,"/dbscale/dynamic_operation/session_info")
				add(conn,"/dbscale/dynamic_operation/block_info")
				add(conn,"/dbscale/dynamic_operation/dynamic_set_option_info")
				add(conn,"/dbscale/configuration")
				add(conn,"/dbscale/configuration/master_config_info")
				add(conn,"/dbscale/configuration/changed_config_info")
			fmt.Println("创建zk信息 完成")
		}
		
		//删除path
		if command == "deletePath"{
			// 创建zk连接地址
					hosts := []string{ZK_IP[0]+":2181"}
					// 连接zk
					conn, _, err := zk.Connect(hosts, time.Second*5)
					defer conn.Close()
					if err != nil {
						fmt.Println(err)
						return
					}
			fmt.Println("删除zk信息")
			del(conn,"/dbscale/configuration/changed_config_info")
			del(conn,"/dbscale/configuration/master_config_info")
			del(conn,"/dbscale/configuration")
			del(conn,"/dbscale/dynamic_operation/dynamic_set_option_info")
			del(conn,"/dbscale/dynamic_operation/block_info")
			del(conn,"/dbscale/dynamic_operation/session_info")
			del(conn,"/dbscale/dynamic_operation/dynamic_config_info")
			del(conn,"/dbscale/dynamic_operation/dynamic_cluster_management_init_info")
			del(conn,"/dbscale/dynamic_operation/dynamic_cluster_management_info")
			del(conn,"/dbscale/dynamic_operation/dynamic_operation_info")
			del(conn,"/dbscale/dynamic_operation")
			del(conn,"/dbscale/keepalive/mul_sync_info")
			del(conn,"/dbscale/keepalive/keepalive_update_info")
			del(conn,"/dbscale/keepalive/keepalive_init_info")
			del(conn,"/dbscale/keepalive")
			del(conn,"/dbscale/cluster/online_nodes")
			del(conn,"/dbscale/cluster/nodes")
			del(conn,"/dbscale/cluster")
			del(conn,"/dbscale")
			
			fmt.Println("删除zk信息完成")
		}
		
					
		//启动所有节点的zk			
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
				if strings.Contains(value,"ZKPTH"){
					ZKPTH = strings.Split(value, ":")[1]
					
				}
		//		过滤本次部署的工具包名称
				if strings.Contains(value, "ZKNM"){
					ZKNM = strings.Split(value, ":")[1]
				}
				//过滤HOST_IP
				if strings.Contains(value, "ZK_HOST"){
					zk_host = strings.Split(value,":")[1]
					ZK_IP = append(ZK_IP,zk_host)
					
				}
			}
			
//			fmt.Println("cd "+BASEPTH+ZKPTH+"/"+ZKNM+"/bin ; ./zkServer.sh start;sleep 2")
			for index,_ := range ZK_IP{
				fmt.Println("zk_ip index "+strconv.Itoa(index))
				fmt.Println("启动： "+ZK_IP[index]+" 上的zk.")
				remote.Nodemission(user, password, ZK_IP[index], 22, BASEPTH+ZKPTH+"/"+ZKNM+"/bin/zkServer.sh start "+BASEPTH+ZKPTH+"/"+ZKNM+"/conf/zoo.cfg;sleep 2")
			}
			fmt.Println("完成所有节点zookeeper的启动.")
		}
		
		//停止所有节点的zk	
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
				if strings.Contains(value,"ZKPTH"){
					ZKPTH = strings.Split(value, ":")[1]
				}
		//		过滤本次部署的工具包名称
				if strings.Contains(value, "ZKNM"){
					ZKNM = strings.Split(value, ":")[1]
				}
				//过滤HOST_IP
				if strings.Contains(value, "ZK_HOST"){
					zk_host = strings.Split(value,":")[1]
					ZK_IP = append(ZK_IP,zk_host)
				}
			}
			
//			fmt.Println("cd "+BASEPTH+ZKPTH+"/"+ZKNM+"/bin ; ./zkServer.sh start;sleep 2")
			for index,_ := range ZK_IP{
				
				remote.Nodemission(user, password, ZK_IP[index], 22, "cd "+BASEPTH+ZKPTH+"/"+ZKNM+"/bin ; ./zkServer.sh stop "+BASEPTH+ZKPTH+"/"+ZKNM+"/conf/zoo.cfg;sleep 2")
			}
			fmt.Println("完成所有节点zookeeper的停止.")
	}

		
}
//增加内容
func add(conn *zk.Conn,path string) {
	var data = []byte("create zk info")
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	s, err := conn.Create(path, data, flags, acls)
	if err != nil {
		fmt.Printf("创建失败: %v\n", err)
		return
	}
	fmt.Printf("创建: %s 成功", s)
}

//删除内容
func del(conn *zk.Conn,path string) {
	_, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		fmt.Printf("数据删除失败: %v\n", err)
		return
	}
	fmt.Println("数据删除成功")
}
