<script setup>
// vue
import { ref } from "vue";

// common
import { objectIsEmpty } from "../components/common/algo.js";

// components
import SearchField from "../components/SearchField.vue";
import TitleLabel from "../components/TitleLabel.vue";
import NavBar from "../components/NavBar.vue";
import XYList from "../components/XYList.vue";
import PopularItems from "../components/PopularItems/PopularItems.vue";

// data
let listAnalyses = ref({})

</script>

<template>
  <div class="main px-5">
    <!--  Выбор города (реализовать)  -->
    <nav-bar/>
    <!--  Заголовок  -->
    <title-label/>
    <!--  Чаще всего ищут  -->
<!--    <popular-items/>-->
    <!--  Поисковая строка и кнопка поиска  -->
    <search-field v-model="listAnalyses"/>
    <!--  Вывод всех анализов  -->
    <x-y-list
        v-if="!objectIsEmpty(listAnalyses)"
        :list="listAnalyses"
    />
  </div>
</template>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;

.main {
  @apply md:container md:mx-auto mx-4;
}

.container_border {
  @apply my-6 p-4 border-2 p-4
  shadow-md border-red-500 rounded-xl;
}

.font_and_color {
  @apply text-cyan-600;
  font-family: 'Comfortaa', cursive;
}

</style>
