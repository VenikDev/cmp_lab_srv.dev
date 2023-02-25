<script setup>

const { analysisData } = defineProps({
  analysisData: Object
})

</script>


<template>
  <div class="container_border">
    <h2 class="text">
      Название: <b> {{ analysisData.name }} </b>
    </h2>
    <span class="text">
      Цена: <b> {{ analysisData.price }} </b> рублей
    </span>
  </div>
</template>

<style scoped>

.text {
  @apply text-cyan-600;
  font-family: 'Comfortaa', cursive;
}

.container_border:focus {
  @apply shadow-lg shadow-cyan-600
}

</style>