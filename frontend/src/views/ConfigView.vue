<template>
    <div>
        <h2 class="text-3xl font-bold my-4">Config</h2>

        <GamePathItem />
        <MainModeItem />
    </div>
</template>

<script lang="ts" setup>
import GamePathItem from "../components/config/GamePathItem.vue";
import MainModeItem from "../components/config/MainModeItem.vue";
</script>
