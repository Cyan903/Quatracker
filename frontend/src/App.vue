<template>
    <NavbarItem />

    <router-view v-slot="{ Component }">
        <transition name="router" mode="out-in">
            <component :is="Component"></component>
        </transition>
    </router-view>
</template>

<script lang="ts" setup>
import { RouterView } from "vue-router";
import NavbarItem from "./components/NavbarItem.vue";
</script>

<style scoped>
.router-enter-from {
    opacity: 0;
}

.router-enter-active {
    transition: all 0.1s ease-in-out;
}
</style>
