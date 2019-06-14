package main

import "fmt"

type tags []tag

type tag struct {
	Tag   uint8
	Value interface{}
}

func (t *tag) SetValue(tagType string, val []byte) error {
	var v tagParser

	switch tagType {
	case "uint":
		v = new(uintTag)
	case "string":
		v = new(stringTag)
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
