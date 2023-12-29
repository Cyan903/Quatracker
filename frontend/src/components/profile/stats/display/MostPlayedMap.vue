<template>
    <div
        class="border sm:h-[60px] my-4 rounded-md border-base-300 border-[1px] flex most-map shadow-[0_1px_0_0_#000]"
        :style="img"
    >
        <div
            class="bg-gray-700 h-full w-[80px] rounded-l-md max-sm:hidden mr-2"
            :style="img"
        ></div>

        <div class="max-sm:p-2">
            <h2 class="font-bold text-lg max-sm:mb-1">
                {{ map.Title }} [{{ map.DifficultyName }}]
            </h2>

            <div class="text-xs">
                <span class="mr-1 font-bold">Creator:</span>
                <span>{{ map.Creator }}</span>
            </div>

            <div class="text-xs mb-2 sm:hidden">
                <span class="mr-1 font-bold">Playcount:</span>
                <span>{{ useComma(props.map.PlayCount || 0) }}</span>
            </div>
        </div>

        <div
            class="ml-auto min-w-[80px] font-bold flex justify-center items-center text-xl mr-2 max-sm:hidden"
        >
            {{ useComma(props.map.PlayCount || 0) }}x
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { MostPlayed } from "@/types/stats";
import { ref, onMounted } from "vue";
import { useImage } from "@/use/useImage";
import { useComma } from "@/use/useUtil";

const props = defineProps<{ map: MostPlayed }>();
const img = ref("");

onMounted(async () => {
    const i = await useImage(props.map.MapID);

    img.value = `
        background: linear-gradient(rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.8)), url("${i}") no-repeat top;
        background-size: cover;
    `;
});
</script>

<style scoped>
/* sm */
@media (min-width: 640px) {
    .most-map {
        background: transparent !important;
    }
}
</style>
