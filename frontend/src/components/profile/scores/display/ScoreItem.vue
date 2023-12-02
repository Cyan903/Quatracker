<template>
    <div>
        <pre>{{ score }}</pre>

        <button class="btn btn-outline" @click="setDetailId(score.ScoreID)">
            Set
        </button>
        <button
            class="btn btn-outline"
            @click="setDetailId(score.ScoreID + 1000)"
        >
            Set
        </button>
        <BackgroundMap :id="props.score.MapID" />
    </div>
</template>

<script lang="ts" setup>
import type { Scores } from "../../../../types/scores";
import { Ref, inject } from "vue";
import BackgroundMap from "./BackgroundMap.vue";

interface Detail {
    scoreDetailID: Ref<number>;
    setDetailId: (id: number) => void;
}

const props = defineProps<{
    score: Scores;
}>();

const { setDetailId } = inject<Detail>("scoreDetail") as Detail;
</script>
