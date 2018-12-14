package util

import (
	"encoding/json"
	"github.com/segmentio/ksuid"
	"math/rand"
	"strings"
	"time"
)

/**
*截取字符串
 */
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

/**
* 获取当前时区时间
 */
func GetNowTime() time.Time {
	l, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(l)
}

/**
* 获取一个uuid
 */
func GetUUID() string {
	return ksuid.New().String()
}

/**
* 取随机数
 */
func Random(ranLength int) int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randIndex := r.Intn(ranLength)
	return randIndex
}

/**
* 过滤数组
 */
func FilterArray(originArray []string, filterArray []string) []string {
	var newArray []string
	filterStr := ArrayToString(filterArray)
	for _, value := range originArray {
		if !strings.Contains(filterStr, value) {
			newArray = append(newArray, value)
		}
	}
	return newArray
}

/**
* 数组转string
 */
func ArrayToString(array []string) string {
	data, _ := json.Marshal(array)
	return string(data)
}

func StringToArray(str string) []string {
	var showArray []string
	if str == "" {
		return showArray
	}
	err := json.Unmarshal([]byte(str), &showArray)
	if err != nil {
		panic(err)
	}
	return showArray
}

/**
* 数组转string
 */
func IntArrayToString(array []int) string {
	data, _ := json.Marshal(array)
	return string(data)
}

func StringToIntArray(str string) []int {
	var showArray []int
	if str == "" {
		return showArray
	}
	err := json.Unmarshal([]byte(str), &showArray)
	if err != nil {
		panic(err)
	}
	return showArray
}

func MapToString(cardMap map[int]int) string {
	data, _ := json.Marshal(cardMap)
	return string(data)
}

func StringToMap(str string) map[int]int {
	var cardMap map[int]int
	if str == "" {
		return cardMap
	}
	err := json.Unmarshal([]byte(str), &cardMap)
	if err != nil {
		panic(err)
	}
	return cardMap
}
