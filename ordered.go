package main

type OrderedInt struct{}

func NewOrderedInt() OrderedInt {
	return OrderedInt{}
}

func (o OrderedInt) IsValidType(data interface{}) bool {
	switch data.(type) {
	case int:
		return true
	default:
		return false
	}
}

func (o OrderedInt) IsLeft(s, t interface{}) bool {
	return s.(int) > t.(int)
}

func (o OrderedInt) IsEqual(s, t interface{}) bool {
	return s.(int) == t.(int)
}
