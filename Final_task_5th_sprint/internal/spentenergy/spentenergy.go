package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight < 0 || height < 0 {
		return 0, fmt.Errorf("Invalid Format")
	}

	if duration < 0 {
		return 0, fmt.Errorf("Invalid Format")
	}

	meanSpeed := MeanSpeed(steps, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH, nil
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight < 0 {
		return 0, fmt.Errorf("Invalid Format")
	}

	if duration < 0 {
		return 0, fmt.Errorf("Invalid Format")
	}

	meanSpeed := MeanSpeed(steps, duration)
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight, nil
}

// МeanSpeed возвращает значение средней скорости движения во время тренировки.
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration < 0 {
		return 0
	}
	distance := Distance(steps)
	hours := duration.Hours()
	return distance / hours
}

// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func Distance(steps int) float64 {
	return float64(steps) * lenStep / kmhInMsec
}
