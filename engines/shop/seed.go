package shop

import (
	"encoding/xml"
	"net/http"
)

func (p *Engine) _initPaymentMethods(t, n string) error {

	var count int
	if err := p.Db.
		Model(&PaymentMethod{}).
		Where("name = ?", n).
		Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return p.Db.Create(&PaymentMethod{Name: n, Type: n, Active: true}).Error
	}
	return nil
}

func (p *Engine) _initShippingMethod(name, tracking string) error {

	var count int
	if err := p.Db.Model(&ShippingMethod{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return p.Db.Create(&ShippingMethod{
			Name:     name,
			Tracking: tracking,
			Active:   true,
		}).Error
	}

	return nil
}

func (p *Engine) _initCountry(name string, states ...string) error {
	var count int

	if err := p.Db.Model(&Country{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		c := Country{Name: name, Active: true}
		if err := p.Db.Create(&c).Error; err != nil {
			return err
		}

		for _, sn := range states {
			var s State
			s.CountryID = c.ID
			s.Name = sn
			if err := p.Db.Create(&s).Error; err != nil {
				return err
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

	res, err := http.Get(iso4217)
	// res.Header.Add("Accept", "application/xml")
	// res.Header.Add("Content-Type", "application/xml; charset=utf-8")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var val ISO4217
	dec := xml.NewDecoder(res.Body)
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

//Seed Insert seed data
func (p *Engine) Seed() error {
	if err := p._initCountry(
		"United States of America",
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
	); err != nil {
		return err
	}
	if err := p._initCountry("中华人民共和国",
		"北京市", "天津市", "上海市", "重庆市",
		"河北省", "山西省", "辽宁省", "吉林省", "黑龙江省", "江苏省", "浙江省", "安徽省", "福建省", "江西省", "山东省", "河南省", "湖北省", "湖南省", "广东省", "海南省", "四川省", "贵州省", "云南省", "陕西省", "甘肃省", "青海省", "台湾省", "内蒙古自治区", "广西壮族自治区", "西藏自治区", "宁夏回族自治区", "新疆维吾尔自治区", "香港特别行政区", "澳门特别行政区",
	); err != nil {
		return err
	}

	//https://www.paypal.com/
	if err := p._initPaymentMethods("paypal", "Paypal"); err != nil {
		return err
	}
	//https://intl.alipay.com/
	if err := p._initPaymentMethods("alipay", "支付宝"); err != nil {
		return err
	}
	//https://pay.weixin.qq.com/
	if err := p._initPaymentMethods("wechat", "微信支付"); err != nil {
		return err
	}

	if err := p._initShippingMethod(
		"UPS",
		"https://www.ups.com/WebTracking/track",
	); err != nil {
		return nil
	}
	if err := p._initShippingMethod(
		"USPS",
		"https://tools.usps.com/go/TrackConfirmAction_input",
	); err != nil {
		return nil
	}
	if err := p._initShippingMethod(
		"Fedex",
		"https://www.fedex.com/apps/fedextrack/",
	); err != nil {
		return nil
	}
	if err := p._initShippingMethod(
		"中国邮政速递",
		"http://www.ems.com.cn/mailtracking/you_jian_cha_xun.html",
	); err != nil {
		return nil
	}

	return p._initCurrencies()
}
