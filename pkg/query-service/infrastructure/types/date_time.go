package types

import "time"

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type DateTime time.Time

func Now() DateTime {
	return DateTime(time.Now())
}

func NewDateTime(value *time.Time) *DateTime {
	if value == nil {
		return nil
	}
	t := *value
	res := DateTime(t)
	return &res
}

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
