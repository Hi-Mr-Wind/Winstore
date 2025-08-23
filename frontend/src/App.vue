<script setup lang="ts">
import StoreHome from "@/components/StoreHome.vue";
</script>

<template>
  <div style="width: 500px">
    <router-link to="/">Home  </router-link>

    <router-link to="/store">  Store</router-link>

    <router-link to="/about">  About</router-link>
  </div>
  <router-view/>
<!--  <StoreHome/>-->
</template>
<style scoped></style>
