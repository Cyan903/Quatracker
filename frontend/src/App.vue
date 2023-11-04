<template>
    <Navbar />
    <RouterView />
</template>

<script lang="ts" setup>
import Navbar from "./components/Navbar.vue";
import { RouterView } from "vue-router";
</script>
