package card

import "fmt"

func ExampleTranslateMCC() {

	fmt.Println(TranslateMCC("5411"))
	fmt.Println(TranslateMCC("5533"))
	fmt.Println(TranslateMCC("5812"))
	fmt.Println(TranslateMCC("5912"))
	fmt.Println(TranslateMCC("1234"))
	// Output:
	// Супермаркеты
	// Автоуслуги
	// Рестораны
	// Аптеки
	// Категория не указана
}
