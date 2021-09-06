package utils

import "log"

func ProtectRun(entry func(),handle func()) {
	defer func() {
		err := recover()
		//switch err.(type) {
		//case runtime.Error:
		//	fmt.Println("runtime error:", err)
		//default:
		//	fmt.Println("error:", err)
		if err != nil {
			log.Println(err)
			handle()
		}
	}()
	entry()
}