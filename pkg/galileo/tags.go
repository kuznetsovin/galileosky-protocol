package galileo

import "fmt"

type tags []tag

type tag struct {
	Tag   uint8       `json:"tag"`
	Value interface{} `json:"value"`
}

func (t *tag) SetValue(tagType string, val []byte) error {
	var v tagParser

	switch tagType {
	case "uint":
		v = &uintTag{}
	case "string":
		v = &stringTag{}
	case "time":
		v = &timeTag{}
	case "coord":
		v = &coordTag{}
	case "speed":
		v = &speedTag{}
	case "int":
		v = &intTag{}
	default:
		return fmt.Errorf("Неизвестный тип данных: %s. Значение: %x", tagType, val)
	}

	if v == nil {
		return fmt.Errorf("Некорректный указатель тэга")
	}

	if err := v.Parse(val); err != nil {
		return err
	}

	t.Value = v

	return nil
}
