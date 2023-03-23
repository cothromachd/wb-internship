package main

import (
    "fmt"
)
/*
Применимость:

1. Когда нужно представить простой или урезанный
интерфейс к сложной подсистеме.

2. Когда нужно разложить подсистему на отдельные слои.
   Пример:
   Процесс готовки вкуснейшой пиццы в мире имеет сложную систему.
   Мы хотим разбить её на слои работы с тестом, с ингредиентами, и так далее.
   Для каждой из этих частей можно попытаться создать фасад и
   заставить классы теста и ингридиентов общаться друг с другом через эти фасады,
   а не напрямую.

Преимущества и недостатки:

+ Изолирует клиентов от компонентов сложной подсистемы.
- Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
*/
  

// Пример иллюстрирует реализацию фасада над работой компьютера.

type CPU struct{}

func (*CPU) Freeze() {
    fmt.Println("CPU Freeze")
}

func (*CPU) Jump(position int) {
    fmt.Printf("CPU Jump to %d\n", position)
}

func (*CPU) Execute() {
    fmt.Println("CPU Execute")
}

type Memory struct{}

func (*Memory) Load(position int, data string) {
    fmt.Printf("Memory Load data '%s' to position %d\n", data, position)
}

type HardDrive struct{}

func (*HardDrive) Read(position int, size int) string {
    data := fmt.Sprintf("HardDrive Read data from position %d with size %d", position, size)
    fmt.Println(data)
    return data
}

type ComputerFacade struct {
    cpu       *CPU
    memory    *Memory
    hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:       &CPU{},
        memory:    &Memory{},
        hardDrive: &HardDrive{},
    }
}

func (c *ComputerFacade) Start() {
    c.cpu.Freeze()
    c.memory.Load(0, "boot_loader")
    c.cpu.Jump(0)
    c.cpu.Execute()
}

func main() {
    computer := NewComputerFacade()
    computer.Start()
}