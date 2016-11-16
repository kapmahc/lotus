package shop

/**
https://baya.github.io/2015/09/17/%E7%94%B5%E5%AD%90%E5%95%86%E5%8A%A1%E7%B3%BB%E7%BB%9F%E5%9F%BA%E7%A1%80%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E5%92%8C%E6%B5%81%E7%A8%8B.html
*/

import (
	"time"

	"github.com/kapmahc/lotus/engines/auth"
	"github.com/kapmahc/lotus/web"
)

//Product product
type Product struct {
	web.Model

	Name        string `json:"name"`
	Description string `json:"description"`

	Variants []Variant `json:"variants"`
}

//TableName table name
func (p *Product) TableName() string {
	return "shop_products"
}

//Variant variant
type Variant struct {
	web.Model

	SKU       string  `json:"sku"`
	Weight    float64 `json:"weight"`
	Height    float64 `json:"height"`
	Width     float64 `json:"width"`
	Length    float64 `json:"length"`
	CostPrice float64 `json:"cost_price"`
	Price     float64 `json:"price"`

	ProductID uint       `json:"product_id"`
	Product   Product    `json:"product"`
	Poperties []Property `json:"properties"`
}

//TableName table name
func (p *Variant) TableName() string {
	return "shop_variants"
}

//Country country
type Country struct {
	web.Model

	Name   string  `json:"name"`
	States []State `json:"states"`
	Active bool    `json:"active"`
}

//TableName table name
func (p *Country) TableName() string {
	return "shop_countries"
}

//State state
type State struct {
	web.Model

	Name string

	CountryID uint    `json:"country_id"`
	Country   Country `json:"country"`
}

//TableName table name
func (p *State) TableName() string {
	return "shop_states"
}

//Address address
type Address struct {
	web.Model

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	Zip       string `json:"zip"`
	Phone1    string `json:"phone1"`
	Phone2    string `json:"phone2"`
	Company   string `json:"company"`

	UserID  uint      `json:"user_id"`
	User    auth.User `json:"user"`
	StateID uint      `json:"state_id"`
	State   string    `json:"state"`
}

//TableName table name
func (p *Address) TableName() string {
	return "shop_addresses"
}

//TaxRate TaxRate
type TaxRate struct {
	web.Model

	Value float64 `json:"value"`

	StateID   uint   `json:"state_id"`
	State     string `json:"state"`
	VariantID uint   `json:"variant_id"`
	Variant   string `json:"variant"`
}

//TableName table name
func (p *TaxRate) TableName() string {
	return "shop_tax_rates"
}

//Property property
type Property struct {
	web.Model

	Name        string `json:"name"`
	Description string `json:"description"`

	Variants []Variant `json:"variants"`
}

//TableName table name
func (p *Property) TableName() string {
	return "shop_properties"
}

//Order order
type Order struct {
	web.Model

	UID             string  `json:"uid"`
	ItemTotal       float64 `json:"item_total"`
	Total           float64 `json:"total"`
	AdjustmentTotal float64 `json:"adjustment_total"`
	PaymentTotal    float64 `json:"payment_total"`

	//:cart, :address, :delivery, :payment, :confirm, :complete
	State string `json:"state"`
	//:ready, :pending, :partial, :shipped, :backorder, :canceled
	ShipmentState string `json:"shipment_state"`
	//:balance_due, :paid, :credit_owed, :failed, :void
	PaymentState string `json:"payment_state"`

	CompletedAt *time.Time `json:"completed_at"`

	AddressID uint    `json:"address_id"`
	Address   Address `json:"address"`
	User      string  `json:"user"`
	UserID    uint    `json:"user_id"`
}

//TableName table name
func (p *Order) TableName() string {
	return "shop_orders"
}

//LineItem line item
type LineItem struct {
	web.Model

	Quantity uint    `json:"quantity"`
	Price    float64 `json:"price"`

	VariantID uint    `json:"variant_id"`
	Variant   Variant `json:"variant"`
	OrderID   uint    `json:"order_id"`
	Order     Order   `json:"order"`
}

//TableName table name
func (p *LineItem) TableName() string {
	return "shop_line_items"
}

//PaymentMethod payment method
type PaymentMethod struct {
	web.Model

	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

//TableName table name
func (p *PaymentMethod) TableName() string {
	return "shop_payment_methods"
}

//Payment payment
type Payment struct {
	web.Model

	Amount float64 `json:"amount"`
	// :checkout, :pending, :processing, :failed, :completed
	State        string `json:"state"`
	ResponseCode string `json:"response_code"`
	AvsResponse  string `json:"avs_response"`

	OrderID         uint          `json:"order_id"`
	Order           Order         `json:"order"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
}

//TableName table name
func (p *Payment) TableName() string {
	return "shop_payments"
}

//ShipmentMethod shipment method
type ShipmentMethod struct {
	web.Model

	Name     string `json:"name"`
	Active   bool   `json:"active"`
	Tracking string `json:"tracking"`
}

//TableName table name
func (p *ShipmentMethod) TableName() string {
	return "shop_shipment_methods"
}

//Shipment shipment
type Shipment struct {
	web.Model

	Tracking  string     `json:"tracking"`
	UID       string     `json:"uid"`
	Cost      float64    `json:"cost"`
	ShippedAt *time.Time `json:"shipped_at"`
	//:ready, :pending, :assemble, :cancelled, :shipped
	State string `json:"state"`

	OrderID         uint          `json:"order_id"`
	Order           Order         `json:"order"`
	PaymentMethodID uint          `json:"payment_method_id"`
	PaymentMethod   PaymentMethod `json:"payment_method"`
}

//TableName table name
func (p *Shipment) TableName() string {
	return "shop_shipments"
}

//ReturnAuthorization return authorization
type ReturnAuthorization struct {
	web.Model
	UID     string     `json:"uid"`
	State   string     `json:"state"`
	Amount  float64    `json:"amount"`
	Reason  string     `json:"reason"`
	EnterAt *time.Time `json:"enter_at"`

	EnterByID uint       `json:"enter_by_id"`
	EnterBy   *auth.User `json:"enter_by"`
	OrderID   uint       `json:"order_id"`
	Order     Order      `json:"order_id"`
}

//TableName table name
func (p *ReturnAuthorization) TableName() string {
	return "shop_return_authorizations"
}

//InventoryUnit inventory unit
type InventoryUnit struct {
	web.Model

	LockVersion int    `json:"lock_version"`
	State       string `json:"state"`

	VariantID             uint `json:"variant_id"`
	OrderID               uint `json:"order_id"`
	ShipmentID            uint `json:"shipment_id"`
	ReturnAuthorizationID uint `json:"return_authorization_id"`
}

//TableName table name
func (p *InventoryUnit) TableName() string {
	return "shop_inventory_units"
}

//Chargeback chargeback
type Chargeback struct {
	web.Model

	State string `json:"state"`

	OrderID    uint    `json:"order_id"`
	OperatorID uint    `json:"operator_id"`
	Amount     float64 `json:"amount"`
}

//TableName table name
func (p *Chargeback) TableName() string {
	return "shop_chargebacks"
}

//Promotion promotion
type Promotion struct {
	web.Model

	Type   string `json:"type"`
	Script string `json:"script"`
}

//TableName table name
func (p *Promotion) TableName() string {
	return "shop_promotions"
}
