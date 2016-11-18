<template>
  <app-dashboard>
    <h3>{{$t('buttons.edit')}} {{ $t('shop.models.payment-method') }} [{{id}}]</h3>
    <hr/>
    <form v-on:submit.prevent="onSubmit">
      <div class="form-group">
        <label for="type">{{ $t("attributes.type") }}</label>
        <input v-model="type" type="text" class="form-control" id="type">
      </div>
      <div class="form-group">
        <label for="name">{{ $t("attributes.name") }}</label>
        <input v-model="name" type="text" class="form-control" id="name">
      </div>
      <div class="form-check">
        <label class="form-check-label">
          <input class="form-check-input" type="checkbox" v-model="active">
          {{ $t("attributes.active") }}
        </label>
      </div>
      <div class="form-group">
        <label for="description">{{ $t("attributes.description") }}</label>
        <textarea v-model="description" class="form-control" id="description" rows="3"></textarea>
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
      name: '',
      type: '',
      description: '',
      active: true
    }
  },
  components: {
    AppDashboard
  },
  created () {
    get(`/shop/payment-methods/${this.id}`, null, function (rst) {
      this.name = rst.name
      this.type = rst.type
      this.description = rst.description
      this.active = rst.active
    }.bind(this))
  },
  methods: {
    onSubmit () {
      postForm(
        `/shop/payment-methods/${this.id}`,
        {
          name: this.name,
          type: this.type,
          active: this.active,
          description: this.description
        },
        function (result) {
          this.$router.push({name: 'shop.payment-methods.index'})
        }.bind(this)
      )
    }
  }
}
</script>
