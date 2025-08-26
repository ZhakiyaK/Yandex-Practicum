package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

// parseTraining возвращает данные тренировки(Шаги, продолжительность прогулки)
func parseTraining(data string) (int, string, time.Duration, error) {
	str := strings.Split(data, ",")

	if len(str) != 3 {
		return 0, "", 0, fmt.Errorf("Неверный формат ввода")
	}

	steps, err := strconv.Atoi(str[0])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Неверные данные: %v", err)
	}

	dur, err := time.ParseDuration(str[2])

	if err != nil {
		return 0, "", 0, fmt.Errorf("Неверные данные: %v", err)
	}
	return steps, str[1], dur, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(steps int) float64 {
	return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps)

	if dist == 0 {
		return 0
	}

	return dist / float64(duration.Hours())
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func TrainingInfo(data string, weight, height float64) string {
	steps, trainingType, duration, _ := parseTraining(data)

	/*if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}*/

	var result string
	switch trainingType {
	case "Бег":
		calories := RunningSpentCalories(steps, weight, duration)
		result = fmt.Sprintf("Тип тренировки: Бег\nДлительность: %v\nДистанция: %.2f км\nСожгли калорий: %.2f\n",
			duration, distance(steps), calories)
	case "Ходьба":
		calories := WalkingSpentCalories(steps, weight, height, duration)
		result = fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %v\nДистанция: %.2f км\nСожгли калорий: %.2f\n",
			duration, distance(steps), calories)
	default:
		return "Неизвестный тип тренировки\n"
	}

	return result

}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	meanSpeed := meanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * float64(duration.Hours()) * minInH
}
