// +build Config

package Config

type FrequencyType int
const(
	FrequencyType_ONCE FrequencyType = 1+iota
	FrequencyType_DAILY
	FrequencyType_WEEKELY
	FrequencyType_MONTHLY
)
