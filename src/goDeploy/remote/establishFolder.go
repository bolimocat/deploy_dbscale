//到各个节点上创建目录（shell）
package remote

import (
	 "fmt"
//    "os/exec"
//    remote "goDeploy/remote"
)


//所有节点上的基础目录
func Efolder(user string,password string,host_ip string,port int,basepth string){
	Nodemission(user, password, host_ip, port, "mkdir -p "+basepth)
	
//	command := "shfile/establishFolder.sh "+host_ip+" "+basepth
//	cmd := exec.Command("/bin/bash", "-c", command)
//
//    output, err := cmd.Output()
//    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
//        return
//    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//在需要的节点上创建zookeeper目录
func Zkfolder(user string,password string,zk_ip string,port int,basepth string,zkpth string,myid string){
	fmt.Println("在 "+zk_ip+" 上创建zookeeper目录")
	Nodemission(user, password, zk_ip, port, "mkdir -p "+basepth+zkpth)
	fmt.Println("在 "+zk_ip+" 上创建zookeeper的data目录")
	Nodemission(user, password, zk_ip, port, "mkdir -p "+basepth+zkpth+"/data")
	fmt.Println("在 "+zk_ip+" 上创建zookeeper的logs目录")
	Nodemission(user, password, zk_ip, port, "mkdir -p "+basepth+zkpth+"/logs")
	fmt.Println("在 "+zk_ip+" 上创建zookeeper的data目录生成myid")
	Nodemission(user, password, zk_ip, port, "echo  "+myid+" > "+basepth+zkpth+"/data/myid")
	

}

//在需要的节点上创建dbscale目录
func DBSCALEfolder(user string,password string,dbscale_ip string,port int,basepth string,dbscalepth string){
	fmt.Println("在 "+dbscale_ip+" 上创建dbscale目录")
	Nodemission(user,password,dbscale_ip,port,"mkdir -p "+basepth+dbscalepth)

}

//在所有节点上创建mysql目录
func MYSQLfolder(user string,password string,host_ip string,port int,basepth string,mysqlpth string){
	fmt.Println("在 "+host_ip+" 上创建mysql目录")
	Nodemission(user,password,host_ip,port,"mkdir -p "+basepth+mysqlpth)
	

}

//修改本地配置文件的格式
func FormatCfg(cfgFile string){
	fmt.Println("修改本地的mysql配置分发调度文件")
	Nodemission("lming","123456","localhost",22,"sed -i \"s/\r//\" "+cfgFile)

}