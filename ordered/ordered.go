package ordered

type Int struct{}

func NewOrderedInt() Int {
	return Int{}
}

func (o Int) IsValidType(data interface{}) bool {
	switch data.(type) {
	case int:
		return true
	default:
		return false
	}
}

func (o Int) IsLeft(s, t interface{}) bool {
	return s.(int) > t.(int)
}

func (o Int) IsEqual(s, t interface{}) bool {
	return s.(int) == t.(int)
}

type String struct{}

func NewOrderedString() String {
	return String{}
}

func (o String) IsValidType(data interface{}) bool {
	switch data.(type) {
	case string:
		return true
	default:
		return false
	}
}

func (o String) IsLeft(s, t interface{}) bool {
	return s.(string) > t.(string)
}

func (o String) IsEqual(s, t interface{}) bool {
	return s.(string) == t.(string)
}
