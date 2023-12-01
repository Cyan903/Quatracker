<template>
    <PageNotAvailable v-if="!cfg.validConfig" />
    <div v-else>
        <pre>{{ cfg.data.MainMode }}</pre>
        <div>
            <UserSwitcher
                :id="user.id"
                :username="user.username"
                @set-user="setUser"
            />

            <ModeSwitcher
                :disabled="user.id == -1"
                :mode="mode"
                @set-mode="(g) => (mode = g)"
            />
        </div>

        <div v-if="user.id != -1">
            <BestScores :id="user.id" :mode="mode" />
        </div>
        <h4 v-else>No users were found.</h4>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref, watch } from "vue";
import { useConfigStore } from "../store/config";

import PageNotAvailable from "../components/config/PageNotAvailable.vue";

import UserSwitcher from "../components/profile/UserSwitcher.vue";
import ModeSwitcher from "../components/profile/ModeSwitcher.vue";

import BestScores from "../components/profile/scores/BestScores.vue";

const cfg = useConfigStore();
const mode = ref(false);
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

onMounted(setMode);
watch(() => cfg.data.MainMode, setMode);
</script>
