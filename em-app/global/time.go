package global

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type MyTime time.Time

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(timeTemplates[2], timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(timeTemplates[2]))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(timeTemplates[2]), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
