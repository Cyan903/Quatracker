<template>
    <div class="tabs">
        <a
            @click="vist('/')"
            :class="visited('/')"
            class="tab tab-lg tab-lifted"
        >
            Profiles
        </a>

        <a
            @click="vist('/about')"
            :class="visited('/about')"
            class="tab tab-lg tab-lifted"
        >
            About
        </a>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();
const active = ref(location.pathname);

const vist = (n: string) => {
    active.value = n;
    router.push(n);
};

const visited = (n: string) => (active.value == n ? "tab-active" : "");
</script>
