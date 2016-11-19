import CountriesIndex from './countries/Index'
import CountriesNew from './countries/New'
import CountriesEdit from './countries/Edit'
import StatesIndex from './states/Index'
import StatesNew from './states/New'
import StatesEdit from './states/Edit'
import PaymentMethodsIndex from './payment-methods/Index'
import PaymentMethodsNew from './payment-methods/New'
import PaymentMethodsEdit from './payment-methods/Edit'
import ShippingMethodsIndex from './shipping-methods/Index'
import ShippingMethodsNew from './shipping-methods/New'
import ShippingMethodsEdit from './shipping-methods/Edit'
import CurrenciesIndex from './currencies/Index'

export default [
  {name: 'shop.currencies.index', path: '/shop/currencies', component: CurrenciesIndex},
  {name: 'shop.shipping-methods.index', path: '/shop/shipping-methods', component: ShippingMethodsIndex},
  {name: 'shop.shipping-methods.new', path: '/shop/shipping-methods/new', component: ShippingMethodsNew},
  {name: 'shop.shipping-methods.edit', path: '/shop/shipping-methods/:id/edit', component: ShippingMethodsEdit},
  {name: 'shop.payment-methods.index', path: '/shop/payment-methods', component: PaymentMethodsIndex},
  {name: 'shop.payment-methods.new', path: '/shop/payment-methods/new', component: PaymentMethodsNew},
  {name: 'shop.payment-methods.edit', path: '/shop/payment-methods/:id/edit', component: PaymentMethodsEdit},
  {name: 'shop.states.index', path: '/shop/states', component: StatesIndex},
  {name: 'shop.states.new', path: '/shop/states/new', component: StatesNew},
  {name: 'shop.states.edit', path: '/shop/states/:id/edit', component: StatesEdit},
  {name: 'shop.countries.index', path: '/shop/countries', component: CountriesIndex},
  {name: 'shop.countries.new', path: '/shop/countries/new', component: CountriesNew},
  {name: 'shop.countries.edit', path: '/shop/countries/:id/edit', component: CountriesEdit}
]
