<template>
  <app-dashboard>
    <admin-settings />
    <br/>
    <h3>{{ $t('shop.models.currency') }}</h3>
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('shop.attributes.currency-cid')}}</th>
          <th>{{$t('attributes.name')}}</th>
          <th>{{$t('shop.models.country')}}</th>
          <th>{{$t('shop.attributes.currency-rate')}}</th>
          <th>{{$t('attributes.active')}}?</th>
          <th>
            {{$t('buttons.manage')}}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.cid}}</th>
          <td>{{i.name}}</td>
          <td>{{i.country}}</td>
          <td>{{i.rate}}</td>
          <td>{{i.active}}</td>
          <td>
            <router-link class="btn btn-warning btn-sm" :to="{name: 'shop.currencies.edit', params:{id: i.id}}">
              {{$t('buttons.edit')}}
            </router-link>
          </td>
        </tr>
      </tbody>
    </table>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import AdminSettings from '../Settings'
import {get} from '../../../utils'

export default {
  name: 'shop-admin-currencies',
  data () {
    return {
      items: []
    }
  },
  created () {
    this.refresh()
  },
  components: {
    AppDashboard,
    AdminSettings
  },
  methods: {
    refresh () {
      get('/shop/currencies', null, function (rst) {
        this.items = rst
      }.bind(this))
    }
  }
}
</script>
