<script setup lang="ts">
import { getProfile, logout } from "@/api/login";
import { onMounted, ref } from "vue";
import { useRouter } from "vue-router";
import HStack from "./stacks/HStack.vue";
import UButton from "./UButton.vue";

const router = useRouter();

const username = ref("");

onMounted(() => {
  getProfile(router).then((response) => {
    username.value = response.data.username;
  });
});

const onLogout = () => {
  localStorage.setItem("access_token", "");
  logout(router);
};
</script>
<template>
  <header class="w-[1000px] mx-auto mt-2">
    <HStack class="justify-between">
      <span>task manager</span>
      <HStack class="items-center gap-2">
        <span>{{ username }}</span>
        <UButton @click="onLogout">Выйти</UButton>
      </HStack>
    </HStack>
  </header>
</template>
