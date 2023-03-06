<script setup>

const sizeBlockPopular = "200";
const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

</script>

<template>
  <div
      class="list_populars_blocks"
  >
    <div
        class="popular_block w-[200px] h-[200px] relative"
        key="idx"
        v-for="(idx, value) in arr"
    >
      <b>
        {{ value }}
      </b>
    </div>
  </div>
</template>

<style scoped>

.list_populars_blocks {
  @apply w-auto border-red-500 text-red-500 max-h-[250px];
}

.popular_block {
  @apply inline-block border-2 border-red-500 rounded-2xl mr-5 text-center;
}

</style>