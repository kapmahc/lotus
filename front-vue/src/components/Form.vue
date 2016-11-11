<template>
  <div class="row">
    <h3>{{ title }}</h3>
    <hr/>
    <form :method="method" :action="action">
      <div class="form-group" v-for="field in fields">
        <label for="email">{{ $t("attributes.email") }}</label>
        <input type="email" class="form-control" id="email">
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'ajax-form',
  props: ['title', 'method', 'action', 'fields']
}
</script>
