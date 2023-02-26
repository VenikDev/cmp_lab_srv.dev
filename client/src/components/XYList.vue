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
          class="mx-4 size_col"
          v-for="nameLab in Object.keys(list)"
          key="item"
      >
        <div class="font_and_color title_lab">
          {{ nameLab }}
        </div>
        <div v-for="analyses in list[nameLab]">
          <list-item :analysisData="analyses"/>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.scroll {
  overflow: scroll;
  height: auto;
  width: 100%;
}

.size_col {
  min-width: 300px;
  width: auto;
}

.title_lab {
  @apply text-2xl bg-PompPower text-white p-4 rounded-xl mt-2
}

</style>