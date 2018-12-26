package main

import (
	"fmt"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// modify it to yours
const (
	ACCESSID  = ""
	ACCESSKEY = ""
)

func main() {
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(ACCESSID, ACCESSKEY) // dysms.New(ACCESSID, ACCESSKEY)

	// send to one person
	respSendSms, err := dysms.SendSms(uuid.New(), "xxxxxx", "data", "SMS_143714xxx", `{"code":"DataHunter"}`).DoActionWithException()
	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())
		return
	}
	fmt.Println("send sms succeed", respSendSms.GetRequestID())
}
