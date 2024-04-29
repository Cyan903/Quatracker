<template>
    <div>
        <div class="font-bold">
            <AdjustmentsHorizontalIcon class="w-5 h-5 inline mr-1" />
            <span class="text-lg align-middle">Default key mode:</span>
        </div>

        <p class="my-2">Remember the preferred key mode.</p>

        <select
            class="select select-bordered max-sm:w-full max-sm:my-2 sm:min-w-[250px]"
            v-model="mode"
            :disabled="!cfg.validConfig"
        >
            <option :value="false">4k</option>
            <option :value="true">7k</option>
        </select>

        <button
            class="btn sm:ml-1 max-sm:w-full"
            @click="setMode"
            :disabled="!cfg.validConfig"
        >
            Update
        </button>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "@/store/config";

import { AdjustmentsHorizontalIcon } from "@heroicons/vue/24/solid";

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
