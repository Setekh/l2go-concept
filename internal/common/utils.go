package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPublicIp() string {
	resp, _ := http.Get("https://api.ipify.org")
	ip, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println("Failed response from ipify, using localhost")
		return "127.0.0.1"
	}

	return string(ip)
}
