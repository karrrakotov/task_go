package rest

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"backend_task/internal/service"
	"backend_task/internal/transport"
)

type parserHandler struct {
	parserService *service.ParserService
}

func (h *parserHandler) Init(router *http.ServeMux) {
	// TODO? --- GET запросы
	router.HandleFunc("/parse", h.Parse)
}

// TODO? GET - /parse
func (h *parserHandler) Parse(w http.ResponseWriter, r *http.Request) {
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

	// Создание ссылки для запроса на сайт
	url := "https://hypeauditor.com/top-instagram-all-russia/"
	ans, err := http.Get(url)
	if err != nil {
		responseError := ResponseError{
			Status:  500,
			IsOk:    false,
			Message: "Ошибка при совершении запроса: " + err.Error(),
		}
		ResponseJson(w, 500, responseError)
		return
	}
	defer ans.Body.Close()

	// Если статус не 200, тогда смотрим причину
	if ans.StatusCode != 200 {
		responseError := ResponseError{
			Status:  ans.StatusCode,
			IsOk:    false,
			Message: fmt.Sprintf("Ошибка: %v", ans.Status),
		}
		ResponseJson(w, ans.StatusCode, responseError)
		return
	}

	// Если все ок, тогда парсим весь html документа
	doc, err := goquery.NewDocumentFromReader(ans.Body)
	if err != nil {
		responseError := ResponseError{
			Status:  500,
			IsOk:    false,
			Message: fmt.Sprintf("Ошибка: %v", err),
		}
		ResponseJson(w, 500, responseError)
		return
	}

	// Обращаемся к логике сервера
	if err := h.parserService.SaveCSV(doc); err != nil {
		responseError := ResponseError{
			Status:  500,
			IsOk:    false,
			Message: fmt.Sprintf("Ошибка: %v", err),
		}
		ResponseJson(w, 500, responseError)
		return
	}

	// Ответ
	responseOk := ResponseOk{
		Data:    "Парсинг данных прошел успешно.",
		Status:  200,
		IsOk:    true,
		Message: "Success!",
	}
	ResponseJson(w, 200, responseOk)
}

func NewHandlerParser(parserService *service.ParserService) transport.ParserHandler {
	return &parserHandler{
		parserService: parserService,
	}
}
