package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

// RemoteSSH 远程连接服务器 返回操作指令集结果/**
func RemoteSSH() string {
	config := &ssh.ClientConfig{
		User: "用户名",
		Auth: []ssh.AuthMethod{
			ssh.Password("密码"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 3,
	}
	client, err := ssh.Dial("tcp", "服务器地址:端口", config)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	output, _ := session.CombinedOutput("cd /;ls;pwd")
	return string(output)
}
