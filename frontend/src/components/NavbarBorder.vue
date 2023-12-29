<template>
    <div
        class="sticky w-full z-[10] h-[1px] top-[80px]"
        :class="background"
    ></div>
</template>

<script lang="ts" setup>
import { computed } from "vue";

const props = defineProps<{ bg?: string }>();
const background = computed(() => props.bg || "bg-base-content");
</script>
