package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// Training струтура описывает тренировку
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse() метод парсит данные о тренировке и возвращает ошибку парсинга(Если ошибочные данные были)
func (t *Training) Parse(datastring string) (err error) {
	str := strings.Split(datastring, ",")

	if len(str) != 3 {
		return fmt.Errorf("invalid format: expected 3 arguments or wrong format\n")
	}

	t.Steps, err = strconv.Atoi(str[0])

	if err != nil {
		return fmt.Errorf("invalid steps format %w", err)
	}

	t.TrainingType = strings.TrimSpace(str[1])

	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return fmt.Errorf("invalid training type format: %w", err)
	}

	t.Duration, err = time.ParseDuration(str[2])

	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	return nil
}

// ActionInfo() метод возвращает данные о тренировке(тип тренировки, длительность, скорость и сколько ккал сожшли) и ошибку о парсинге
func (t Training) ActionInfo() (string, error) {

	if t.Duration < 0 {
		return "", fmt.Errorf("Invalid duration format")
	}

	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Height, t.Weight, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч.\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return result, nil
}
