package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParserService struct{}

func (s *ParserService) SaveCSV(doc *goquery.Document) error {
	// Открываем файл для записи данных
	file, err := os.Create("instagram_top_russia.csv")
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()

	// Создаем файл .csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Создаем шапки столбцов в .csv документе
	writer.Write([]string{"Rank", "Nick", "Name", "Category", "Followers", "Country", "Eng. (Auth.)", "Eng. (Avg.)"})

	// Ищем нужные данные в html документе
	doc.Find(".row .row__top").Each(func(i int, s *goquery.Selection) {
		rank := s.Find(".row-cell.rank span").First().Text()
		nick := s.Find(".row-cell.contributor .contributor__name-content").Text()
		name := s.Find(".row-cell.contributor .contributor__title").Text()

		categories := []string{}
		s.Find(".row-cell.category .tag__content.ellipsis").Each(func(i int, selection *goquery.Selection) {
			categories = append(categories, selection.Text())
		})
		category := strings.Join(categories, ",")

		followers := s.Find(".row-cell.subscribers").Text()
		country := s.Find(".row-cell.audience").Text()
		engAuth := s.Find(".row-cell.authentic").Text()
		engAvg := s.Find(".row-cell.engagement").Text()

		// Сохраняем в .csv
		writer.Write([]string{rank, nick, name, category, followers, country, engAuth, engAvg})
	})

	return nil
}

func NewServiceParser() *ParserService {
	return &ParserService{}
}
