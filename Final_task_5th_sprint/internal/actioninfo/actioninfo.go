package actioninfo

import (
	"fmt"
)

// DataParser интерфейс в котором объявлены сигнатуры методов Parse() и ActionInfo().
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info() функция, которая принимает слайс строк с данными о тренировках или прогулках и экземпляр одной из структур Training или DaySteps
func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Printf("Ошибка записи %d: %v\n", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()

		if err != nil {
			fmt.Printf("Ошибка записи %d: %s\n", i+1, err)
			continue
		}

		fmt.Println(i+1, info)
	}
}
