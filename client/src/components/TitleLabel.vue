<script setup>

import { Translate } from "./translate"

</script>

<template>

<div class="container_border">
  <label class="label_level_1">
    {{ Translate("title") }}
  </label>
  <br/>
  <label class="label_level_2">
    {{ Translate("label1") }}
  </label>
</div>

</template>

<style scoped>

  .label_level_1 {
    @apply text-cyan-600 text-5xl;
    font-family: 'Comfortaa', cursive;
  }

  .label_level_2 {
    @apply text-black text-xl;
    font-family: 'Comfortaa', cursive;
  }
</style>