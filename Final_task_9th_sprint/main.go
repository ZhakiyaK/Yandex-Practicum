package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements функция возвращает слайс в котором был сгенерирован рандомное число
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	randomElement := rand.NewSource(time.Now().UnixNano())
	randomRange := rand.New(randomElement)
	data := make([]int, size)

	for i := range data {
		data[i] = randomRange.Int()
	}

	return data
}

// maximum фунция, которая возвращает максимальный элемент в слайсе
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	return slices.Max(data)

	/*
		maxValue := data[0]
		for i := 0;i < len(data); i++ {
			if data[i] > maxValue {
				maxValue = data[i]
			}
		}

		return maxValue
	*/
}

// maxChunks фунция возвращает максимальное значение слайса
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}
	var wg sync.WaitGroup
	maxSlices := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		firstIdx := i * chunkSize
		lastIdx := firstIdx + chunkSize

		if i == CHUNKS-1 {
			lastIdx = len(data)
		}

		chunk := data[firstIdx:lastIdx]

		go func(index int, c []int) {
			defer wg.Done()
			maxSlices[index] = maximum(chunk)

		}(i, chunk)

	}
	wg.Wait()

	maxVal := maximum(maxSlices)
	return maxVal
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	generated := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(generated)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(generated)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
