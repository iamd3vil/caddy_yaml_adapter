package yamladapter

import (
	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/ghodss/yaml"
)

func init() {
	caddyconfig.RegisterAdapter("yaml", Adapter{})
}

// Adapter adapts YAML to Caddy JSON.
type Adapter struct{}

// Adapt converts the YAML config in body to Caddy JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) (result []byte, warnings []caddyconfig.Warning, err error) {
	result, err = yaml.YAMLToJSON(body)
	return
}
