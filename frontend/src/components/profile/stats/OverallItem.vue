<template>
    <h2>Stats</h2>
    <pre>{{ data }}</pre>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";

const props = defineProps<{
    id: number;
    mode: boolean;
}>();

const data = ref([]);
const update = () => {
    data.value = [];
};

onMounted(update);
watch(() => props.id, update);
watch(() => props.mode, update);
</script>
