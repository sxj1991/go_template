package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
	"time"
)

// RemoteSSH 远程连接服务器 返回操作指令集结果/**
func RemoteSSH() string {
	// ssh 客户端配置信息
	config := &ssh.ClientConfig{
		User: "用户名",
		Auth: []ssh.AuthMethod{
			ssh.Password("密码"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 3,
	}
	// 获取client
	client, err := ssh.Dial("tcp", "服务器地址:端口", config)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(client)
	// 获取session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(session)
	// 执行命令
	output, _ := session.CombinedOutput("cd /;ls;pwd")
	return string(output)
}
