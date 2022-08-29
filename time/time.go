package time

import "time"

/*TimeStr
 * @Author Hex575A
 * @Description 返回字符串类型 2006-01-02 15:04:05 格式时间
 * @Date 16:55 2022/8/29
 * @Param
 * @return string 2006-01-02 15:04:05
 */
func TimeStr() string {
	NowTime := time.Now().Format("2006-01-02 15:04:05")
	return NowTime
}

/*Time
 * @Author Hex575A
 * @Description 返回字符串类型 20060102150405 格式时间
 * @Date 16:55 2022/8/29
 * @Param
 * @return string 20060102150405
 */
func Time() string {
	NowTime := time.Now().Format("20060102150405")
	return NowTime
}

/*Year
 * @Author Hex575A
 * @Description 获取当前年份字符串格式
 * @Date 16:55 2022/8/29
 * @Param
 * @return string 20060102150405
 */
func Year() string {
	NowYear := time.Now().Format("2006")
	return NowYear
}

/*Month
 * @Author Hex575A
 * @Description 获取当前月份字符串格式
 * @Date 16:55 2022/8/29
 * @Param
 * @return string 20060102150405
 */
func Month() string {
	NowMonth := time.Now().Format("01")
	return NowMonth
}

/*Date
 * @Author Hex575A
 * @Description 获取当前年月日字符串格式
 * @Date 16:54 2022/8/29
 * @Param
 * @return string 2006-01-02
 */
func Date() string {
	NowDate := time.Now().Format("2006-01-02")
	return NowDate
}

/*FirstOfMonth
 * @Author Hex575A
 * @Description 获取当月第一天日期
 * @Date 16:23 2022/8/29
 * @Param
 * @return 返回当前月份第一天 string格式
 */
func FirstOfMonth() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	layout := "2006-01-02"
	return firstOfMonth.Format(layout)
}

/*LastOfMonth
 * @Author Hex575A
 * @Description 获取本月最后一天日期
 * @Date 16:24 2022/8/29
 * @Param
 * @return 返回当前月份最后一天 string格式
 */
func LastOfMonth() string {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	layout := "2006-01-02"
	return lastOfMonth.Format(layout)

}

/*CalculateDate
 * @Author Hex575A
 * @Description 以特定日期进行随意的加减运算
 * @Date 16:26 2022/8/29
 * @Param Date time.Time
 * @Param Years int 要加的年数
 * @Param Month int 要加的月数
 * @Param days  int 要加的天数
 * @return
 */
func CalculateDate(Date time.Time, Years int, Months int, days int) string {
	Newday := Date.AddDate(Years, Months, days)
	Newday2 := Newday.Format("2006-01-02")
	return Newday2
}

/*Str2Unix
 * @Author Hex575A
 * @Description 字符串型日期转Unix时间戳
 * @Date 16:26 2022/8/29
 * @Param str 字符串型日期
 * @return int64 时间戳
 */
func Str2Unix(str string) (int64, error) {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, err := time.ParseInLocation(timeLayout, str, loc)
	sr := theTime.Unix()
	return sr, err
}
