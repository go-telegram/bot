package bot

import "testing"

type TestParams struct {
	ChatID any `json:"chat_id" rules:"required,type:string|int"`
	Age    int `json:"age" rules:"required"`
	Foo    any `json:"foo" rules:"type:string"`
}

func Test_paramsValidate(t *testing.T) {
	tests := []struct {
		name     string
		args     *TestParams
		wantErr  bool
		errValue string
	}{
		{
			name:     "required",
			args:     &TestParams{ChatID: "", Age: 10},
			wantErr:  true,
			errValue: "Foo: is required",
		},
		{
			name:     "not required, bad type",
			args:     &TestParams{ChatID: "foo", Age: 10, Foo: 3},
			wantErr:  true,
			errValue: "Foo: invalid type 'int', expected one of string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := paramsValidate(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errValue {
				t.Errorf("paramsValidate() error = %v, wantErr %v", err, tt.errValue)
			}
		})
	}
}
