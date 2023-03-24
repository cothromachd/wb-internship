package main

import "fmt"

/*
Применимость:
1. Когда вы хотите параметризовать объекты выполняемым
   действием.

2. Когда вы хотите ставить операции в очередь,
   выполнять их по расписанию или передавать по сети.

3. Когда вам нужна операция отмены.

Преимущества и недостатки
+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
+ Позволяет реализовать простую отмену и повтор операций.
+ Позволяет реализовать отложенный запуск операций.
+ Позволяет собирать сложные команды из простых.
+ Реализует принцип открытости/закрытости.

- Усложняет код программы из-за введения множества дополнительных классов.
*/
type Button struct {
    command Command
}

func (b *Button) press() {
    b.command.execute()
}

type Command interface {
    execute()
}


type OnCommand struct {
    device Device
}

func (c *OnCommand) execute() {
    c.device.on()
}

type OffCommand struct {
    device Device
}

func (c *OffCommand) execute() {
    c.device.off()
}

type Device interface {
    on()
    off()
}

type Tv struct {
    isRunning bool
}

func (t *Tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *Tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}

func main() {
    tv := &Tv{}

    onCommand := &OnCommand{
        device: tv,
    }

    offCommand := &OffCommand{
        device: tv,
    }

    onButton := &Button{
        command: onCommand,
    }
    onButton.press()

    offButton := &Button{
        command: offCommand,
    }
    offButton.press()
}