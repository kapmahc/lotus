<template>
  <div id="root">
    <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
      <router-link class="navbar-brand" :to="{ name: 'home' }">
        {{ info.subTitle }}
      </router-link>

      <ul class="nav navbar-nav">
        <li class="nav-item active">
          <router-link class="nav-link" :to="{ name: 'dashboard' }">
            {{ $t('auth.pages.dashboard') }}
            <span class="sr-only">(current)</span>
          </router-link>
        </li>
        <drop-link v-if="user.uid"
          v-for="(item, index) in member"
          v-bind:item="item"
          v-bind:id="index" />
        <drop-link v-if="user.roles && user.roles.includes('admin')"
          v-for="(item, index) in admin"
          v-bind:item="item"
          v-bind:id="index" />
      </ul>

      <ul class="nav navbar-nav float-xs-right">
        <lang-bar/>
        <personal-bar/>
      </ul>
     </nav>
    <div class="container">
      <slot>
        <nav-pane v-if="user.uid" :items="member" />
        <nav-pane v-if="user.roles && user.roles.includes('admin')" :items="admin" />
      </slot>
      <layout-footer/>
    </div>
  </div>
</template>

<script>
import LayoutFooter from '../components/Footer'
import LangBar from '../components/LangBar'
import PersonalBar from '../components/PersonalBar'
import NavPane from '../components/NavPane'
import DropLink from '../components/DropLink'

export default {
  name: 'app-dashboard',
  data () {
    var member = []
    var admin = []

    member.push({
      id: 'auth-users',
      title: 'auth.pages.profile',
      links: [
        {label: 'auth.pages.info', href: 'users.info'},
        {label: 'auth.pages.change-password', href: 'users.change-password'},
        {label: 'auth.pages.logs', href: 'users.logs'}
      ]
    })

    admin.push({
      id: 'auth-site',
      title: 'auth.pages.admin-profile',
      links: [
        {label: 'auth.pages.admin-author', href: 'admin.author'},
        {label: 'auth.pages.admin-base', href: 'admin.base'},
        {label: 'auth.pages.admin-locales', href: 'admin.locales'},
        {label: 'auth.pages.admin-seo', href: 'admin.seo'},
        {label: 'auth.pages.admin-smtp', href: 'admin.smtp'},
        {label: 'auth.pages.admin-status', href: 'admin.status'},
        {label: 'auth.pages.admin-users', href: 'admin.users'},
        {label: 'auth.pages.leavewords', href: 'leavewords.index'},
        {label: 'auth.pages.notices', href: 'notices.admin'}
      ]
    })

    member.push({
      id: 'shop-member',
      title: 'shop.pages.self-profile',
      links: [
        {label: 'shop.pages.self-addresses', href: 'home'},
        {label: 'shop.pages.self-cart', href: 'home'},
        {label: 'shop.pages.self-orders', href: 'home'},
        {label: 'shop.pages.self-returns', href: 'home'},
        {label: 'shop.pages.self-messages', href: 'home'},
        {label: 'shop.pages.self-history', href: 'home'}
      ]
    })
    admin.push({
      id: 'shop-admin',
      title: 'shop.pages.admin-profile',
      links: [
        {label: 'shop.pages.admin-countries', href: 'shop.countries.index'},
        {label: 'shop.pages.admin-payments', href: 'shop.payment-methods.index'},
        {label: 'shop.pages.admin-shipments', href: 'home'},
        {label: 'shop.pages.admin-products', href: 'home'},
        {label: 'shop.pages.admin-promotions', href: 'home'},
        {label: 'shop.pages.admin-orders', href: 'home'},
        {label: 'shop.pages.admin-returns', href: 'home'}
      ]
    })

    return {
      member,
      admin
    }
  },
  computed: {
    info () {
      return this.$store.state.siteInfo
    },
    user () {
      return this.$store.state.currentUser
    }
  },
  components: {
    LayoutFooter,
    LangBar,
    PersonalBar,
    NavPane,
    DropLink
  }
}
</script>
