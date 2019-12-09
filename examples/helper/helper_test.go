package helper

import (
	"encoding/json"
	"os"
	"path"
	"testing"
)

func TestFooBar(t *testing.T) {
	expected := "42, 4.20"
	s := SomeStruct{}
	loadFromJSON2(t, "ok.json", &s)

	got := s.FooBar()
	if got != expected {
		t.Errorf("expecting: %q got: %q", expected, got)
	}
}

func TestFooBarError(t *testing.T) {
	//START OMIT
	expected := "42, 4.20"
	s := SomeStruct{}
	loadFromJSON(t, "ok", &s) // HL

	got := s.FooBar()
	if got != expected {
		t.Errorf("expecting: %q got: %q", expected, got)
	}
	//END OMIT
}

func TestFooBarError2(t *testing.T) {
	//START2 OMIT
	expected := "42, 4.20"
	s := SomeStruct{}
	loadFromJSON2(t, "ok", &s) // HL

	got := s.FooBar()
	if got != expected {
		t.Errorf("expecting: %q got: %q", expected, got)
	}
	//END2 OMIT
}

//START-LoadJSON OMIT
func loadFromJSON(t *testing.T, name string, dst interface{}) {
	filename := path.Join("testdata", name)
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.NewDecoder(file).Decode(dst); err != nil {
		t.Fatal(err)
	}
}

//END-LoadJSON OMIT

//START-LoadJSON2 OMIT
func loadFromJSON2(t *testing.T, name string, dst interface{}) {
	t.Helper() // HL
	filename := path.Join("testdata", name)
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	if err := json.NewDecoder(file).Decode(dst); err != nil {
		t.Fatal(err)
	}
}

//END-LoadJSON2 OMIT
