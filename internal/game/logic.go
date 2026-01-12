package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Web-serfer/app/internal/constants"
)

// RunGame — запускает основной процесс игры
func RunGame() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	// --- ЭКРАН ПРИВЕТСТВИЯ ---
	fmt.Println(constants.Cyan + "\n╔════════════════════════════════════════╗")
	fmt.Println("║      ДОБРО ПОЖАЛОВАТЬ В ГЕО-КВИЗ!      ║")
	fmt.Println("╚════════════════════════════════════════╝" + constants.Reset)
	fmt.Println("Отгадайте столицы 10 стран. У вас 3 жизни.")
	fmt.Println(constants.Gray + "Подсказка появится автоматически после 2-й ошибки." + constants.Reset)
	fmt.Print("\nГотовы начать? (1 - Да / 2 - Выход): ")

	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	if choice != "1" && !strings.EqualFold(choice, "да") && !strings.EqualFold(choice, "y") {
		fmt.Println(constants.Yellow + "До встречи! Возвращайтесь за победой." + constants.Reset)
		return
	}

	questions := GetQuestions()
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	// --- ПЕРЕМЕННЫЕ СОСТОЯНИЯ ---
	lives := 3
	score := 0

	for i, q := range questions {
		if lives <= 0 {
			fmt.Println(constants.Red + "\n💔 Упс! У вас закончились все жизни..." + constants.Reset)
			break
		}

		fmt.Printf(constants.Blue+"\n---------- [ ВОПРОС %d из %d ] ----------\n"+constants.Reset, i+1, len(questions))
		fmt.Printf("Страна: "+constants.Yellow+"%s\n"+constants.Reset, q.CountryGenitive)

		attempt := 0
		for attempt < 4 {
			// Статус-бар (Жизни и Счет)
			status := fmt.Sprintf(constants.Gray+"[Жизни: %s%s"+constants.Gray+"] [Счет: %d]"+constants.Reset,
				constants.Red+strings.Repeat("❤️", lives), constants.Gray+strings.Repeat("🖤", 3-lives), score)
			fmt.Println(status)

			fmt.Print("Ваш ответ: ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			// 1. Валидация
			if input == "" {
				fmt.Println(constants.Gray + "(!) Введите название города..." + constants.Reset)
				continue
			}

			// 2. Обработка сдачи
			if strings.EqualFold(input, "не знаю") || strings.EqualFold(input, "сдаюсь") {
				fmt.Printf(constants.Yellow+"Очень жаль. Это был город %s.\n"+constants.Reset, q.Capital)
				lives--
				break
			}

			// 3. Проверка ответа
			if strings.EqualFold(input, q.Capital) {
				fmt.Println(constants.Green + "✅ ВЕРНО! " + constants.Reset)
				fmt.Printf(constants.Cyan+"📖 Факт: %s\n"+constants.Reset, q.Fact)

				// Начисление баллов (зависит от попытки)
				score += (4 - attempt) * 10
				break
			} else {
				// Ошибка
				attempt++

				switch attempt {
				case 1:
					fmt.Println(constants.Red + "❌ Не совсем так. Попробуйте еще раз!" + constants.Reset)
				case 2:
					fmt.Println(constants.Red + "❌ Снова мимо." + constants.Reset)
					if len(q.Hints) > 0 {
						fmt.Printf(constants.Yellow+"💡 ПОДСКАЗКА: %s\n"+constants.Reset, q.Hints[0])
					}
				case 3:
					fmt.Println(constants.Red + "❌ Ой-ой! Последняя попытка!" + constants.Reset)
					if len(q.Hints) > 1 {
						fmt.Printf(constants.Yellow+"💡 ВТОРАЯ ПОДСКАЗКА: %s\n"+constants.Reset, q.Hints[1])
					}
				case 4:
					fmt.Printf(constants.Red+"💀 Ошибок слишком много! Правильный ответ: %s\n"+constants.Reset, q.Capital)
					fmt.Printf(constants.Cyan+"📖 Факт: %s\n"+constants.Reset, q.Fact)
					lives--
				}
			}
		}
		fmt.Println(constants.Gray + "------------------------------------------" + constants.Reset)
	}

	// ВЫЗОВ ФУНКЦИИ ОКОНЧАНИЯ ИГРЫ
	printFinalResults(score, lives)
}

// printFinalResults — вспомогательная функция для красивого вывода итогов
func printFinalResults(score int, lives int) {
	fmt.Println(constants.Cyan + "\n╔════════════════════════════════════════╗")
	fmt.Println("║            ИГРА ЗАВЕРШЕНА!             ║")
	fmt.Printf("║       Ваш итоговый счет: %-13d ║\n", score)
	fmt.Println("╚════════════════════════════════════════╝" + constants.Reset)

	if lives > 0 {
		fmt.Println(constants.Green + "Поздравляем! Вы настоящий географ! 🌍" + constants.Reset)
		fmt.Println(constants.Gray + "Вы прошли игру, сохранив жизни. Это крутой результат!" + constants.Reset)
	} else {
		fmt.Println(constants.Red + "💔 Жизни закончились." + constants.Reset)
		fmt.Println(constants.Yellow + "Хорошая попытка! Попробуйте еще раз, чтобы улучшить результат." + constants.Reset)
	}
	fmt.Println() // Пустая строка в конце
}
