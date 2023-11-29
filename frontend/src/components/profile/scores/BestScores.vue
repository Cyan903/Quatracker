<template>
    <div>
        <ModalItem :show="modal" @hide="modal = false">
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
        </ModalItem>

        <button class="btn" @click="showModal">Settings</button>
        <div>
            <pre>{{ scores }}</pre>
            <button class="btn btn-secondary" @click="best(true)">Load</button>
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

const best = async (pageAdd: boolean) => {
    if (pageAdd) page.value++;

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

    if (pageAdd) {
        scores.value = [...scores.value, ...(data.result || [])];

        return;
    }

    scores.value = data.result || [];
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

const init = () => {
    if (!cfg.validConfig) return;

    options.judgement = "";
    options.status = "";
    options.lnPercent = 101;

    best(false);
};

watch(() => props.id, init);
watch(() => props.mode, init);
watch(() => options, init, { deep: true });
onMounted(init);
</script>
