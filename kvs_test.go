package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	for _, scenario := range []struct {
		about string
		key   string
		value string
	}{
		{
			about: "add key foo value bar",
			key:   "foo",
			value: "bar",
		},
		{
			about: "add key bar value foo",
			key:   "bar",
			value: "foo",
		},
	} {
		mw := mockWriter{
			storage: "",
		}

		kvs := new(&mw)
		t.Run(scenario.about, func(t *testing.T) {
			if err := kvs.add(scenario.key, scenario.value); err != nil {
				t.Fatalf("Unexpected error when adding value '%s' to key '%s': %s", scenario.value, scenario.key, err.Error())
			}

			if value, exists := kvs.entries[scenario.key]; exists {
				if value != scenario.value {
					t.Fatalf("Expected value '%s' for key '%s', got '%s'", scenario.key, scenario.key, value)
				}
			} else {
				t.Fatalf("Could not find expected entry for key '%s'", scenario.key)
			}

			expectedStored := fmt.Sprintf("%s:%s\n", scenario.key, scenario.value)
			if mw.storage != expectedStored {
				t.Fatalf("Expected '%s' string stored, found '%s'", expectedStored, mw.storage)
			}

		})
	}
}

type mockWriter struct {
	storage string
}

func (mw *mockWriter) Write(p []byte) (int, error) {
	mw.storage += string(p)

	return len(p), nil
}
