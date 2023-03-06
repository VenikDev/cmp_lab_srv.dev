<script setup lang="ts">
import {ref, watch} from "vue";
import axios from "axios";
import {URL} from "./common/url.js";
import {Translate} from "./common/translate.js";

let query = ref("");

const props = defineProps({
  modelValue: {
    type: String,
  }
});

const emit = defineEmits(['update:modelValue']);

const list = ref<string>(props.modelValue);

watch(list, (newValue) => {
  emit('update:modelValue', newValue);
});

watch(() => props.modelValue, (newValue) => {
  list.value = newValue;
});

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

}

</script>

<template>
  <div class="container_border">
    <div class="my-2 content-center">
      <label class="label_level_2">
        <b>
          {{ Translate("label2") }}
        </b>
      </label>
    </div>
    <div class="search_line">
      <div>
        <svg xmlns="http://www.w3.org/2000/svg"
             fill="none"
             viewBox="0 0 24 24"
             stroke-width="1.5"
             stroke="currentColor"
             class="w-8 h-8 text-black text-center mx-4"
        >
          <path stroke-linecap="round"
                stroke-linejoin="round"
                d="M9.75 3.104v5.714a2.25 2.25 0 01-.659 1.591L5 14.5M9.75 3.104c-.251.023-.501.05-.75.082m.75-.082a24.301 24.301 0 014.5 0m0 0v5.714c0 .597.237 1.17.659 1.591L19.8 15.3M14.25 3.104c.251.023.501.05.75.082M19.8 15.3l-1.57.393A9.065 9.065 0 0112 15a9.065 9.065 0 00-6.23-.693L5 14.5m14.8.8l1.402 1.402c1.232 1.232.65 3.318-1.067 3.611A48.309 48.309 0 0112 21c-2.773 0-5.491-.235-8.135-.687-1.718-.293-2.3-2.379-1.067-3.61L5 14.5" />
        </svg>

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
  </div>
</template>

<style scoped>
.input {
  @apply shadow appearance-none border rounded w-9/12 py-2 px-3
         text-gray-700 leading-tight focus:outline-none
         h-10 rounded-xl mr-2;
  font-family: 'Comfortaa', cursive;
}

.button_for_search {
  @apply text-red-500 w-1/5 border-2 border-red-500 h-10 rounded-xl duration-300 active:bg-red-500 active:text-white
}

.label_level_2 {
  @apply text-red-500 text-2xl;
  font-family: 'Comfortaa', cursive;
}

.search_line {
  @apply flex
}
</style>