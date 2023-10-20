package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robfig/cron"

	"backend_task/internal/models"
	"backend_task/internal/service"
	"backend_task/internal/transport"
)

// Создание расписание
var c *cron.Cron = cron.New()

type clientHandler struct {
	serviceClient *service.ClientService
}

func (h *clientHandler) Init(router *http.ServeMux) {
	// TODO? --- GET запросы
	router.HandleFunc("/getCources", h.getCources)
}

// TODO? GET - /getCources
func (h *clientHandler) getCources(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodGet {
		responseError := ResponseError{
			Status:  405,
			IsOk:    false,
			Message: "Метод не разрешен",
		}
		ResponseJson(w, 405, responseError)
		return
	}

	// Получение id и action с входящего запроса
	id := r.URL.Query().Get("id")
	action := r.URL.Query().Get("action")

	// Проверка, парсим определенный id, или все данные
	var existId bool
	if id != "" {
		existId = true
	}

	// Проверяем действие, которое было вызвано для клиента
	if action == "start" {
		// Останавливаем предыдущее расписание, чтобы не перезаписывать документ по ошибке
		c.Stop()
		// Запускаем клиет
		if err := h.StartWorkClient(id, existId); err != nil {
			responseError := ResponseError{
				Status:  500,
				IsOk:    false,
				Message: "Ошибка: " + err.Error(),
			}
			ResponseJson(w, 500, responseError)
			return
		}
	} else if action == "stop" {
		// Останавливаем предыдущее расписание
		c.Stop()
	} else {
		// В случае, если action был передан неверно, возвращаем ошибку
		responseError := ResponseError{
			Status:  400,
			IsOk:    false,
			Message: "Неверный action",
		}
		ResponseJson(w, 400, responseError)
		return
	}

	// Ответ
	var responseOk ResponseOk
	if action == "start" {
		responseOk = ResponseOk{
			Data:    "Клиент успешно запущен!",
			Status:  200,
			IsOk:    true,
			Message: "Success!",
		}
	} else if action == "stop" {
		responseOk = ResponseOk{
			Data:    "Клиент успешно остановлен!",
			Status:  200,
			IsOk:    true,
			Message: "Success!",
		}
	}
	ResponseJson(w, 200, responseOk)
}

func (h *clientHandler) StartWorkClient(id string, existId bool) error {
	// Создание ссылки для запроса
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
	ans, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка при совершении запроса: %v", err.Error())
	}
	defer ans.Body.Close()

	// Проверяем статус ответа с сервера
	if ans.StatusCode != 200 {
		var body map[string]interface{}
		if err := json.NewDecoder(ans.Body).Decode(&body); err != nil {
			return fmt.Errorf("ошибка при декодировании полученных данных: %v", err.Error())
		}

		// Если статус не 200, тогда смотрим причину
		return fmt.Errorf("ошибка: %v", body)
	}

	// Если все ок, тогда делаем декодирование входных данных, полученных с тела запроса
	var clientResponseDTO []models.ClientGetDTO
	if err := json.NewDecoder(ans.Body).Decode(&clientResponseDTO); err != nil {
		return fmt.Errorf("ошибка при декодировании полученных данных: %v", err.Error())
	}

	if existId {
		// Ищем объект с нужным id
		for idx := range clientResponseDTO {
			// Если находим то возвращаем его
			if clientResponseDTO[idx].ID == id {
				if err := h.serviceClient.SaveOneData(existId, clientResponseDTO[idx]); err != nil {
					return err
				}
				return nil
			}
		}

		// Если не находим то возвращаем информацию о том, что искомого объекта нет
		// Ответ
		return fmt.Errorf("искомый объект с id = %v, не найден", id)
	}

	// Обращаемся к логике сервера, чтобы сохранить спаршенные данные в json
	if err := h.serviceClient.SaveOneData(existId, clientResponseDTO); err != nil {
		return err
	}

	// Запускаем расписание cron
	if err := h.Every(id, existId); err != nil {
		return err
	}

	return nil
}

func (h *clientHandler) Every(id string, existId bool) error {
	// Добавляем функцию, которая будет срабатывать каждые 10 минут
	err := c.AddFunc("@every 10m", func() {
		h.StartWorkClient(id, existId)
	})

	if err != nil {
		return fmt.Errorf("ошибка при запуске функции cron")
	}
	c.Start()

	return nil
}

func NewHandlerClient(serviceClient *service.ClientService) transport.ClientHandler {
	return &clientHandler{
		serviceClient: serviceClient,
	}
}
