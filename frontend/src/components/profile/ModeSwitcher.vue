<template>
    <div class="form-control">
        <label class="label cursor-pointer">
            <span class="label-text">{{ mode ? "7k" : "4k" }}</span>
            <input
                type="checkbox"
                class="toggle toggle-primary"
                :checked="mode"
                :disabled="disabled"
                @input="
                    $emit(
                        'setMode',
                        ($event.target as HTMLInputElement).checked,
                    )
                "
            />
        </label>
    </div>
</template>

<script lang="ts" setup>
defineProps<{ mode: boolean; disabled: boolean }>();
defineEmits<{
    (e: "setMode", game: boolean): void;
}>();
</script>
