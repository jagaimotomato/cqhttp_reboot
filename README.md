##### cqhttp_reboot

###### ✨ 简介
- 遇到风控时自动重启go-cqhttp (适用于centos/linux 平台)
- 原理：监控go-cqhttp控制台输出判断是否重启(go-cqhttp 需使用账号密码登录)

###### 🎬 使用
- 在服务器中修改go-cqhttp中的device.json中的protocol为1，使用命令（nohup ./go-cqhttp > logs/runtime.log 2>&1 &）启动go-cqhttp
- 下载代码，使用go get获取依赖
- 修改main.go中的logPath路径以及shell中的go-cqhttp路径，如果需要邮件通知email改为true，并修改邮件配置
- 编译后把文件丢服务器运行，命令：cd到编译文件所在的目录，使用 nohup ./cqhttp_reboot > logs/reboot.log 2>&1 & 启动
