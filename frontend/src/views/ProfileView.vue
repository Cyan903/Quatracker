<template>
    <PageNotAvailable v-if="!cfg.validConfig" />
    <template v-else>
        <UsersDisplay />

        <button @click="best" class="btn">c</button>
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
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "../store/config";
import { Scores, getBest } from "../use/useScores";
import { user } from "../use/useUsers";

import PageNotAvailable from "../components/config/PageNotAvailable.vue";
import UsersDisplay from "../components/UsersDisplay.vue";

const cfg = useConfigStore();
const toast = useToast();

const scores = ref<Scores[]>([]);
const id = ref(0);

const best = async () => {
    const data = await getBest(user.id, 1, 0, "", "", -1.0);

    if (data.error) {
        toast.error("Could not get best scores!");
        return;
    }

    scores.value = data.result || [];
};

const r = () => {
    // GetRecentScores(1, 2, 0)
    //     .then((s) => (scores.value = s))
    //     .catch((e) => {
    //         toast.error(e);
    //     });
};

const d = () => {
    // GetScoreDetails(id.value)
    //     .then((s) => (scores.value = s))
    //     .catch((e) => {
    //         toast.error(e);
    //     });
};

const u = () => {
    // GetUsers()
    //     .then((s) => (scores.value = s))
    //     .catch((e) => {
    //         toast.error(e);
    //     });
};

const j = () => {
    // GetUserJudgements()
    //     .then((s) => (scores.value = s))
    //     .catch((e) => {
    //         toast.error(e);
    //     });
};

const init = () => {
    if (!cfg.validConfig) return;

    best();
};

watch(() => user.id, init);
watch(() => cfg.validConfig, init);
onMounted(init);
</script>
