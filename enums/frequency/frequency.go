package frequency

type Type uint8

const (
	DAILY Type = iota + 1
	WEEKLY
	BIWEEKLY
	MONTHLY
	QUARTERLY
	ANNUALLY
)

// TODO: check if this assumption is ok
var toValue = map[Type]int{
	DAILY:     365,
	WEEKLY:    52,
	BIWEEKLY:  26,
	MONTHLY:   12,
	QUARTERLY: 4,
	ANNUALLY:  1,
}

func (t *Type) Value() int {
	return toValue[*t]
}
