package yamladapter

import (
	"bytes"
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/ghodss/yaml"
)

// Adapter adapts YAML to Caddy JSON.
type Adapter struct{}

// Adapt converts the YAML config in body to Caddy JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) (result []byte, warnings []caddyconfig.Warning, err error) {
	result, err = yaml.YAMLToJSON(body)
	if err != nil {
		return
	}

	// The JSON returned is not indented, so we can indent the JSON
	result, err = jsonPrettyPrint(result)
	return
}

// jsonPrettyPrint indents JSON with 4 spaces
func jsonPrettyPrint(j []byte) ([]byte, error) {
	out := bytes.NewBuffer([]byte{})
	err := json.Indent(out, j, "", "    ")
	return out.Bytes(), err
}

// Interface guard
var _ caddyconfig.Adapter = (*Adapter)(nil)
