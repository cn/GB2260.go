package gb2260

import (
	"strings"
	"testing"
)

func Test_DivisionCountry(t *testing.T) {
	// 110101 北京市/市辖区/东城区
	gb := NewGB2260("")

	division := gb.Get("110101")
	if division == nil {
		t.Error("division not exist")
	}

	if division.IsProvince() {
		t.Error("expect not province, got province")
	}

	if division.IsPrefecture() {
		t.Error("expect not prefecture, got prefecture")
	}

	if !division.IsCountry() {
		t.Error("expect country, got not country")
	}

	var names []string
	stacks := division.Stack()
	for _, div := range stacks {
		names = append(names, div.Name)
	}

	stackName := strings.Join(names, "/")
	if stackName != "北京市/市辖区/东城区" {
		t.Error("export 北京市/市辖区/东城区, got ", stackName)
	}
}

func Test_DivisionPrefecture(t *testing.T) {
	// 110100 北京市/市辖区
	gb := NewGB2260("")
	division := gb.Get("110100")
	if division == nil {
		t.Error("division not exist")
	}

	if division.IsProvince() {
		t.Error("expect not province, got province")
	}

	if !division.IsPrefecture() {
		t.Error("expect prefecture, got not prefecture")
	}

	if division.IsCountry() {
		t.Error("expect not country, got country")
	}

	var names []string
	stacks := division.Stack()
	for _, div := range stacks {
		names = append(names, div.Name)
	}

	stackName := strings.Join(names, "/")
	if stackName != "北京市/市辖区" {
		t.Error("export 北京市/市辖区, got ", stackName)
	}

}

func Test_DivisionProvince(t *testing.T) {
	// 110000 北京市
	gb := NewGB2260("")
	division := gb.Get("110000")
	if division == nil {
		t.Error("division not exist")
	}

	if !division.IsProvince() {
		t.Error("expect province, got not province")
	}

	if division.IsPrefecture() {
		t.Error("expect not prefecture, got prefecture")
	}

	if division.IsCountry() {
		t.Error("expect not country, got country")
	}

	var names []string
	stacks := division.Stack()
	for _, div := range stacks {
		names = append(names, div.Name)
	}

	stackName := strings.Join(names, "/")
	if stackName != "北京市" {
		t.Error("export 北京市, got ", stackName)
	}
}

func Test_Compare(t *testing.T) {
	div := Division{Code: "110101", Name: "东城区", Year: "2014"}
	gb := NewGB2260("")

	p := gb.Get("110101")
	if p == nil {
		t.Error("division not exist")
	}

	if !p.Equal(div) {
		t.Error("expect equal division, got not equal")
	}

	div = Division{Code: "110000", Name: "东城区", Year: "2014"}
	p = gb.Get("110101")
	if p == nil {
		t.Error("division not exist")
	}

	if p.Equal(div) {
		t.Error("expect not equal division, go equal")
	}
}

func Test_Provinces(t *testing.T) {
	gb := NewGB2260("")
	p := gb.Provinces()
	if p == nil {
		t.Error("provinces is nil")
	}

	if len(p) != 34 {
		t.Error("expect provinces length 34, but not")
	}
}

func Test_Prefectures(t *testing.T) {
	gb := NewGB2260("")

	prefectures := gb.Prefectures("110101")
	if prefectures == nil {
		t.Error("prefectures is nil")
	}

	if len(prefectures) != 2 {
		t.Error("expect prefectures length 2, but not")
	}
}

func Test_Counties(t *testing.T) {
	gb := NewGB2260("")

	countries := gb.Counties("110101")
	if countries == nil {
		t.Error("prefectures is nil")
	}

	if len(countries) != 14 {
		t.Error("expect counties length 14, but not")
	}
}
