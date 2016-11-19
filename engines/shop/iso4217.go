package shop

import "encoding/xml"

const iso4217 = "http://www.currency-iso.org/dam/downloads/lists/list_one.xml"

//ISO4217 iso-4217
type ISO4217 struct {
	XMLName xml.Name `xml:"ISO_4217"`
	CcyTbl  CcyTbl
}

//CcyTbl CcyTbl
type CcyTbl struct {
	CcyNtry []CcyNtry
}

//CcyNtry CcyNtry
type CcyNtry struct {
	CtryNm     string
	CcyNm      string
	Ccy        string
	CcyNbr     string
	CcyMnrUnts string
}
