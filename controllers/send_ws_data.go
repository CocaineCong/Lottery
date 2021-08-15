package controllers

func init() {
	go SendWsData()
}

func SendWsData() {
	for { 			//保证一次协程调用完成
		accounts := <- accountsChan
		for client := range clients {
			err := client.WriteJSON(accounts)
			if err != nil {
				_ = client.Close()
				delete(clients,client)
			}
		}
	}
}
