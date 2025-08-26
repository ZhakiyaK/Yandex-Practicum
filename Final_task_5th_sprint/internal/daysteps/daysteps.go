package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// DaySteps структура описывает прогулку
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse() функция парсит данные о прогулке
func (ds *DaySteps) Parse(datastring string) (err error) {

	str := strings.Split(datastring, ",")

	if len(str) != 2 {
		return fmt.Errorf("invalid format: expected 2 arguments or wrong format\n")
	}

	ds.Steps, err = strconv.Atoi(str[0])

	if err != nil {
		return fmt.Errorf("invalid steps format: %w\n", err)
	}

	ds.Duration, err = time.ParseDuration(str[1])

	if err != nil {
		return fmt.Errorf("invalid duration format: %w\n", err)
	}

	return nil
}

// ActionInfo() метод возвращает данные с прогулки(кол-во шагов, пройденная дистанция, сколько ккал сожгли)
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration < 0 {
		return "", fmt.Errorf("Invalid duration")
	}

	distance := spentenergy.Distance(ds.Steps)

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км\nВы сожгли %.2f ккал\n", ds.Steps, distance, spentCalories)

	return result, nil
}
