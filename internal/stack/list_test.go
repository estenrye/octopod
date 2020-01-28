package stack

import(
	//"octopod/internal/types"
	"octopod/internal/types"
	"sort"
	"testing"
)

func TestListStacks(t *testing.T) {
	var expected = []types.Service{
		{
			Name:   "testing_whoami_1",
			Labels: []types.KeyValue{
				{
					Name:  "com.docker.stack.image",
					Value: "containous/whoami",
				},
				{
					Name:  "com.docker.stack.namespace",
					Value: "testing",
				},
				{
					Name:  "io.company.variables.vargroup",
					Value: "2019.08.09",
				},
			},
		},
		{
			Name:   "testing_whoami_2",
			Labels: []types.KeyValue{
				{
					Name:  "com.docker.stack.image",
					Value: "containous/whoami",
				},
				{
					Name:  "com.docker.stack.namespace",
					Value: "testing",
				},
				{
					Name:  "io.company.variables.vargroup",
					Value: "2019.08.09",
				},
			},
		},
	}

	var got, err = ListServices("testing")

	if nil != err {
		t.Errorf("Error: %v", err)
	}

	if nil == got {
		t.Errorf("Expected not nil result.")
	}

	if len(expected) != len(got) {
		t.Errorf("Length not equal expected.")
	}

	if len(got) > 0 && len(expected[0].Labels) != len(got[0].Labels) {
		t.Errorf("Labels length for service 0 not equal expected.")
	}

	sort.Sort(types.OrderByKeyValueName(expected[0].Labels))
	sort.Sort(types.OrderByKeyValueName(got[0].Labels))

	for i, label := range got[0].Labels {
		if expected[0].Labels[i].Name != label.Name {
			t.Errorf("Label Name `%s` not expected.  Expected `%s`", label.Name, expected[0].Labels[i].Name)
		}
		if expected[0].Labels[i].Value != label.Value {
			t.Errorf("Label Value `%s` not expected.  Expected `%s`", label.Value, expected[0].Labels[i].Value)
		}
	}

	sort.Sort(types.OrderByKeyValueName(expected[1].Labels))
	sort.Sort(types.OrderByKeyValueName(got[1].Labels))

	if len(got) > 1 && len(expected[1].Labels) != len(got[1].Labels) {
		t.Errorf("Labels length for service 1 not equal expected.")
	}

	for i, label := range got[1].Labels {
		if expected[1].Labels[i].Name != label.Name {
			t.Errorf("Label Name `%s` not expected.  Expected `%s`", label.Name, expected[1].Labels[i].Name)
		}
		if expected[0].Labels[i].Value != label.Value {
			t.Errorf("Label Value `%s` not expected.  Expected `%s`", label.Value, expected[1].Labels[i].Value)
		}
	}
}