
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table shop_currencies(
  id serial primary key,
  cid char(3) not null,
  code char(3) not null,
  name varchar(255) not null,
  country varchar(255) not null,
  rate numeric(12,6) not null,
  units varchar(8) not null,
  active boolean not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_currencies_cid on shop_currencies(cid);
create index idx_shop_currencies_code on shop_currencies(code);
create index idx_shop_currencies_name on shop_currencies(name);
create index idx_shop_currencies_country on shop_currencies(country);

create table shop_products(
  id serial primary key,
  name varchar(255) not null,
  description text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_products_name on shop_products(name);

create table shop_categories(
  id serial primary key,
  name varchar(255) not null,
  parent_id int,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_categories_name on shop_categories(name);

create table shop_products_categories(
  id serial primary key,
  product_id int not null,
  category_id int not null,
  updated_at timestamp without time zone not null
);
create unique index idx_shop_products_categories_ids on shop_products_categories(category_id, product_id);

create table shop_variants(
  id serial primary key,
  sku varchar(255) not null,
  price numeric(12,2) not null,
  cost numeric(12,2) not null,
  length numeric(12,2) not null,
  weight numeric(12,2) not null,
  height numeric(12,2) not null,
  product_id int not null,
  currency_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_variants_sku on shop_variants(sku);

create table shop_countries(
  id serial primary key,
  name varchar(255) not null,
  active boolean not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_countries_name on shop_countries(name);

create table shop_states(
  id serial primary key,
  name varchar(255) not null,
  country_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_states_name on shop_states(name);
create unique index idx_shop_states_name_country on shop_states(name, country_id);

create table shop_tax_rates(
  id serial primary key,
  value numeric(12,2) not null,
  state_id int not null,
  variant_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_tax_rates_state_variant on shop_tax_rates(state_id, variant_id);

create table shop_properties(
  id serial primary key,
  name varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_properties_name on shop_properties(name);

create table shop_variants_properties(
  id serial primary key,
  variant_id int not null,
  property_id int not null,
  created_at timestamp without time zone not null default now()
);
create unique index idx_shop_variants_properties on shop_variants_properties(variant_id, property_id);

create table shop_addresses(
  id serial primary key,
  firstname varchar(255) not null,
  lastname varchar(255) not null,
  address1 varchar(255) not null,
  address2 varchar(255),
  phone1 varchar(32) not null,
  phone2 varchar(32),
  zipcode char(6) not null,
  company varchar(255),
  state_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);

create table shop_orders(
  id serial primary key,
  uid varchar(255) not null,
  item_total numeric(12,2) not null,
  total numeric(12,2) not null,
  adjustment_total numeric(12,2) not null,
  payment_total numeric(12,2) not null,
  state varchar(16) not null,
  shipment_state varchar(16) not null,
  payment_state varchar(16) not null,
  user_id int not null,
  completed_at timestamp without time zone,
  address_id int not null,
  currency_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_orders_uid on shop_orders(uid);
create index idx_shop_orders_state on shop_orders(state);
create index idx_shop_orders_shipment_state on shop_orders(shipment_state);
create index idx_shop_orders_payment_state on shop_orders(payment_state);

create table shop_line_items(
  id serial primary key,
  variant_id int not null,
  order_id int not null,
  quantity int not null,
  price numeric(8,2) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_line_items_variant_order on shop_line_items(variant_id, order_id);

create table shop_payment_methods(
  id serial primary key,
  type varchar(16) not null,
  name varchar(255) not null,
  description text,
  active boolean not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_payment_methods_type on shop_payment_methods(type);
create unique index idx_shop_payment_methods_name on shop_payment_methods(name);

create table shop_payments(
  id serial primary key,
  amount numeric(8,2) not null,
  order_id int not null,
  payment_method_id int not null,
  state varchar(16) not null,
  response_code varchar(255) not null,
  avs_response text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_payments_state on shop_payments(state);

create table shop_shipping_methods(
  id serial primary key,
  name varchar(255) not null,
  active boolean not null,
  tracking varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_shipping_methods_name on shop_shipping_methods(name);

create table shop_shipments(
  id serial primary key,
  tracking varchar(255) not null,
  uid varchar(255) not null,
  cost numeric(8,2) not null,
  shipped_at timestamp without time zone,
  order_id int not null,
  shipping_method_id int not null,
  address_id int not null,
  state varchar(16) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_shipments_uid on shop_shipments(uid);
create index idx_shop_shipments_state on shop_shipments(state);

create table shop_return_authorizations(
  id serial primary key,
  uid varchar(255) not null,
  state varchar(16) not null,
  amount numeric(8,2) not null,
  order_id int not null,
  reason text not null,
  enter_by_id int not null,
  enter_at timestamp without time zone,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_shop_return_authorizations_uid on shop_return_authorizations(uid);
create index idx_shop_return_authorizations_state on shop_return_authorizations(state);

create table shop_inventory_units(
  id serial primary key,
  lock_version int not null,
  state varchar(16) not null,
  variant_id int not null,
  shipment_id int not null,
  return_authorization_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_inventory_units_state on shop_inventory_units(state);

create table shop_chargebacks(
  id serial primary key,
  state varchar(16) not null,
  order_id int not null,
  operator_id int not null,
  amount numeric(8,2) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_shop_chargebacks_state on shop_chargebacks(state);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table shop_chargebacks;
drop table shop_inventory_units;
drop table shop_return_authorizations;
drop table shop_shipments;
drop table shop_shipping_methods;
drop table shop_payments;
drop table shop_payment_methods;
drop table shop_line_items;
drop table shop_orders;
drop table shop_addresses;
drop table shop_variants_properties;
drop table shop_properties;
drop table shop_tax_rates;
drop table shop_states;
drop table shop_countries;
drop table shop_variants;
drop table shop_products_categories;
drop table shop_categories;
drop table shop_products;
drop table shop_currencies;
