<template>
    <PageNotAvailable v-if="!cfg.validConfig" />
    <div v-else>
        <NavbarBorder bg="bg-primary" />

        <div
            class="p-4 bg-base-200 shadow-lg sticky w-full z-[10] top-[81px] rounded-b-xl"
        >
            <div class="flex">
                <StatsSwitcher
                    :disabled="user.id == -1"
                    :stats="stats"
                    @set-stats="(n) => (stats = n)"
                />

                <ModeSwitcher
                    :disabled="user.id == -1"
                    :mode="mode"
                    @set-mode="(g) => (mode = g)"
                />

                <UserSwitcher
                    :id="user.id"
                    :username="user.username"
                    @set-user="setUser"
                />
            </div>
        </div>

        <div v-if="user.id != -1">
            <DetailPopup
                :id="scoreDetailID"
                :open="detailOpen"
                @hide="detailOpen = false"
            />

            <div v-if="stats" class="px-4">
                <OverallItem :id="user.id" :mode="mode ? 2 : 1" />
                <MapsItem :id="user.id" :mode="mode ? 2 : 1" />
            </div>
            <div v-else class="px-4">
                <BestScores :id="user.id" :mode="mode" />
                <RecentScores class="my-[80px]" :id="user.id" :mode="mode" />
            </div>
        </div>
        <h4 v-else class="font-bold text-center my-4">No users were found.</h4>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, provide, reactive, ref, watch } from "vue";
import { useConfigStore } from "@/store/config";

import NavbarBorder from "@/components/NavbarBorder.vue";
import PageNotAvailable from "@/components/config/PageNotAvailable.vue";

import UserSwitcher from "@/components/profile/switcher/UserSwitcher.vue";
import ModeSwitcher from "@/components/profile/switcher/ModeSwitcher.vue";
import StatsSwitcher from "@/components/profile/switcher/StatsSwitcher.vue";

import DetailPopup from "@/components/profile/scores/display/DetailPopup.vue";
import BestScores from "@/components/profile/scores/BestScores.vue";
import RecentScores from "@/components/profile/scores/RecentScores.vue";

import OverallItem from "@/components/profile/stats/OverallItem.vue";
import MapsItem from "@/components/profile/stats/MapsItem.vue";

const cfg = useConfigStore();
const user = reactive({
    id: -1,
    username: "",
});

const mode = ref(false);
const stats = ref(false);
const detailOpen = ref(false);
const scoreDetailID = ref(0);

const setUser = (id: number, name: string) => {
    user.id = id;
    user.username = name;
};

const setMode = () => {
    if (cfg.data?.MainMode) {
        mode.value = cfg.data.MainMode;
    }
};

const setDetailId = (id: number) => {
    detailOpen.value = true;
    scoreDetailID.value = id;
};

provide("scoreDetail", {
    scoreDetailID,
    setDetailId,
});

onMounted(setMode);
watch(() => cfg.data.MainMode, setMode);
</script>
