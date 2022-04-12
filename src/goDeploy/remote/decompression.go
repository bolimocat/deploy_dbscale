//借用shell到远程目录解压tar包
package remote

import (
	"fmt"
//    "os/exec"
)

//远程解压zookeeper
func DecoZK(user string,password string,zk_ip string,port int,basepth string,zkpth string,zkname string){
	
	fmt.Println("远程解压zookeeper" )
	Nodemission(user, password, zk_ip, port, "cd "+basepth+zkpth+"; tar vzxf "+zkname)
//	command := "shfile/decoZK.sh "+zk_ip+" "+basepth+""+zkpth+" "+zkname
//	fmt.Println("command : "+command)
//	cmd := exec.Command("/bin/bash", "-c", command)
//
//    output, err := cmd.Output()
//    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
//        return
//    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//远程解压dbscale
func DecoDBSCALE(user string,password string,dbscale_ip string,port int,basepth string,dbscalepth string,dbscalename string){
	
	fmt.Println("远程解压dbscale" )
	Nodemission(user,password,dbscale_ip,port,"cd "+basepth+dbscalepth+" ; tar vzxf "+dbscalename)
	//	command := "shfile/decoDBSCALE.sh "+dbscale_ip+" "+basepth+""+dbscalepth+" "+dbscalename
//	cmd := exec.Command("/bin/bash", "-c", command)
//
//    output, err := cmd.Output()
//    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
//        return
//    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//远程解压mysql
func DecoMYSQL(user string,password string,mysql_ip string,port int,basepth string,dbpth string,dbname string){
	fmt.Println("解压mysql")
	Nodemission(user,password,mysql_ip,port,"cd "+basepth+dbpth+" ; tar vzxf "+dbname)
	//	command := "shfile/decoMYSQL.sh "+mysql_ip+" "+basepth+""+dbpth+" "+dbname
//	cmd := exec.Command("/bin/bash", "-c", command)
//
//    output, err := cmd.Output()
//    if err != nil {
//        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
//        return
//    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}