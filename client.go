package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Record представляет запись для отправки на сервер.
type Record struct {
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

const serverURL = "http://localhost:8080"

func main() {
	for {
		fmt.Println("Выберите операцию:")
		fmt.Println("1. Добавить запись")
		fmt.Println("2. Получить записи")
		fmt.Println("3. Обновить запись")
		fmt.Println("4. Удалить запись")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			record := inputRecordData()
			sendRequest("/record/add", record)
		case 2:
			record := getRecordData()
			sendRequest("/records/get", record)
		case 3:
			record := inputRecordData()
			sendRequest("/record/update", record)
		case 4:
			record := getRecordData()
			sendRequest("/record/delete", record)
		default:
			fmt.Println("Неверный выбор операции")
		}

		// Пауза перед следующей итерацией
		time.Sleep(5 * time.Second)
	}
}

func inputRecordData() Record {
	record := Record{}
	fmt.Print("Введите имя: ")
	fmt.Scanln(&record.Name)
	fmt.Print("Введите фамилию: ")
	fmt.Scanln(&record.LastName)
	fmt.Print("Введите отчество: ")
	fmt.Scanln(&record.MiddleName)
	fmt.Print("Введите номер телефона: ")
	fmt.Scanln(&record.Phone)
	fmt.Print("Введите адрес: ")
	fmt.Scanln(&record.Address)
	return record
}

func getRecordData() Record {
	record := Record{}
	fmt.Print("Введите номер телефона: ")
	fmt.Scanln(&record.Phone)
	return record
}

func sendRequest(endpoint string, data interface{}) {
	url := serverURL + endpoint

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Ошибка кодирования JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка отправки запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Статус ответа:", resp.Status)
}
