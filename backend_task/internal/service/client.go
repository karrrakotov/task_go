package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type ClientService struct{}

func (s *ClientService) SaveOneData(existId bool, data interface{}) error {
	// Открываем файл для записи данных
	file, err := os.Create("client.json")
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Создаем JSON-кодировщик
	encoder := json.NewEncoder(file)

	// Записываем данные в файл в формате JSON
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("ошибка при кодировании JSON: %v", err)
	}

	return nil
}

func NewServiceClient() *ClientService {
	return &ClientService{}
}
