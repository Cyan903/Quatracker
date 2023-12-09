<template>
    <div>
        <ModalItem :show="modal" @hide="modal = false">
            <template #default>
                <h4 class="text-lg font-bold">Filter Options</h4>
                <select
                    class="select select-bordered w-full mt-4"
                    v-model="options.judgement"
                >
                    <option selected value="">Judgement preset (any)</option>
                    <option v-for="judge in judges" :key="judge" :value="judge">
                        {{ judge }}
                    </option>
                </select>

                <select
                    class="select select-bordered w-full my-2"
                    v-model="options.status"
                >
                    <option selected value="">Map rank status (any)</option>
                    <option value="Not Submitted">Not Submitted</option>
                    <option value="Unranked">Unranked</option>
                    <option value="Ranked">Ranked</option>
                    <option value="Dan Course">Dan Course</option>
                </select>

                <h4 class="text-sm font-bold my-2">% of long notes:</h4>
                <div class="flex items-center my-2">
                    <DelayedRange
                        min="0"
                        max="101"
                        class="range"
                        v-model="options.lnPercent"
                    />

                    <h4 class="text-xs mx-2 w-[100px] text-center">
                        {{ lnDisplay }}
                    </h4>
                </div>
            </template>

            <template #actions>
                <button class="btn btn-error btn-outline" @click="confirmReset">
                    Reset
                </button>
            </template>
        </ModalItem>

        <div class="divider">
            <StarIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Best Scores</span>
        </div>

        <div
            class="flex lg:max-w-[800px] lg:mx-auto justify-between align-middle"
        >
            <div class="font-bold text-xs [&>span]:after:content-['/'] [&>span]:after:mx-2">
                <span>{{ options.judgement || "Any Judgement" }}</span>
                <span>{{ options.status || "Any Rank Status" }}</span>
                <span class="after:!content-none">{{ lnDisplay }}</span>
            </div>

            <Cog6ToothIcon
                v-if="noScores"
                class="w-[1rem] h-[1rem] opacity-60"
            />
            <Cog6ToothIcon
                v-else
                class="w-[1rem] h-[1rem] hover:opacity-60 active:opacity-80"
                @click="showModal"
            />
        </div>

        <div>
            <h4 v-if="noScores" class="text-center my-[100px]">
                No scores found!
            </h4>
            <template v-else>
                <ScoreItem
                    v-for="score in scores"
                    :key="score.ScoreID"
                    :score="score"
                />

                <div class="text-center">
                    <button
                        :disabled="noScores"
                        class="btn btn-primary btn-outline max-md:w-full w-[50%] md:max-w-[700px]"
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
import type { Scores } from "@/types/scores";
import { StarIcon, Cog6ToothIcon } from "@heroicons/vue/24/solid";
import { ref, reactive, watch, onMounted, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { getBest } from "@/use/useScores";
import { getJudgements } from "@/use/useUsers";

import DelayedRange from "@/components/util/DelayedRange.vue";
import ModalItem from "@/components/util/ModalItem.vue";
import ScoreItem from "@/components/profile/scores/display/ScoreItem.vue";

const cfg = useConfigStore();
const toast = useToast();

const props = defineProps<{
    id: number;
    mode: boolean;
}>();

const scores = ref<Scores[]>([]);
const judges = ref<string[]>([]);
const modal = ref(false);
const page = ref(0);
const options = reactive({
    judgement: "",
    status: "",
    lnPercent: 101,
});

const gamemode = computed(() => {
    return props.mode ? 2 : 1;
});

const lnPercent = computed(() => {
    return options.lnPercent == 101 ? -1.0 : options.lnPercent;
});

const lnDisplay = computed(() => {
    return options.lnPercent >= 101 ? "Any LN%" : options.lnPercent + "%";
});

const noScores = computed(() => {
    return scores.value?.length == 0;
});

const best = async () => {
    const data = await getBest(
        props.id,
        gamemode.value,
        page.value,
        options.judgement,
        options.status,
        lnPercent.value,
    );

    if (data.error) {
        toast.error("Could not get best scores!");
        return;
    }

    if (!data.result) {
        toast.error("Reached end of scores listing!");
    }

    scores.value = [...scores.value, ...(data.result || [])];
};

const judgements = async () => {
    const data = await getJudgements();

    if (data.error) {
        toast.error("Could not get user judgements!");
        return;
    }

    judges.value = data.result || [];
};

const showModal = async () => {
    await judgements();
    modal.value = true;
};

// Data fetching...
const init = () => {
    if (!cfg.validConfig) return;

    options.judgement = "";
    options.status = "";
    options.lnPercent = 101;

    scores.value = [];
    judges.value = [];

    page.value = 0;

    best();
};

const reset = () => {
    scores.value = [];
    judges.value = [];
    page.value = 0;

    best();
    judgements();
};

const paginate = () => {
    page.value++;
    best();
};

// TODO: Don't use confirm
const confirmReset = () => {
    if (confirm("Are you sure?")) {
        init();
        modal.value = false;
    }
};

// Refreshing...
watch(() => props.id, init);
watch(() => props.mode, init);
watch(() => options, reset, { deep: true });
onMounted(init);
</script>
