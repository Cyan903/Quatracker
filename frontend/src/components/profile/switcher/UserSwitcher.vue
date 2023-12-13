<template>
    <div
        v-if="props.id != -1"
        class="dropdown md:w-[calc(100%-105px)] max-w-[700px]"
        :class="{ 'dropdown-open': open }"
    >
        <div class="btn btn-sm" @click="usersFetch(true)">
            <UserIcon class="w-3 h-3" />
            {{ props.username }} #{{ props.id }}
        </div>

        <ul
            class="dropdown-content z-[5] menu p-2 shadow bg-base-200 rounded-box w-full max-md:left-[-100px] max-md:absolute max-md:w-[calc(100vw-2rem)]"
        >
            <li v-for="user in list.Local" :key="user.Id">
                <a @click="setUserData(user.Id, user.Username)">
                    {{ user.Username }}
                </a>
            </li>

            <div v-if="list.Unknown?.length" class="divider">Unknown</div>

            <li v-for="user in list.Unknown" :key="user.Id">
                <a @click="setUserData(user.Id, user.Username)">
                    #{{ user.Id }}
                </a>
            </li>
        </ul>
    </div>
    <button v-else class="btn btn btn-sm" disabled>No users found!</button>
</template>

<script lang="ts" setup>
import type { UsersList } from "@/use/useUsers";
import { onMounted, ref } from "vue";
import { useToast } from "vue-toastification";
import { getUsers } from "@/use/useUsers";

import { UserIcon } from "@heroicons/vue/24/solid";

const emits = defineEmits<{
    (e: "setUser", id: number, name: string): void;
}>();

const props = defineProps<{
    id: number;
    username: string;
}>();

const toast = useToast();
const open = ref(false);
const list = ref<UsersList>({ Local: [], Unknown: [] });

const usersFetch = async (toggle: boolean) => {
    const data = await getUsers();
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
    await usersFetch(false);

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
