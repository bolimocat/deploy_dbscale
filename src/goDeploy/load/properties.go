package load

import (
	"fmt"
	"os"
    "log"
    "bufio"
    "strings"
)

func Loadinfo(profile string) []string{
//	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleText(profile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
	

}

func HandleText(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,3)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    //声明一个切片，用来存放所有配置信息
    
//    	var idx int = 0
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text()  // or
      //line := scanner.Bytes()
	     
	     //加载操作用户
	     if strings.Contains(line,"USER"){
	     	//加载用户信息
	     	user := strings.Split(line,"=")
	     	infoSlice = append(infoSlice,"user:"+user[1])

	     }
	     
	     //加载用户密码
	     if strings.Contains(line, "PASSWORD"){
	     	password := strings.Split(line, "=")
	     	infoSlice = append(infoSlice,"password:"+password[1])
	     }
	     
	     //加载所有使用的节点IP
	     if strings.Contains(line, "ALL_HOST"){
	     	//分割掉字符ALL_HOST
	     	ALL_HOST := strings.Split(line, "=")
	     	//切割HOST_IP
	     	HOST_IP := strings.Split(ALL_HOST[1], ",")
	     	for key := range HOST_IP{
	     				infoSlice = append(infoSlice,"HOST_IP:"+HOST_IP[key])
			}
	     	
	     }
	     
	     //判断当前部署是否使用GLOBAL
	      if strings.Contains(line, "GLOBAL_ENABLE") {
		      	 //扫描到GLOBAL_ENABLE行
			   GLOBAL_ENABLE := strings.Split(line,"=")
			   if GLOBAL_ENABLE[1] == "Y" {
//			   	fmt.Printf("部署使用GLOBAL角色\n")
//				   	infoSlice = append(infoSlice,"GLOBAL_ENABLE")
			   	//切割GLOBAL_HOST的信息
				   	GLOBAL_HOST := strings.Split(GLOBAL_ENABLE[2], ",")
				   	for key := range GLOBAL_HOST{
//				   		fmt.Println("GLOBAL HOST : ",GLOBAL_HOST[key])
				   		infoSlice = append(infoSlice,"GLOBAL_HOST:"+GLOBAL_HOST[key])
				   	}
				   	
			   }else{
				   	//ENABLE_GLOBAL没有选
					infoSlice = append(infoSlice,"GLOBAL_DISABLE")
			   }
			   
	      	 
	      }
	      //得到要作为GLOBAL节点的HOST
//			if strings.Contains(line, "GLOBAL_HOST"){
//				 GLOBAL_HOST := strings.Split(line, "=")
//				 fmt.Println("GLOBAL_HOST : ",GLOBAL_HOST[1])
//			   	}
			
		//得到zookeeper节点IP数组
		if strings.Contains(line, "ZK_HOST"){
			//得到配置文件中ZK行的内容
			ZK_LINE := strings.Split(line, "=")
			//切割ZK的HOST信息
				ZK_HOST := strings.Split(ZK_LINE[1], ",")
				for key := range ZK_HOST{
//					fmt.Println(" zk_host : ",ZK_HOST[key])
					infoSlice = append(infoSlice,"ZK_HOST:"+ZK_HOST[key])
				} 
		}
			
       //得到DBSCALE节点的IP数组
     if strings.Contains(line, "DBSCALE_HOST"){
     	//得到配置文件中DBSCALE行的内容
     	DBSCALE_LINE := strings.Split(line, "=")
     	//切割DBSCALE的HOST信息
	     	DBSCALE_HOST := strings.Split(DBSCALE_LINE[1], ",")
	     	for key := range DBSCALE_HOST{
//	     		fmt.Println(" DBSCALE_HOST : ",DBSCALE_HOST[key])
				infoSlice = append(infoSlice,"DBSCALE_HOST:"+DBSCALE_HOST[key])
	     	}
     }
       
      //得到NORMAL节点的IP数组 
      if strings.Contains(line, "NORMAL_HOST"){
      	//得到配置文件中的NORMAL行的内容
      	NORMAL_LINE := strings.Split(line, "=")
      	//切割normal的host信息
      	NORMAL_HOST := strings.Split(NORMAL_LINE[1], ",")
	      for key := range NORMAL_HOST{
//		     fmt.Println("NORMAL HOST : ",NORMAL_HOST[key])
		     infoSlice = append(infoSlice,"NORMAL_HOST:"+NORMAL_HOST[key])
	      }
      }
      
      //得到分片使用状态
      if strings.Contains(line, "SHARD_ENABLE"){
      	//得到SHARD_ENABLE状态行
	      SHARD_ENABLE := strings.Split(line, "=")
		      if SHARD_ENABLE[1] == "Y"{
//		      	fmt.Print("使用 分片\n")
//		      	infoSlice = append(infoSlice,"SHARD_ENABLE")
		      	//切割SHARD HOST信息
			      	SHARD_HOST := strings.Split(SHARD_ENABLE[2], ",")
			      	for key := range SHARD_HOST{
//			      		fmt.Println("SHARD HOST : ",SHARD_HOST[key])
			      		infoSlice = append(infoSlice,"SHARD_HOST:"+SHARD_HOST[key])
			      	}
		      }else{
		      	fmt.Println("不使用 分片")
		      	infoSlice = append(infoSlice,"SHARD_DISABLE")
		      }
      }
      
      //获得slave节点数量信息
      if strings.Contains(line, "SLAVE_NUM"){
      SLAVE_INFO := strings.Split(line, "=")
//      fmt.Println("本次部署的slave数量： ",SLAVE_INFO[1])
      	infoSlice = append(infoSlice,"SLAVENUM:"+SLAVE_INFO[1])
      }
     

	//获得BASE_PATH
	if strings.Contains(line, "BASE_PATH"){
		BASE_PATH := strings.Split(line, "=")
//		fmt.Println("BASE_PAHT : ",BASE_PATH[1])
		infoSlice = append(infoSlice,"BASTPTH:"+BASE_PATH[1])
	}

	//获得ZKPTH
	if strings.Contains(line,"ZKPTH"){
		ZKPTH := strings.Split(line,"=")
//		fmt.Println("ZKPTH : ",ZKPTH[1])
		infoSlice = append(infoSlice,"ZKPTH:"+ZKPTH[1])
	}
	//获得DBSCALEPTH
	if strings.Contains(line, "DBSCLPTH"){
		DBSCALEPTH := strings.Split(line, "=")
//		fmt.Println("DBSCALEPTH : ",DBSCALEPTH[1])
		infoSlice = append(infoSlice,"DBSCALEPTH:"+DBSCALEPTH[1])
	}
      
	if strings.Contains(line,"DBPTH"){
		DBPTH := strings.Split(line, "=")
//		fmt.Println("DBPTH : ",DBPTH[1])
		infoSlice = append(infoSlice,"DBPTH:"+DBPTH[1])
	}
	
	if strings.Contains(line, "ZKNM"){
		ZKNM := strings.Split(line,"=")
		infoSlice = append(infoSlice,"ZKNM:"+ZKNM[1])
	}
	
	if strings.Contains(line, "DBSCALNM"){
		DBSCALNM := strings.Split(line,"=")
		infoSlice = append(infoSlice,"DBSCALNM:"+DBSCALNM[1])
	}
	
	if strings.Contains(line, "DBNAME"){
		DBNAME := strings.Split(line,"=")
		infoSlice = append(infoSlice,"DBNAME:"+DBNAME[1])
	}
	
//	if strings.Contains(line, "CLUSTER_UPDATE_USER"){
//		CLUSTER_UPDATE_USER := strings.Split(line,"=")
//		infoSlice = append(infoSlice,"CLUSTER_UPDATE_USER:"+CLUSTER_UPDATE_USER[1])
//	}
    	
    }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

//    return nil
//	fmt.Println("infoSliceinfo: ",len(infoSlice),cap(infoSlice),infoSlice)
	return infoSlice,nil
}