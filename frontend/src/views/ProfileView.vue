<template>
    <div>
        <h2>Profile</h2>
        <PageNotAvailable v-if="!cfg.validConfig" />
        <template v-else>
            <button @click="c">c</button>
            <pre>{{ scores }}</pre>
            <pre>{{ cfg.data }}</pre>
            <pre>{{ cfg.validConfig }}</pre>
        </template>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useToast } from "vue-toastification";

import { GetScores } from "../../wailsjs/go/app/App";
import { useConfigStore } from "../store/config";
import { database } from "../../wailsjs/go/models";

import PageNotAvailable from "../components/config/PageNotAvailable.vue";

const cfg = useConfigStore();
const toast = useToast();
const scores = ref<database.ScoreBoard[]>([]);

const c = () => {
    GetScores()
        .then((s) => (scores.value = s))
        .catch((e) => {
            toast.error(e);
        });
};
</script>
