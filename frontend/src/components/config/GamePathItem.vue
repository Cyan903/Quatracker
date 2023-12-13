<template>
    <div class="leading-[2rem] break-all">
        <div class="font-bold">
            <FolderIcon class="w-5 h-5 inline mr-1" />
            <span class="text-lg align-middle">
                Path to Quaver's game files:
            </span>
        </div>

        <div class="my-2">
            <span class="max-sm:hidden">Example:</span>
            <span class="font-mono text-sm">
                /home/user/.local/share/Steam/steamapps/common/Quaver/
            </span>
        </div>

        <input
            type="text"
            class="input input-bordered max-sm:w-full max-sm:my-2 sm:min-w-[250px]"
            :class="{ 'input-error': invalid }"
            placeholder="Path to Quaver game files..."
            v-model="path"
        />

        <button
            class="btn sm:ml-1 max-sm:w-full"
            @click="setPath"
            :disabled="invalid"
        >
            Update
        </button>
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "@/store/config";

import { FolderIcon } from "@heroicons/vue/24/solid";

const cfg = useConfigStore();
const toast = useToast();
const path = ref("");

const invalid = computed(() => path.value.length === 0);

const setPath = async () => {
    const pathTest = await cfg.setGamePath(path.value);

    if (!pathTest) {
        toast.success(`Quaver's path has been set to ${path.value}`);
    }
};

onMounted(() => {
    if (cfg.data?.GamePath) {
        path.value = cfg.data.GamePath;
    }
});
</script>
