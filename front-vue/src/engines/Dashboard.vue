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

        <li v-for="nv in header" class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" :id="nv.id" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            {{ $t(nv.title) }}
          </a>
          <div class="dropdown-menu" :aria-labelledby="nv.id">
            <router-link v-for="lk in nv.links" class="dropdown-item" :to="{ name: lk.href }">
              {{ $t(lk.label) }}
            </router-link>
          </div>
        </li>

      </ul>

      <ul class="nav navbar-nav float-xs-right">
        <lang-bar/>
        <personal-bar/>
      </ul>
     </nav>
    <div class="container">
      <slot>
        <div v-for="nv in header" class="row">
          <br/>
          <h4>{{ $t(nv.title) }}</h4>
          <hr/>
          <div class="btn-group" role="group">
            <router-link v-for="lk in nv.links" class="btn btn-secondary" :to="{ name: lk.href }">
              {{ $t(lk.label) }}
            </router-link>
          </div>
        </div>
      </slot>
      <layout-footer/>
    </div>
  </div>
</template>

<script>
import LayoutFooter from '../components/Footer'
import LangBar from '../components/LangBar'
import PersonalBar from '../components/PersonalBar'

export default {
  name: 'app-dashboard',
  data () {
    return {
      header: [
        {
          id: 'auth-users',
          title: 'auth.pages.profile',
          links: [
            {
              label: 'auth.pages.info',
              href: 'users.info'
            },
            {
              label: 'auth.pages.change-password',
              href: 'users.change-password'
            },
            {
              label: 'auth.pages.logs',
              href: 'users.logs'
            }
          ]
        },
        {
          id: 'auth-site',
          title: 'auth.pages.admin-profile',
          links: [
            {
              label: 'auth.pages.admin-author',
              href: 'admin.author'
            },
            {
              label: 'auth.pages.admin-base',
              href: 'admin.base'
            },
            {
              label: 'auth.pages.admin-i18n',
              href: 'admin.i18n'
            },
            {
              label: 'auth.pages.admin-seo',
              href: 'admin.seo'
            },
            {
              label: 'auth.pages.admin-smtp',
              href: 'admin.smtp'
            },
            {
              label: 'auth.pages.admin-status',
              href: 'admin.status'
            },
            {
              label: 'auth.pages.admin-users',
              href: 'admin.users'
            },
            {
              label: 'auth.pages.leavewords',
              href: 'leavewords.index'
            },
            {
              label: 'auth.pages.notices',
              href: 'notices.admin'
            }
          ]
        }
      ]
    }
  },
  computed: {
    info () {
      return this.$store.state.siteInfo
    }
  },
  components: {
    LayoutFooter,
    LangBar,
    PersonalBar
  }
}
</script>
