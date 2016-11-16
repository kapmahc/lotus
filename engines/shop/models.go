package shop

import (
	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

//Country country
type Country struct {
	web.Model

	Name string
}

//TableName table name
func (p *Country) TableName() string {
	return "shop_countries"
}

//State state
type State struct {
	web.Model

	Name string

	CountryID uint   `json:"country_id"`
	Country   string `json:"country"`
}

//TableName table name
func (p *State) TableName() string {
	return "shop_states"
}

//TaxRate TaxRate
type TaxRate struct {
	web.Model

	Value float64 `json:"value"`

	StateID   uint   `json:"state_id"`
	State     string `json:"state"`
	ProductID uint   `json:"product_id"`
	Product   string `json:"product"`
}

//TableName table name
func (p *TaxRate) TableName() string {
	return "shop_tax_rates"
}

//Address address
type Address struct {
	web.Model

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Content   string `json:"content"`
	Zip       string `json:"zip"`
	Phone     string `json:"phone"`

	UserID  uint      `json:"user_id"`
	User    auth.User `json:"user"`
	StateID uint      `json:"state_id"`
	State   string    `json:"state"`
}

//TableName table name
func (p *Address) TableName() string {
	return "shop_addresses"
}

//Property property
type Property struct {
	web.Model

	Name    string `json:"name"`
	Content string `json:"content"`
}

//TableName table name
func (p *Property) TableName() string {
	return "shop_properties"
}

//Product product
type Product struct {
	web.Model
}

//TableName table name
func (p *Product) TableName() string {
	return "shop_products"
}

//Option product option
type Option struct {
	web.Model
}

//TableName table name
func (p *Option) TableName() string {
	return "shop_options"
}

//Promotion promotion
type Promotion struct {
	web.Model
}

//TableName table name
func (p *Promotion) TableName() string {
	return "shop_promotions"
}

//Order order
type Order struct {
	web.Model
}

//TableName table name
func (p *Order) TableName() string {
	return "shop_orders"
}

//Payment payment
type Payment struct {
	web.Model
}

//TableName table name
func (p *Payment) TableName() string {
	return "shop_payments"
}

//Refund refund
type Refund struct {
	web.Model
}

//TableName table name
func (p *Refund) TableName() string {
	return "shop_refunds"
}

//Price price
type Price struct {
	web.Model
}

//TableName table name
func (p *Price) TableName() string {
	return "shop_price"
}

//Shipment shipment
type Shipment struct {
	web.Model
}

//TableName table name
func (p *Shipment) TableName() string {
	return "shop_shipments"
}
