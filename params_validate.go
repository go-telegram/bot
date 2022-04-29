package bot

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
Validation rules describes in struct tag 'rules'
format: <rule_name>[:<rule_value>[|<rule_value>|...]]

allowed validation rules:

- chat_id								must be string or int
- required								must be present and not empty
- required_if_empty:<field_name>,...	must be present and not empty if params fields is empty
- min:<value>							support: int, slice, string
- max:<value>							support: int, slice, string
- equals:<value>|<value>|...			support: string

An example:

type User struct {
	Name 	string `json"name" rules:"required,min:3,max:10"`
	Type 	string `json"name" rules:"equals:foo|bar|other"`
}

*/

func paramsValidate(params any) error {
	if params == nil {
		return nil
	}

	v := reflect.ValueOf(params).Elem()

	var errors []string

	for i := 0; i < v.NumField(); i++ {
		rulesTag := v.Type().Field(i).Tag.Get("rules")
		if rulesTag == "" {
			continue
		}

		rules := strings.Split(rulesTag, ",")

		var fieldErrors []string

		for _, rule := range rules {
			ruleParts := strings.Split(rule, ":")
			ruleName := ruleParts[0]
			var ruleParams []string
			if len(ruleParts) > 1 {
				ruleParams = strings.Split(ruleParts[1], "|")
			}

			ruleFunc, ok := validationRules[ruleName]
			if !ok {
				return fmt.Errorf("%s: validation rule '%s' is not found", v.Type().Field(i).Name, ruleName)
			}

			fieldErrors = append(fieldErrors, ruleFunc(v, v.Field(i), ruleParams))
		}

		var errorsStrings []string

		for _, e := range fieldErrors {
			if e != "" {
				errorsStrings = append(errorsStrings, e)
			}
		}

		if len(errorsStrings) > 0 {
			errors = append(errors, fmt.Sprintf("%s: %s", v.Type().Field(i).Name, strings.Join(errorsStrings, ", ")))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}

	return nil
}

type validationRuleFunc func(value, field reflect.Value, ruleParams []string) string

var validationRules = map[string]validationRuleFunc{
	"chat_id":           paramsValidateChatID,
	"required":          paramsValidateRequired,
	"required_if_empty": paramsValidateRequiredIfEmpty,
	"min":               paramsValidateMin,
	"max":               paramsValidateMax,
	"equals":            paramsValidateEquals,
}

// paramsValidateMin validates that value is greater than min
func paramsValidateMin(_, field reflect.Value, params []string) string {
	if len(params) != 1 {
		return "validation rule 'min' must have one parameter"
	}

	min, err := strconv.Atoi(params[0])
	if err != nil {
		return "validation rule 'min' parameter must be a number"
	}

	if field.Kind() == reflect.String {
		if field.Len() < min {
			return fmt.Sprintf("string length must be greater than %d", min)
		}
	} else if field.Kind() == reflect.Int {
		if field.Int() < int64(min) {
			return fmt.Sprintf("value must be greater than %d", min)
		}
	} else if field.Kind() == reflect.Slice {
		if field.Len() < min {
			return fmt.Sprintf("slice length must be greater than %d", min)
		}
	} else {
		return fmt.Sprintf("validation rule 'min' is not supported for type '%s'", field.Kind().String())
	}

	return ""
}

// paramsValidateMax validates that value is less than max
func paramsValidateMax(_, field reflect.Value, params []string) string {
	if len(params) != 1 {
		return "validation rule 'max' must have one parameter"
	}

	max, err := strconv.Atoi(params[0])
	if err != nil {
		return "validation rule 'max' parameter must be a number"
	}

	if field.Kind() == reflect.String {
		if field.Len() > max {
			return fmt.Sprintf("string length must be less than %d", max)
		}
	} else if field.Kind() == reflect.Int {
		if field.Int() > int64(max) {
			return fmt.Sprintf("value must be less than %d", max)
		}
	} else if field.Kind() == reflect.Slice {
		if field.Len() > max {
			return fmt.Sprintf("slice length must be less than %d", max)
		}
	} else {
		return fmt.Sprintf("validation rule 'max' is not supported for type '%s'", field.Kind().String())
	}

	return ""
}

// paramsValidateEquals validates that value is equals to one of the params
func paramsValidateEquals(_, field reflect.Value, params []string) string {
	if len(params) == 0 {
		return "validation rule 'equals' must have at least one parameter"
	}

	for _, param := range params {
		if field.String() == param {
			return ""
		}
	}

	return fmt.Sprintf("value must be one of: %s", strings.Join(params, ", "))
}

// paramsValidateRequiredIfEmpty validates that field is required if params fields is empty
func paramsValidateRequiredIfEmpty(value, field reflect.Value, params []string) string {
	if len(params) < 1 {
		return "validation rule 'required_if_empty' must have at least one parameter"
	}

	fieldIsEmpty := field.IsZero()

	for _, param := range params {
		if value.FieldByName(param).IsZero() && fieldIsEmpty {
			return fmt.Sprintf("field is required, because '%s' is empty", param)
		}
	}

	return ""
}

// paramsValidateChatID validates that field is a string or int
func paramsValidateChatID(_, field reflect.Value, _ []string) string {
	if field.IsNil() || field.IsZero() {
		return ""
	}
	typeName := reflect.TypeOf(field.Interface()).Kind().String()
	for _, t := range []string{"string", "int"} {
		if typeName == t {
			return ""
		}
	}
	return fmt.Sprintf("invalid type '%s', expected one of: string, int", typeName)
}

// paramsValidateRequired validates that field is not nil and not empty
func paramsValidateRequired(_, field reflect.Value, _ []string) string {
	if field.IsZero() {
		return "is required"
	}
	return ""
}
