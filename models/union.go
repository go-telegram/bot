package models

import (
	"encoding/json"
	"fmt"
)

// marshalVariant encodes the populated variant of a tagged union such as
// RichBlock, InputRichBlock or RichText.
//
// setTag stamps the discriminator on a copy of the variant rather than on the
// value the caller handed us, so encoding has no side effects and the same value
// can be encoded from several goroutines at once. A variant that was never set is
// reported as an error instead of panicking inside encoding/json.
func marshalVariant[T any, K ~string](union string, tag K, variant *T, setTag func(*T)) ([]byte, error) {
	if variant == nil {
		return nil, fmt.Errorf("nil variant for %s type %q", union, tag)
	}

	v := *variant
	setTag(&v)

	return json.Marshal(&v)
}
