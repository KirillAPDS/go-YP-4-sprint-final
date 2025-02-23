package daysteps

import (
	"time"
	"fmt"
	"strings"
	"strconv"
	"errors"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {

	var (
		steps int
		d     time.Duration
	)

	s := strings.Split(data, ",")

	switch len(s) {
	case 2: {
			step, err := strconv.Atoi(s[0])

			if err != nil {
				return 0, 0, err
			}
			steps = step

			duration, err := time.ParseDuration(s[1])

			if err != nil {
				return 0, 0, err
			}
			d = duration
		}

	default: {
			err := errors.New("Invalid arguments count")
			return 0, 0, err
		}
	}

	return steps, d, nil
}


// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	/*Функция должна парсить строку с данными с помощью parsePackage(), 
	вычислять дистанцию в километрах и количество 
	потраченных калорий и возвращать строку в таком виде: 

	Количество шагов: 792.
	Дистанция составила 0.51 км.
	Вы сожгли 221.33 ккал. */
	steps, duration, err := parsePackage(data)
	if err != nil {
		errorBytes := []byte(fmt.Sprintf("%v\n", err))
		return string(errorBytes)
	}
	dist := float64(steps) * StepLength / 1000
	//speed := dist / duration.Hours()
	//ccal := 0.035 * weight + (speed * speed / height) * 0.029 * weight
	ccal := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, ccal)
}
