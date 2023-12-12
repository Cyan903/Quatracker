<template>
    <div class="modes" @click="toggle">
        <template v-if="!disabled">
            <button class="mode-1" :class="{ active: !mode }">4k</button>
            <button class="mode-2" :class="{ active: mode }">7k</button>
        </template>
        <div v-else class="btn btn-sm w-[90%]">Disabled</div>
    </div>
</template>

<script lang="ts" setup>
const props = defineProps<{ mode: boolean; disabled: boolean }>();
const emits = defineEmits<{
    (e: "setMode", game: boolean): void;
}>();

const toggle = () => {
    if (props.disabled) return;
    emits("setMode", !props.mode);
};
</script>
