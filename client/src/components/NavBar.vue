<script setup>

import {Translate} from "./common/translate.js";
</script>


<template>
  <div class="nav_bar">
    <div>
      <button class="select_city_btn">
        {{ Translate("default_city") }}
      </button>
    </div>
    <div class="name_firm">
      <b class="text-red-600">Z</b>драв<b class="text-red-600">RU</b>
    </div>
  </div>
</template>

<style scoped>

.select_city_btn {
  @apply text-white text-left p-2;
  font-family: 'Comfortaa', cursive;
  text-decoration: underline
}

.name_firm {
  @apply text-right text-white text-3xl grow pl-2;
  font-family: 'Comfortaa', cursive;
}

.nav_bar {
  @apply flex my-3 shadow-md p-4 bg-ChineseViolet rounded-xl;
}

</style>