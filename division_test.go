package gb2260

import (
	"strings"
	"testing"
)

func Test_DivisionCountry(t *testing.T) {
	// 110101 北京市/市辖区/东城区
	division := Get("110101")

	if division.IsProvince() {
		t.Error("expect not province, got province")
	}

	if division.IsPrefecture() {
		t.Error("expect not prefecture, got prefecture")
	}

	if !division.IsCountry() {
		t.Error("expect country, got not country")
	}

	names := make([]string, 0)
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
	division := Get("110100")

	if division.IsProvince() {
		t.Error("expect not province, got province")
	}

	if !division.IsPrefecture() {
		t.Error("expect prefecture, got not prefecture")
	}

	if division.IsCountry() {
		t.Error("expect not country, got country")
	}

	names := make([]string, 0)
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
	division := Get("110000")

	if !division.IsProvince() {
		t.Error("expect province, got not province")
	}

	if division.IsPrefecture() {
		t.Error("expect not prefecture, got prefecture")
	}

	if division.IsCountry() {
		t.Error("expect not country, got country")
	}

	names := make([]string, 0)
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
	if Get("110101") != div {
		t.Error("expect equal division, got not equal")
	}

	div = Division{Code: "110000", Name: "东城区", Year: "2014"}
	if Get("110101") == div {
		t.Error("expect not equal division, go equal")
	}
}
