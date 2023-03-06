<script setup>

import { Translate } from "./common/translate.js"

</script>

<template>

<div class="my-10">
  <h1 class="text-center">
    <div class="title">
      Инвестируйте
    </div>
    <div class="title mr-[-200px]">
      <span class="text-red-500">
        <b class="underline-offset-8 hover:underline-offset-4">в свое</b>
      </span>
    </div>
    <div class="title ml-[-200px]">
      здоровье
    </div>
  </h1>
  <br/>
  <div class="label_level_2">
    <h1 class="tag">
      <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="w-9 h-9 mr-2">
        <path
            stroke-linecap="round"
            d="M16.5 12a4.5 4.5 0 11-9 0 4.5 4.5 0 019 0zm0 0c0 1.657 1.007 3 2.25 3S21 13.657 21 12a9 9 0 10-2.636 6.364M16.5 12V8.25" />
      </svg>

      <b>
        {{ Translate("label1") }}
      </b>
    </h1>
  </div>
</div>

</template>

<style scoped>
  .title {
    @apply text-black text-8xl;
    font-family: 'Comfortaa', cursive;
  }

  .label_level_2 {
    @apply text-black text-3xl text-center mt-[-30px]
           inline-block;
    font-family: 'Comfortaa', cursive;
  }

  .tag {
    @apply bg-red-500 text-white rounded-3xl p-5 flex
           shadow-xl shadow-red-300 hover:shadow-none duration-300
  }
</style>