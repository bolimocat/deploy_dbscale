//go生成zookeeper的配置文件
package configure

import (
	"fmt"
	"os"
	"strconv"
)

func GenZKCfg(zkcfgfile string,zk_ip string,basepth string,zkpth string,ip []string){
	f, err := os.OpenFile(zkcfgfile+"_"+zk_ip+".cfg", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte("autopurge.purgeInterval=120\n"))
		f.Write([]byte("initLimit=5\n"))
		f.Write([]byte("syncLimit=2\n"))
		f.Write([]byte("autopurge.snapRetainCount=10\n"))
		f.Write([]byte("4lw.commands.whitelist=*\n"))
		f.Write([]byte("tickTime=6000\n"))
		f.Write([]byte("dataDir="+basepth+""+zkpth+"/data\n"))
		f.Write([]byte("dataLogDIr="+basepth+""+zkpth+"/logs\n"))
		f.Write([]byte("clientPort=2181\n"))
		f.Write([]byte("reconfigEnabled=true\n"))
		f.Write([]byte("admin.enableServer=false\n"))
		f.Write([]byte("standaloneEnabled=false\n"))
		for index,_ := range ip{
			f.Write([]byte("server."+strconv.Itoa(index+1)+"="+ip[index]+":2888:3888:participant;0.0.0.0:2181\n"))

		}
	}
}







