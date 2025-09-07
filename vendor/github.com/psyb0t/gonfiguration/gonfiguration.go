package gonfiguration

import (
	"reflect"
	"strconv"
	"sync"

	"github.com/pkg/errors"
)

type gonfiguration struct {
	sync.RWMutex
	defaults map[string]interface{}
	envVars  map[string]string
}

func (g *gonfiguration) reset() {
	g.Lock()
	defer g.Unlock()

	gonfig = &gonfiguration{
		defaults: map[string]interface{}{},
		envVars:  map[string]string{},
	}
}

//nolint:gochecknoglobals
var gonfig *gonfiguration

//nolint:gochecknoinits
func init() {
	if gonfig == nil {
		gonfig = &gonfiguration{
			defaults: map[string]interface{}{},
			envVars:  map[string]string{},
		}
	}
}

func Parse(dst interface{}) error {
	envVars, err := getEnvVars()
	if err != nil {
		return errors.Wrap(err, "wtf.. Parse can't get env vars")
	}

	gonfig.setEnvVars(envVars)

	dstVal, err := getDstStructValue(dst)
	if err != nil {
		return errors.Wrap(err, "wtf.. Parse can't get dst struct val")
	}

	if err := parseDstFields(dstVal, envVars); err != nil {
		return errors.Wrap(err, "wtf.. Parse can't parse dst fields")
	}

	return nil
}

func GetAllValues() map[string]interface{} {
	defaults := gonfig.getDefaults()
	envVars := gonfig.getEnvVars()

	allValues := map[string]interface{}{}

	for key, val := range defaults {
		allValues[key] = val
	}

	for key, val := range envVars {
		allValues[key] = val
	}

	return allValues
}

func Reset() {
	gonfig.reset()
}

func parseDstFields(dstVal reflect.Value, envVars map[string]string) error {
	for i := 0; i < dstVal.NumField(); i++ {
		fieldType := dstVal.Type().Field(i)

		tag, ok := fieldType.Tag.Lookup("env")
		if !ok {
			continue
		}

		fieldValue := dstVal.Field(i)
		if !isSupportedType(fieldValue.Kind()) {
			return errors.New("wtf.. Type not supported")
		}

		if err := fillFieldValue(fieldValue, tag, envVars); err != nil {
			return errors.Wrap(err, "wtf.. Parse can't fill field value")
		}
	}

	return nil
}

func fillFieldValue(fieldValue reflect.Value, tag string, envVars map[string]string) error {
	if err := setDefaultValue(fieldValue, tag); err != nil {
		return err // wraps the error inside the function
	}

	envVal, ok := envVars[tag]
	if !ok {
		return nil
	}

	return setEnvVarValue(fieldValue, envVal)
}

func setDefaultValue(fieldValue reflect.Value, tag string) error {
	defaultValue := gonfig.getDefault(tag)
	if defaultValue == nil {
		return nil
	}

	if reflect.TypeOf(defaultValue).Kind() != fieldValue.Kind() {
		return errors.New("wtf.. Default value type mismatch")
	}

	fieldValue.Set(reflect.ValueOf(defaultValue))

	return nil
}

func setEnvVarValue(fieldValue reflect.Value, envVal string) error {
	switch fieldValue.Kind() { //nolint:exhaustive
	case reflect.String:
		fieldValue.SetString(envVal)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(fieldValue, envVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(fieldValue, envVal)
	case reflect.Float32, reflect.Float64:
		return setFloat(fieldValue, envVal)
	case reflect.Bool:
		return setBool(fieldValue, envVal)
	default:
		return errors.Wrapf(
			ErrUnsupportedFieldType,
			"FieldName: %s FieldType %s",
			fieldValue.Type(),
			fieldValue.Kind(),
		)
	}

	return nil
}

func setInt(fieldValue reflect.Value, envVal string) error {
	num, err := strconv.ParseInt(envVal, 10, 64)
	if err != nil {
		return errors.Wrap(err, "wtf.. Failed to parse int")
	}

	fieldValue.SetInt(num)

	return nil
}

func setUint(fieldValue reflect.Value, envVal string) error {
	num, err := strconv.ParseUint(envVal, 10, 64)
	if err != nil {
		return errors.Wrap(err, "wtf.. Failed to parse uint")
	}

	fieldValue.SetUint(num)

	return nil
}

func setFloat(fieldValue reflect.Value, envVal string) error {
	num, err := strconv.ParseFloat(envVal, fieldValue.Type().Bits())
	if err != nil {
		return errors.Wrap(err, "wtf.. Failed to parse float")
	}

	fieldValue.SetFloat(num)

	return nil
}

func setBool(fieldValue reflect.Value, envVal string) error {
	b, err := strconv.ParseBool(envVal)
	if err != nil {
		return errors.Wrap(err, "wtf.. Failed to parse bool")
	}

	fieldValue.SetBool(b)

	return nil
}

func getDstStructValue(dst interface{}) (reflect.Value, error) {
	val := reflect.ValueOf(dst)
	if val.Kind() != reflect.Ptr {
		return reflect.Value{}, ErrTargetNotPointer
	}

	val = val.Elem()
	if val.Kind() != reflect.Struct {
		return reflect.Value{}, ErrDestinationNotStruct
	}

	return val, nil
}

func isSupportedType(kind reflect.Kind) bool {
	switch kind { //nolint:exhaustive
	case reflect.String,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Bool:
		return true
	default:
		return false
	}
}
