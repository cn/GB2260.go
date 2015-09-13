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

var (
	_EmptyDivision Division
)

// Province 省份
func (div Division) Province() Division {
	code := div.Code[0:2] + "0000"
	return Get(code)
}

// Prefecture 地区
func (div Division) Prefecture() Division {
	if div.IsProvince() {
		return _EmptyDivision
	}

	code := div.Code[:4] + "00"
	return Get(code)
}

// Country 县
func (div Division) Country() Division {
	if div.IsProvince() || div.IsPrefecture() {
		return _EmptyDivision
	}
	return div
}

// Description 描述
func (div Division) Description() string {
	stack := div.Stack()
	var names []string
	names = make([]string, 0)

	for _, s := range stack {
		names = append(names, s.Name)
	}

	return strings.Join(names, " ")
}

// IsProvince 是否省份
func (div Division) IsProvince() bool {
	return div.Province() == div
}

// IsPrefecture 是否地区
func (div Division) IsPrefecture() bool {
	return div.Prefecture() == div
}

// IsCountry 是否县
func (div Division) IsCountry() bool {
	return div.Country() != _EmptyDivision
}

// Stack 省，地区，县
func (div Division) Stack() []Division {
	var stacks []Division
	stacks = make([]Division, 0)

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

// Get 获取Division
func Get(code string) Division {
	key, err := strconv.Atoi(code)
	if err != nil {
		return _EmptyDivision
	}

	div, ok := divisions[key]
	if !ok {
		return _EmptyDivision
	}

	return div
}

// Search not implemenet
func Search(code string) *Division {
	return nil
}
