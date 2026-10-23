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

// The Google sample reads Docs, Sheets and Slides freely and asks before
// changing them.
func TestGoogleReadsAndChanges(t *testing.T) {
	raw, err := os.ReadFile("connectors/google.json")
	if err != nil {
		t.Fatal(err)
	}
	var c connector.Connector
	if err := json.Unmarshal(raw, &c); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ method, path, want string }{
		{"GET", "/v1/documents/DOC1", connector.Read},
		{"POST", "/v1/documents/DOC1:batchUpdate", connector.Write},
		{"GET", "/v4/spreadsheets/SHEET1/values/Sheet1!A1:F50", connector.Read},
		{"POST", "/v4/spreadsheets/SHEET1/values/Sheet1!A1:append", connector.Write},
		{"PUT", "/v4/spreadsheets/SHEET1/values/Sheet1!A1", connector.Write},
		{"GET", "/v1/presentations/PRES1", connector.Read},
		{"POST", "/v1/presentations/PRES1:batchUpdate", connector.Write},
		{"GET", "/drive/v3/files/FILE1/export", connector.Read},
	} {
		if got := c.Classify(tc.method, tc.path, nil); got != tc.want {
			t.Errorf("%s %s = %q, want %q", tc.method, tc.path, got, tc.want)
		}
	}
}
