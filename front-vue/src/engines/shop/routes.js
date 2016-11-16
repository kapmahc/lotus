import CountriesIndex from './countries/Index'
import CountriesNew from './countries/New'
import CountriesEdit from './countries/Edit'

export default [
  {name: 'shop.countries.index', path: '/shop/countries', component: CountriesIndex},
  {name: 'shop.countries.new', path: '/shop/countries/new', component: CountriesNew},
  {name: 'shop.countries.edit', path: '/shop/countries/:id/edit', component: CountriesEdit}
]
