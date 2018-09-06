package main

import (
	"fmt"

	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/tobyzxj/uuid"
)

// modify it to yours
const (
	ACCESSID  = "LTAI8CLg4899HZgu"
	ACCESSKEY = "AFGyVIFCaILfa4Q3u9uX8kmBnS9QvV"
)

func main() {
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(ACCESSID, ACCESSKEY) // dysms.New(ACCESSID, ACCESSKEY)

	// send to one person
	respSendSms, err := dysms.SendSms(uuid.New(), "13161332258", "数猎天下", "SMS_143714198", `{"code":"DataHunter"}`).DoActionWithException()
	if err != nil {
		fmt.Println("send sms failed", err, respSendSms.Error())
		return
	}
	fmt.Println("send sms succeed", respSendSms.GetRequestID())
}
