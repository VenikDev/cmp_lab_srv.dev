<script setup>

import {Translate} from "./translate.js";
import {ref} from "vue";
import axios from "axios";

let query = ref("");
let list = ref([])

const key = "url"
const value = `https://gemotest.ru/search/?q=${query}`
const url = `api/v1/analysis?${key}=${value}`
const getResponse = async () => {
  list = await axios.get(url)
      .then((response) => {
        return response.data
      })
}

</script>

<template>
  <div>
    <div class="container_border">
      <div class="my-2 content-center">
        <label class="label_level_2">
          {{ Translate("label2") }}
        </label>
      </div>
      <input
          v-model="query"
          class="input"
          id="search"
          type="text"
          :placeholder="Translate('field_for_search')"
      >
      <button class="button_for_search" @click="getResponse">
        {{ Translate("field_for_search") }}
      </button>
    </div>
    <!-- print all response -->
<!--    <div v-for="(idx, value) in list" key="idx">-->
<!--      <div class="container_border">-->
<!--        <h3>-->
<!--          {{ idx }}-->
<!--        </h3>-->
<!--        <span>-->
<!--        {{ value }}-->
<!--      </span>-->
<!--      </div>-->
<!--    </div>-->
  </div>
</template>

<style scoped>
.input {
  @apply shadow appearance-none border rounded w-9/12 py-2 px-3 text-gray-700 leading-tight focus:outline-none h-10 rounded-xl mr-2;
  font-family: 'Comfortaa', cursive;
}

.button_for_search {
  @apply text-cyan-600 w-1/5 border-2 border-cyan-600 h-10 rounded-xl active:bg-cyan-600 active:text-white
}

.label_level_2 {
  @apply text-black text-xl;
  font-family: 'Comfortaa', cursive;
}
</style>