package utils

type Set struct {
	setMap map[string]SetInterface
}

func NewSet() *Set {
	return &Set{setMap: make(map[string]SetInterface, 0)}
}

func NewSetFromStrings(values []string) *Set {
	set := NewSet()
	for _, value := range values {
		set.AddString(value)
	}
	return set
}

func NewSetWithData(values []SetInterface) *Set {
	set := NewSet()
	for _, value := range values {
		set.Add(value)
	}
	return set
}

func (set Set) Add(value SetInterface) {
	if set.setMap[value.ToString()] == nil {
		set.setMap[value.ToString()] = value
	}
}

func (set Set) GetValue(key string) string {
	v, ok := set.setMap[key]
	if !ok {
		return ""
	}
	return v.ToString()
}

func (set Set) AddString(value string) {
	if set.setMap[value] == nil {
		set.setMap[value] = SetString(value)
	}
}

func (set Set) AddValues(values []SetInterface) {
	for _, value := range values {
		set.Add(value)
	}
}

func (set Set) AddStringValues(values []string) {
	for _, value := range values {
		set.AddString(value)
	}
}

func (set Set) Update(value SetInterface) {
	set.setMap[value.ToString()] = value
}

func (set Set) Remove(value SetInterface) {
	if set.setMap[value.ToString()] != nil {
		delete(set.setMap, value.ToString())
	}
}
func (set Set) Size() int {
	return len(set.setMap)
}

func (set Set) Values() []SetInterface {
	var values []SetInterface = make([]SetInterface, 0)
	for _, value := range set.setMap {
		values = append(values, value)
	}
	return values
}

func (set Set) StringValues() []string {
	values := []string{}
	for _, value := range set.setMap {
		values = append(values, value.ToString())
	}
	return values
}

type SetInterface interface {
	ToString() string
}

type SetString string

func (s SetString) ToString() string {
	return string(s)
}

type SetId struct {
	ID string
}

func (id SetId) ToString() string {
	return id.ID
}

func RemoveDuplicatedObjectIds(ids []string) []string {
	set := NewSet()
	for _, id := range ids {
		set.Add(SetId{id})
	}

	res := make([]string, set.Size())
	for i, v := range set.Values() {
		res[i] = v.(SetId).ToString()
	}
	return res
}
