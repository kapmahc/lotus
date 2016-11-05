<template>
  <nav class="navbar navbar-fixed-top navbar-dark bg-inverse">
    <a class="navbar-brand" href="#">Project name</a>
    <ul class="nav navbar-nav">
      <li class="nav-item active">
        <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">About</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">Contact</a>
      </li>
    </ul>

    <ul class="nav navbar-nav float-xs-right">
      <lang-bar/>
      <personal-bar/>
    </ul>
   </nav>
</template>

<script>
import LangBar from './LangBar'
import PersonalBar from './PersonalBar'

export default {
  name: 'layout-header',
  data () {
    return {
    }
  },
  components: {
    LangBar,
    PersonalBar
  }
}
</script>
