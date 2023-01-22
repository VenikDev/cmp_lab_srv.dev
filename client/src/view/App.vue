<script setup>
import SearchField from "../components/SearchField.vue";
import TitleLabel from "../components/TitleLabel.vue";
import NavBar from "../components/NavBar.vue";
import LeftSideBar from "../components/LeftSideBar.vue";
</script>

<template>
  <div class="main px-5">
    <NavBar/>
    <TitleLabel/>
    <SearchField/>
    <LeftSideBar/>
  </div>
</template>

<style scoped>

.main {
  @apply md:container md:mx-auto
}

</style>
