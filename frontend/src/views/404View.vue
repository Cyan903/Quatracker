<template>
    <div class="text-center">
        <NavbarBorder bg="bg-error" />

        <div class="my-4 p-4 inline-block">
            <h2 class="text-3xl font-bold">Not found!</h2>
            <p class="my-4">Page was not found. Go home?</p>

            <router-link class="btn w-full" to="/">Back</router-link>
        </div>
    </div>
</template>

<script lang="ts" setup>
import NavbarBorder from "@/components/NavbarBorder.vue";
</script>
