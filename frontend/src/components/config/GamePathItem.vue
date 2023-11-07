<template>
    <div>
        <p class="my-4">Path to Quaver's game files...</p>
        <code>/home/hwk/.local/share/Steam/steamapps/common/Quaver/</code>

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
import { computed, onMounted, ref, watch } from "vue";
import { useToast } from "vue-toastification";
import { useConfigStore } from "../../store/config";

const cfg = useConfigStore();
const toast = useToast();
const path = ref("");

const invalid = computed(() => {
    // TODO: Better path checks
    if (path.value.length === 0) {
        return true;
    }

    return false;
});

const setPath = () => {
    try {
        if (!cfg.setGamePath(path.value)) {
            toast.error("Could not set game path!");
            return;
        }

        toast.success(`Quaver's path has been set to ${path.value}`);
    } catch (e) {
        toast.error(`Could not set game path - ${e}`);
    }
};

onMounted(() => {
    if (cfg.data?.GamePath) {
        path.value = cfg.data.GamePath;
    }
});
</script>
