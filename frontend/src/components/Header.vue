<script setup lang="ts">
import { getProfile, logout } from "@/api/login";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import HStack from "./stacks/HStack.vue";
import UButton from "./UButton.vue";
import { useTasks } from "@/composables/useTasks";

const { tasks } = useTasks();
const router = useRouter();

const username = ref("");

onMounted(() => {
  getProfile().then((response) => {
    username.value = response.data.username;
  });
});

const onLogout = () => {
  localStorage.setItem("access_token", "");
  logout();
  router.push("/login");
  tasks.value = [];
};
</script>
<template>
  <header class="w-[1000px] mx-auto mt-2">
    <HStack class="justify-between">
      <router-link to="/" class="text-2xl font-medium">
        task manager
      </router-link>
      <HStack class="items-center gap-2">
        <span>{{ username }}</span>
        <UButton @click="onLogout">Выйти</UButton>
      </HStack>
    </HStack>
  </header>
</template>
