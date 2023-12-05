<template>
    <div
        :style="img"
        class="flex items-center my-4 rounded-md min-h-[80px] py-2"
    >
        <div class="text-center mx-4">
            <h2
                class="text-4xl font-bold"
                :class="`grade-${score.Score.Grade}`"
            >
                {{ score.Score.Grade }}
            </h2>

            <h4 class="text-sm font-bold mt-2 md:hidden">
                {{ score.Score.PerformanceRating.toFixed(2) }}
            </h4>
        </div>

        <div class="w-full">
            <h2
                class="font-bold hover:underline hover:opacity-80 md:hidden"
                :title="fullTitle"
                @click="setDetailId(score.ScoreID)"
            >
                {{ shorten(score.Map.Title, 25) }} [{{
                    shorten(score.Map.DifficultyName, 15)
                }}]
            </h2>

            <h2
                class="font-bold hover:underline hover:opacity-80 max-md:hidden"
                :title="fullTitle"
                @click="setDetailId(score.ScoreID)"
            >
                {{ shorten(fullTitle, 50) }}
            </h2>

            <div class="text-xs my-1 md:hidden">
                <span
                    class="font-bold text-sm"
                    :title="`${score.Score.RankedAccuracy.toFixed(2)}%`"
                >
                    {{ score.Score.Accuracy.toFixed(2) }}% ({{
                        score.Score.JudgementWindowPreset
                    }})
                </span>

                <span>+ {{ mods.join(", ") }}</span>
            </div>

            <div class="text-xs md:my-2">
                <span
                    class="font-bold"
                    :class="useDifficulty(score.Map.Rating)"
                >
                    {{ score.Map.Rating.toFixed(2) }}
                </span>

                <span class="font-bold">
                    <span class="mx-1"> / </span>
                    <span
                        class="map-rank"
                        :class="useRank(score.Map.RankedStatus)"
                    >
                        {{ score.Map.RankedStatus }}</span
                    >
                    <span class="mx-1"> / </span>
                </span>

                <span>
                    <span class="font-bold">LN:</span>
                    {{ score.Map.LNPercent.toFixed(1) }}%
                </span>
            </div>

            <div class="divider my-0 md:hidden"></div>

            <h4
                class="text-xs opacity-80 text-[0.6em] md:mt-2"
                :title="score.Score.DateTime"
            >
                {{ score.Score.DateTime }}
            </h4>
        </div>

        <div class="max-md:hidden w-[15%] text-center">
            <h4 v-for="m in mods" :key="m" class="font-bold text-sm">
                {{ m }}
            </h4>
        </div>

        <div class="max-md:hidden w-[25%] px-8 text-center">
            <h2 class="text-lg font-bold mt-2 w-[100%]">
                {{ score.Score.PerformanceRating.toFixed(2) }}
            </h2>

            <div class="font-bold text-xs">
                <h4
                    class="my-1"
                    :title="`${score.Score.RankedAccuracy.toFixed(2)}%`"
                >
                    {{ score.Score.Accuracy.toFixed(2) }}%
                </h4>

                <h4>{{ score.Score.JudgementWindowPreset }}</h4>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { Scores } from "../../../../types/scores";
import { Ref, inject, onMounted, ref, computed } from "vue";
import { useImage } from "../../../../use/useImage";
import { useDifficulty, useRank } from "../../../../use/useColors";

interface Detail {
    scoreDetailID: Ref<number>;
    setDetailId: (id: number) => void;
}

const { setDetailId } = inject<Detail>("scoreDetail") as Detail;
const props = defineProps<{ score: Scores }>();
const img = ref("");

const fullTitle = computed(
    () =>
        `${props.score.Map.Artist} - ${props.score.Map.Title} [${props.score.Map.DifficultyName}] (${props.score.Map.Creator})`,
);

// TODO: Convert mods
const mods = computed(() => {
    return [props.score.Score.Mods];
});

const shorten = (str: string, n: number) => {
    if (str.length - 3 > n) {
        return str.slice(0, n) + "...";
    }

    return str;
};

onMounted(async () => {
    const i = await useImage(props.score.MapID);

    img.value = `
        background: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.8)), url("${i}");
    `;
});
</script>

<style scoped>
tr {
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;
    border-radius: 10px;
}
</style>
../../../../use/useColors
