package bot

import (
	"fmt"
	"reflect"
	"strings"
)

func paramsValidate(params any) error {
	// todo: check nil params
	v := reflect.ValueOf(params).Elem()

	var errors []string

	for i := 0; i < v.NumField(); i++ {
		jsonTag := v.Type().Field(i).Tag.Get("json")
		if jsonTag == "-" || jsonTag == "" {
			continue
		}
		omitempty := strings.Contains(jsonTag, ",omitempty")

		if omitempty && v.Field(i).IsZero() {
			continue
		}

		rulesTag := v.Type().Field(i).Tag.Get("rules")
		if rulesTag == "" {
			continue
		}

		rules := strings.Split(rulesTag, ",")

		fieldName := v.Type().Field(i).Name
		value := v.Field(i)

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
				return fmt.Errorf("rule %s not found", ruleName)
			}

			fieldErrors = append(fieldErrors, ruleFunc(value, ruleParams))
		}

		var errorsStrings []string

		for _, e := range fieldErrors {
			if e != "" {
				errorsStrings = append(errorsStrings, e)
			}
		}

		if len(errorsStrings) > 0 {
			errors = append(errors, fmt.Sprintf("%s: %s", fieldName, strings.Join(errorsStrings, ", ")))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}

	return nil
}

type validationRuleFunc func(value reflect.Value, ruleParams []string) string

var validationRules = map[string]validationRuleFunc{
	"required":          paramsValidateRequired,
	"type":              paramsValidateType,
	"required_if_empty": paramsValidateRequiredIfEmpty,
}

// todo: check slice len
// todo: child struct
// todo: equals
// todo: min,max

func paramsValidateRequiredIfEmpty(value reflect.Value, params []string) string {
	// todo:
	return "not implemented"
}

func paramsValidateType(value reflect.Value, ruleParams []string) string {
	if value.IsNil() {
		return ""
	}
	typeName := reflect.TypeOf(value.Interface()).Kind().String()
	for _, t := range ruleParams {
		if typeName == t {
			return ""
		}
	}
	return fmt.Sprintf("invalid type '%s', expected one of %s", typeName, strings.Join(ruleParams, ", "))
}

func paramsValidateRequired(value reflect.Value, _ []string) string {
	if value.IsZero() {
		return "is required"
	}
	return ""
}
