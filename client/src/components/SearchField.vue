<script setup>
import {computed, ref, watch} from "vue";
import axios from "axios";
import {URL} from "./common/url.js";
import {Translate} from "./common/translate.js";

let query = ref("");

const props = defineProps({
  list: Object,
})
let list = ref(props.list)

const emits = defineEmits([
  "update:modalValue"
])

watch(list, () => {
  emits("update:modalValue", list)
})

/**
 * Установить новое ключевое слово в url
 */
const getURL = () => {
  return URL + `api/v1/analysis?key=${query.value}`
}

/**
 * Получить список анализов по искомому слову
 */
const getResponse = async () => {
  list.value = (await axios.get(getURL())).data
  console.log("list: ", list.value)
}

</script>

<template>
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