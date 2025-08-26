package personaldata

import "fmt"

// Personal структура которая хранит данные о человеке
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print() метод ввыводит данные о человеке
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
