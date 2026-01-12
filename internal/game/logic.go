package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Web-serfer/app/internal/constants"
	"github.com/Web-serfer/app/internal/types"
)

// Запуск основного цикла игры
func RunGame() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	// --- ЭКРАН ПРИВЕТСТВИЯ ---
	fmt.Println(constants.Cyan + "========================================")
	fmt.Println("    ДОБРО ПОЖАЛОВАТЬ В ГЕО-КВИЗ!")
	fmt.Println("========================================" + constants.Reset)
	fmt.Println("Вы готовы проверить свои знания и")
	fmt.Println("отгадать столицы 10 стран?")
	fmt.Println("\n1. Да, поехали!")
	fmt.Println("2. Нет, я еще не готов.")
	fmt.Print("\nВыберите вариант (1 или 2): ")

	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	// Используем switch для обработки выбора в меню
	switch choice {
	case "1", "Да", "да", "д", "y", "yes":
		fmt.Println(constants.Green + "\nОтлично! Начинаем игру...\n" + constants.Reset)
		time.Sleep(1 * time.Second) // Небольшая пауза для эффекта
	case "2", "Нет", "нет", "н", "n", "no":
		fmt.Println(constants.Yellow + "Жаль! Возвращайтесь, когда будете готовы. Пока!" + constants.Reset)
		return // Завершаем программу
	default:
		fmt.Println(constants.Red + "Не совсем понял ответ, но, кажется, вы не готовы. До свидания!" + constants.Reset)
		return
	}

	// Получаем вопросы для игры
	questions := GetQuestions()

	// Используем тип для демонстрации, что импорт нужен
	var _ types.Question

	// Перемешиваем вопросы
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	// --- ОСНОВНОЙ ЦИКЛ ИГРЫ ---
	lives := 3
	score := 0

	for i, q := range questions {
		if lives <= 0 {
			break
		}

		fmt.Printf(constants.Blue+"[Вопрос %d/10]"+constants.Reset+" Назовите столицу %s?\n", i+1, constants.Yellow+q.CountryGenitive+constants.Reset)

		attempt := 0
		for attempt < 4 {
			// Статус
			fmt.Printf("Жизни: %s | Попытка: %d/4\n", strings.Repeat("❤️", lives), attempt+1)
			fmt.Print("Ваш ответ: ")

			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			// Валидация
			if input == "" {
				fmt.Println(constants.Red + "(!) Пожалуйста, введите название. Пустой ответ не принимается." + constants.Reset)
				continue
			}

			// Проверка (через switch для интереса)
			isCorrect := strings.ToLower(input) == strings.ToLower(q.Capital)

			switch {
			case isCorrect:
				fmt.Println(constants.Green + "🌟 Абсолютно верно!" + constants.Reset)
				fmt.Printf(constants.Cyan+"Интересный факт: %s\n\n"+constants.Reset, q.Fact)
				score += (4 - attempt) * 10
				goto nextQuestion // Переходим к следующему вопросу

			case !isCorrect:
				attempt++
				switch attempt {
				case 1, 2, 3:
					fmt.Printf(constants.Red+"❌ Ошибка! Подсказка %d: %s\n"+constants.Reset, attempt, q.Hints[attempt-1])
				case 4:
					fmt.Printf(constants.Red+"💀 Вы не справились. Правильный ответ: %s\n"+constants.Reset, q.Capital)
					lives--
				}
			}
		}
	nextQuestion: // Метка для быстрого перехода
	}

	// --- ФИНАЛ ---
	fmt.Printf("\n" + constants.Cyan + "========================================")
	fmt.Printf("\nИГРА ЗАВЕРШЕНА!")
	fmt.Printf("\nВаш итоговый счет: %d баллов", score)
	fmt.Printf("\n========================================\n" + constants.Reset)
}