package load

import (
	"fmt"
	"os"
    "log"
    "bufio"
//    "strings"
)

func Distributemysql() []string{
	fmt.Println("读取mysql配置文件的分配记录")
	var sliceinfo []string
	sliceinfo,err := HandlePlan("file/distribute_mysql")
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }

}

func HandlePlan(textfile string) ([]string,error) {	//返回切片和error信息
    file, err := os.Open(textfile)
    var infoSlice []string =make([]string,3)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    //声明一个切片，用来存放所有配置信息
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
  
//        line := scanner.Text()  // or
		
				 infoSlice = append(infoSlice,scanner.Text())
	}

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }


	return infoSlice,nil
}