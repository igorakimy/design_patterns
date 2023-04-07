/*
	Цепочка обязанностей - это поведенческий паттерн, позволяющий передавать запрос
	по цепочке потенциальных обработчиков, пока один из них не обработает запрос.

	Избавляет от жесткой привязки отправителя запроса к его получателю, позволяя
	выстраивать цепь из различных обработчиков динамически.
*/
package main

func main() {
	cashier := &Cashier{}

	// Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	// Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	// Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	// Patient visiting
	reception.execute(patient)
}

// Reception registering patient
// Doctor checking patient
// Medical giving medicine to patient
// Cashier getting money from patient
