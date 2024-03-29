package functions

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

func Handle(conn net.Conn, mutex *sync.Mutex) {
	timeClient := time.Now().Format("2006-01-02 15:04:05")
	Welcome(conn)
	// получение имени пользователя
	username := GetName(conn)
	// запись в переменную темп структуры пользователя
	tempClient := Client{
		name: username,
		addr: conn.RemoteAddr().String(),
		conn: conn,
	}
	mutex.Lock()
	// запись в мапу пользователя(ключ: имя пользователя, значение: имя пользователя, айпи адрес пользователя, интерфейс конекшена)
	clients[username] = tempClient
	// проверка на количество пользователей
	if len(clients) > 10 {
		mutex.Unlock()
		conn.Write([]byte("Chat is full!"))
		delete(clients, username)
		conn.Close()
		return
	}
	mutex.Unlock()
	mutex.Lock()
	// новое сообщение отправленное на канал джойн(нового подключение)
	join <- NewMessage("has joined our chat...", conn, tempClient, timeClient)
	mutex.Unlock()
	// шаблон с временем и именем пользователя
	fmt.Fprintf(conn, "%v", len(clients))
	fmt.Fprintf(conn, "[%s][%s]:", timeClient, username)
	// переменная принимающая вводимый текст для одного пользователя
	input := bufio.NewScanner(conn)
	// цикл сканера будет идти пока сканер true
	for input.Scan() {
		time := time.Now().Format("2006-01-02 15:04:05")
		// проверка сообщения на пустоту
		if input.Text() == "" {
			fmt.Fprintf(conn, "you can't send empty messages\n")
			fmt.Fprintf(conn, "[%s][%s]:", time, username)
			continue
		}
		// проверка на текста на валидность
		if IsCorrect(input.Text(), conn, time, username) == false {
			fmt.Fprintln(conn, "Incorrect input")
			fmt.Fprintf(conn, "[%s][%s]:", time, username)
			continue
		}
		// запись в переменую время, имя пользователя, введенный текст
		text := fmt.Sprintf("[%s][%s]:%s\n", time, username, input.Text())
		mutex.Lock()
		// запись сообщения в срез строк
		historytext = append(historytext, text)
		// новое сообщение отправленное на канал сообщения
		messages <- NewMessage(input.Text(), conn, tempClient, time)
		mutex.Unlock()
	}
	mutex.Lock()
	// удаление пользователя с мапы
	delete(clients, username)
	leaving <- NewMessage("has left our chat...", conn, tempClient, timeClient)
	conn.Close()
	mutex.Unlock()
}
