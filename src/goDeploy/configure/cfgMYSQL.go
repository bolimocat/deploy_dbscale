//go生成mysql的配置文件
package configure

import (
	"fmt"
	"os"
	"strconv"
)

func CountMysqlNum(global_num int,normal_num int,shard_num int,slave_num string) int{
	var mysql_num int
	auth := 1
	slave,err := strconv.Atoi(slave_num)
	mysql_num = auth + slave * auth + global_num + global_num * slave + normal_num + normal_num * slave + shard_num + shard_num * slave
	if err != nil {
		fmt.Println(err.Error())
	}
	return mysql_num
}

//mysql_ip []string,basepth string,dbpth string,auth_ip string,global_num int,global_ip string,normal_num int,normal_ip []string,shard_num int,shard_ip []string,slave_num string
func GenMYSQLCfg(mysqlcfgfile string,mysql_port string,basepth string,dbpth string,dbname string,server_id int) {
	//根据配置文件计算要生成的mysql配置文件个数
	f, err := os.OpenFile(mysqlcfgfile+"_"+mysql_port+".conf", os.O_CREATE|os.O_WRONLY, 0644)
	
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		f.Write([]byte("[client] \r\n"))
		f.Write([]byte("socket = "+basepth+"/mysqldata"+mysql_port+"/mysql.sock \r\n"))
		f.Write([]byte("\r\n"))
		f.Write([]byte("[mysqld] \r\n"))
		f.Write([]byte("transaction_write_set_extraction='XXHASH64' \r\n"))
		f.Write([]byte("binlog_transaction_dependency_tracking='WRITESET' \r\n"))
		f.Write([]byte("basedir = "+basepth+""+dbpth+" \r\n"))
		f.Write([]byte("datadir = "+basepth+"/mysqldata"+mysql_port+" \r\n"))
		f.Write([]byte("port = "+mysql_port+" \r\n"))
		f.Write([]byte("socket = "+basepth+"/mysqldata"+mysql_port+"/mysql.sock \r\n"))
		f.Write([]byte("server-id = 54"+mysql_port+""+ strconv.Itoa(server_id)+ "\r\n"))
		f.Write([]byte("#tmpdir = "+basepth+"/mysqldata"+mysql_port+"/tmp \r\n"))
		f.Write([]byte("general-log-file = "+basepth+"/mysqldata"+mysql_port+"/logfile/mysqld.log \r\n"))
		f.Write([]byte("slow-query-log-file = "+basepth+"/mysqldata"+mysql_port+"/logfile/mysqld-slow.log \r\n"))
		f.Write([]byte("pid-file = "+basepth+"/mysqldata"+mysql_port+"/mysqld.pid \r\n"))
		f.Write([]byte("log-error = "+basepth+"/mysqldata"+mysql_port+"/mysqld.err \r\n"))
		f.Write([]byte("plugin-dir = "+basepth+""+dbpth+"/"+dbname+"/lib/plugin \r\n"))
		f.Write([]byte("lc-messages = en_US \r\n"))
		f.Write([]byte("lc-messages-dir = "+basepth+""+dbpth+"/"+dbname+"/share \r\n"))
		f.Write([]byte("enforce_gtid_consistency = on \r\n"))
		f.Write([]byte("binlog_ignore_db = dbscale_tmp \r\n"))
		f.Write([]byte("innodb_buffer_pool_instances = 2 \r\n"))
		f.Write([]byte("innodb_buffer_pool_size = 1073741824 \r\n"))
		f.Write([]byte("innodb_doublewrite = 0 \r\n"))
		f.Write([]byte("innodb_flush_log_at_trx_commit = 0 \r\n"))
		f.Write([]byte("innodb_flush_method = O_DIRECT \r\n"))
		f.Write([]byte("innodb_flush_neighbors = 0 \r\n"))
		f.Write([]byte("innodb_io_capacity = 10000 \r\n"))
		f.Write([]byte("innodb_io_capacity_max = 40000 \r\n"))
		f.Write([]byte("innodb_lock_wait_timeout = 20 \r\n"))
		f.Write([]byte("innodb_max_dirty_pages_pct = 90 \r\n"))
		f.Write([]byte("innodb_max_dirty_pages_pct_lwm = 10 \r\n"))
		f.Write([]byte("gtid_mode = on \r\n"))
		f.Write([]byte("innodb_print_all_deadlocks = on \r\n"))
		f.Write([]byte("innodb_read_io_threads = 32 \r\n"))
		f.Write([]byte("innodb_thread_concurrency = 128 \r\n"))
		f.Write([]byte("innodb_use_native_aio = 1 \r\n"))
		f.Write([]byte("innodb_write_io_threads = 32 \r\n"))
		f.Write([]byte("interactive_timeout = 31536000 \r\n"))
		f.Write([]byte("local_infile = 1 \r\n"))
		f.Write([]byte("lock_wait_timeout = 600 \r\n"))
		f.Write([]byte("log_bin_trust_function_creators = 1 \r\n"))
		f.Write([]byte("max_connections = 20480 \r\n"))
		f.Write([]byte("net_read_timeout = 10000 \r\n"))
		f.Write([]byte("net_write_timeout = 10000 \r\n"))
		f.Write([]byte("performance_schema = off \r\n"))
		f.Write([]byte("secure_file_priv =  \r\n"))
		f.Write([]byte("slave_skip_errors = 1396,1032,1062,1050 \r\n"))
		f.Write([]byte("slow_query_log = off \r\n"))
		f.Write([]byte("sort_buffer_size = 2097152 \r\n"))
		f.Write([]byte("sync_binlog = 0 \r\n"))
		f.Write([]byte("table_definition_cache = 5000 \r\n"))
		f.Write([]byte("table_open_cache = 5000 \r\n"))
		f.Write([]byte("table_open_cache_instances = 64 \r\n"))
		f.Write([]byte("wait_timeout = 31536000 \r\n"))
		f.Write([]byte("default_authentication_plugin = mysql_native_password \r\n"))
		f.Write([]byte("lower_case_table_names = 1 \r\n"))
		f.Write([]byte("skip_slave_start = 1 \r\n"))
		f.Write([]byte("transaction_isolation = READ-COMMITTED \r\n"))
		f.Write([]byte("ssl = off \r\n"))
		f.Write([]byte("expire_logs_days = 30 \r\n"))
		f.Write([]byte("log_bin = greatdb-bin \r\n"))
		f.Write([]byte("plugin-load = rpl_semi_sync_master=semisync_master.so;rpl_semi_sync_slave=semisync_slave.so \r\n"))
		f.Write([]byte("log-slave-updates = 1 \r\n"))
		f.Write([]byte("innodb_log_file_size = 1G \r\n"))
		f.Write([]byte("innodb_log_files_in_group = 4 \r\n"))
		
	}


}


