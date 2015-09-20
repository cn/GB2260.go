package gb2260

import (
	"strconv"
	"strings"
)

// Division 地区
type Division struct {
	Code string // The six-digit number of the specific administrative division.
	Name string // The Chinese name of the specific administrative division.
	Year string // Optional. The revision year, and empty means "latest".
}

// TODO:
// If year is not specified, use the latest data.
func NewDivision(year string) *Division {
	if year == "" {
		return nil
	}

	return nil
}

// TODO: use hash to compare equal
func (div Division) equal(other Division) bool {
	return div.Code == other.Code && div.Name == other.Name && div.Year == other.Year
}

// Province 省份
func (div Division) Province() *Division {
	code := div.Code[0:2] + "0000"
	return Get(code)
}

// Prefecture 地区
func (div Division) Prefecture() *Division {
	if div.IsProvince() {
		return nil
	}

	code := div.Code[:4] + "00"
	return Get(code)
}

// Country 县
func (div Division) Country() *Division {
	if div.IsProvince() || div.IsPrefecture() {
		return nil
	}
	return div
}

// Description 描述
func (div Division) Description() string {
	stack := div.Stack()
	var names []string
	names = make([]string, 0)

	for _, s := range stack {
		if s != nil {
			names = append(names, s.Name)
		}
	}

	return strings.Join(names, " ")
}

// IsProvince 是否省份
func (div Division) IsProvince() bool {
	pro := div.Province()
	if pro == nil {
		return false
	}

	return pro.equal(div)
}

// IsPrefecture 是否地区
func (div Division) IsPrefecture() bool {
	pre := div.Prefecture()
	if pre == nil {
		return false
	}

	return pre.equal(div)
}

// IsCountry 是否县
func (div Division) IsCountry() bool {
	country := div.Country()
	if country == nil {
		return false
	}

	return ountry.equal(div)
}

// Stack 省，地区，县
func (div Division) Stack() []*Division {
	var stacks []*Division
	stacks = make([]*Division, 0)

	province := div.Province()
	stacks = append(stacks, province)

	if div.IsPrefecture() || div.IsCountry() {
		prefecture := div.Prefecture()
		stacks = append(stacks, prefecture)
	}

	if div.IsCountry() {
		stacks = append(stacks, div)
	}

	return stacks
}

// Get Obtain Division
func Get(code string) *Division {
	key, err := strconv.Atoi(code)
	if err != nil {
		return nil
	}

	div, ok := divisions[key]
	if !ok {
		return nil
	}

	return div
}

// Provinces Return a list of provinces in Division data structure.
func Provinces() *[]Division {
	//
}

// Prefectures  Return a list of prefecture level cities in Division data structure.
// A province_code is a 6-length province code. It can also be:
// 2-length code
// 4-length code that endswith 00
func Prefectures(provinceCode string) []*Division {

}

// Counties Return a list of counties in Division data structure.
// A prefecture_code is a 6-length code that endswith 00. It can also be a 4-length code.
func Counties(prefectureCode string) []*Division {

}

// Search not implemenet
func Search(code string) *Division {
	return nil
}
