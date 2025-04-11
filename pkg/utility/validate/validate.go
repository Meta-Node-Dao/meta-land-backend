package validate

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	en "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

const (
	FUNCSOURCESELEF  = "self"
	FUNCSOURCEPARENT = "parent"
)

var (
	Default      = validator.New()
	DefaultTrans = func() ut.Translator {
		var found bool
		en := en.New()
		trans, found := ut.New(en, en).GetTranslator("en_US")
		if !found {
			panic(found)
		}
		err := en_translations.RegisterDefaultTranslations(Default, trans)
		if err != nil {
			panic(err)
		}
		return trans
	}()

	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,18}[a-zA-Z0-9]{1}$`)
)

func init() {
	Default.RegisterValidation("address", validateEthAddress)
	Default.RegisterValidation("func", ValidFunc)
	Default.RegisterTranslation("usernane", DefaultTrans, func(ut ut.Translator) error {
		if err := ut.Add("username", "{0} is a invalid field", false); err != nil {
			return err
		}
		return nil
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(fe.Tag(), fe.Field())

		return t
	})
	Default.RegisterValidation(`required_with_eq`, validateRequiredWithEqual)
}

func validateEthAddress(fl validator.FieldLevel) bool {
	return usernameRegex.MatchString(fl.Field().String())
}

func validateRequiredWithEqual(fl validator.FieldLevel) bool {
	param := strings.Split(fl.Param(), `:`)
	paramField := param[0]
	paramValue := param[1]

	if paramField == `` {
		return true
	}

	// param field reflect.Value.
	var paramFieldValue reflect.Value

	if fl.Parent().Kind() == reflect.Ptr {
		paramFieldValue = fl.Parent().Elem().FieldByName(paramField)
	} else {
		paramFieldValue = fl.Parent().FieldByName(paramField)
	}

	if isEq(paramFieldValue, paramValue) == false {
		return true
	}

	return hasValue(fl)
}

/*
1.valid func value must be styled "self.func" or "parent.func"
2.func type must be "func()bool"
*/
func ValidFunc(fl validator.FieldLevel) bool {
	funcString := strings.Split(fl.Param(), ".")
	if len(funcString) != 2 {
		panic(`valid func:value must be styled "self.func" or "parent.func"`)
	}
	var method reflect.Value
	var typeStr string
	switch funcString[0] {
	case FUNCSOURCESELEF:
		method = fl.Field().MethodByName(funcString[1])
		typeStr = fl.Field().Type().String()
	case FUNCSOURCEPARENT:
		method = fl.Parent().MethodByName(funcString[1])
		typeStr = fl.Parent().Type().String()
	default:
		panic(`valid func:value must be styled "self.func" or "parent.func"`)
	}
	if method.Kind() == reflect.Invalid {
		panic(fmt.Sprintf(`valid func:%s of %s does not exit`, funcString[1], typeStr))
	}
	if method.Type().NumIn() != 0 {
		panic(fmt.Sprintf(`valid func:%s of %s must be "func()bool"`, funcString[1], typeStr))
	}
	if method.Type().NumOut() != 1 {
		panic(fmt.Sprintf(`valid func:%s of %s must be "func()bool"`, funcString[1], typeStr))
	}
	if method.Type().Out(0).Kind() != reflect.Bool {
		panic(fmt.Sprintf(`valid func:%s of %s must be "func()bool"`, funcString[1], typeStr))
	}
	return method.Call(nil)[0].Bool()
}

func hasValue(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:

		_, _, nullable := fl.ExtractType(field)
		if nullable && field.Interface() != nil {
			return true
		}
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func isEq(field reflect.Value, value string) bool {
	switch field.Kind() {

	case reflect.String:
		return field.String() == value

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(value)

		return int64(field.Len()) == p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asInt(value)

		return field.Int() == p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(value)

		return field.Uint() == p

	case reflect.Float32, reflect.Float64:
		p := asFloat(value)

		return field.Float() == p
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func asInt(param string) int64 {

	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)
	panicIf(err)

	return i
}

func asFloat(param string) float64 {

	i, err := strconv.ParseFloat(param, 64)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
