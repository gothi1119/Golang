package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {
	//	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	k, err := registry.OpenRemoteKey("C:\\Users\\Gothi\\Desktop\\Regiparser_Exam\\config", registry.LOCAL_MACHINE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	s, _, err := k.GetIntegerValue("InstallDate")
	if err != nil {
		log.Fatal(err)
	}
	tm := time.Unix(int64(s), 0)
	fmt.Printf("Windows system root is %v\n", tm)
}
