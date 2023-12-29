<template>
    <div>
        <div class="divider my-8">
            <BackwardIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Playcount History</span>
        </div>

        <div class="m-auto sm:w-[75%] lg:max-w-[800px] stat-pop h-[350px]">
            <div
                v-if="loading"
                class="skeleton h-[350px] bg-base-200 rounded-lg"
            ></div>
            <Line v-else :data="chartStats" :options="chartOptions" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { PlayHistory } from "@/types/stats";
import { BackwardIcon } from "@heroicons/vue/24/solid";
import { onMounted, watch, ref, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { getHistory } from "@/use/useStats";

import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from "chart.js";
import { Line } from "vue-chartjs";

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
);
const cfg = useConfigStore();
const toast = useToast();
const loading = ref(true);

const props = defineProps<{
    id: number;
    mode: number;
}>();

const chartOptions = ref({
    responsive: true,
    maintainAspectRatio: false,
});

const stats = ref<PlayHistory>({});

const history = async () => {
    const data = await getHistory(props.id, props.mode);

    if (data.error) {
        toast.error("Could not get history!");
        return;
    }

    if (!data.result) {
        toast.error("No history|");
        return;
    }

    stats.value = data.result;
};

const init = async () => {
    if (!cfg.validConfig) return;

    loading.value = true;
    stats.value = {};

    await history();
    loading.value = false;
};

const chartStats = computed(() => {
    return {
        labels: Object.keys(stats.value),
        datasets: [
            {
                label: "Playcount",
                backgroundColor: "#f87979",
                data: Object.values(stats.value),
            },
        ],
    };
});

watch(() => props.id, init);
watch(() => props.mode, init);

onMounted(init);
</script>
