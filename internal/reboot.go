package internal

import "os/exec"

/**
*  @Author: tkykc
*  @github: https://github.com/takoyakiccc
*  @date: 2021/12/29
 */

type RebootService struct {
	Ch chan *Result
	Shell string
}

type Result struct {
	R bool
	Text string
	Err error
}


func (r *RebootService) Reboot()  {
	cmd := exec.Command("/bin/bash", "-c", r.Shell)
	if err := cmd.Run(); err != nil {
		r.Ch <- &Result{
			R: false,
			Text: "执行shell失败",
			Err: err,
		}
	} else {
		r.Ch <- &Result{
			R: true,
			Text: "执行shell成功",
			Err: nil,
		}
	}
}