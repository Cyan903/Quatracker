<template>
    <div
        v-if="user.id != -1"
        class="dropdown w-full"
        :class="{ 'dropdown-open': open }"
    >
        <div class="btn" @click="getUsersData">
            {{ user.id }} - {{ user.username }}
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
        <p class="m-4">No users found!</p>
    </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { UsersList, user, getUsers, setUser } from "../use/useUsers";

const open = ref(false);
const list = ref<UsersList>({ Local: [], Unknown: [] });

const getUsersData = async () => {
    list.value = await getUsers();
    open.value = !open.value;
};

const setUserData = (id: string, name: string) => {
    setUser(parseInt(id), name);
    open.value = false;
};
</script>
