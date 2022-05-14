/*Package sulution
Задана строка S из малых латинских букв,
требуется узнать длину наибольшей подстроки
состоящей не более чем из двух различных символов.
*/
package sulution

func Run(src string) int {
	// строки с 0,1,2 символами пройдут всегда
	if len(src) < 3 {
		return len(src)
	}

	var max, left, right int
	runeSrc := []rune(src)    // приводим для удобства
	var symbols []int         // слайс для индексов первых включений символов
	abc := make(map[rune]int) // мапа для проверки уникальности

	for i := 0; i < len(src); i++ {
		if _, ok := abc[runeSrc[i]]; ok { // уже имеющийся символ
			right = i
		} else if !ok && len(abc) < 2 { // новый символ
			right = i
			abc[runeSrc[i]] = i
			symbols = append(symbols, i)
		} else { // новый символ, но выходящий за словие не более 2
			length := right - left + 1
			if length > max {
				max = length
			}
			// надо перевести каретку на более позднее включение одного из символов
			to := symbols[0]
			if symbols[1] > to {
				to = symbols[1]
			}
			// приводим параметры к новому началу анализа
			i, left, right = to, to, to
			// обнуляем мапу и слайс
			abc = make(map[rune]int)
			abc[runeSrc[i]] = i
			// готовим контейнеры к новой итерации
			symbols = make([]int, 0)
			symbols = append(symbols, i)
		}
	}
	// если прошли всю строку, а она состояла не более чем из двух символов
	if right-left+1 > max {
		return right - left + 1
	}
	return max
}
