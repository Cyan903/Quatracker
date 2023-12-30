<template>
    <div v-if="score && score.ScoreID">
        <ModalItem :show="open" :full="true" @hide="hide">
            <template #default>
                <div
                    :style="img"
                    class="p-4 bg-img top-0 left-0 w-full min-h-[25%] overflow-hidden"
                >
                    <h2 class="text-2xl font-bold" :title="fullTitle">
                        {{ useShorten(fullTitle, 150) }}
                    </h2>

                    <h4
                        class="font-bold text-sm my-1"
                        :class="useDifficulty(score.Map.DifficultyInfo.Rating)"
                        :title="fullDiff"
                    >
                        {{ useShorten(fullDiff, 150) }}
                    </h4>

                    <span
                        class="text-sm font-bold btn btn-xs mt-2"
                        :class="useRank(score.Map.RankedStatus)"
                    >
                        {{ score.Map.RankedStatus }}
                    </span>

                    <div class="divider my-2"></div>

                    <div class="text-xs">
                        <div>
                            <span class="font-bold">Mapped by: </span>
                            <span>{{ score.Map.Creator }}</span>
                        </div>

                        <div>
                            <span class="font-bold">Played by: </span>
                            <span>
                                ({{ score.Score.LocalProfileId }})
                                {{ score.Score.Name }}
                            </span>
                        </div>
                    </div>
                </div>

                <div
                    class="p-4 bg-base-300 rounded-b-xl font-bold flex space-between justify-around"
                >
                    <div class="text-center w-[24%]">
                        <h2 class="text-sm">
                            {{ score.Score.PerformanceRating.toFixed(2) }}
                        </h2>

                        <div class="btn btn-xs">Rating</div>
                    </div>

                    <div class="text-center w-[24%]">
                        <h2 class="text-sm">
                            {{ useComma(score.Score.TotalScore) }}
                        </h2>

                        <div class="btn btn-xs">
                            <span class="max-sm:hidden">Total</span>
                            Score
                        </div>
                    </div>

                    <div class="text-center w-[24%]">
                        <h2 class="text-sm">
                            {{ useComma(score.Score.MaxCombo) }}x
                        </h2>

                        <div class="btn btn-xs">
                            <span class="max-sm:hidden">Max</span>
                            Combo
                        </div>
                    </div>
                </div>

                <div class="divider p-4 font-bold">Judgements</div>
                <div class="p-4">
                    <div
                        class="font-bold text-center mb-[3em]"
                        :class="{ italic: !isStandard && !showJudges }"
                    >
                        <span
                            class="text-5xl align-middle"
                            :class="`grade-${
                                showJudges
                                    ? score.Score.JudgedHits.RankedGrade
                                    : score.Score.JudgedHits.Grade
                            }`"
                        >
                            {{
                                getGrade(
                                    showJudges
                                        ? score.Score.JudgedHits.RankedGrade
                                        : score.Score.JudgedHits.Grade,
                                )
                            }}
                        </span>

                        <span class="mx-2 align-middle">&#x2022;</span>
                        <span class="align-middle text-xl">
                            {{
                                (showJudges
                                    ? score.Score.JudgedHits.RankedAccuracy || 0
                                    : score.Score.JudgedHits.Accuracy || 0
                                ).toFixed(2)
                            }}%
                        </span>

                        <span v-if="!isStandard && !showJudges">*</span>
                    </div>

                    <div
                        class="grid sm:grid-cols-2 grid-rows-3 gap-4 md:w-[75%] m-auto"
                    >
                        <div v-for="judge in judges" :key="judge">
                            <span
                                class="btn btn-outline btn-sm align-middle mt-[1px] min-w-[80px] max-sm:w-[60px]"
                                :class="`grade-${
                                    grades[judges.indexOf(judge)]
                                }`"
                            >
                                {{ judge }}
                            </span>

                            <div
                                class="inline-block text-left md:min-w-[150px]"
                            >
                                <span
                                    class="ml-2 align-middle text-xl"
                                    :class="`grade-${
                                        grades[judges.indexOf(judge)]
                                    }`"
                                >
                                    {{
                                        showJudges
                                            ? `${getPreset(judge)}ms`
                                            : `${useComma(getJudge(judge))}x`
                                    }}
                                </span>

                                <span
                                    v-if="!showJudges"
                                    class="text-[0.6em] align-bottom"
                                >
                                    ({{ getPercent(judge) }})
                                </span>
                            </div>
                        </div>
                    </div>

                    <div class="text-center mt-8">
                        <div class="my-4">
                            <span class="font-bold">Judge: </span>
                            <span>{{ score.Score.JudgementWindowPreset }}</span>
                        </div>

                        <button
                            class="btn btn-primary btn-outline btn-sm md:w-[75%] w-full"
                            @click="showJudges = !showJudges"
                        >
                            Display
                            {{
                                showJudges
                                    ? "Score Information"
                                    : "Judgement Preset"
                            }}
                        </button>
                    </div>
                </div>

                <div class="divider p-4 font-bold">Score Information</div>
                <div class="p-4 overflow-x-auto md:w-[75%] m-auto">
                    <table class="table">
                        <tbody>
                            <tr>
                                <th>Score ID</th>
                                <td>{{ score.ScoreID }}</td>
                            </tr>
                            <tr>
                                <th>Personal Best?</th>
                                <td>
                                    {{
                                        score.Score.PersonalBest ? "Yes" : "No"
                                    }}
                                </td>
                            </tr>
                            <tr>
                                <th>Date</th>
                                <td>{{ fullDate }}</td>
                            </tr>
                            <tr>
                                <th>Mods</th>
                                <td>{{ mods.join(", ") }}</td>
                            </tr>
                            <tr>
                                <th>Pause Count</th>
                                <td>
                                    {{ score.Score.PauseCount }}
                                </td>
                            </tr>
                            <tr>
                                <th>Scroll Speed</th>
                                <td>
                                    {{ score.Score.ScrollSpeed }}
                                </td>
                            </tr>
                            <tr>
                                <th>Rating Processor Version</th>
                                <td>
                                    {{ score.Score.RatingVersion }}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="divider p-4 font-bold">Map Information</div>
                <div class="p-4 overflow-x-auto md:w-[75%] m-auto">
                    <table class="table">
                        <tbody>
                            <tr>
                                <th>Map ID</th>
                                <td>{{ score.MapID }}</td>
                            </tr>
                            <tr>
                                <th>Song Length</th>
                                <td>{{ songLength }}</td>
                            </tr>
                            <tr>
                                <th>BPM</th>
                                <td>
                                    {{
                                        score.Map.DifficultyInfo.BPM.toFixed(2)
                                    }}
                                </td>
                            </tr>
                            <tr>
                                <th>LN Percent</th>
                                <td>
                                    {{
                                        score.Map.DifficultyInfo.LNPercent.toFixed(
                                            2,
                                        )
                                    }}%
                                </td>
                            </tr>
                            <tr>
                                <th>Mode</th>
                                <td>
                                    {{ score.Map.ModeInfo.Mode ? "7k" : "4k" }}
                                </td>
                            </tr>
                            <tr>
                                <th>Has Scratch Key?</th>
                                <td>
                                    {{
                                        score.Map.ModeInfo.HasScratchKey
                                            ? "Yes"
                                            : "No"
                                    }}
                                </td>
                            </tr>
                            <tr>
                                <th>Difficulty Processor Version</th>
                                <td>
                                    {{
                                        score.Map.ModeInfo
                                            .DifficultyProcessorVersion
                                    }}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </template>
        </ModalItem>
    </div>
</template>

<script lang="ts" setup>
import type { ScoreDetails, Judges } from "@/types/scores";
import { computed, ref, watch } from "vue";
import { useToast } from "vue-toastification";
import { useDifficulty, useRank } from "@/use/useColors";
import { useShorten, useComma } from "@/use/useUtil";
import { useMods } from "@/use/useMods";
import { useImage } from "@/use/useImage";
import { getDetails } from "@/use/useScores";

import ModalItem from "@/components/util/ModalItem.vue";
import moment from "moment";

const props = defineProps<{
    id: number;
    open: boolean;
}>();

const emits = defineEmits<{
    (e: "hide"): void;
}>();

const toast = useToast();
const score = ref<ScoreDetails>();
const loaded = ref(false);
const showJudges = ref(false);
const img = ref("");

const judges = ["Marv", "Perf", "Great", "Good", "Okay", "Miss"];
const grades = ["SS", "S", "A", "B", "C", "F"];
const standardJudgementConfig = {
    CountMarv: 18,
    CountPerf: 43,
    CountGreat: 76,
    CountGood: 106,
    CountOkay: 127,
    CountMiss: 164,
};

const init = async () => {
    const data = await getDetails(props.id);

    if (data.error || !data.result || !data.result?.ScoreID) {
        toast.error("Could not get score details!");
        emits("hide");

        return;
    }

    score.value = data.result;
    img.value = `
        background: linear-gradient(rgba(0, 0, 0, 0.8), rgba(0, 0, 0, 0.9)), url("${await useImage(
            score.value.MapID,
        )}") no-repeat center;
    `;
};

const getJudge = (judge: string) => {
    const judgeItem =
        score.value?.Score.JudgedHits[`Count${judge}` as keyof Judges];

    if (judgeItem === undefined) {
        return "";
    }

    return judgeItem;
};

const getPreset = (preset: string) => {
    const judgeItem =
        score.value?.Score.JudgementConfig[`Count${preset}` as keyof Judges];

    if (judgeItem === undefined) {
        return "";
    }

    return judgeItem;
};

const getPercent = (judge: string) => {
    if (!score.value) return "0.00%";

    const judgeCount = getJudge(judge);
    let sum = 0;

    if (judgeCount === "") {
        console.warn("[DetailPopup] Could not get judge", judge);
        return "0.00%";
    }

    for (const jitem of judges) {
        const jj = getJudge(jitem);

        if (jj === "") {
            console.warn("[DetailPopup] Could not get judge", jitem);
            continue;
        }

        sum += jj;
    }

    return `${((judgeCount / sum) * 100).toFixed(2)}%`;
};

const getGrade = (grade: string) => {
    if (!grade) return "F";
    return grade == "SS" ? "S+" : grade;
};

const fullTitle = computed(
    () => `${score.value?.Map?.Artist} - ${score.value?.Map?.Title}`,
);

const fullDiff = computed(
    () =>
        ` [${score.value?.Map
            ?.DifficultyName}] - (${score.value?.Map?.DifficultyInfo?.Rating.toFixed(
            2,
        )})`,
);

const fullDate = computed(() => {
    const d = new Date(score.value?.Score.DateTime || 0);
    return moment(d).format("MMMM Do YYYY, h:mma");
});

const songLength = computed(() =>
    moment
        .utc(score.value?.Map?.DifficultyInfo?.SongLength || 0)
        .format("mm:ss"),
);

const mods = computed(() => {
    if (!score.value?.Score.Mods || score.value.Score.Mods === "0") {
        return ["None"];
    }

    return useMods(score.value.Score.Mods, true);
});

const isStandard = computed(() => {
    if (
        !score.value?.Score.JudgedHits ||
        !score.value?.Score.JudgementConfig ||
        !score.value?.Score.JudgementWindowPreset
    ) {
        return false;
    }

    const acc =
        score.value.Score.JudgementWindowPreset === "Standard*" &&
        score.value.Score.JudgedHits.Accuracy ===
            score.value.Score.JudgedHits.RankedAccuracy;

    const hits =
        score.value.Score.JudgementConfig.CountMarv ===
            standardJudgementConfig.CountMarv &&
        score.value.Score.JudgementConfig.CountPerf ===
            standardJudgementConfig.CountPerf &&
        score.value.Score.JudgementConfig.CountGreat ===
            standardJudgementConfig.CountGreat &&
        score.value.Score.JudgementConfig.CountGood ===
            standardJudgementConfig.CountGood &&
        score.value.Score.JudgementConfig.CountOkay ===
            standardJudgementConfig.CountOkay &&
        score.value.Score.JudgementConfig.CountMiss ===
            standardJudgementConfig.CountMiss;

    return acc && hits;
});

const hide = () => {
    emits("hide");

    loaded.value = false;
    showJudges.value = false;
    img.value = "";
};

watch(
    () => props.open,
    async () => {
        loaded.value = false;

        if (props.open) {
            await init();
        }

        loaded.value = true;
    },
);
</script>
