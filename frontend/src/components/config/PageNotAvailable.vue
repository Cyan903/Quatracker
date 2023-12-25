<template>
    <div class="text-center">
        <NavbarBorder bg="bg-error" />
        <div class="p-4 inline-block">
            <h2 class="font-bold text-3xl">Page is not available.</h2>
            <p class="my-4">
                Please update your configuration in the configuration tab.
            </p>

            <router-link to="/config" class="btn w-full">
                Configuration
            </router-link>
        </div>
    </div>
</template>

<script lang="ts" setup>
import NavbarBorder from "@/components/NavbarBorder.vue";
</script>
