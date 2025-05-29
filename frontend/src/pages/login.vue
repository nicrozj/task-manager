<script setup lang="ts">
import UButton from "@/components/stacks/UButton.vue";
import VStack from "@/components/stacks/VStack.vue";
import UInput from "@/components/UInput.vue";
</script>

<template>
  <VStack class="h-screen w-full justify-center items-center">
    <VStack class="gap-4">
      <h1 class="text-2xl text-center">Вход</h1>
      <VStack class="gap-2">
        <UInput placeholder="Логин" />
        <UInput placeholder="Пароль" type="password" />
        <UButton>Войти</UButton>
      </VStack>
    </VStack>
  </VStack>
</template>
