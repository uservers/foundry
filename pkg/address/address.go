package address

// Address es una estructura de datos de una direccion
type Address struct {
	Name  Name   `json:"name"`
	Email string `json:"email"`

	Org string `json:"org"`

	Addr1      StreetRef `json:"addr1"`
	Addr2      string    `json:"addr2"`
	City       string    `json:"city"`
	Region     string    `json:"region"`
	PostalCode string    `json:"postalcode"`
	Country    string    `json:"country"`

	Phone Phone `json:"phone"`
}

// Name es una estructura de un nombre
type Name struct {
	Name    string `json:"given"`
	Middle  string `json:"middle"`
	Family  string `json:"family"`
	Family2 string `json:"family2"`
}

// StreetRef es una calle con sus numeros
type StreetRef struct {
	Street string `json:"street"`
	ExtNr  string `json:"ext"`
	IntNr  string `json:"int"`
}

// Phone a telephone number abstraction
type Phone struct {
	CountryCode int   `json:"cc"`
	AreaCode    int   `json:"area"`
	Number      int64 `json:"number"`
}
