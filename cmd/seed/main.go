package main

import (
	"log"

	"gorm.io/datatypes"

	"github.com/Rodco-off/checkCode/internal/config"
	"github.com/Rodco-off/checkCode/internal/db"
	"github.com/Rodco-off/checkCode/internal/domain"
)

func main() {
	cfg := config.Load()
	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}

	tasks := []domain.Task{
		{
			Title:       "Сумма двух чисел",
			Description: "Напишите функцию Solve(input string) string, которая принимает два числа через пробел и возвращает их сумму.",
			Difficulty:  "easy",
			Constraints: "1 <= a, b <= 1000",
			TestCases:   datatypes.JSON([]byte(`[{"input": "2 3", "output": "5"}, {"input": "-1 5", "output": "4"}, {"input": "0 0", "output": "0"}]`)),
		},
		{
			Title:       "Реверс строки",
			Description: "Напишите функцию Solve(input string) string, которая принимает строку и возвращает её в обратном порядке. Используйте пакет strings или цикл.",
			Difficulty:  "easy",
			Constraints: "Длина строки не превышает 100 символов.",
			TestCases:   datatypes.JSON([]byte(`[{"input": "hello", "output": "olleh"}, {"input": "Go", "output": "oG"}, {"input": "12345", "output": "54321"}]`)),
		},
		{
			Title:       "Палиндром",
			Description: "Напишите функцию Solve(input string) string, которая проверяет, является ли строка палиндромом (читается одинаково слева направо и справа налево). Верните \"true\" или \"false\".",
			Difficulty:  "medium",
			Constraints: "Строка состоит только из строчных латинских букв.",
			TestCases:   datatypes.JSON([]byte(`[{"input": "racecar", "output": "true"}, {"input": "hello", "output": "false"}, {"input": "a", "output": "true"}]`)),
		},
		{
			Title:       "Максимальное число",
			Description: "Напишите функцию Solve(input string) string, которая принимает числа через пробел и возвращает максимальное из них.",
			Difficulty:  "medium",
			Constraints: "Количество чисел от 1 до 100.",
			TestCases:   datatypes.JSON([]byte(`[{"input": "10 5 20 8", "output": "20"}, {"input": "-5 -1 -10", "output": "-1"}, {"input": "42", "output": "42"}]`)),
		},
		{
			Title:       "Четное или нечетное",
			Description: "Напишите функцию Solve(input string) string, которая принимает одно число и возвращает \"even\", если оно четное, и \"odd\", если нечетное.",
			Difficulty:  "hard",
			Constraints: "Число от -10000 до 10000.",
			TestCases:   datatypes.JSON([]byte(`[{"input": "4", "output": "even"}, {"input": "7", "output": "odd"}, {"input": "0", "output": "even"}]`)),
		},
	}

	for _, task := range tasks {
		var count int64
		database.Model(&domain.Task{}).Where("title = ?", task.Title).Count(&count)
		if count == 0 {
			if err := database.Create(&task).Error; err != nil {
				log.Printf("Ошибка добавления '%s': %v", task.Title, err)
			} else {
				log.Printf("Успешно добавлена: %s", task.Title)
			}
		}
	}

	log.Println("Готово!")
}
