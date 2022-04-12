//go生成dbscale的配置文件
package configure

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	tool "goDeploy/tool"
	
)

func GenDBSCALEfg(dbscalecfgfile string,dbscale_ip string,basepth string,dbscalepth string,host_ip []string,zk_ip []string,normal_ip []string,shard_ip []string,slave_num string,GLOBAL_ENABLE string,BASEPTH string,DBPTH string,DBNAME string){
	
	var zkstr string
	gid := 102
	port := 16310
		
	slave , err :=  strconv.Atoi(slave_num)
	
	
	
	f, err := os.OpenFile(dbscalecfgfile+"_"+dbscale_ip+".conf", os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//删除旧文件最终保留一份
		tool.DeleteFile("file/distribute_mysql")
		//生成mysql配置文件的分发规则文件
		tool.RecordMysqlDistribute("file/distribute_mysql", "")
		f.Write([]byte("[main]\n"))
		
		f.Write([]byte("admin-password = 123456\n"))
		f.Write([]byte("admin-user = dbscale_internal\n"))
		f.Write([]byte("allow-dot-in-ident = 0\n"))
		f.Write([]byte("allow-modify-server-directly = 1\n"))
		f.Write([]byte("authenticate-source = auth\n"))
		f.Write([]byte("auto-master-failover-flashback = 1\n"))
		f.Write([]byte("auto-space-level = 0\n"))
		f.Write([]byte("auto-start-io-thread = 1\n"))
		f.Write([]byte("backend-sql-net-timeout = 60\n"))
		f.Write([]byte("backlog = 10000\n"))
		f.Write([]byte("catch-sigsegv = 1\n"))
		f.Write([]byte("close-load-conn = 1\n"))
		f.Write([]byte("cluster-password = 123456\n"))
		f.Write([]byte("cluster-user = dbscale_internal\n"))
//		f.Write([]byte("config-version = 8\n"))
		f.Write([]byte("conn-always-set-session-var = 1\n"))
		f.Write([]byte("conn-pool-num = 10\n"))
		f.Write([]byte("connection-pool-admin-interval = 60\n"))
		f.Write([]byte("cross-node-join-method = 0\n"))
		f.Write([]byte("cursor-use-free-conn = 0\n"))
		f.Write([]byte("datasource-in-one = 1\n"))
		f.Write([]byte("dbscale-acl-strict-mode = 0\n"))
		f.Write([]byte("dbscale-hosts = %\n"))
		f.Write([]byte("dbscale-internal-user = dbscale_internal\n"))
		f.Write([]byte("dbscale-master-rescramble-delay = 1\n"))
		f.Write([]byte("dbscale-safe-sql-mode = 2\n"))
		f.Write([]byte("default-login-schema = information_schema\n"))
		f.Write([]byte("default-session-variables = character_set_database:TX_READ_ONLY\n"))
		f.Write([]byte("disable-parallel-modify = 1\n"))
		f.Write([]byte("do-audit-log = NONE\n"))
		f.Write([]byte("driver = mysql\n"))
		f.Write([]byte("enable-acl = 1\n"))
		f.Write([]byte("enable-get-rep-connection = 1\n"))
		f.Write([]byte("enable-last-insert-id = 1\n"))
		f.Write([]byte("enable-multiple-stmt-check = 1\n"))
		f.Write([]byte("enable-node-follower = 0\n"))
		f.Write([]byte("enable-oracle-sequence = 1\n"))
		f.Write([]byte("enable-session-get-max-inc-value = 1\n"))
		f.Write([]byte("enable-session-swap = 0\n"))
		f.Write([]byte("enable-session-swap-during-execution = 1\n"))
		f.Write([]byte("enable-simplified-float-number = 1\n"))
		f.Write([]byte("enable-xa-transaction = 1\n"))
		f.Write([]byte("force-execute-partial-set = 0\n"))
		f.Write([]byte("function-type-file = function_type.txt\n"))
		f.Write([]byte("get-connection-retry-times = 1\n"))
		f.Write([]byte("insert-select-sql-size = 65536\n"))
		f.Write([]byte("load-analysis-num = 6\n"))
//		f.Write([]byte("load-insert-select-fields-term = ^A\n"))
//		f.Write([]byte("load-insert-select-lines-term = ^X\n"))
		f.Write([]byte("location-id = 0\n"))
		f.Write([]byte("log-file = "+basepth+""+dbscalepth+"/logs/dbscale.log\n"))
		f.Write([]byte("log-level = INFO\n"))
		f.Write([]byte("lower-case-table-names = 1\n"))
		f.Write([]byte("max-connect-by-result-num = 100000\n"))
		f.Write([]byte("max-cross-join-moved-rows = 50000000\n"))
		f.Write([]byte("max-dataserver-monitor = 40\n"))
		f.Write([]byte("max-federated-cross-join-rows = 10000\n"))
		f.Write([]byte("max-fetchnode-ready-rows-size = 10000\n"))
		f.Write([]byte("max-load-analysis-wait-size = 12\n"))
		f.Write([]byte("max-load-once-packet-num = 65535\n"))
		f.Write([]byte("max-load-packet-size = 16777216\n"))
		f.Write([]byte("max-load-ready-packets = 64\n"))
		f.Write([]byte("max-mergenode-ready-rows-size = 40000000\n"))
		f.Write([]byte("max-replication-delay = 30\n"))
		f.Write([]byte("max-single-sort-rows = 100000\n"))
		f.Write([]byte("max-slave-retrived-binlog-pos-delay-for-wakeup = 1000000\n"))
		f.Write([]byte("max-wise-group-size = 1000000\n"))
		f.Write([]byte("migrate-method = 1\n"))
		f.Write([]byte("migrate-write-packet-size = 16777216\n"))
		f.Write([]byte("migrate-write-thread = 10\n"))
		f.Write([]byte("monitor-interval = 2\n"))
		f.Write([]byte("monitor-net-timeout = 15\n"))
		f.Write([]byte("monitor-retry-count = 1\n"))
		f.Write([]byte("monitor-retry-count-stable = 2\n"))
		f.Write([]byte("mul-dbscale-forward-timeout = 100\n"))
		f.Write([]byte("multiple-mode = 1\n"))
		f.Write([]byte("node-host-addr = "+dbscale_ip+"\n"))
		f.Write([]byte("on-view = 1\n"))
		f.Write([]byte("pid-file = dbscale.pid\n"))
		f.Write([]byte("record-auto-increment-delete-value = 0\n"))
		f.Write([]byte("restrict-create-table = 0\n"))
		f.Write([]byte("session-init-charset = utf8\n"))
		f.Write([]byte("spark-dbscale-url = jdbc:mysql://127.0.0.1:3307/information_schema?zeroDateTimeBehavior=CONVERT_TO_NULL&tinyInt1isBit=false\n"))
		f.Write([]byte("support-show-warning = 1\n"))
		f.Write([]byte("support-tokudb = 0\n"))
		f.Write([]byte("thread-pool-low = 900\n"))
		f.Write([]byte("thread-pool-max = 10001\n"))
		f.Write([]byte("thread-pool-min = 10\n"))
		f.Write([]byte("update-delete-quick-limit = 0\n"))
		f.Write([]byte("use-alias-host = 0\n"))
		f.Write([]byte("use-load-data-for-insert-select = 0\n"))
		f.Write([]byte("use-partial-parse = 1\n"))
		f.Write([]byte("use-spark = 0\n"))
		f.Write([]byte("use-table-for-one-column-subquery = 1\n"))
		f.Write([]byte("wait-timeout = 172800\n"))
		f.Write([]byte("zk-log-file = "+basepth+""+dbscalepth+"/logs/zookeeper.log\n"))
		 for _,value := range zk_ip {
		 	zkstr = value+":2181," + zkstr
		 }
		 zkstr = strings.TrimRight(zkstr, ",")
		f.Write([]byte("zookeeper-host = "+zkstr+"\n"))
		
		f.Write([]byte("\n"))
		f.Write([]byte("\n[driver mysql]\n"))
		f.Write([]byte("type = MySQLDriver\n"))
		f.Write([]byte("port = "+strconv.Itoa(port)+"\n"))
		f.Write([]byte("admin-port = 23399\n"))
		f.Write([]byte("bind-address = 0.0.0.0\n"))
		f.Write([]byte("\n"))
		f.Write([]byte("\n[catalog default]\n"))
		
		for index,_ := range normal_ip {
			f.Write([]byte("data-source = normal_"+strconv.Itoa(index)+"\n"))
		}
		
		f.Write([]byte("\n[data-source auth]\n"))
		gid = gid + 1
		f.Write([]byte("group-id = "+strconv.Itoa(gid)+"\n"))
		f.Write([]byte("user = dbscale_internal\n"))
		f.Write([]byte("password = 123456\n"))
		f.Write([]byte("type = replication\n"))
		f.Write([]byte("semi-sync-on = 1\n"))
		f.Write([]byte("load-balance-strategy = MASTER\n"))
		f.Write([]byte("master = auth_0_0-1-1000-400-800\n"))
		
		for i := 0;i<slave;i++ {
			f.Write([]byte("slave = auth_0_"+strconv.Itoa(i+1)+"-1-1000-400-800\n"))
		}
		
		if GLOBAL_ENABLE == "Y" {
			//有全局角色
			f.Write([]byte("\n[data-source global_ds]\n"))
			gid = gid + 1
			f.Write([]byte("group-id =  "+strconv.Itoa(gid)+"\n"))
			f.Write([]byte("user = dbscale_internal\n"))
			f.Write([]byte("password = 123456\n"))
			f.Write([]byte("type = replication\n"))
			f.Write([]byte("semi-sync-on = 1\n"))
			f.Write([]byte("load-balance-strategy = MASTER\n"))
			f.Write([]byte("master = global_0_0-1-1000-400-800\n"))
			for  i := 0;i<slave;i++ {
				f.Write([]byte("slave = global_0_"+strconv.Itoa(i+1)+"-1-1000-400-800\n"))
			} 
			for index,_ := range normal_ip {
				f.Write([]byte("slave-source = normal_"+strconv.Itoa(index)+"\n"))
			}
			for index,_ := range shard_ip {
				f.Write([]byte("slave-source = part_"+strconv.Itoa(index)+"\n"))
			}
		}
		
		for index,_ := range normal_ip {
				gid = gid +1
				f.Write([]byte("\n[data-source normal_"+strconv.Itoa(index)+"]\n"))
				f.Write([]byte("group-id = "+strconv.Itoa(gid)+"\n"))
				f.Write([]byte("user = dbscale_internal\n"))
				f.Write([]byte("password = 123456\n"))
				f.Write([]byte("type = replication\n"))
				f.Write([]byte("semi-sync-on = 1\n"))
				f.Write([]byte("load-balance-strategy = MASTER\n"))
				f.Write([]byte("master = normal_"+strconv.Itoa(index)+"_"+strconv.Itoa(index)+"-1-1000-400-800\n"))
				for  i:= slave-1;i>=0;i-- {
					f.Write([]byte("slave = normal_"+strconv.Itoa(index)+"_"+strconv.Itoa(i+1)+"-1-1000-400-800\n"))
				} 
				f.Write([]byte("\n"))
			}
		
		for index,_ := range shard_ip {
			gid = gid + 1
			f.Write([]byte("\n[data-source part_"+strconv.Itoa(index)+"]\n"))
			f.Write([]byte("group-id =  "+strconv.Itoa(gid)+"\n"))
			f.Write([]byte("user = dbscale_internal\n"))
			f.Write([]byte("password = 123456\n"))
			f.Write([]byte("type = replication\n"))
			f.Write([]byte("semi-sync-on = 1\n"))
			f.Write([]byte("load-balance-strategy = MASTER\n"))
			f.Write([]byte("master = part_"+strconv.Itoa(index)+"_0-1-1000-400-800\n"))
				for num := slave - 1 ; num >= 0 ; num -- {
					f.Write([]byte("slave = part_"+strconv.Itoa(index)+"_"+strconv.Itoa(num+1)+"-1-1000-400-800\n"))
				} 
		}
		
		port = port + 1
		for i := slave; i >= 0 ; i -- {
			f.Write([]byte("\n[data-server auth_0_"+strconv.Itoa(i)+"]\n"))
			f.Write([]byte("host = "+host_ip[i]+"\n"))
			f.Write([]byte("port = "+strconv.Itoa(port)+"\n"))
			f.Write([]byte("user = dbscale_internal\n"))
			f.Write([]byte("password = 123456\n"))
			f.Write([]byte("remote-user = root\n"))
			f.Write([]byte("remote-port = 22\n"))
			tool.AppendMysqlDistribute("file/distribute_mysql", "\nauth_0_"+strconv.Itoa(i)+":"+host_ip[i]+":"+strconv.Itoa(port))
			GenMYSQLCfg("file/cfg/mysql_"+host_ip[i], strconv.Itoa(port),BASEPTH,DBPTH,DBNAME,i)
//			tool.FormatCfg("file/cfg/mysql_"+host_ip[i]+"_"+strconv.Itoa(port)+".conf")
		}
//		auth_0_0:172.16.70.193:16311
		port = port + 1
		if GLOBAL_ENABLE == "Y"{
			fmt.Println("当前集群有全局角色")
			for i := 0 ; i <= slave ; i++ {
				f.Write([]byte("\n[data-server global_0_"+strconv.Itoa(i)+"]\n"))
				f.Write([]byte("host = "+host_ip[i]+"\n"))
				f.Write([]byte("port = "+strconv.Itoa(port)+"\n"))
				f.Write([]byte("user = dbscale_internal\n"))
				f.Write([]byte("password = 123456\n"))
				f.Write([]byte("remote-user = root\n"))
				f.Write([]byte("remote-port = 22\n"))
				tool.AppendMysqlDistribute("file/distribute_mysql", "\nglobal_0_"+strconv.Itoa(i)+":"+host_ip[i]+":"+strconv.Itoa(port))
				GenMYSQLCfg("file/cfg/mysql_"+host_ip[i], strconv.Itoa(port),BASEPTH,DBPTH,DBNAME,i)
//				tool.FormatCfg("file/cfg/mysql_"+host_ip[i]+"_"+strconv.Itoa(port)+".conf")
			}
		}
		
		//一般数据源
		for index,_ := range normal_ip {
			port = port + 1
			for i := slave ; i >= 0; i--{
				f.Write([]byte("\n[data-server normal_"+strconv.Itoa(index)+"_"+strconv.Itoa(i)+"]\n"))
				f.Write([]byte("host = "+host_ip[i]+"\n"))
				f.Write([]byte("port = "+strconv.Itoa(port)+"\n"))
				f.Write([]byte("user = dbscale_internal\n"))
				f.Write([]byte("password = 123456\n"))
				f.Write([]byte("remote-user = root\n"))
				f.Write([]byte("remote-port = 22\n"))
				tool.AppendMysqlDistribute("file/distribute_mysql", "\nnormal_"+strconv.Itoa(index)+"_"+strconv.Itoa(i)+":"+host_ip[i]+":"+strconv.Itoa(port))
				GenMYSQLCfg("file/cfg/mysql_"+host_ip[i], strconv.Itoa(port),BASEPTH,DBPTH,DBNAME,i)
//				tool.FormatCfg("file/cfg/mysql_"+host_ip[i]+"_"+strconv.Itoa(port)+".conf")
			}
		}
		
		
		
		//分片数据源
		for index,_ := range shard_ip {
			port = port + 1
			for i := 0;i <= slave;i++{
				f.Write([]byte("\n[data-server part_"+strconv.Itoa(index)+"_"+strconv.Itoa(i)+"]\n"))
				f.Write([]byte("host = "+host_ip[i]+"\n"))
				f.Write([]byte("port = "+strconv.Itoa(port)+"\n"))
				f.Write([]byte("user = dbscale_internal\n"))
				f.Write([]byte("password = 123456\n"))
				f.Write([]byte("remote-user = root\n"))
				f.Write([]byte("remote-port = 22\n"))
				tool.AppendMysqlDistribute("file/distribute_mysql", "\npart_"+strconv.Itoa(index)+"_"+strconv.Itoa(i)+":"+host_ip[i]+":"+strconv.Itoa(port))
				GenMYSQLCfg("file/cfg/mysql_"+host_ip[i], strconv.Itoa(port),BASEPTH,DBPTH,DBNAME,i)
			}
		}
	}


}

