<template>
    <div>
        <div class="divider my-8">
            <ArrowsUpDownIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Judges Count</span>
        </div>

        <div class="m-auto sm:w-[75%] lg:max-w-[800px] stat-pop">
            <div class="h-[350px]">
                <div v-if="loading">
                    <div
                        class="skeleton h-[3rem] mb-4 bg-base-200 rounded-lg"
                    ></div>
                    <div
                        class="skeleton h-[300px] bg-base-200 rounded-lg"
                    ></div>
                </div>
                <Pie v-else :data="chartStats" :options="chartOptions" />
            </div>

            <div class="overflow-x-auto stat-pop my-4">
                <table class="table m-auto sm:w-[75%]">
                    <tbody>
                        <tr v-for="(i, j) in stats" :key="j">
                            <th>{{ j.replace("Count", "Total ") }}</th>
                            <td>{{ useComma(i || 0) }}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { JudgesCount } from "@/types/stats";
import { ArrowsUpDownIcon } from "@heroicons/vue/24/solid";
import { onMounted, watch, ref, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { useComma } from "@/use/useUtil";
import { getJudges } from "@/use/useStats";

import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Pie } from "vue-chartjs";

ChartJS.register(ArcElement, Tooltip, Legend);

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

const stats = ref<JudgesCount>({
    CountMarv: 0,
    CountPerf: 0,
    CountGreat: 0,
    CountGood: 0,
    CountOkay: 0,
    CountMiss: 0,
});

const judges = async () => {
    const data = await getJudges(props.id, props.mode);

    if (data.error) {
        toast.error("Could not get judges!");
        return;
    }

    if (!data.result) {
        toast.error("No judges to count!");
        return;
    }

    stats.value = data.result;
};

const init = async () => {
    if (!cfg.validConfig) return;

    loading.value = true;
    stats.value = {
        CountMarv: 0,
        CountPerf: 0,
        CountGreat: 0,
        CountGood: 0,
        CountOkay: 0,
        CountMiss: 0,
    };

    await judges();
    loading.value = false;
};

const chartStats = computed(() => {
    return {
        labels: [
            "CountMarv",
            "CountPerf",
            "CountGreat",
            "CountGood",
            "CountOkay",
            "CountMiss",
        ].map((t) => {
            const percent =
                (stats.value[t as keyof typeof stats.value] /
                    Object.values(stats.value).reduce((s, a) => s + a, 0)) *
                100;

            return `Total ${t.replace("Count", "")} (${(isNaN(percent)
                ? 0
                : percent
            ).toFixed(2)}%)`;
        }),

        datasets: [
            {
                backgroundColor: [
                    "#f9f9da",
                    "#feff81",
                    "#369c4d",
                    "#29a5cd",
                    "#e547a0",
                    "#d62020",
                ],

                data: Object.values(stats.value),
            },
        ],
    };
});

watch(() => props.id, init);
watch(() => props.mode, init);

onMounted(init);
</script>
