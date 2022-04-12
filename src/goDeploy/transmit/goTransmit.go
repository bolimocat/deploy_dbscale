package transmit

import (
	  "fmt"
	  "log"
	  "os"
	  "path"
	  "net"
	  "time"
	  "github.com/pkg/sftp"
     "golang.org/x/crypto/ssh"
)


//transmitContent：本次操作的说明
func Transmit(user string, password string,hostip string,sshport int ,localFilePath string, remoteDir string,remoteName string,transmitContent string){
	fmt.Println(transmitContent)
	  var err   error
	  var  sftpClient  *sftp.Client
  
    // 这里换成实际的 SSH 连接的 用户名，密码，主机名或IP，SSH端口
  sftpClient, err = connect(user, password, hostip, sshport)
  if err != nil {
    log.Fatal(err)
  }
  defer sftpClient.Close()
  
  srcFile, err := os.Open(localFilePath)
  if err != nil {
    log.Fatal(err)
  }
  defer srcFile.Close()
 
  var remoteFileName = path.Base(remoteName)
  dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
  if err != nil {
    log.Fatal(err)
  }
  defer dstFile.Close()
//  var chunk []byte //新增
 
  buf := make([]byte, 20240)
  for {
    n, _ := srcFile.Read(buf)
    if n == 0 {
      break
    }
//    chunk = append(chunk, buf[:n]...) //新增
    dstFile.Write(buf)
//	dstFile.Write(chunk)
  }
 
  fmt.Println(transmitContent+" over!")
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
  var (
    auth         []ssh.AuthMethod
    addr         string
    clientConfig *ssh.ClientConfig
    sshClient    *ssh.Client
    sftpClient   *sftp.Client
    err          error
  )
  // get auth method
  auth = make([]ssh.AuthMethod, 0)
  auth = append(auth, ssh.Password(password))
 
  clientConfig = &ssh.ClientConfig{
    User:    user,
    Auth:    auth,
    Timeout: 30 * time.Second,
    HostKeyCallback:func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
  }
 
  // connet to ssh
  addr = fmt.Sprintf("%s:%d", host, port)
 
  if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
    return nil, err
  }
 
  // create sftp client
  if sftpClient, err = sftp.NewClient(sshClient); err != nil {
    return nil, err
  }
 
  return sftpClient, nil
}
  
