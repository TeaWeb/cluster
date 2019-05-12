package configs

import "testing"

func TestMapVars(t *testing.T) {
	t.Log(mapVars("a.${id}.${name}.yml", map[string]string{
		"id":   "abc",
		"name": "ABC",
	}))
}
