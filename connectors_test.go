package connectors

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/luckylhb90/bnbloc-connections/connector"
)

// Every sample is a valid connector, with no field the schema does not
// know, whose first read rule reads.
func TestSamplesAreValidConnectors(t *testing.T) {
	files, _ := filepath.Glob("connectors/*.json")
	if len(files) == 0 {
		t.Fatal("no samples")
	}
	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		var c connector.Connector
		dec := json.NewDecoder(bytesReader(raw))
		dec.DisallowUnknownFields()
		if err := dec.Decode(&c); err != nil {
			t.Fatalf("%s: %v", f, err)
		}
		if err := c.Validate(); err != nil {
			t.Fatalf("%s: %v", f, err)
		}
		if m, body := firstReadRequest(c); c.Classify(m, firstReadPath(c), body) != connector.Read {
			t.Errorf("%s: a request of its first read rule is not a read", f)
		}
		if c.Enabled {
			t.Errorf("%s: a sample is added disabled", f)
		}
	}
}
