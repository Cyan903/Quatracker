<template>
    <div>
        <pre>{{ score }}</pre>

        <BackgroundMap :id="score.MapID" />
    </div>
</template>

<script lang="ts" setup>
import type { Scores } from "../../../../use/useScores";
import BackgroundMap from "./BackgroundMap.vue";

defineProps<{ score: Scores }>();
</script>
