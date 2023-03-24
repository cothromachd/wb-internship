package main

import (
	"fmt"
)
/*
Применимость:
1. Когда нужно использовать разные вариации какого-то
алгоритма внутри одного объекта.

2. Когда у есть множество похожих классов, 
   отличающихся только некоторым поведением.

3. Когда нужно не обнажать детали реализации алгоритмов для других классов.

4. Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. 
   Каждая ветка такого оператора представляет собой вариацию алгоритма.

Преимущества и недостатки
+ Горячая замена алгоритмов на лету.
+ Изолирует код и данные алгоритмов от остальных классов.
+ Уход от наследования к делегированию.
+ Реализует принцип открытости/закрытости.
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между страте-
гиями, чтобы выбрать подходящую.
*/

type Strategy interface {
	Route(stP, enP int)
}

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(s Strategy) {
	n.Strategy = s
}

type CarNavigator struct {}

func (cn *CarNavigator) Route(stP, enP int)  {
	fmt.Printf("Navigating %d kilometres by car\n", enP - stP)
}

type BoatNavigator struct{}

func (bn *BoatNavigator) Route(stP, enP int)  {
	fmt.Printf("Navigating %d kilometres by boat\n", enP - stP)
}

func main() {
	nav := Navigator{}
	cn := CarNavigator{}

	nav.SetStrategy(&cn)
	nav.Route(1, 4)

	bn := BoatNavigator{}

	nav.SetStrategy(&bn)

	nav.Route(1, 5)
}
 