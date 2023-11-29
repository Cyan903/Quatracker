<template>
    <input
        type="range"
        class="range"
        :value="modelValue"
        @input="updateInput"
    />
</template>

<script lang="ts" setup>
let timer = 0;

defineProps<{
    modelValue: number;
}>();

const emit = defineEmits<{
    (e: "update:modelValue", evt: number): void;
}>();

const updateInput = (evt: Event) => {
    const val = parseInt((evt.target as HTMLInputElement).value);

    clearInterval(timer);

    timer = setTimeout(() => {
        emit("update:modelValue", val);
    }, 200);
};
</script>
