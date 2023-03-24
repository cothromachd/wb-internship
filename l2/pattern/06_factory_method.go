package main

import (
	"fmt"
)
/*
Применимость:
1. Когда заранее неизвестны типы и зависимости объектов, с
	которыми должен работать ваш код.
2. Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
3. Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

Преимущества и недостатки
+ Избавляет класс от привязки к конкретным классам
  продуктов.
+ Выделяет код производства продуктов в одно место, упрощая поддержку кода.
+ Упрощает добавление новых продуктов в программу.
+ Реализует принцип открытости/закрытости.
- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
*/
type IGun interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

type Gun struct {
    name  string
    power int
}

func (g *Gun) setName(name string) {
    g.name = name
}

func (g *Gun) getName() string {
    return g.name
}

func (g *Gun) setPower(power int) {
    g.power = power
}

func (g *Gun) getPower() int {
    return g.power
}

type Ak47 struct {
    Gun
}

func newAk47() IGun {
    return &Ak47{
        Gun: Gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}

type musket struct {
    Gun
}

func newMusket() IGun {
    return &musket{
        Gun: Gun{
            name:  "Musket gun",
            power: 1,
        },
    }
}

func getGun(gunType string) (IGun, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "musket" {
        return newMusket(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
    ak47, _ := getGun("ak47")
    musket, _ := getGun("musket")

    printDetails(ak47)
    printDetails(musket)
}

func printDetails(g IGun) {
    fmt.Printf("Gun: %s", g.getName())
    fmt.Println()
    fmt.Printf("Power: %d", g.getPower())
    fmt.Println()
}