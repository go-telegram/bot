package bot

import (
	"testing"
)

func Test_paramsValidate_min(t *testing.T) {
	type TestParams struct {
		MinInt    int      `rules:"min:2"`
		MinString string   `rules:"min:2"`
		MinSlice  []string `rules:"min:2"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name:     "nil",
			args:     nil,
			wantErr:  false,
			errValue: "",
		},
		{
			name: "min int",
			args: &TestParams{
				MinInt:    0,
				MinString: "xxx",
				MinSlice:  []string{"", "", ""},
			},
			wantErr:  true,
			errValue: "MinInt: value must be greater than 2",
		},
		{
			name: "min string",
			args: &TestParams{
				MinInt:    3,
				MinString: "x",
				MinSlice:  []string{"", "", ""},
			},
			wantErr:  true,
			errValue: "MinString: string length must be greater than 2",
		},
		{
			name: "min slice",
			args: &TestParams{
				MinInt:    3,
				MinString: "xxx",
				MinSlice:  []string{""},
			},
			wantErr:  true,
			errValue: "MinSlice: slice length must be greater than 2",
		},
		{
			name: "ok",
			args: &TestParams{
				MinInt:    3,
				MinString: "xxx",
				MinSlice:  []string{"", "", ""},
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_max(t *testing.T) {
	type TestParams struct {
		MinInt    int      `rules:"max:2"`
		MinString string   `rules:"max:2"`
		MinSlice  []string `rules:"max:2"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name: "max int",
			args: &TestParams{
				MinInt:    3,
				MinString: "x",
				MinSlice:  []string{""},
			},
			wantErr:  true,
			errValue: "MinInt: value must be less than 2",
		},
		{
			name: "min string",
			args: &TestParams{
				MinInt:    1,
				MinString: "xxx",
				MinSlice:  []string{""},
			},
			wantErr:  true,
			errValue: "MinString: string length must be less than 2",
		},
		{
			name: "min slice",
			args: &TestParams{
				MinInt:    1,
				MinString: "x",
				MinSlice:  []string{"", "", ""},
			},
			wantErr:  true,
			errValue: "MinSlice: slice length must be less than 2",
		},
		{
			name: "ok",
			args: &TestParams{
				MinInt:    1,
				MinString: "x",
				MinSlice:  []string{""},
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_equals(t *testing.T) {
	type TestParams struct {
		Equals string `rules:"equals:foo|bar"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name: "baz",
			args: &TestParams{
				Equals: "baz",
			},
			wantErr:  true,
			errValue: "Equals: value must be one of: foo, bar",
		},
		{
			name: "ok 1",
			args: &TestParams{
				Equals: "foo",
			},
			wantErr:  false,
			errValue: "",
		},
		{
			name: "ok 2",
			args: &TestParams{
				Equals: "bar",
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_chatID(t *testing.T) {
	type TestParams struct {
		CharID any `rules:"chat_id"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name: "bad type",
			args: &TestParams{
				CharID: true,
			},
			wantErr:  true,
			errValue: "CharID: invalid type 'bool', expected one of: string, int",
		},
		{
			name: "string",
			args: &TestParams{
				CharID: "foo",
			},
			wantErr:  false,
			errValue: "",
		},
		{
			name: "int",
			args: &TestParams{
				CharID: 42,
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_required(t *testing.T) {
	type TestParams struct {
		RequiredIface  any    `rules:"required"`
		RequiredString string `rules:"required"`
		RequiredInt    int    `rules:"required"`
		RequiredSlice  []int  `rules:"required"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name: "iface",
			args: &TestParams{
				RequiredIface:  nil,
				RequiredString: "x",
				RequiredInt:    1,
				RequiredSlice:  []int{1},
			},
			wantErr:  true,
			errValue: "RequiredIface: is required",
		},
		{
			name: "string",
			args: &TestParams{
				RequiredIface:  1,
				RequiredString: "",
				RequiredInt:    1,
				RequiredSlice:  []int{1},
			},
			wantErr:  true,
			errValue: "RequiredString: is required",
		},
		{
			name: "int",
			args: &TestParams{
				RequiredIface:  1,
				RequiredString: "x",
				RequiredInt:    0,
				RequiredSlice:  []int{1},
			},
			wantErr:  true,
			errValue: "RequiredInt: is required",
		},
		{
			name: "slice",
			args: &TestParams{
				RequiredIface:  1,
				RequiredString: "x",
				RequiredInt:    1,
				RequiredSlice:  nil,
			},
			wantErr:  true,
			errValue: "RequiredSlice: is required",
		},
		{
			name: "ok",
			args: &TestParams{
				RequiredIface:  1,
				RequiredString: "x",
				RequiredInt:    1,
				RequiredSlice:  []int{1},
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_requiredIfEmpty(t *testing.T) {
	type TestParams struct {
		Field1 string `rules:"required_if_empty:Field2|Field3"`
		Field2 string `rules:"required_if_empty:Field1"`
		Field3 string `rules:"required_if_empty:Field1"`
	}

	tests := []struct {
		name     string
		args     any
		wantErr  bool
		errValue string
	}{
		{
			name: "empty 1,2",
			args: &TestParams{
				Field1: "",
				Field2: "",
				Field3: "x",
			},
			wantErr:  true,
			errValue: "Field1: field is required, because 'Field2' is empty; Field2: field is required, because 'Field1' is empty",
		},
		{
			name: "ok 1",
			args: &TestParams{
				Field1: "",
				Field2: "x",
				Field3: "x",
			},
			wantErr:  false,
			errValue: "",
		},
		{
			name: "ok 2",
			args: &TestParams{
				Field1: "x",
				Field2: "",
				Field3: "",
			},
			wantErr:  false,
			errValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
				return
			}
		})
	}
}

func Test_paramsValidate_no_tag(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1"`
	}{}

	err := paramsValidate(&noTag)
	if err != nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}
}

func Test_paramsValidate_bad_rule(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"unknown"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'unknown' is not found"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_min_value(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"min:a"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'min' parameter must be a number"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_min_params_count_no_params(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"min"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'min' must have one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_min_params_count_two_params(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"min:4|5"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'min' must have one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_min_unsupported_type(t *testing.T) {
	noTag := struct {
		Field1 bool `json:"field_1" rules:"min:3"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'min' is not supported for type 'bool'"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_max_value(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"max:a"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'max' parameter must be a number"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_max_params_count_no_params(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"max"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'max' must have one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_max_params_count_two_params(t *testing.T) {
	noTag := struct {
		Field1 string `json:"field_1" rules:"max:4|5"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'max' must have one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_max_unsupported_type(t *testing.T) {
	noTag := struct {
		Field1 bool `json:"field_1" rules:"max:3"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'max' is not supported for type 'bool'"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_equals_no_params(t *testing.T) {
	noTag := struct {
		Field1 bool `json:"field_1" rules:"equals"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'equals' must have at least one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_bad_required_if_empty_no_params(t *testing.T) {
	noTag := struct {
		Field1 bool `json:"field_1" rules:"required_if_empty"`
	}{}

	err := paramsValidate(&noTag)
	if err == nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}

	wantErr := "Field1: validation rule 'required_if_empty' must have at least one parameter"

	if err.Error() != wantErr {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, wantErr)
		return
	}
}

func Test_paramsValidate_chat_id_empty(t *testing.T) {
	type params struct {
		ChatID          any    `json:"chat_id" rules:"required_if_empty:InlineMessageID,chat_id"`
		MessageID       int    `json:"message_id,omitempty" rules:"required_if_empty:InlineMessageID"`
		InlineMessageID string `json:"inline_message_id,omitempty" rules:"required_if_empty:ChatID|MessageID"`
	}

	err := paramsValidate(&params{
		ChatID:          nil,
		InlineMessageID: "42",
	})
	if err != nil {
		t.Errorf("paramsValidate() error = %v, wantErr %v", err, nil)
		return
	}
}
