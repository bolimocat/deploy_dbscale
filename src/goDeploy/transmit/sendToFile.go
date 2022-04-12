//临时使用shell代替go的gossh包做文件传送
package transmit

import (
	"fmt"
    "os/exec"
)

func Sendzk(zk_ip string,basepth string,zkpth string){
	command := "shfile/sendzk.sh "+zk_ip+" "+basepth+""+zkpth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func Senddbscale(dbscale_ip string,basepth string,dbscalepth string){
	command := "shfile/senddbscale.sh "+dbscale_ip+" "+basepth+""+dbscalepth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func Sendmysql(mysql_ip string,basepth string,mysqlpth string){
	command := "shfile/sendmysql.sh "+mysql_ip+" "+basepth+""+mysqlpth
	cmd := exec.Command("/bin/bash", "-c", command)

    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}