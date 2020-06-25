package main

import "time"

type MyTime time.Time

// const timeFormat = "2016-01-02 15:04:05"

// func (t *MyTime) UnmarshalJSON(data []byte) (err error) {
// 	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
// 	*t = MyTime(now)
// 	return err
// }

// func (t *MyTime) MarshalJSON() ([]byte, error) {
// 	b := make([]byte, 0, len(timeFormat)+2)
// 	b = append(b, '"')
// 	b = time.Time(t).AppendFormat(b, timeFormat)
// 	b = append(b, '"')
// 	return b, nil
// }

// func (t *MyTime) String() string {
// 	return time.Time(t).Format(timeFormat)
// }
