//生成mysql的分发调度文件
package configure

import (
	"fmt"
	"os"
)

func RecordMysqlDistribute(distributefile string,line string){
	f, err := os.OpenFile(distributefile, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte(line+" \r\n"))
	}
}

