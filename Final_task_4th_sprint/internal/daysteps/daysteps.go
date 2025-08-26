package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength       = 0.65                                // длина шага в метрах
	ErrInvalidFormat = errors.New("Неверный формат ввода") // Обработка ошибки
)

// parsePackage возвращает данные дневной активности(Шаги за прогулку и продолжительность прогулки)
func parsePackage(data string) (int, time.Duration, error) {
	str := strings.Split(data, ",")
	if len(str) != 2 {
		return 0, 0, ErrInvalidFormat
	}

	steps, err := strconv.Atoi(str[0])

	if err != nil {
		return 0, 0, ErrInvalidFormat
	}

	dur, err := time.ParseDuration(str[1])

	if err != nil {
		return 0, 0, ErrInvalidFormat
	}
	return steps, dur, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в виде строки в параметре data
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}
	distMeters := float64(steps) * StepLength
	distKm := distMeters / 1000

	cal := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	var result string

	result = fmt.Sprintf("Количество шагов: %v.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distKm, cal)

	return result
}
