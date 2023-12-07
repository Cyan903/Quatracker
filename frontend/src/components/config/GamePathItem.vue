<template>
    <div>
        <p class="my-4">Path to Quaver's game files...</p>
        <code class="block my-2">
            /home/hwk/.local/share/Steam/steamapps/common/Quaver/
        </code>

        <input
            type="text"
            class="input input-bordered"
            :class="{ 'input-error': invalid }"
            placeholder="Path to Quaver game files..."
            v-model="path"
        />

        <button class="btn" @click="setPath" :disabled="invalid">Update</button>
        <div class="my-4 divider"></div>
    </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "@/store/config";

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
