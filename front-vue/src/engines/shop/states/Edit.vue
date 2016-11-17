<template>
  <app-dashboard>
    <h3>{{$t('buttons.edit')}} {{ $t('shop.models.state') }} [{{id}}]</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <button type="submit" class="btn btn-primary">{{ $t('buttons.submit') }}</button>
      <button type="reset" class="btn btn-secondary">{{ $t('buttons.reset') }}</button>
    </form>
  </app-dashboard>
</template>

<script>
import AppDashboard from '../../Dashboard'
import {postForm, get} from '../../../utils'

export default {
  name: 'edit-state',
  data () {
    return {
      id: this.$route.params.id,
      name: ''
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get(`/shop/states/${this.id}`, null, function (rst) {
      this.name = rst.name
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        `/shop/states/${this.id}`,
        {
          name: this.name
        },
        function (result) {
          this.$router.push({name: 'shop.states.index'})
        }.bind(this)
      )
    }
  }
}
</script>
