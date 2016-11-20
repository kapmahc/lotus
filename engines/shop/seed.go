package shop

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

func (p *Engine) _initPaymentMethods() error {
	var count int
	if err := p.Db.
		Model(&PaymentMethod{}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	fd, err := os.Open("db/seed/payment_methods.csv")
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := csv.NewReader(fd)
	for {
		line, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		if line[0] != "Type" {
			if err := p.Db.Create(&PaymentMethod{
				Type:        line[0],
				Name:        line[1],
				Description: line[2],
			}).Error; err != nil {
				return err
			}
		}

	}

	return nil
}

func (p *Engine) _initShippingMethods() error {
	var count int
	if err := p.Db.
		Model(&ShippingMethod{}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	fd, err := os.Open("db/seed/shipping_methods.csv")
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := csv.NewReader(fd)
	for {
		line, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		if line[0] != "Name" {
			if err := p.Db.Create(&ShippingMethod{
				Name:     line[0],
				Tracking: line[1],
			}).Error; err != nil {
				return err
			}
		}

	}

	return nil
}

func (p *Engine) _initCountries() error {
	var count int
	if err := p.Db.
		Model(&Country{}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	fd, err := os.Open("db/seed/countries.csv")
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := csv.NewReader(fd)
	for {
		line, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}

		if line[0] != "Name" {
			c := Country{Name: line[0]}
			if err := p.Db.Create(&c).Error; err != nil {
				return err
			}
			for _, n := range strings.Split(line[1], ", ") {
				if err := p.Db.Create(&State{CountryID: c.ID, Name: n}).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (p *Engine) _initCurrencies() error {
	var count int
	if err := p.Db.Model(&Currency{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	// res, err := http.Get(iso4217)
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()

	fd, err := os.Open("db/seed/list_one.xml")
	if err != nil {
		return err
	}

	var val ISO4217
	dec := xml.NewDecoder(fd)
	if err = dec.Decode(&val); err != nil {
		return err
	}

	for _, it := range val.CcyTbl.CcyNtry {
		if err = p.Db.Create(&Currency{
			Country: it.CtryNm,
			Name:    it.CcyNm,
			Cid:     it.Ccy,
			Code:    it.CcyNbr,
			Units:   it.CcyMnrUnts,
		}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (p *Engine) _initPostalCodes() error {
	var count int
	if err := p.Db.
		Model(&PostalCode{}).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	fd, err := os.Open("db/seed/us_postal_codes.csv")
	if err != nil {
		return err
	}
	defer fd.Close()

	rd := csv.NewReader(fd)
	for {
		line, err := rd.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		if line[0] != "Postal Code" {
			pc := PostalCode{
				Country:           "United States of America",
				Cid:               line[0],
				PlaceName:         line[1],
				State:             line[2],
				StateAbbreviation: line[3],
				County:            line[4],
			}
			if pc.Latitude, err = strconv.ParseFloat(line[5], 64); err != nil {
				return err
			}
			if pc.Longitude, err = strconv.ParseFloat(line[6], 64); err != nil {
				return err
			}
			if err = p.Db.Create(&pc).Error; err != nil {
				return err
			}
		}

	}

	return nil
}

//Seed Insert seed data
func (p *Engine) Seed() error {
	if err := p._initCountries(); err != nil {
		return err
	}
	if err := p._initPaymentMethods(); err != nil {
		return err
	}
	if err := p._initShippingMethods(); err != nil {
		return err
	}
	if err := p._initCurrencies(); err != nil {
		return err
	}
	if err := p._initPostalCodes(); err != nil {
		return err
	}

	return nil
}
