package gb2260

import "strconv"

type Division struct {
	Code string
	Name string
	Year string
}

var (
	EMPTY_DIVISION Division
)

func (this Division) Province() Division {
	code := this.Code[0:2] + "0000"
	return Get(code)
}

func (this Division) Prefecture() Division {
	if this.IsProvince() {
		return EMPTY_DIVISION
	}

	code := this.Code[:4] + "00"
	return Get(code)
}

func (this Division) Country() Division {
	if this.IsProvince() || this.IsPrefecture() {
		return EMPTY_DIVISION
	}
	return this
}

func (this Division) IsProvince() bool {
	return this.Province() == this
}

func (this Division) IsPrefecture() bool {
	return this.Prefecture() == this
}

func (this Division) IsCountry() bool {
	return this.Country() != EMPTY_DIVISION
}

func (this Division) Stack() []Division {
	stacks := make([]Division, 0)

	province := this.Province()
	stacks = append(stacks, province)

	if this.IsPrefecture() || this.IsCountry() {
		prefecture := this.Prefecture()
		stacks = append(stacks, prefecture)
	}

	if this.IsCountry() {
		stacks = append(stacks, this)
	}

	return stacks
}

func Get(code string) Division {
	key, err := strconv.Atoi(code)
	if err != nil {
		return EMPTY_DIVISION
	}

	div, ok := divisions[key]
	if !ok {
		return EMPTY_DIVISION
	}

	return div
}

// not implemenet
func Search(code string) *Division {
	return nil
}
