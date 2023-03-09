package slice

/*Chunk
 * @Author Hex575A
 * @Description lodash _.Chunk
 * @Date 16:21 2023/3/9
 * @Param
 * @return
 */
func Chunk(arr []interface{}, length int) [][]interface{} {
	var arrLength = len(arr)
	if length <= 0 {
		panic("the length is invalid")
	}
	if arrLength == 0 {
		s := make([][]interface{}, 0)
		return s
	}
	// create new slice
	s := make([][]interface{}, 0)
	for i := 0; i < arrLength; i = i + length {
		ss := make([]interface{}, length)
		for j := 0; j < length; j++ {
			if i+j < arrLength {
				ss[j] = arr[i+j]
			}
		}
		s = append(s, ss)
	}
	// deal last slice
	if arrLength%length > 0 {
		lstss := make([]interface{}, 0)
		for i := arrLength - arrLength%length; i < arrLength; i++ {
			lstss = append(lstss, arr[i])
		}
		s[len(s)-1] = lstss
	}
	return s
}

/*Compact
 * @Author Hex575A
 * @Description lodash _.compact
 * @Date 16:22 2023/3/9
 * @Param
 * @return
 */
func Compact(arr []interface{}) []interface{} {
	s := make([]interface{}, 0)
	for _, v := range arr {
		if v != nil && v != 0 && v != "" {
			s = append(s, v)
		}
	}
	return s
}

/*Concat
 * @Author Hex575A
 * @Description lodash _.concat
 * @Date 16:22 2023/3/9
 * @Param
 * @return
 */
func Concat(arr []interface{}, value interface{}) []interface{} {
	return append(arr, value)
}

/*Difference
 * @Author Hex575A
 * @Description lodash _.difference
 * @Date 16:41 2023/3/9
 * @Param
 * @return
 */
func Difference(arr []interface{}, value []interface{}) []interface{} {

	var arrLength = len(arr)
	var valLength = len(value)
	var s = make([]interface{}, 0)
	for i := 0; i < valLength; i++ {
		for j := 0; j < arrLength; j++ {
			if arr[j] == value[i] {
				arr[j] = nil
			}
		}
	}

	for _, v := range arr {
		if v != nil {
			s = append(s, v)
		}
	}
	return s
}

/*Drop
 * @Author Hex575A
 * @Description lodash _.drop
 * @Date 17:04 2023/3/9
 * @Param
 * @return
 */
func Drop(arr []interface{}, args ...int) []interface{} {
	var value = 1
	var arrLength = len(arr)
	if len(args) != 0 {
		value = args[0]
	}
	s := make([]interface{}, 0)
	if value >= arrLength {
		return s
	}
	if value <= 0 {
		return arr
	}
	for i := value; i < arrLength; i++ {
		s = append(s, arr[i])
	}
	return s
}
