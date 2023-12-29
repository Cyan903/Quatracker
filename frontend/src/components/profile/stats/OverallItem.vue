<template>
    <div>
        <div class="divider my-8">
            <PresentationChartLineIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Overall Stats</span>
        </div>

        <div class="overflow-x-auto stat-pop">
            <table class="table m-auto sm:w-[75%] lg:max-w-[800px]">
                <tbody>
                    <tr v-if="diffAcc">
                        <th>Ranked Accuracy</th>
                        <td>
                            {{
                                stats.overall.RankedAccuracy.toFixed(2) + "%" ||
                                "Unknown"
                            }}
                        </td>
                    </tr>

                    <tr>
                        <th>Overall Accuracy</th>
                        <td>
                            {{
                                stats.overall.Accuracy.toFixed(2) + "%" ||
                                "Unknown"
                            }}
                        </td>
                    </tr>

                    <tr>
                        <th>Overall Performance</th>
                        <td>
                            {{
                                useComma(
                                    stats.overall.Performance.toFixed(2),
                                ) || "Unknown"
                            }}
                        </td>
                    </tr>

                    <tr>
                        <th>Total Hits</th>
                        <td>
                            {{ useComma(stats.total.Hits) || "Unknown" }}
                        </td>
                    </tr>

                    <tr>
                        <th>Total Score</th>
                        <td>{{ useComma(stats.total.Score) || "Unknown" }}</td>
                    </tr>

                    <tr>
                        <th>Total Playcount</th>
                        <td>
                            {{ useComma(stats.plays.Playcount) || "Unknown" }}
                        </td>
                    </tr>

                    <tr>
                        <th>Passed/Failed</th>
                        <td
                            v-if="
                                stats.plays.Failed !== 0 &&
                                stats.plays.Playcount !== 0
                            "
                        >
                            <span>{{
                                useComma(stats.plays.Passed) || "Unknown"
                            }}</span>

                            <span>/</span>

                            <span>{{
                                useComma(stats.plays.Failed) || "Unknown"
                            }}</span>

                            <span>
                                ({{
                                    (
                                        (stats.plays.Failed /
                                            stats.plays.Playcount) *
                                        100
                                    ).toFixed(2)
                                }}%)</span
                            >
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { OverallStats, PlayStats, TotalStats } from "@/types/stats";
import { PresentationChartLineIcon } from "@heroicons/vue/24/solid";
import { onMounted, watch, reactive, computed } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { useComma } from "@/use/useUtil";
import { getOverall, getPlays, getTotal } from "@/use/useStats";

const cfg = useConfigStore();
const toast = useToast();

const props = defineProps<{
    id: number;
    mode: number;
}>();

const stats = reactive({
    overall: {
        Accuracy: 0,
        RankedAccuracy: 0,
        Performance: 0,
    } as OverallStats,

    total: {
        Hits: 0,
        Score: 0,
    } as TotalStats,

    plays: {
        Playcount: 0,
        Passed: 0,
        Failed: 0,
    } as PlayStats,
});

const overall = async () => {
    const data = await getOverall(props.id, props.mode);

    if (data.error && data.error !== "no scores") {
        toast.error("Could not get overall stats!");

        return;
    }

    if (!data.result) {
        toast.error("No overall stats!");
        return;
    }

    stats.overall = data.result;
};

const total = async () => {
    const data = await getTotal(props.id, props.mode);

    if (data.error && data.error !== "no scores") {
        toast.error("Could not get total count!");

        return;
    }

    if (!data.result) {
        toast.error("No total count!");
        return;
    }

    stats.total = data.result;
};

const plays = async () => {
    const data = await getPlays(props.id, props.mode);

    if (data.error) {
        toast.error("Could not get play stats!");
        return;
    }

    if (!data.result) {
        toast.error("No play stats!");
        return;
    }

    stats.plays = data.result;
};

const init = () => {
    if (!cfg.validConfig) return;

    stats.overall = { Accuracy: 0, RankedAccuracy: 0, Performance: 0 };
    stats.total = { Hits: 0, Score: 0 };
    stats.plays = { Playcount: 0, Passed: 0, Failed: 0 };

    overall();
    total();
    plays();
};

const diffAcc = computed(() => {
    return (
        stats.overall.Accuracy.toFixed(2) !==
        stats.overall.RankedAccuracy.toFixed(2)
    );
});

watch(() => props.id, init);
watch(() => props.mode, init);

onMounted(init);
</script>
