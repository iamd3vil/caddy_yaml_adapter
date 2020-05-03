package yamladapter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestYamlAdapter(t *testing.T) {
	yamlConf, err := ioutil.ReadFile("testdata/mock.yml")
	if err != nil {
		t.Fatalf("error while reading mock yaml file")
	}

	jsonConf, err := ioutil.ReadFile("testdata/caddy.json")
	if err != nil {
		t.Fatalf("error while reading mock json file")
	}

	a := Adapter{}

	j, _, err := a.Adapt(yamlConf, map[string]interface{}{})
	if err != nil {
		t.Fatalf("error while parsing yaml config: %v", err)
	}

	if !bytes.Equal(jsonConf, j) {
		t.Fatalf("expected result: %s, got: %s", jsonConf, j)
	}
}
