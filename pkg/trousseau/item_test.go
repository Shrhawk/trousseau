package trousseau

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keyValuePairUnmarshalTests = map[string]struct {
	in  []byte
	out KeyValuePair
	err error
}{
	"valid json": {
		[]byte(`{"abc":"123"}`),
		KeyValuePair{"abc", "123"},
		nil,
	},
	"multiple keys": {
		[]byte(`{"abc": "123", "easy as": "do re mi"}`),
		KeyValuePair{},
		fmt.Errorf("invalid item data section; reason: invalid number of values"),
	},
	"no keys": {
		[]byte(`{}`),
		KeyValuePair{},
		fmt.Errorf("invalid item data section; reason: invalid number of values"),
	},
}

func TestKeyValuePairUnmarshal(t *testing.T) {
	for name, test := range keyValuePairUnmarshalTests {
		t.Run(name, func(t *testing.T) {
			var kv KeyValuePair
			err := json.Unmarshal(test.in, &kv)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.out, kv)
		})
	}
}
