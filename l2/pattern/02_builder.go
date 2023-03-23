package pattern

/*
Применимость:
1. Когда нужно избавиться от «телескопического конструктора»,
   который объемляет в себе много параметров.

2. Когда ваш код должен создавать разные представления какого-то объекта. 
   Например, деревянные и железобетонные дома.

3. Когда вам нужно собирать сложные составные объекты.

Преимущества и недостатки:
+ Позволяет создавать продукты пошагово.
+ Позволяет использовать один и тот же код для создания раз- личных продуктов.
+ Изолирует сложный код сборки продукта от его основной бизнес-логики.

- Усложняет код программы из-за введения дополнительных классов.
- Клиент будет привязан к конкретным классам строителей, 
  так как в интерфейсе директора может не быть метода полу- чения результата.
*/
import (
	"fmt"
)

var (
	HpCollectorType   = "hp"
	AsusCollectorType = "asus"
)
// Интерфейс строителя
type Collector interface {
	SetBrand()
	SetCore()
	SetGraphicCard()
	SetMemory()
	GetComputer() Computer
}

type Computer struct {
	Brand       string
	Core        int
	GraphicCard int
	Memory      int
}

func (pc *Computer) Print() {
	fmt.Printf("[%s], Core: [%d], Graphic card: [%d], Memory: [%d]\n", pc.Brand, pc.Core, pc.GraphicCard, pc.Memory)
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	case HpCollectorType:
		return &HpCollector{}
	case AsusCollectorType:
		return &AsusCollector{}
	default:
		return nil
	}
}

// Строитель Hp компьютера
type HpCollector struct {
	Brand       string
	Core        int
	GraphicCard int
	Memory      int
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{Brand: collector.Brand,
		Core:        collector.Core,
		GraphicCard: collector.GraphicCard,
		Memory:      collector.Memory}
}

// Строитель Asus компьютера
type AsusCollector struct {
	Brand       string
	Core        int
	GraphicCard int
	Memory      int
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{Brand: collector.Brand,
		Core:        collector.Core,
		GraphicCard: collector.GraphicCard,
		Memory:      collector.Memory}
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

// CreateComputer - функция-директор, позволяет задавать последовательности команд стройщику
func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetBrand()
	factory.Collector.SetCore()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMemory()
	return factory.Collector.GetComputer()
}

func main() {
	hpCollector := new(HpCollector)
	asusCollector := new(AsusCollector)

	factory := NewFactory(hpCollector)

	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)

	asusComputer := factory.CreateComputer()
	asusComputer.Print()
}
