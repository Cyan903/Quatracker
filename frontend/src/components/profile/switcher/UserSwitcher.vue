<template>
    <div
        v-if="props.id != -1"
        class="dropdown w-full"
        :class="{ 'dropdown-open': open }"
    >
        <div class="btn" @click="getUsers(true)">
            {{ props.id }} - {{ props.username }}
        </div>

        <ul
            class="dropdown-content z-[1] menu p-2 shadow bg-base-200 rounded-box w-full"
        >
            <li v-for="user in list.Local" :key="user.Id">
                <a @click="setUserData(user.Id, user.Username)">
                    {{ user.Id }} - {{ user.Username }}
                </a>
            </li>

            <div v-if="list.Unknown?.length" class="divider"></div>

            <li v-for="user in list.Unknown" :key="user.Id">
                <a @click="setUserData(user.Id, user.Username)">
                    {{ user.Id }} - {{ user.Username }}
                </a>
            </li>
        </ul>
    </div>
    <div v-else class="w-full">
        <button class="btn" disabled>No users found!</button>
    </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { useToast } from "vue-toastification";

import { useWails } from "../../../use/useWails";
import { GetUsers } from "../../../../wailsjs/go/app/App";

const emits = defineEmits<{
    (e: "setUser", id: number, name: string): void;
}>();

const props = defineProps<{
    id: number;
    username: string;
}>();

interface UsersList {
    Local: {
        Id: string;
        Username: string;
    }[];

    Unknown: {
        Id: string;
        Username: string;
    }[];
}

const toast = useToast();
const open = ref(false);
const list = ref<UsersList>({ Local: [], Unknown: [] });

const getUsers = async (toggle: boolean) => {
    const data = await useWails<UsersList>(GetUsers);
    const none = { Local: [], Unknown: [] };

    if (toggle) {
        open.value = !open.value;
    }

    if (data.error) {
        toast.error("Could not get users!");
        list.value = none;

        return;
    }

    list.value = data.result || none;
};

const initUsers = async () => {
    emits("setUser", -1, "");
    await getUsers(false);

    if (!list.value.Local?.length && !list.value.Unknown?.length) {
        toast.error("No users found!");

        return;
    }

    if (list.value.Local?.length !== 0) {
        emits(
            "setUser",

            parseInt(list.value.Local[0].Id),
            list.value.Local[0].Username,
        );

        return;
    }

    emits(
        "setUser",

        parseInt(list.value.Unknown[0].Id),
        list.value.Unknown[0].Username,
    );
};

const setUserData = (id: string, name: string) => {
    emits("setUser", parseInt(id), name);
    open.value = false;
};

onMounted(initUsers);
</script>
