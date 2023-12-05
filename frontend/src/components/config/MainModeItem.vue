<template>
    <div>
        <p class="my-4">Default key mode:</p>

        <select
            class="select select-bordered"
            v-model="mode"
            :disabled="!cfg.validConfig"
        >
            <option :value="false">4k</option>
            <option :value="true">7k</option>
        </select>

        <button class="btn" @click="setMode" :disabled="!cfg.validConfig">
            Update
        </button>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "@/store/config";

const cfg = useConfigStore();
const toast = useToast();
const mode = ref(false);

const setMode = async () => {
    const modeTest = await cfg.setMainMode(mode.value);

    if (!modeTest) {
        toast.success(`Default mode set to ${mode.value ? "7k" : "4k"}`);
    }
};

onMounted(() => {
    if (cfg.data?.MainMode) {
        mode.value = cfg.data.MainMode;
    }
});
</script>
