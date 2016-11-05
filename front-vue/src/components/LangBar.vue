<template>
  <li class="nav-item dropdown">
    <a class="nav-link dropdown-toggle" id="switch-lang-bar" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
      {{ $t('lang-bar.switch') }}
    </a>
    <div class="dropdown-menu dropdown-menu-right" aria-labelledby="switch-lang-bar">
      <a class="dropdown-item" href="/?locale=en-US">
        english
      </a>
      <a class="dropdown-item" href="/?locale=en-US">
        zh-CN
      </a>
      <a class="dropdown-item" href="/?locale=en-US">
        zh-TW
      </a>
    </div>
  </li>
</template>

<script>
export default {
  name: 'lang-bar',
  data () {
    return {
    }
  }
}
</script>
