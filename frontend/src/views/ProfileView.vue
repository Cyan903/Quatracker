<template>
    <div>
        <h2>Profile</h2>
        <PageNotAvailable v-if="!cfg.validConfig" />
        <template v-else>
            <button @click="c" class="btn">c</button>
            <button @click="r" class="btn">r</button>
            <button @click="d" class="btn">d</button>
            <button @click="u" class="btn">u</button>
            <button @click="j" class="btn">j</button>
            <input type="number" class="input" v-model="id" />
            <hr />

            <pre>{{ scores }}</pre>
            <pre>{{ cfg.data }}</pre>
            <pre>{{ cfg.validConfig }}</pre>
        </template>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "../store/config";
import {
    GetBestScores,
    GetRecentScores,
    GetScoreDetails,
    GetUsers,
    GetUserJudgements
} from "../../wailsjs/go/app/App";

import PageNotAvailable from "../components/config/PageNotAvailable.vue";

const cfg = useConfigStore();
const toast = useToast();
const scores = ref<any>([]);
const id = ref(0);

const c = () => {
    GetBestScores(1, 2, 0)
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};

const r = () => {
    GetRecentScores(1, 2, 0)
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};

const d = () => {
    GetScoreDetails(id.value)
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};

const u = () => {
    GetUsers()
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};

const j = () => {
    GetUserJudgements()
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};
</script>
