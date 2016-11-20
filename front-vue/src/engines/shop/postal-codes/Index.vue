<template>
  <app-dashboard>
    <admin-settings />
    <br/>
    <h3>{{ $t('shop.models.postal-code') }}</h3>
    <paginator :pager="pager" />
    <table class="table table-bordered table-hover">
      <thead>
        <tr>
          <th>{{$t('shop.attributes.postal-code-cid')}}</th>
          <th>{{$t('shop.attributes.postal-code-place-name')}}</th>
          <th>{{$t('shop.attributes.postal-code-county')}}</th>
          <th>{{$t('shop.attributes.postal-code-state')}}</th>
          <th>{{$t('shop.models.country')}}</th>
          <th>{{$t('shop.attributes.postal-code-latitude')}}</th>
          <th>{{$t('shop.attributes.postal-code-longitude')}}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in items">
          <th scope="row">{{i.cid}}</th>
          <td>{{i.place_name}}</td>
          <td>{{i.county}}</td>
          <td>{{i.state}}({{i.state_abbreviation}})</td>
          <td>{{i.country}}</td>
          <td>{{i.latitude}}</td>
          <td>{{i.longitude}}</td>
        </tr>
      </tbody>
    </table>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import AdminSettings from '../Settings'
import {get} from '../../../utils'
import Paginator from '../../../components/Paginator'

export default {
  name: 'shop-admin-postal-codes',
  data () {
    return {
      items: [],
      pager: {
        size: 120,
        current: 1
      }
    }
  },
  created () {
    this.refresh(this.pager.current)
  },
  components: {
    AppDashboard,
    AdminSettings,
    Paginator
  },
  methods: {
    onPager (i, e) {
      e.preventDefault()
      this.refresh(i)
    },
    refresh (page) {
      get(
        '/shop/postal-codes',
        {
          page,
          size: this.pager.size
        },
        function (rst) {
          this.items = rst.items
          var pager = rst.pager
          pager.href = 'shop.postal-codes.index'
          pager.click = this.onPager
          this.pager = pager
        }.bind(this)
      )
    }
  }
}
</script>
