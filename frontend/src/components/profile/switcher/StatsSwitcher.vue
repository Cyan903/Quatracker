<template>
    <button
        class="btn btn-sm btn-secondary btn-outline w-[45px]"
        :disabled="disabled"
        @click="toggle()"
    >
        <ChartBarIcon v-if="stats" class="w-4 h-4" />
        <ChartPieIcon v-else class="w-4 h-4" />
    </button>
</template>

<script lang="ts" setup>
import { ChartBarIcon, ChartPieIcon } from "@heroicons/vue/24/solid";

const props = defineProps<{ stats: boolean; disabled: boolean }>();
const emits = defineEmits<{
    (e: "setStats", stats: boolean): void;
}>();

// https://vue-chartjs.org/examples/#vue-3-charts
const toggle = () => {
    if (props.disabled) return;
    emits("setStats", !props.stats);
};
</script>
