<script setup>
import SearchField from "../components/SearchField.vue";
import TitleLabel from "../components/TitleLabel.vue";
import NavBar from "../components/NavBar.vue";
// import LeftSideBar from "../components/LeftSideBar.vue";
</script>

<template>
  <div class="main px-5">
    <NavBar/>
    <TitleLabel/>
    <SearchField/>
<!--    <LeftSideBar/>-->
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

</style>
