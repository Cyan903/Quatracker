<template>
    <PageNotAvailable v-if="!cfg.validConfig" />
    <div v-else>
        <div class="bg-base-100 sticky w-full z-[10] top-[80px]">
            <div class="p-4 flex">
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

            <div class="w-full h-[1px] bg-accent"></div>
        </div>

        <div class="relative top-4">
            <div v-if="user.id != -1">
                <DetailPopup
                    :id="scoreDetailID"
                    :open="detailOpen"
                    @hide="detailOpen = false"
                />

                <div class="px-4">
                    <BestScores :id="user.id" :mode="mode" />
                    <RecentScores
                        class="my-[80px]"
                        :id="user.id"
                        :mode="mode"
                    />
                </div>
            </div>
            <h4 v-else class="font-bold text-center my-4">
                No users were found.
            </h4>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, provide, reactive, ref, watch } from "vue";
import { useConfigStore } from "@/store/config";

import PageNotAvailable from "@/components/config/PageNotAvailable.vue";

import UserSwitcher from "@/components/profile/switcher/UserSwitcher.vue";
import ModeSwitcher from "@/components/profile/switcher/ModeSwitcher.vue";

import DetailPopup from "@/components/profile/scores/display/DetailPopup.vue";
import BestScores from "@/components/profile/scores/BestScores.vue";
import RecentScores from "@/components/profile/scores/RecentScores.vue";

const cfg = useConfigStore();
const mode = ref(false);
const scoreDetailID = ref(0);
const detailOpen = ref(false);
const user = reactive({
    id: -1,
    username: "",
});

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
