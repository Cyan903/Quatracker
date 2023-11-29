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
                    <option v-for="judge in judges" :key="judge">
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

                    <h4 class="text-xs mx-2">
                        {{
                            options.lnPercent >= 101
                                ? "any"
                                : options.lnPercent + "%"
                        }}
                    </h4>
                </div>
            </template>

            <template #actions>
                <button class="btn btn-error btn-outline" @click="confirmReset">
                    Reset
                </button>
            </template>
        </ModalItem>

        <div>
            <h4 class="font-bold text-md">Best Scores</h4>
            <button class="btn" @click="showModal">Settings</button>
        </div>
        <div>
            <h4 v-if="noScores" class="text-center my-4">No scores found!</h4>
            <template v-else>
                <ScoreItem
                    v-for="score in scores"
                    :key="score.ScoreID"
                    :score="score"
                />
            </template>

            <button
                :disabled="noScores"
                class="btn btn-secondary"
                @click="paginate"
            >
                Load
            </button>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, watch, onMounted, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "../../../store/config";
import { Scores, getBest } from "../../../use/useScores";
import { getJudgements } from "../../../use/useUsers";

import DelayedRange from "../../util/DelayedRange.vue";
import ModalItem from "../../util/ModalItem.vue";
import ScoreItem from "./ScoreItem.vue";

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
    return options.lnPercent == 101 ? -1.0 : options.lnPercent / 100;
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
