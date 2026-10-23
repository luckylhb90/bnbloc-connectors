package connectors

import (
	"bytes"
	"io"
	"strings"

	"github.com/luckylhb90/bnbloc-connections/connector"
)

func bytesReader(b []byte) io.Reader { return bytes.NewReader(b) }

// firstReadPath is a concrete path under the connector's first read rule.
func firstReadPath(c connector.Connector) string {
	p := c.Policy.Read[0].Paths[0]
	return strings.NewReplacer("**", "x", "*", "x").Replace(p)
}

// firstReadRequest is a method and body the first read rule matches (a
// GraphQL read is a POST of a query).
func firstReadRequest(c connector.Connector) (string, []byte) {
	r := c.Policy.Read[0]
	m := "GET"
	if len(r.Methods) > 0 {
		m = r.Methods[0]
	}
	if r.Body != "" {
		return m, []byte(`{"query":"{ viewer { id } }"}`)
	}
	return m, nil
}
