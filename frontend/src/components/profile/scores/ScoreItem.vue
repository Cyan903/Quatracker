<template>
    <div>
        <!-- <input :value="JSON.stringify(score)" /> -->
        <pre>{{ score }}</pre>
    </div>
</template>

<script lang="ts" setup>
import type { Scores } from "../../../use/useScores";

defineProps<{ score: Scores }>();
</script>
