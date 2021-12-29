package main

import (
	"cqhttp_reboot/internal"
	"cqhttp_reboot/pkg"
	"fmt"
	"github.com/hpcloud/tail"
	"strings"
)

const (
	logPath = "runtime.log的路径"
	shell = `
			ps -ef | grep go-cqhttp | grep -v grep | awk '{print $2}' | xargs -i -t kill {}
			cd  gocqhttp编译后文件的绝对路径
			rm -f session.token
			nohup ./go-cqhttp > logs/runtime.log 2>&1 &
			`
	email = false
)

func main()  {
	t, err := tail.TailFile(logPath, tail.Config{
		Follow: true,
		Logger: nil,
		ReOpen: false,
		MustExist: true,
	})

	if err != nil {
		fmt.Printf("reboot error, err: %v", err.Error())
		panic(err)
	}

	rs := internal.RebootService{Ch: make(chan *internal.Result), Shell: shell}

	for line := range t.Lines {
		if strings.Contains(line.Text, "账号可能被风控") {
			go rs.Reboot()
			select {
			case result := <-rs.Ch:
				if result.R {
					fmt.Println(result.Text)
				} else {
					fmt.Println(result.Text, err.Error())
					if email {
						e := pkg.EMail{
							User: "发送者邮箱",
							Password: "发送者邮箱密码",
							Host: "发送者host",
							Port: 233,
							MailTo: []string{"接收者邮箱"},
							Alias: "别名",
							Subject: "主题",
							Body: "内容",
						}
						if err := e.Send(); err != nil {
							fmt.Printf("邮件发送错误：%v", err)
						}
					}
				}
			}
		}
	}
}