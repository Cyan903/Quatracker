<template>
    <div
        class="my-4 rounded-lg lg:max-w-[800px] lg:mx-auto min-h-[80px] p-2 longest-map sm:flex items-center shadow-lg"
        :style="img"
    >
        <div class="m-4 text-center max-sm:hidden">
            <h2 class="font-bold text-2xl">{{ useComma(map.Total || 0) }}x</h2>
            <h4>{{ text }}</h4>
        </div>
        <div class="w-full ml-2">
            <h2 class="text-xl font-bold">{{ map.Map.Title }}</h2>
            <h4 class="font-bold">[{{ map.Map.DifficultyName }}]</h4>
            <div class="sm:hidden text-sm mt-2">
                <span class="font-bold mr-1">{{ text }}:</span>
                <span>{{ useComma(map.Total || 0) }}x</span>
            </div>
            
            <button
                class="btn btn-xs btn-info mt-2"
                @click="setDetailId(props.map.ScoreID)"
            >
                Details
            </button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { LongestStatsCombo } from "@/types/stats";
import { type Ref, inject, onMounted, ref } from "vue";
import { useImage } from "@/use/useImage";
import { useComma } from "@/use/useUtil";

interface Detail {
    scoreDetailID: Ref<number>;
    setDetailId: (id: number) => void;
}

const { setDetailId } = inject<Detail>("scoreDetail") as Detail;
const props = defineProps<{ map: LongestStatsCombo; text: string }>();
const img = ref("");

onMounted(async () => {
    const i = await useImage(props.map.MapID);

    img.value = `
        background: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.8)), url("${i}") no-repeat top;
        background-size: cover;
    `;
});
</script>
