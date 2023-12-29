<template>
    <div>
        <div class="divider my-8">
            <DocumentChartBarIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Map Stats</span>
        </div>

        <div class="m-auto sm:w-[75%] lg:max-w-[800px] stat-pop">
            <h2 class="my-4">
                <ArrowPathIcon
                    class="w-[1.3rem] h-[1.3rem] mr-1 inline-block"
                />
                <span class="font-bold align-middle text-xl">Most Played</span>
            </h2>

            <LoadingScore
                v-if="loading.most || stats.most.length == 0"
                :n="5"
            />
            <template v-else>
                <MostPlayedMap
                    v-for="m in stats.most"
                    :key="m.MapID"
                    :map="m"
                />
            </template>

            <div class="max-sm:text-center text-right my-4">
                <PaginateItem
                    color="btn-primary"
                    :disable="!cfg.validConfig"
                    :page="page"
                    @update="paginate"
                />
            </div>

            <div class="divider"></div>
        </div>

        <div class="m-auto sm:w-[75%] lg:max-w-[800px] stat-pop">
            <div class="tabs tabs-lifted">
                <a
                    class="tab"
                    :class="{ 'tab-active font-bold ': !tabs }"
                    @click="tabs = false"
                >
                    <span class="max-sm:hidden mr-1">Highest</span>
                    <span>Combo</span>
                </a>
                <a
                    class="tab"
                    :class="{ 'tab-active font-bold': tabs }"
                    @click="tabs = true"
                >
                    <span class="max-sm:hidden mr-1">Most</span>
                    <span>Hits</span>
                </a>
            </div>

            <div
                v-if="!tabs"
                class="bg-base-100 border-base-300 border-[1px] border-t-0 border-r-[1px] rounded-b-box p-6"
            >
                <LoadingScore
                    v-if="loading.longest || stats.longest.Combo.Total == -1"
                    :n="1"
                />
                <MapsItem v-else :map="stats.longest.Combo" text="Combo" />
            </div>
            <div
                v-else
                class="bg-base-100 border-base-300 border-[1px] border-t-0 border-r-[1px] rounded-b-box p-6"
            >
                <LoadingScore
                    v-if="loading.longest || stats.longest.Hits.Total == -1"
                    :n="1"
                />
                <MapsItem v-else :map="stats.longest.Hits" text="Hits" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { LongestStats, MostPlayed } from "@/types/stats";
import { DocumentChartBarIcon, ArrowPathIcon } from "@heroicons/vue/24/solid";
import { onMounted, watch, reactive, ref } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { getLongest, getMost } from "@/use/useStats";

import MapsItem from "@/components/profile/stats/display/LongestMap.vue";
import MostPlayedMap from "@/components/profile/stats/display/MostPlayedMap.vue";
import LoadingScore from "@/components/profile/scores/display/LoadingScore.vue";
import PaginateItem from "@/components/util/PaginateItem.vue";

const cfg = useConfigStore();
const toast = useToast();

const props = defineProps<{
    id: number;
    mode: number;
}>();

const emptyLongest = {
    Hits: {
        ScoreID: -1,
        MapID: -1,
        Map: {
            Title: "Unknown",
            DifficultyName: "Unknown",
        },

        Total: -1,
    },

    Combo: {
        ScoreID: -1,
        MapID: -1,
        Map: {
            Title: "Unknown",
            DifficultyName: "Unknown",
        },

        Total: -1,
    },
};

const page = ref(0);
const tabs = ref(false);
const loading = reactive({
    longest: true,
    most: true,
});

const stats = reactive({
    longest: structuredClone(emptyLongest) as LongestStats,
    most: [] as MostPlayed[],
});

const longest = async () => {
    const data = await getLongest(props.id, props.mode);

    if (data.error && data.error !== "no scores") {
        toast.error("Could not get longest maps!");

        return;
    }

    if (!data.result) {
        toast.error("Nothing found for longest maps!");
        return;
    }

    stats.longest = data.result;
};

const most = async () => {
    const data = await getMost(props.id, props.mode, page.value);

    if (data.error && data.error !== "no scores") {
        toast.error("Could not get most played!");

        return;
    }

    if (!data.result) {
        toast.error("Reached end of most played!");
    }

    stats.most = data.result || [];
};

const init = async () => {
    if (!cfg.validConfig) return;

    loading.longest = true;
    loading.most = true;

    stats.longest = structuredClone(emptyLongest);
    stats.most = [];

    page.value = 0;

    await longest();
    await most();

    loading.longest = false;
    loading.most = false;
};

const paginate = async (n: number) => {
    page.value += n;
    loading.most = true;

    await most();

    loading.most = false;
};

watch(() => props.id, init);
watch(() => props.mode, init);

onMounted(init);
</script>
