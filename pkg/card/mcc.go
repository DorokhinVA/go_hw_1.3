package card

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5812": "Рестораны",
		"5912": "Аптеки",
	}

	value, ok := mcc[code]
	if ok {
		return value
	} else {
		return "Категория не указана"
	}
}
