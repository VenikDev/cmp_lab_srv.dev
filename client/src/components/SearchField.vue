<template>
  <div class="w-full my-10 shadow-md p-4">
    <div class="my-2">
      <label class="label_level_2">Найти интересующие анализы</label>
    </div>
    <input class="input" id="search" type="text" placeholder="Найти">
  </div>
</template>

<script setup>

</script>

<style scoped>
@tailwind base;
@tailwind components;
@tailwind utilities;

.input {
  @apply shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none ;
  font-family: 'Comfortaa', cursive;
}

.label_level_2 {
  @apply text-black text-xl;
  font-family: 'Comfortaa', cursive;
}
</style>