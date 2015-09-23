package gb2260

import (
	"fmt"
	"regexp"
	"strings"
)

// GB2260 GB2260
type GB2260 struct {
	Store map[string]string
	Year  string
}

var (
	_LatestYear = "2014"
)

// NewGB2260 If year is not specified, use the latest data.
func NewGB2260(year string) GB2260 {
	if year == "" {
		year = _LatestYear
	}

	return GB2260{
		Store: divisions[year],
		Year:  year,
	}
}

// Get Obtain Division
func (g GB2260) Get(code string) *Division {
	name, ok := g.Store[code]
	if !ok {
		return nil
	}

	return &Division{
		Code: code,
		Name: name,
		Year: g.Year,
		gb:   g,
	}
}

// Search Division
func Search(code string, years []string) *Division {
	for _, year := range years {
		gb := NewGB2260(year)
		division := gb.Get(code)
		if division != nil {
			return division
		}
	}

	return nil
}

// Provinces Return a list of provinces in Division data structure.
func (g GB2260) Provinces() []*Division {
	var divisions []*Division

	for code, name := range g.Store {
		if strings.HasSuffix(code, "0000") {
			d := &Division{
				Code: code,
				Name: name,
				Year: g.Year,
				gb:   g,
			}

			divisions = append(divisions, d)
		}
	}

	return divisions
}

// Prefectures  Return a list of prefecture level cities in Division data structure.
// A province_code is a 6-length province code. It can also be:
// 2-length code
// 4-length code that endswith 00
func (g GB2260) Prefectures(provinceCode string) []*Division {
	if len(provinceCode) < 2 {
		return nil
	}

	var divisions []*Division
	province := provinceCode[:2]

	prefecutres, err := regexp.Compile(province + "\\d{2}" + "00$")
	if err != nil {
		return nil
	}

	realProvinceCode := province + "0000"

	for code, name := range g.Store {
		if prefecutres.MatchString(code) && code != realProvinceCode {
			d := Division{
				Code: code,
				Name: name,
				Year: g.Year,
				gb:   g,
			}
			divisions = append(divisions, &d)
		}
	}

	return divisions

}

// Counties Return a list of counties in Division data structure.
// A prefecture_code is a 6-length code that endswith 00. It can also be a 4-length code.
func (g GB2260) Counties(prefectureCode string) []*Division {
	if len(prefectureCode) < 4 {
		return nil
	}

	prefecutre := prefectureCode[:4]

	var divisions []*Division
	country, err := regexp.Compile(prefecutre + "\\d{2}$")
	if err != nil {
		return nil
	}

	realPrefectureCode := prefecutre + "00"

	for code, name := range g.Store {
		if country.MatchString(code) && code != realPrefectureCode {
			d := Division{
				Code: code,
				Name: name,
				Year: g.Year,
				gb:   g,
			}
			divisions = append(divisions, &d)
		}
	}

	return divisions
}

// Division 地区
type Division struct {
	Code string `json:"code"` // The six-digit number of the specific administrative division.
	Name string `json:"name"` // The Chinese name of the specific administrative division.
	Year string `json:"year"` // Optional. The revision year, and empty means "latest".
	gb   GB2260
}

func (div Division) String() string {
	return fmt.Sprintf("code:%s, name:%s, year:%s", div.Code, div.Name, div.Year)
}

// Equal use hash to compare equal
func (div Division) Equal(other Division) bool {
	return div.Code == other.Code && div.Name == other.Name && div.Year == other.Year
}

// Province 省份
func (div Division) Province() *Division {
	code := div.Code[0:2] + "0000"
	return div.gb.Get(code)
}

// Prefecture 地区
func (div Division) Prefecture() *Division {
	if div.IsProvince() {
		return nil
	}

	code := div.Code[:4] + "00"
	return div.gb.Get(code)
}

// Country 县
func (div Division) Country() *Division {
	if div.IsProvince() || div.IsPrefecture() {
		return nil
	}
	return &div
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

	return pro.Equal(div)
}

// IsPrefecture 是否地区
func (div Division) IsPrefecture() bool {
	pre := div.Prefecture()
	if pre == nil {
		return false
	}

	return pre.Equal(div)
}

// IsCountry 是否县
func (div Division) IsCountry() bool {
	country := div.Country()
	if country == nil {
		return false
	}

	return country.Equal(div)
}

// Stack 省，地区，县
func (div Division) Stack() []*Division {
	var stacks []*Division
	stacks = make([]*Division, 0)

	province := div.Province()
	if province != nil {
		stacks = append(stacks, province)
	}

	if div.IsPrefecture() || div.IsCountry() {
		prefecture := div.Prefecture()
		if prefecture != nil {
			stacks = append(stacks, prefecture)
		}
	}

	if div.IsCountry() {
		stacks = append(stacks, &div)
	}

	return stacks
}
