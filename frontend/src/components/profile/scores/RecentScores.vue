<template>
    <div>
        <div class="divider">
            <ArrowUturnLeftIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Recent Scores</span>
        </div>

        <div
            class="flex lg:max-w-[800px] lg:mx-auto justify-between align-middle"
        >
            <div class="text-xs">
                <span class="font-bold">Total Scores:</span>
                <span class="ml-[3px]">{{ useComma(scores.Total) }}</span>
            </div>

            <button
                class="btn btn-xs btn-outline"
                :class="hide ? 'btn-success' : 'btn-error'"
                :disabled="noScores"
                @click="hide = !hide"
            >
                Failed
            </button>
        </div>

        <div>
            <h4 v-if="noScores" class="text-center my-[100px]">
                No scores found!
            </h4>
            <template v-else>
                <ScoreItem
                    v-for="score in scores.Scores"
                    :key="score.ScoreID"
                    :score="score"
                />

                <div class="text-center">
                    <button
                        :disabled="noScores"
                        class="btn btn-secondary btn-outline max-md:w-full w-[50%] md:max-w-[700px]"
                        @click="paginate"
                    >
                        Load
                    </button>
                </div>
            </template>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { CountedScores } from "@/types/scores";
import { ArrowUturnLeftIcon } from "@heroicons/vue/24/solid";
import { ref, watch, onMounted, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { useComma } from "@/use/useUtil";
import { getRecent } from "@/use/useScores";

import ScoreItem from "@/components/profile/scores/display/ScoreItem.vue";

const cfg = useConfigStore();
const toast = useToast();

const props = defineProps<{
    id: number;
    mode: boolean;
}>();

const scores = ref<CountedScores>({
    Total: 0,
    Scores: [],
});

const page = ref(0);
const hide = ref(false);

const gamemode = computed(() => {
    return props.mode ? 2 : 1;
});

const noScores = computed(() => {
    return scores.value?.Scores.length == 0;
});

const recent = async () => {
    const data = await getRecent(
        props.id,
        gamemode.value,
        page.value,
        hide.value,
    );

    if (data.error) {
        toast.error("Could not get recent scores!");
        return;
    }

    if (!data.result?.Scores) {
        toast.error("Reached end of scores listing!");
    }

    scores.value.Total = data.result?.Total || 0;
    scores.value.Scores = [
        ...scores.value.Scores,
        ...(data.result?.Scores || []),
    ];
};

// Data fetching...
const init = () => {
    if (!cfg.validConfig) return;

    page.value = 0;
    scores.value = {
        Total: 0,
        Scores: [],
    };

    recent();
};

const paginate = () => {
    page.value++;
    recent();
};

// Refreshing...
watch(() => props.id, init);
watch(() => props.mode, init);
watch(hide, () => {
    page.value = 0;
    scores.value = {
        Total: 0,
        Scores: [],
    };

    recent();
});

onMounted(init);
</script>
