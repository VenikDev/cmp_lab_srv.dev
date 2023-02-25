<script setup>
import SearchField from "../components/SearchField.vue";
import TitleLabel from "../components/TitleLabel.vue";
import NavBar from "../components/NavBar.vue";
import {ref} from "vue";
import ListItemAnalyses from "../components/ListItemAnalyses.vue";
// import LeftSideBar from "../components/LeftSideBar.vue";

let listAnalyses = ref({
  "citilab": {
    "Что-то там": 2000,
    "И что там еще": 1000
  },
  "gemotest": {
    "Что-то там": 1500,
    "И что там еще": 3000
  }
})

</script>

<template>
  <div class="main px-5">
    <!--  Выбор города (реализовать)  -->
    <nav-bar/>
    <!--  Заголовок  -->
    <title-label/>
    <!--  Поисковая строка и кнопка поиска  -->
    <search-field v-model="listAnalyses"/>
    <!--  Вывод всех анализов  -->
    <div v-for="item in listAnalyses"
         key="item"
    >
      <list-item-analyses
          :analysisData="item"
      />
    </div>
  </div>
</template>

<style scoped>
@tailwind base;
@tailwind components;
@tailwind utilities;

.main {
  @apply md:container md:mx-auto mx-4
}

.container_border {
  @apply my-6 shadow-md p-4 border-2 p-4 border-cyan-600 rounded-xl
}

.font_and_color {
  @apply text-cyan-600;
  font-family: 'Comfortaa', cursive;
}

</style>
