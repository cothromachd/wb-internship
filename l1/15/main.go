package main

import "fmt"

/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
   Приведите корректный пример реализации.*/

// Глобальную переменную нужно убрать	

// Перевести передачу среза с string на rune, ибо
// по передаче среза из 100 элеметов, передасться 100 байтов, 
// что может быть не равно 100 символов, поэтому операция каста
// в rune обезопасит процесс
func createHugeString(v int) string {
	a := ""

	for i:=0; i<v; i++ {
		a = a + "b"
	}

	return a
}

func someFunc() string {
	v := createHugeString(1 << 10)
	runev := []rune(v)
	return string(runev[:100])
}	

func main() {
   justString := someFunc()
   fmt.Println(justString, len([]rune(justString)))
}