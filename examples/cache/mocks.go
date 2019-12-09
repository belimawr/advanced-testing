package cache

//START-CACHE-MOCK OMIT
type MockCache struct {
	GetFn func(key string) (string, error)
	SetFn func(key, value string) error
}

func (m MockCache) Get(key string) (string, error) {
	return m.GetFn(key)
}

func (m MockCache) Set(key, value string) error {
	return m.SetFn(key, value)
}

//END-CACHE-MOCK OMIT

//START-CACHE-MOCK2 OMIT
type MockCache2 struct {
	GetFn    func(key string) (string, error)
	GetCount int
	SetFn    func(key, value string) error
	SetCount int
}

func (m *MockCache2) Get(key string) (string, error) {
	m.GetCount++
	return m.GetFn(key)
}

func (m *MockCache2) Set(key, value string) error {
	m.SetCount++
	return m.SetFn(key, value)
}

//END-CACHE-MOCK2 OMIT
