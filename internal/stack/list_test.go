package stack

import (
	"encoding/json"
	"octopod/internal/types"
	"testing"
)

func TestListStacks(t *testing.T) {
	var expected = []types.Service{
		{
			Name:  "testing_whoami_1",
			Image: "containous/whoami:latest@sha256:c0d68a0f9acde95c5214bd057fd3ff1c871b2ef12dae2a9e2d2a3240fdd9214b",
			Labels: map[string]string{
				"com.docker.stack.image":        "containous/whoami",
				"com.docker.stack.namespace":    "testing",
				"io.company.variables.vargroup": "2019.08.09",
			},
			Environment: []string{
				"ENV_Var=value",
			},
		},
		{
			Name:  "testing_whoami_2",
			Image: "containous/whoami:latest@sha256:c0d68a0f9acde95c5214bd057fd3ff1c871b2ef12dae2a9e2d2a3240fdd9214b",
			Labels: map[string]string{
				"com.docker.stack.image":        "containous/whoami",
				"com.docker.stack.namespace":    "testing",
				"io.company.variables.vargroup": "2019.08.09",
			},
			Environment: []string{
				"ENV_Var=value",
			},
		},
	}

	var got, err = ListServices()

	if nil != err {
		t.Errorf("Error: %v", err)
	}

	if nil == got {
		t.Errorf("Expected not nil result.")
	}

	if len(expected) != len(got) {
		t.Errorf("Length not equal expected.")
	}

	var expectedJson, expectedJsonErr = json.MarshalIndent(expected, "", "  ")
	if expectedJsonErr != nil {
		t.Errorf("Failed to prepare expected json.")
	}
	var gotJson, gotJsonErr = ListServicesJson()
	if gotJsonErr != nil {
		t.Errorf("Failed to Get Json list of services.")
	}

	if gotJson != string(expectedJson) {
		t.Errorf("Json output not equal.\n\nExpected:\n%s\n\nActual:\n%s\n", string(expectedJson), gotJson)
	}
}
