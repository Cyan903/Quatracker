<template>
    <main>
        <div id="result" class="result">{{ data.resultText }}</div>
        <div id="input" class="input-box">
            <input
                id="name"
                v-model="data.name"
                autocomplete="off"
                class="input"
                type="text"
            />

            <button class="btn" @click="greet">Greet</button>
        </div>
    </main>
</template>

<script lang="ts" setup>
import { reactive } from "vue";
import { useToast } from "vue-toastification";
import { Greet } from "../../wailsjs/go/app/App";

const toast = useToast();
const data = reactive({
    name: "",
    resultText: "Please enter your name below 👇",
});

function greet() {
    Greet(data.name).then((result) => {
        data.resultText = result;
        toast.success(`Hello, ${result}!`);
    });
}
</script>
