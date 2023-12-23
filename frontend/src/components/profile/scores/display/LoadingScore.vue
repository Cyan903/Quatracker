<template>
    <div
        v-for="_ in 10"
        :key="_"
        class="my-4 rounded-lg lg:max-w-[800px] lg:mx-auto min-h-[80px] py-2 skeleton"
    >
        <div class="sm:flex items-center justify-between p-3">
            <div class="flex">
                <div
                    class="h-10 rounded-full skeleton bg-gray-700 w-10 mr-4"
                ></div>
                <div>
                    <div
                        class="h-2.5 max-sm:w-[50%] bg-gray-300 rounded-full skeleton dark:bg-gray-600 w-24 mb-2.5"
                    ></div>
                    <div
                        class="sm:w-[150px] w-[75%] h-2 bg-gray-200 rounded-full skeleton dark:bg-gray-700 mb-2.5"
                    ></div>
                    <div
                        class="sm:w-[200px] w-full h-2 bg-gray-200 rounded-full skeleton dark:bg-gray-700"
                    ></div>
                </div>
            </div>
            <div class="text-center max-sm:hidden">
                <div
                    class="h-2.5 bg-gray-300 rounded-full dark:bg-gray-600 w-12 skeleton"
                ></div>

                <div
                    class="h-1 bg-gray-300 rounded-full dark:bg-gray-700 w-12 skeleton my-2"
                ></div>
                <div
                    class="h-1 bg-gray-300 rounded-full dark:bg-gray-700 w-12 skeleton"
                ></div>
            </div>
        </div>
    </div>
</template>
