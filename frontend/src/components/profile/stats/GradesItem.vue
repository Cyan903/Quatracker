<template>
    <div>
        <div class="divider my-8">
            <AcademicCapIcon class="w-[2rem] h-[2rem]" />
            <span class="font-bold">Grades Count</span>
        </div>

        <div class="m-auto sm:w-[75%] lg:max-w-[800px] stat-pop my-4">
            <div class="text-center">
                <template v-for="(_, grade) in placeholderGrades" :key="grade">
                    <div class="text-lg inline-block m-2 join">
                        <button
                            class="join-item btn btn-sm min-w-[3rem]"
                            :class="`grade-${grade}`"
                        >
                            {{ grade == "SS" ? "S+" : grade }}
                        </button>

                        <button
                            class="join-item btn btn-sm min-w-[3rem]"
                            :class="`grade-${grade}`"
                        >
                            {{ useComma(stats[gradeMode][grade] || 0) }}
                        </button>
                    </div>

                    <div v-if="grade == 'S'" class="w-full max-sm:hidden"></div>
                </template>
            </div>

            <div class="my-8 w-[75%] m-auto max-sm:w-full">
                <button
                    class="btn btn-outline btn-sm btn-primary w-full mb-2"
                    @click="toggleModes"
                >
                    {{ gradeMode == "AllGrades" ? "All" : "Ranked" }}
                </button>

                <button
                    class="btn btn-outline btn-sm btn-secondary w-full"
                    @click="pb = !pb"
                >
                    PB: {{ pb ? "Yes" : "No" }}
                </button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import type { GradesCount } from "@/types/stats";
import { AcademicCapIcon } from "@heroicons/vue/24/solid";
import { onMounted, watch, ref } from "vue";
import { useToast } from "vue-toastification";

import { useConfigStore } from "@/store/config";
import { useComma } from "@/use/useUtil";
import { getGrades } from "@/use/useStats";

const cfg = useConfigStore();
const toast = useToast();
const loading = ref(true);
const gradeMode = ref<"AllGrades" | "RankedGrades">("AllGrades");
const pb = ref(false);

const props = defineProps<{
    id: number;
    mode: number;
}>();

const placeholderGrades = {
    X: 0,
    SS: 0,
    S: 0,
    A: 0,
    B: 0,
    C: 0,
    D: 0,
    F: 0,
};

const stats = ref<GradesCount>({
    AllGrades: structuredClone(placeholderGrades),
    RankedGrades: structuredClone(placeholderGrades),
});

const grades = async () => {
    const data = await getGrades(props.id, props.mode, pb.value);

    if (data.error) {
        toast.error("Could not get grades!");
        return;
    }

    if (!data.result) {
        toast.error("No grades found!");
        return;
    }

    stats.value = data.result;
};

const toggleModes = () => {
    gradeMode.value =
        gradeMode.value == "AllGrades" ? "RankedGrades" : "AllGrades";
};

const init = async () => {
    if (!cfg.validConfig) return;

    loading.value = true;
    pb.value = false;
    gradeMode.value = "AllGrades";
    stats.value = {
        AllGrades: structuredClone(placeholderGrades),
        RankedGrades: structuredClone(placeholderGrades),
    };

    await grades();
    loading.value = false;
};

watch(pb, async () => {
    if (!cfg.validConfig) return;

    loading.value = true;
    stats.value = {
        AllGrades: structuredClone(placeholderGrades),
        RankedGrades: structuredClone(placeholderGrades),
    };

    await grades();
    loading.value = false;
});

watch(() => props.id, init);
watch(() => props.mode, init);

onMounted(init);
</script>
