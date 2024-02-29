package functions

import (
	"fmt"
	"sync"
)

func Broadcast(mutex *sync.Mutex) {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for _, client := range clients {
				if client.name != msg.name {
					fmt.Fprintf(client.conn, "\n[%s][%s]:%s\n", msg.time, msg.name, msg.text)
				}
				fmt.Fprintf(client.conn, "[%s][%s]:", msg.time, client.name)
			}
			mutex.Unlock()

		case msg := <-leaving:
			mutex.Lock()
			for _, client := range clients {
				if client.name != msg.name {
					fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", msg.name, msg.text, msg.time, client.name)
				}
			}
			mutex.Unlock()
		case msg := <-join:
			mutex.Lock()
			for _, client := range clients {
				// если имя пользователя совпадает с отправленным сообщением имени пользователя то  возвращает старую переписку сообщения
				if client.name == msg.name {
					for _, w := range historytext {
						fmt.Fprintf(client.conn, "%s", "\r")
						fmt.Fprintf(client.conn, "%s[%s][%s]:", w, msg.time, client.name)
					}
				}
				if client.name != msg.name {
					fmt.Fprintf(client.conn, "\n%s %s\n[%s][%s]:", msg.name, msg.text, msg.time, client.name)
				}
			}
			mutex.Unlock()

		}
	}
}
