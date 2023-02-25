<script setup>
import ScrollArea from "./ScrollArea.vue";
import ListItem from "./ListItem.vue";

const {} = defineProps({
  list: Object
})

</script>

<template>
  <!-- для бордюра и отступов -->
  <div class="container_border scroll">
    <div class="grid grid-cols-3">
      <div
          class="container_border mx-4"
          v-for="item in Object.keys(list)"
          key="item"
      >
        <h1 class="font_and_color">
          {{ item }}
        </h1>
        <div v-for="analyses in list[item]">
          <list-item :analysisData="analyses"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.scroll {
  width: 1000px;
  overflow: scroll;
  max-height: 1000px;
}

.scroll::-webkit-scrollbar-track {
  @apply bg-cyan-600
}

</style>