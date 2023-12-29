<template>
    <div
        v-for="_ in n || 10"
        :key="_"
        class="my-4 rounded-lg lg:max-w-[800px] lg:mx-auto min-h-[80px] py-2 bg-base-200 skeleton"
    >
        <div class="sm:flex items-center justify-between p-3">
            <div class="flex">
                <div
                    class="h-10 rounded-full skeleton bg-gray-700 w-10 mr-4"
                ></div>
                <div>
                    <div
                        class="h-2.5 max-sm:w-[50%] rounded-full skeleton bg-gray-600 w-24 mb-2.5"
                    ></div>
                    <div
                        class="sm:w-[150px] w-[75%] h-2 rounded-full skeleton bg-gray-700 mb-2.5"
                    ></div>
                    <div
                        class="sm:w-[200px] w-full h-2 rounded-full skeleton bg-gray-700"
                    ></div>
                </div>
            </div>
            <div class="text-center max-sm:hidden">
                <div class="h-2.5 rounded-full bg-gray-600 w-12 skeleton"></div>
                <div
                    class="h-1 rounded-full bg-gray-700 w-12 skeleton my-2"
                ></div>
                <div class="h-1 rounded-full bg-gray-700 w-12 skeleton"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
defineProps<{ n?: number }>();
</script>
