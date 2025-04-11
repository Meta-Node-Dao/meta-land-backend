package tool

import (
	"ceres/pkg/router"
	"ceres/pkg/utility/jwt"
	"errors"
	"github.com/google/uuid"
	"reflect"
)

func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not found")
}

func SliceIntersection(sli1 []string, sli2 []string) []string {
	var slice []string
	for _, item1 := range sli1 {
		if ok, _ := Contain(item1, sli2); ok {
			slice = append(slice, item1)
		}
	}
	return slice
}

func SliceDiff(source []string, compare []string) []string {
	var slice []string
	for _, item1 := range source {
		if ok, _ := Contain(item1, compare); !ok {
			slice = append(slice, item1)
		}
	}
	return slice
}

func GetComerIDByToken(ctx *router.Context) (uint64, error) {
	header := ctx.Request.Header
	token := header.Get("x-comunion-authorization")
	comerID, err := jwt.Verify(token)
	if err != nil {
		return 0, err
	}
	return comerID, err
}

func RandStr() string {
	uuid := uuid.New()
	key := uuid.String()
	return key
}
