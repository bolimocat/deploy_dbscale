//项目中的常用工具函数
package tool

import (
	"fmt"
	"math/rand"
	"os"
	 "log"
    "bufio"
    "io"
	 "os/exec"
)

//顺序内容随机排序少于4个数就不好用:(
func Disoderslice(sourceSlice []string) string{
		 for i := len(sourceSlice)-1; i > 0; i-- {
        num := rand.Intn(i + 1)
        sourceSlice[i], sourceSlice[num] = sourceSlice[num], sourceSlice[i]
    }
		 str := ""
	    for i := 0; i < len(sourceSlice); i++ {
       str += sourceSlice[i]
	    } 
		 return str
	
}

//创建用于记录mysql分发规则的空文件
func RecordMysqlDistribute(distributefile string,line string){
	
	f, err := os.OpenFile(distributefile, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte(line))
//		f.Write([]byte(line+" "))
		
}
}

//在记录mysql分发规则的文件中追加内容
func AppendMysqlDistribute(distributefile string,line string){
	f, err := os.OpenFile(distributefile, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		        content := line
		        	// 查找文件末尾的偏移量
					n, _ := f.Seek(0, 2)
			 		// 从末尾的偏移量开始写入内容
					_, err = f.WriteAt([]byte(content), n)
	}
}


//在记录mysql分发规则的文件中通过集群节点个数过滤需要发送的详细信息：(打印正常，转文件异常
func DisRepeatDistribute(filepth string,linenum int)([]string,error){
	file, err := os.Open(filepth)
    	var infoSlice []string =make([]string,3)
    	i := 0
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", filepth, err)
        return infoSlice,err
    }
    defer file.Close()
	 RecordMysqlDistribute("file/distribute_mysql", "")
    scanner := bufio.NewScanner(file)
     for scanner.Scan() {
     	line := scanner.Text() 
     
     	if i <= linenum {
     		fmt.Println(" -- ",line,i)
     		AppendMysqlDistribute("file/distribute_mysql", line+"\n")
     		i = i + 1
     	} 
     }
    
     if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", filepth, err)
        return infoSlice,err
    }
	return infoSlice,nil
        
}


//删除指定文件
func DeleteFile(deletefile string){
	err := os.Remove(deletefile)
	if err != nil{
		fmt.Print(deletefile+" 删除失败。\n")
	}else{
		fmt.Print(deletefile+" 删除成功。\n")
	}
}

//删除文件中的空行
func DeleteBlankFile(srcFilePah string, destFilePath string) error {
 srcFile, err := os.OpenFile(srcFilePah, os.O_RDONLY, 0666)
 defer srcFile.Close()
 if err != nil {
  return err
 }
 
 srcReader := bufio.NewReader(srcFile)
 destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0666)
 defer destFile.Close()
 if err != nil {
  return err
 }
// var destContent string
 for {
  str, _ := srcReader.ReadString('\n')
  if err != nil {
   if err == io.EOF {
    fmt.Print("The file end is touched.")
    break
   } else {
    return err
   }
  }
  if 0 == len(str) || str == "\r\n" {
   continue
  }
  fmt.Print(str)
  destFile.WriteString(str)
 }
 return nil
}
 
//执行单条sh命令
func ExeShell(command string){
	cmd := exec.Command(command, "")
	err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }   
    log.Printf("Waiting for command to finish...")
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)
} 

//配置文件格式化
//func FormatCfg(cfgFile string){
////	 strline := "sed -i \"s/^M//\""
//		strline := "dos2unix"
//	 cmd := exec.Command( strline+" "+cfgFile)
////	 fmt.Println("cmd -- "+strline+" "+cfgFile)
//	   err := cmd.Run()
//    if err != nil {
//        fmt.Println("Execute Command failed:" + err.Error())
//        return
//    }
//	fmt.Println(cfgFile + " 格式化文件完成。")
//}

