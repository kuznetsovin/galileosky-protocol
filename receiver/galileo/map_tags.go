package galileo

type tagDesc struct {
	Len  uint
	Type string
}

var tagsTable = map[byte]tagDesc{
	// версия железа
	0x01: {1, "uint"},
	// версия прошивки
	0x02: {1, "uint"},
	// IMEI
	0x03: {15, "string"},
	// идентификатор устройства
	0x04: {2, "uint"},
	// номер записи в архиве
	0x10: {2, "uint"},
	// Дата и время
	0x20: {4, "time"},
	// Координаты в градусах, число спутников,
	// признак корректности определения координат и
	// источник координат
	0x30: {9, "coord"},
	// Скорость в км/ч направлене в градусах
	0x33: {4, "speed"},
	// высота, м.
	0x34: {2, "int"},
	// Одно из значений: 1. HDOP (делить на 10) - если истоник координат GPS
	// модуль, 2 погрешность в метрах если источник gsm-сети (умножить на 10)
	0x35: {1, "uint"},
	// Статус устройства
	0x40: {2, "bitstring"},
	// Напряжение питания, мВ
	0x41: {2, "uint"},
	// Напряжение аккумулятора, мВ
	0x42: {2, "uint"},
	// Статусы входов
	0x45: {2, "bitstring"},
	// Статусы выходов
	0x46: {2, "bitstring"},
	// Значение на входе 0.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x50: {2, "uint"},
	// Значение на входе 1.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x51: {2, "uint"},
	// Значение на входе 2.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x52: {2, "uint"},
	// Значение на входе 3.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x53: {2, "uint"},
	// Значение на входе 4.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x54: {2, "uint"},
	// Значение на входе 5.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x55: {2, "uint"},
	// Значение на входе 6.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x56: {2, "uint"},
	// Значение на входе 7.
	// В зависимости от настроек один из вариантов: напряжение,
	// число импульсов, частота Гц
	0x57: {2, "uint"},
	// RS485[0] ДУТ с адресом 0
	0x60: {2, "uint"},
	// RS485[1] ДУТ с адресом 1
	0x61: {2, "uint"},
	// RS485[2] ДУТ с адресом 2
	0x62: {2, "uint"},
	//// RS485[3] ДУТ с адресом 2
	//0x63: {2, "uint"},
	//// RS485[4] ДУТ с адресом 2
	//0x64: {2, "uint"},
}
