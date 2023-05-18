// Package helper wrapper for mapper https://github.com/dranikpg/dto-mapper

package mapper

import (
	"github.com/dranikpg/dto-mapper"
	"reflect"
)

func Mapper(to interface{}, from interface{}) error {
	return dto.Map(to, from)
}

func MapAndAssign(source interface{}, target interface{}) (interface{}, error) {
	responseDto := reflect.New(reflect.TypeOf(target).Elem()).Interface()

	if err := Mapper(responseDto, source); err != nil {
		return nil, err
	}

	return responseDto, nil
}
