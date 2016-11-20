<template>
  <nav aria-label="Page navigation">
    <ul class="pagination">
      <li class="page-item">
        <a class="page-link" href="#" v-on:click="pager.click(pager.prev, $event)" aria-label="Previous">
          <span aria-hidden="true">&laquo;</span>
          <span class="sr-only">Previous</span>
        </a>
      </li>
      <li v-for="i in pager.total" :class="i == pager.current ? 'page-item active' : 'page-item'">
        <a v-if="i==(pager.current-6)" class="page-link">
          ...
        </a>
        <a v-if="i<=(pager.current+5) && i>=(pager.current-5)" class="page-link" href="#" v-on:click="pager.click(i, $event)">
          {{ i }}
        </a>
        <a v-if="i==(pager.current+6)" class="page-link">
          ...
        </a>
      </li>
      <li class="page-item">
        <a class="page-link" href="#" v-on:click="pager.click(pager.next, $event)" aria-label="Next">
          <span aria-hidden="true">&raquo;</span>
          <span class="sr-only">Next</span>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  props: ['pager'],
  name: 'paginator'
}
</script>
