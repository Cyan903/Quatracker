<template>
    <div>
        <NavbarBorder bg="bg-secondary" />
        <div class="p-4">
            <h2 class="text-3xl font-bold my-4">Configuration</h2>

            <GamePathItem />
            <div class="my-4 divider"></div>
            <MainModeItem />
        </div>
    </div>
</template>

<script lang="ts" setup>
import NavbarBorder from "@/components/NavbarBorder.vue";

import GamePathItem from "@/components/config/GamePathItem.vue";
import MainModeItem from "@/components/config/MainModeItem.vue";
</script>
