<script setup>

const {analysisData} = defineProps({
  analysisData: Object
})

</script>


<template>
  <div class="container_border">
    <h2 class="font_and_color">
      {{ analysisData.name }}
    </h2>
    <span class="font_and_color">
      {{ analysisData.price }}
    </span>
  </div>
</template>

<style scoped>

</style>