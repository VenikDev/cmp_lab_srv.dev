<script setup>

import {ref} from "vue";

const { analysisData } = defineProps({
  analysisData: Object
})

let dialogIsOpen = ref(false)

</script>


<template>
  <div
      class="container_border container_border_shadow"
      @click="dialogIsOpen = !dialogIsOpen"
  >
    <div class="title">
      <b> {{ analysisData.name }} </b>
    </div>
    <div class="price">
      Цена: <b> {{ analysisData.price }} </b> рублей
    </div>
  </div>
</template>

<style scoped>

.text {
  @apply text-cyan-600;
  font-family: 'Comfortaa', cursive;
}

.container_border_shadow {
  @apply duration-300 hover:shadow-lg hover:shadow-PompPower;
}

.title {
  @apply text-ChineseViolet text-xl mb-3
}

.price {
  @apply text-black text-xl text-right
}

</style>