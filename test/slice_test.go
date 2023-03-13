package main

import (
	"fmt"
	"github.com/skywalkerOAO/Gotos/slice"
	"reflect"
	"testing"
)

/*TestChunk
 * @Author Hex575A
 * @Description lodash _.chunk Test
 * @Date 15:47 2023/3/9
 * @Param
 * @return
 */
func TestChunk(t *testing.T) {

	var tests = []struct {
		name   string
		input  []interface{}
		length int
		except [][]interface{}
	}{
		// the table itself
		{"all numbers be divided with no remainder ", []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, [][]interface{}{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}},
		{"all numbers be divided with remainder ", []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4, [][]interface{}{{1, 2, 3, 4}, {5, 6, 7, 8}, {9}}},
		{"all string be divided with no remainder ", []interface{}{"1", "2", "3", "4"}, 2, [][]interface{}{{"1", "2"}, {"3", "4"}}},
		{"all string be divided with remainder ", []interface{}{"1", "2", "3", "4"}, 3, [][]interface{}{{"1", "2", "3"}, {"4"}}},
		{"random be divided with no remainder ", []interface{}{"1", 2, "3", []interface{}{1, 3, 4, 5}, 5, 6, 7, "8", 9}, 3, [][]interface{}{{"1", 2, "3"}, {[]interface{}{1, 3, 4, 5}, 5, 6}, {7, "8", 9}}},
		{"random be divided with remainder ", []interface{}{"1", 2, "3", []interface{}{1, 3, 4, 5}, 5, 6, 7, "8", 9}, 4, [][]interface{}{{"1", 2, "3", []interface{}{1, 3, 4, 5}}, {5, 6, 7, "8"}, {9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Chunk(tt.input, tt.length)
			if !reflect.DeepEqual(actual, tt.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
}

/*TestCompact
 * @Author Hex575A
 * @Description lodash _.compact Test
 * @Date 16:12 2023/3/9
 * @Param
 * @return
 */
func TestCompact(t *testing.T) {

	var tests = []struct {
		name   string
		input  []interface{}
		except []interface{}
	}{
		// the table itself
		{"1 ", []interface{}{0, 1, "2", 3, nil, 5, false, ""}, []interface{}{1, "2", 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Compact(tt.input)
			if !reflect.DeepEqual(actual, tt.except) {
				t.Error("Not in line with expectations")
			}
		})
	}
}

/*TestConcat
 * @Author Hex575A
 * @Description lodash _.concat Test
 * @Date 16:20 2023/3/9
 * @Param
 * @return
 */
func TestConcat(t *testing.T) {
	var tests = []struct {
		name   string
		input  []interface{}
		input2 interface{}
		except []interface{}
	}{
		// the table itself
		{"concat test 1", []interface{}{1, "2", 3, 5, 6}, "1", []interface{}{1, "2", 3, 5, 6, "1"}},
		{"concat test 2", []interface{}{1, "2", 3, 5, 6}, []interface{}{1, 2, 3, 4, 5, 6}, []interface{}{1, "2", 3, 5, 6, []interface{}{1, 2, 3, 4, 5, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Concat(tt.input, tt.input2)
			if !reflect.DeepEqual(actual, tt.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
}

/*TestDifferent
 * @Author Hex575A
 * @Description lodash _.difference Test
 * @Date 16:52 2023/3/9
 * @Param
 * @return
 */
func TestDifference(t *testing.T) {
	var tests = []struct {
		name   string
		input  []interface{}
		input2 []interface{}
		except []interface{}
	}{
		// the table itself
		{"different test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, []interface{}{1, 3}, []interface{}{"2", 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Difference(tt.input, tt.input2)
			if !reflect.DeepEqual(actual, tt.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
}

/*TestDrop
 * @Author Hex575A
 * @Description lodash _.drop Test
 * @Date 17:17 2023/3/9
 * @Param
 * @return
 */
func TestDrop(t *testing.T) {
	var tests = []struct {
		name   string
		input  []interface{}
		input2 int
		except []interface{}
	}{
		// the table itself
		{"drop test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, 2, []interface{}{3, 1, 3, 1, 5, 6}},
		{"different test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, 3, []interface{}{1, 3, 1, 5, 6}},
	}

	var test2 = []struct {
		name   string
		input  []interface{}
		except []interface{}
	}{
		// the table itself
		{"no value drop test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, []interface{}{"2", 3, 1, 3, 1, 5, 6}},
		{"no value drop test2", []interface{}{1, 1, 3, 4, 5, 6, 7, 8, "2", 3, 1, 3, 1, 5, 6}, []interface{}{1, 3, 4, 5, 6, 7, 8, "2", 3, 1, 3, 1, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.Drop(tt.input, tt.input2)
			if !reflect.DeepEqual(actual, tt.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
	for _, tt2 := range test2 {
		t.Run(tt2.name, func(t *testing.T) {
			actual := slice.Drop(tt2.input)
			if !reflect.DeepEqual(actual, tt2.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
}

/*TestDropRight
 * @Author Hex575A
 * @Description lodash _.dropRight Test
 * @Date 9:19 2023/3/10
 * @Param
 * @return
 */
func TestDropRight(t *testing.T) {
	var tests = []struct {
		name   string
		input  []interface{}
		input2 int
		except []interface{}
	}{
		// the table itself
		{"drop test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, 2, []interface{}{1, "2", 3, 1, 3, 1}},
		{"different test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, 3, []interface{}{1, "2", 3, 1, 3}},
	}

	var test2 = []struct {
		name   string
		input  []interface{}
		except []interface{}
	}{
		// the table itself
		{"no value drop test", []interface{}{1, "2", 3, 1, 3, 1, 5, 6}, []interface{}{1, "2", 3, 1, 3, 1, 5}},
		{"no value drop test2", []interface{}{1, 1, 3, 4, 5, 6, 7, 8, "2", 3, 1, 3, 1, 5, 6}, []interface{}{1, 1, 3, 4, 5, 6, 7, 8, "2", 3, 1, 3, 1, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := slice.DropRight(tt.input, tt.input2)
			if !reflect.DeepEqual(actual, tt.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
	for _, tt2 := range test2 {
		t.Run(tt2.name, func(t *testing.T) {
			actual := slice.DropRight(tt2.input)
			if !reflect.DeepEqual(actual, tt2.except) {
				fmt.Println(actual)
				t.Error("Not in line with expectations")
			}
		})
	}
}
