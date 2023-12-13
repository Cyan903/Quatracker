<template>
    <div class="[&>*]:align-middle">
        <button
            class="btn btn-sm btn-outline"
            :class="color"
            :disabled="disable || page <= 0"
            @click="$emit('update', -1)"
        >
            <ArrowLeftIcon class="w-4 h-4" />
        </button>

        <div class="text-md font-bold inline mx-2 select-none">
            {{ useComma(page + 1) }}
        </div>

        <button
            class="btn btn-sm btn-outline"
            :class="color"
            :disabled="disable"
            @click="$emit('update', 1)"
        >
            <ArrowRightIcon class="w-4 h-4" />
        </button>
    </div>
</template>

<script lang="ts" setup>
import { useComma } from "@/use/useUtil";
import { ArrowLeftIcon, ArrowRightIcon } from "@heroicons/vue/24/solid";

defineProps<{
    disable: boolean;
    page: number;
    color: string;
}>();

defineEmits<{
    (e: "update", n: number): void;
}>();
</script>
