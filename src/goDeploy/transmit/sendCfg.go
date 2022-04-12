//借用shell，向目标端发送配置文件
package transmit

import (
	 "fmt"
    "os/exec"
//    "strings"
//    "log"
)

//传送zookeeper配置
func SendZKcfg(zk_ip string,basepth string,zkpth string){
	command := "shfile/sendzkcfg.sh "+zk_ip+" "+basepth+""+zkpth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//传送dbscale配置
func SendDBSCALEcfg(dbscale_ip string,basepth string,dbscalepth string){
	command := "shfile/senddbscalecfg.sh "+dbscale_ip+" "+basepth+""+dbscalepth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//传送mysql的配置到指定节点
//func SendMYSQLcfg(distibutefile string,basepth string,dbpth string){
//	
//	var mysql_port string
//	var mysql_ip string
//	
//	
//	file, err := os.Open(distibutefile)
////    	var infoSlice []string =make([]string,3)
//    if err != nil {
//        log.Printf("Cannot open text file: %s, err: [%v]", distibutefile, err)
////        return infoSlice,err
//    }
//    defer file.Close()
//	 scanner := bufio.NewScanner(file)
//	 for scanner.Scan() {
//		 	  line := scanner.Text()
//		 	  mysql_port = strings.Split(line,":")[2]
//		 	  mysql_ip = strings.Split(line, ":")[1]
//		
//		 	  command := "shfile/sendmysqlcfg.sh "+mysql_port+" "+mysql_ip+" "+basepth+""+dbpth
//			  cmd := exec.Command("/bin/bash", "-c", command)
//
//		     output, err := cmd.Output()
//		    if err != nil {
//	        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
//	        return
//		    }
//    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
//	 }
//
//}











