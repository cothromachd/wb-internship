package main

import (
	"encoding/json"
	"fmt"
)
/*
Применимость:
1. Когда вам нужно выполнить какую-то операцию над всеми элементами
   сложной структуры объектов, например, деревом. Посетитель позволяет применять 
   одну и ту же операцию к объектам различных классов.

2. Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные 
   между собой операции, но вы не хотите «засорять» классы такими операциями.
   Посетитель позволяет извлечь родственные операции из классов, составляющих структуру объектов, поместив их в один класс-посетитель.
   Если структура объектов является общей для нескольких приложений, то паттерн позволит в каждое приложение включить только нужные операции.

3. Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
   Посетитель позволяет определить поведение только для этих классов, оставив его пустым для всех остальных.

Преимущества и недостатки:
+ Упрощает добавление операций, работающих со сложными
  структурами объектов.
+ Объединяет родственные операции в одном классе.
+ Посетитель может накапливать состояние при обходе структуры элементов.

- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов
*/
type Visitor interface {
	ConvertPD(pd *PoliceDepartment)
	ConvertMD(md *MedicalDepartment)
}

type JSONConverterVisitor struct {
	ConvertedJSON []byte
}

func (JCV *JSONConverterVisitor) ConvertPD(pd *PoliceDepartment) {
	JCV.ConvertedJSON, _ = json.Marshal(pd)
}

func (JCV *JSONConverterVisitor) ConvertMD(md *MedicalDepartment) {
	JCV.ConvertedJSON, _ = json.Marshal(md)
}

type PoliceDepartment struct {
	DepartmentType string
	Officers int
	OpenHours string
}

func (pd *PoliceDepartment) SetOfficers(c int) {
	pd.Officers = c
}

func (pd *PoliceDepartment) SetOpenHours(oh string) {
	pd.OpenHours = oh
}

func (pd *PoliceDepartment) Print() {
	fmt.Printf("[%s], officers count: %d, open hours: %s", pd.DepartmentType, pd.Officers, pd.OpenHours)
}


func (pd *PoliceDepartment) accept(v Visitor) {
	v.ConvertPD(pd)
}

type MedicalDepartment struct {
	DepartmentType string
	Doctors int
	OpenHours string
}

func (md *MedicalDepartment) SetDoctors(c int) {
	md.Doctors = c
}

func (md *MedicalDepartment) SetOpenHours(oh string) {
	md.OpenHours = oh
}

func (md *MedicalDepartment) Print() {
	fmt.Printf("[%s], doctors count: %d, open hours: %s", md.DepartmentType, md.Doctors, md.OpenHours)
}

func (md *MedicalDepartment) accept(v Visitor) {
	v.ConvertMD(md)
}

func main() {
	md := &MedicalDepartment{DepartmentType: "Medical"}
	pd := &PoliceDepartment{DepartmentType: "Police"}

	md.SetDoctors(10)
	md.SetOpenHours("9:00-21:00")

	pd.SetOfficers(50)
	pd.SetOpenHours("00:00-23:59")

	JSONConverter := new(JSONConverterVisitor)

	pd.accept(JSONConverter)
	fmt.Println(string(JSONConverter.ConvertedJSON))

	md.accept(JSONConverter)
	fmt.Println(string(JSONConverter.ConvertedJSON))
}