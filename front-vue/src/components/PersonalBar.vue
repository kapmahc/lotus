<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      sign in / up
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <a class="dropdown-item" href="/?locale=en-US">
        sign in
      </a>
      <a class="dropdown-item" href="/?locale=en-US">
        sign up
      </a>
      <a class="dropdown-item" href="/?locale=en-US">
        forgot password
      </a>
    </div>
  </li>
</template>

<script>
export default {
  name: 'forgot-password',
  data () {
    return {
    }
  }
}
</script>
