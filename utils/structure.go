package utils

import "github.com/mitchellh/mapstructure"

func Decode(source, dest interface{}) error {
	err := mapstructure.Decode(source, dest)
	if err != nil {
		return err
	}
	return nil
}
