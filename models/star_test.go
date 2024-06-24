package models

import (
	"encoding/json"
	"testing"
)

func TestTransactionPartner_UnmarshalJSON_User(t *testing.T) {
	src := `{"type":"user","user":{"id":42}}`

	m := TransactionPartner{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != TransactionPartnerTypeUser {
		t.Fatalf("invalid type %s, expect %s", m.Type, TransactionPartnerTypeUser)
	}

	if m.User == nil {
		t.Fatal("invalid User")
	}

	if m.User.User.ID != 42 {
		t.Fatalf("invalid id %d, expect 42", m.User.User.ID)
	}
}

func TestTransactionPartner_UnmarshalJSON_Other(t *testing.T) {
	src := `{"type":"other"}`

	m := TransactionPartner{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != TransactionPartnerTypeOther {
		t.Fatalf("invalid type %s, expect %s", m.Type, TransactionPartnerTypeOther)
	}
}

func TestTransactionPartner_UnmarshalJSON_Fragment(t *testing.T) {
	src := `{"type":"fragment","withdrawal_state":{"type":"pending"}}`

	m := TransactionPartner{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != TransactionPartnerTypeFragment {
		t.Fatalf("invalid type %s, expect %s", m.Type, TransactionPartnerTypeFragment)
	}

	if m.Fragment == nil {
		t.Fatal("invalid Fragment")
	}

	if m.Fragment.WithdrawalState == nil {
		t.Fatal("invalid WithdrawalState")
	}

	if m.Fragment.WithdrawalState.Type != RevenueWithdrawalStateTypePending {
		t.Fatalf("invalid type %s, expect %s", m.Fragment.WithdrawalState.Type, RevenueWithdrawalStateTypePending)
	}
}

func TestRevenueWithdrawalState_UnmarshalJSON_Pending(t *testing.T) {
	src := `{"type":"pending"}`

	m := RevenueWithdrawalState{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != RevenueWithdrawalStateTypePending {
		t.Fatalf("invalid type %s, expect %s", m.Type, RevenueWithdrawalStateTypePending)
	}

	if m.Pending == nil {
		t.Fatal("invalid Pending")
	}

	if m.Pending.Type != RevenueWithdrawalStateTypePending {
		t.Fatalf("invalid type %s, expect %s", m.Pending.Type, RevenueWithdrawalStateTypePending)
	}
}

func TestRevenueWithdrawalState_UnmarshalJSON_Succeeded(t *testing.T) {
	src := `{"type":"succeeded","date":1,"url":"xxx"}`

	m := RevenueWithdrawalState{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != RevenueWithdrawalStateTypeSucceeded {
		t.Fatalf("invalid type %s, expect %s", m.Type, RevenueWithdrawalStateTypeSucceeded)
	}

	if m.Succeeded == nil {
		t.Fatal("invalid Succeeded")
	}

	if m.Succeeded.Type != RevenueWithdrawalStateTypeSucceeded {
		t.Fatalf("invalid type %s, expect %s", m.Pending.Type, RevenueWithdrawalStateTypeSucceeded)
	}

	if m.Succeeded.Date != 1 {
		t.Fatalf("invalid date %d, expect 1", m.Succeeded.Date)
	}

	if m.Succeeded.URL != "xxx" {
		t.Fatalf("invalid url %s, expect xxx", m.Succeeded.URL)
	}
}

func TestRevenueWithdrawalState_UnmarshalJSON_Failed(t *testing.T) {
	src := `{"type":"failed"}`

	m := RevenueWithdrawalState{}

	err := json.Unmarshal([]byte(src), &m)
	if err != nil {
		t.Fatal(err)
	}

	if m.Type != RevenueWithdrawalStateTypeFailed {
		t.Fatalf("invalid type %s, expect %s", m.Type, RevenueWithdrawalStateTypeFailed)
	}

	if m.Failed == nil {
		t.Fatal("invalid Failed")
	}

	if m.Failed.Type != RevenueWithdrawalStateTypeFailed {
		t.Fatalf("invalid type %s, expect %s", m.Pending.Type, RevenueWithdrawalStateTypeFailed)
	}
}
