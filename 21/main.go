package main

import "fmt"

// Из плюсов:
// 1. Сокрытие от клиента преобразования различных интерфейсов
// 2. Позволяет менять логику не затрагивая вызов
// 3. Возможность переиспользования
// Из минусов:
// Усложнение кода
// Паттерн применяют для использования объектов с несовместимыми интерфейсами

// Интерфейс, который жидает клиент
type AnalyseData interface {
	SendXmlData()
}

// Структура несовместимая с этим интерфейсом
type JsonDoc struct{}

// Структура, которая реализует этот интерфейс
type XmlDoc struct{}

// Структура адаптера, который преобразует JsonDoc к AnalyseData
type JsonDocAdapter struct {
	jsonDoc *JsonDoc
}

func NewJsonAdapter() *JsonDocAdapter {
	return &JsonDocAdapter{
		jsonDoc: &JsonDoc{},
	}
}

// Метод, который несовместим с целевым интерфейсом
func (j *JsonDoc) ConvertToXml() string {
	return "<xml></xml>"
}

// Реализация метода интерфейса
func (x *XmlDoc) SendXmlData() {
	fmt.Println("Отправка XML документа")
}

// Реализация метода целефого интерфейса через Адаптер
func (a *JsonDocAdapter) SendXmlData() {
	a.jsonDoc.ConvertToXml()
	fmt.Println("Отправка XML данных в аналитику")
}

func main() {
	adapter := NewJsonAdapter()
	adapter.SendXmlData() // Клиент не знает о преобразовании форматов
}
