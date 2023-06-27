package cci

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type IdInfo struct {
	State    string    // 省
	City     string    // 市
	Region   string    // 区、县
	Birthday time.Time // 生日
	Gender   string    // 性别
}

var (
	ErrLength    = errors.New("invalid id, id length must eq 18")
	ErrInvalidID = errors.New("the ID format is invalid")
	ErrInvalid   = errors.New("id invalid")
)

// Check 校验18位身份证号是否正确有效，并返回相关信息
func Check(id string) (*IdInfo, bool, error) {
	if len(id) != 18 {
		return nil, false, ErrLength
	}
	// 正则校验前17位是否为数字
	match, err := regexp.MatchString(`^\d{17}`, id)
	if err != nil {
		return nil, false, err
	}
	if !match {
		return nil, false, ErrInvalidID
	}
	// 转化为数字
	idInts := make([]int, 0)
	for i := 0; i < 17; i++ {
		sd := int(id[i] - '0') // 直接将run转为int
		idInts = append(idInts, sd)
	}
	// 检查地区码是否合法
	file, err := os.Open("data.json")
	if err != nil {
		return nil, false, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	info, err := file.Stat()
	if err != nil {
		return nil, false, err
	}
	localDataB := make([]byte, info.Size())
	_, err = file.Read(localDataB)
	if err != nil {
		return nil, false, err
	}
	localData := make([]*State, 0)
	err = json.Unmarshal(localDataB, &localData)
	if err != nil {
		return nil, false, err
	}
	idInfo := new(IdInfo)
	for _, state := range localData {
		if state.Code == id[:2] {
			idInfo.State = state.Name
			if len(state.City) > 0 {

				for _, city := range state.City {
					if city.Code == id[2:4] {
						idInfo.City = city.Name
						if len(city.Region) > 0 {
							for _, region := range city.Region {
								if region.Code == id[4:6] {
									idInfo.Region = region.Name
									break
								}
							}
						}
						break
					}
				}
			}
			break
		}
	}
	// 检查生日是否合法，使用time.Parse来校验日期是否有效
	birthday, err := time.Parse(
		time.DateOnly, fmt.Sprintf("%s-%s-%s", id[6:10], id[10:12], id[12:14]),
	)
	if err != nil {
		fmt.Println(err)
		return nil, false, ErrInvalidID
	}
	idInfo.Birthday = birthday

	// 性别
	atoi, err := strconv.Atoi(id[16:17])
	if err != nil {
		return nil, false, ErrInvalidID
	}
	if atoi%2 == 1 {
		idInfo.Gender = "男"
	} else {
		idInfo.Gender = "女"
	}
	// 开始校验最后一位是否合法
	// 前17位对应相乘的数，这是固定的
	// 7,9,10,5,8,4,2,1,6,3,7,9,10,5,8,4,2
	multNum := [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 计算乘积和
	multSum := 0
	for i, idInt := range idInts {
		multSum += idInt * multNum[i]
	}

	// 取余11
	afterSurp := multSum % 11

	// 参照值 1, 0, X, 9, 8, 7, 6, 5, 4, 3, 2
	referVal := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	// 计算得到身份证号最后一位应该为：
	lastValueShouldBe := referVal[afterSurp]
	if lastValueShouldBe != strings.ToUpper(id[17:]) {
		return nil, false, ErrInvalid
	}

	return idInfo, true, nil
}
