<template>
    <div class="bg-base-100 sticky w-full z-[10] h-[80px] top-0 p-4">
        <div class="flex justify-between align-middle">
            <img
                class="w-[150px] opacity-80"
                :class="visited('/')"
                @click="vist('/')"
                :src="QuaverIcon"
            />

            <div class="rotateAnimate p-4">
                <Cog8ToothIcon
                    class="w-[25px] h-[25px] opacity-60"
                    :class="visited('/config')"
                    @click="vist('/config')"
                />
            </div>
        </div>
    </div>

    <div
        class="bg-base-content sticky w-full z-[2] h-[1px] top-[80px] opacity-20"
    ></div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";

import { Cog8ToothIcon } from "@heroicons/vue/24/solid";
import QuaverIcon from "@/assets/images/logo.svg";

const router = useRouter();
const route = useRoute();
const active = ref(location.pathname);

const vist = (n: string) => {
    if (n == "/config" && route.fullPath == n) {
        active.value = "/";
        router.push("/");
        return;
    }

    active.value = n;
    router.push(n);
};

const visited = (n: string) => (active.value == n ? "visited" : "");
</script>

<style scoped>
.visited {
    opacity: 100% !important;
}

.rotateAnimate .visited {
    transform: rotate(30deg);
}

.rotateAnimate * {
    transition: all 0.3s ease;
}
</style>
