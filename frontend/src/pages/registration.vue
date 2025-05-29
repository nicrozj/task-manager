<script setup lang="ts">
import UButton from "@/components/stacks/UButton.vue";
import VStack from "@/components/stacks/VStack.vue";
import UInput from "@/components/UInput.vue";
import { computed, reactive, ref } from "vue";

interface FormData {
  username: string;
  password: string;
}

interface RegistrationData extends FormData {
  repeatPassword: string;
}

const registrationData = reactive<RegistrationData>({
  username: "",
  password: "",
  repeatPassword: "",
});

const isPasswordMatch = computed(
  () =>
    (!isRepeatPasswordTouched.value ||
      registrationData.password === registrationData.repeatPassword) &&
    registrationData.repeatPassword !== ""
);
const isRepeatPasswordTouched = ref<boolean>(false);
</script>
<template>
  <VStack class="h-screen w-full justify-center items-center">
    <VStack class="gap-4 h-96">
      <h1 class="text-2xl text-center">Регистрация</h1>
      <VStack class="gap-4">
        <VStack class="gap-2">
          <UInput placeholder="Логин" v-model="registrationData.username" />
          <UInput
            placeholder="Пароль"
            type="password"
            v-model="registrationData.password"
          />
          <UInput
            placeholder="Повторите пароль"
            type="password"
            @focus="isRepeatPasswordTouched = true"
            v-model="registrationData.repeatPassword"
          />
        </VStack>
        <UButton :isActive="isPasswordMatch">Зарегистрироваться</UButton>
        <div
          v-if="!isPasswordMatch"
          class="bg-red-400 py-2 text-white px-4 rounded-md"
        >
          Пароли не совпадают!
        </div>
      </VStack>
    </VStack>
  </VStack>
</template>
