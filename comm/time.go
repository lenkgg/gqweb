package comm

import "time"

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"

)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}



func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}


func(t Time) Unix()int64 {
	return time.Time(t).Unix()
}

func(t Time) Sub(tm time.Time)time.Duration{
	return t.Sub(tm)
}

func(t Time) Add(tm time.Duration)Time {
	return Time(time.Time(t).Add(tm))
}

func(t Time) Before(tm Time)bool{
	return time.Time(t).Before(time.Time(tm))
}

func(t Time) After(tm Time)bool{
	return time.Time(t).After(time.Time(tm))
}

func(t Time) ToTime() time.Time{
	return time.Time(t)
}

func (t Time) AddDate(y int,m int, d int) Time{
	return Time(time.Time(t).AddDate(y,m,d))
}

func Forever() (Time){
	now,_ := time.ParseInLocation(timeFormart, "2050-01-01 00:00:00", time.Local)
	return Time(now)
}

func NewTime(str string)(Time){
	tt, _ := time.ParseInLocation(timeFormart, str, time.Local)
	var t = Time(tt)
	return t
}

func Now() (Time){
	var tt = Time(time.Now())
	return tt
}