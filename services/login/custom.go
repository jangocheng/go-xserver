package main

import "github.com/fananchong/go-xserver/common"

func (login *Login) customVerify(account, password string, userdata []byte) (accountID uint64, errcode common.LoginErrCode) {
	return
}
