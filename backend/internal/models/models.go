package models

type Cards struct { // for db queries
	Bin    int
	Issuer string
}

type CardInfo struct { // for external requests
	Issuer struct {
		Name string `json:"name"`
	} `json:"bank"`
}
