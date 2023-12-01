<template>
    <img v-if="source == ''" src="/src/assets/images/placeholder-bg.png" />
    <img v-else :src="source" />
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";

const props = defineProps<{ id: number }>();
const source = ref("");

onMounted(async () => {
    const req = await fetch(`/id/${props.id}.id`)
        .then((r) => r.blob())
        .catch(() => console.info(`[BackgroundMap] ${props.id} failed`));

    if (!req) return;
    source.value = URL.createObjectURL(req);
});
</script>
