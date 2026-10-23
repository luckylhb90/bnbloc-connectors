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
