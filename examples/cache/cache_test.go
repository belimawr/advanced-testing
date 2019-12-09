package cache

import (
	"fmt"
	"testing"
)

type Fooer struct {
	c Cache
}

func (f Fooer) DoFoo(v string) error {
	// some code
	cached, err := f.c.Get("key-" + v)
	// more code

	fmt.Sprint(cached, err)
	return nil
}

func Test_Mock_Example1(t *testing.T) {
	arg := "some string argument"
	//START-CACHE-EX1 OMIT
	mock := MockCache{
		GetFn: func(key string) (string, error) {
			return "bar", nil
		},
	}

	// use the mock
	f := Fooer{mock}

	if err := f.DoFoo(arg); err != nil {
		t.Fatalf("did not expect an error: %q", err)
	}
	//END-CACHE-EX1 OMIT
}

func Test_Mock_Example2(t *testing.T) {
	arg := "some-arg"
	expectedCacheKey := "key-" + arg
	//START-CACHE-EX2 OMIT
	mock := MockCache{
		GetFn: func(key string) (string, error) {
			if key != "key-some-arg" { // HL
				t.Errorf("expecting: %q, got: %q", // HL
					expectedCacheKey, key) // HL
			}
			return "bar", nil
		},
	}

	// use the mock
	f := Fooer{mock}

	if err := f.DoFoo(arg); err != nil {
		t.Fatalf("did not expect an error: %q", err)
	}
	//END-CACHE-EX2 OMIT
}

func Test_Mock_Example3(t *testing.T) {
	arg := "some string argument"
	expectedGetCalls := 1
	//START-CACHE-EX3 OMIT
	mock := &MockCache2{ // HL
		GetFn: func(key string) (string, error) {
			return "bar", nil
		},
	}

	// use the mock
	f := Fooer{mock}

	if err := f.DoFoo(arg); err != nil {
		t.Fatalf("did not expect an error: %q", err)
	}

	if mock.GetCount != expectedGetCalls { // HL
		t.Errorf("expecting Get to be called %d times, instead of: %d times",
			expectedGetCalls, mock.GetCount)
	}
	//END-CACHE-EX3 OMIT
}
