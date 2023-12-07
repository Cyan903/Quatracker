<template>
    <div v-if="score && score.ScoreID">
        <ModalItem :show="open" @hide="emits('hide')">
            <template #default>
                <h4>{{ props.id }}</h4>
                <pre>{{ score }}</pre>
            </template>
        </ModalItem>
    </div>
</template>

<script lang="ts" setup>
import type { ScoreDetails } from "@/types/scores";
import { ref, watch } from "vue";
import { useToast } from "vue-toastification";
import { getDetails } from "@/use/useScores";

import ModalItem from "@/components/util/ModalItem.vue";

const props = defineProps<{
    id: number;
    open: boolean;
}>();

const emits = defineEmits<{
    (e: "hide"): void;
}>();

const toast = useToast();
const score = ref<ScoreDetails>();
const loaded = ref(false);

const init = async () => {
    const data = await getDetails(props.id);

    if (data.error || !data.result || !data.result?.ScoreID) {
        toast.error("Could not get score details!");
        emits("hide");

        return;
    }

    score.value = data.result;
};

watch(
    () => props.open,
    async () => {
        loaded.value = false;

        if (props.open) {
            await init();
        }

        loaded.value = true;
    },
);
</script>
