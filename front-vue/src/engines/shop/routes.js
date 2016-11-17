import CountriesIndex from './countries/Index'
import CountriesNew from './countries/New'
import CountriesEdit from './countries/Edit'
import StatesIndex from './states/Index'
import StatesNew from './states/New'
import StatesEdit from './states/Edit'

export default [
  {name: 'shop.states.index', path: '/shop/states', component: StatesIndex},
  {name: 'shop.states.new', path: '/shop/states/new', component: StatesNew},
  {name: 'shop.states.edit', path: '/shop/states/:id/edit', component: StatesEdit},
  {name: 'shop.countries.index', path: '/shop/countries', component: CountriesIndex},
  {name: 'shop.countries.new', path: '/shop/countries/new', component: CountriesNew},
  {name: 'shop.countries.edit', path: '/shop/countries/:id/edit', component: CountriesEdit}
]
