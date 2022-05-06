/*Package solution
Задана строка S из малых латинских букв,
требуется узнать длину наибольшей подстроки
состоящей не более чем из двух различных символов.
*/
package sulution

func Run(src string) int {
	var max int
	for i := range src {
		if length := LargestLen(src[i:]); length > max {
			max = length
		}
		// если максимальная длина подстроки больше, чем длина хвоста, останавливаемся
		if max >= len(src[i:]) {
			break
		}
	}
	return max
}

func LargestLen(src string) int {
	// подстроки типа "", "a", "ad", длины меньше 3, не нуждаются в дополнительном поиске
	if len(src) < 3 {
		return len(src)
	}
	// массив рун для хранения искомой подстроки
	var sub []rune
	// мапа для хранения использованных символов
	abc := make(map[rune]interface{})
	// проходим по подстроке с начала
	for _, k := range src {
		// если символ уже в алфавите, просто присоединяем его к искомой подстроке
		if _, ok := abc[k]; ok {
			sub = append(sub, k)
			// если символ новый, то смотрим, какой он по счёту, если не 3+, то запоминаем его и присоединяем к подстроке
		} else if !ok && len(abc) < 2 {
			abc[k] = struct{}{}
			sub = append(sub, k)
			// в противном случае возвращаем длину найденной подстроки
		} else {
			return len(sub)
		}
	}
	return len(sub)
}
