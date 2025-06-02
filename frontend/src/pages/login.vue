<script setup lang="ts">
import UButton from "@/components/UButton.vue";
import VStack from "@/components/stacks/VStack.vue";
import UInput from "@/components/UInput.vue";
import { computed, reactive, ref } from "vue";
import { login } from "@/api/login";
import { router } from "@/main";

interface LoginData {
  username: string;
  password: string;
}

const loginData = reactive<LoginData>({
  username: "",
  password: "",
});

const isLoading = ref(false);
const serverError = ref<string | null>(null);
const isSubmitted = ref(false);

const usernameInputError = computed(() => {
  if (loginData.username.length < 3) {
    return "Логин должен содержать минимум 3 символа";
  } else if (loginData.username.length > 16) {
    return "Логин не должен быть длиннее 16 символов";
  }
});
const passwordInputError = computed(() => {
  if (loginData.password.length < 6) {
    return "Пароль должен содержать минимум 6 символов";
  } else if (loginData.password.length > 16) {
    return "Пароль не должен быть длиннее 16 символов";
  }
});

const isFormValid = computed(
  () => !usernameInputError.value && !passwordInputError.value
);

const submitForm = async () => {
  isSubmitted.value = true;
  serverError.value = null;

  if (!isFormValid.value || isLoading.value) return;

  isLoading.value = true;

  try {
    await login({
      username: loginData.username,
      password: loginData.password,
    });
    router.push("/");
  } catch (error: any) {
    serverError.value =
      error.message || "Неверный логин или пароль. Попробуйте снова.";
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <VStack class="h-screen w-full justify-center items-center">
    <VStack
      class="gap-2 p-8 bg-white w-full max-w-md rounded-md dark:bg-slate-800 dark:text-white"
    >
      <VStack class="gap-1 text-center">
        <h1 class="text-3xl text-gray-800 dark:text-white">Вход</h1>
        <p class="text-gray-500 dark:text-slate-400">Войдите в свой аккаунт</p>
      </VStack>
      <VStack class="gap-2">
        <VStack class="gap-2">
          <VStack class="gap-1">
            <UInput
              placeholder="Логин"
              v-model="loginData.username"
              :hasError="isSubmitted && Boolean(usernameInputError)"
            />
            <div
              v-if="isSubmitted && usernameInputError"
              class="text-red-500 text-sm pl-2"
            >
              Логин должен содержать минимум 3 символа
            </div>
          </VStack>

          <VStack class="gap-1">
            <UInput
              placeholder="Пароль"
              type="password"
              v-model="loginData.password"
              :hasError="isSubmitted && Boolean(passwordInputError)"
            />
            <div
              v-if="isSubmitted && passwordInputError"
              class="text-red-500 text-sm pl-2"
            >
              Пароль должен содержать минимум 6 символов
            </div>
          </VStack>
        </VStack>

        <div
          v-if="serverError"
          class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded-lg dark:bg-red-900 dark:text-white dark:border-red-600"
        >
          <div class="font-medium">Ошибка входа</div>
          <div class="text-sm mt-1">{{ serverError }}</div>
        </div>

        <UButton
          :isActive="isFormValid"
          :is-disabled="!isFormValid"
          :loading="isLoading"
          @click="submitForm"
          class="mt-2"
        >
          Войти
        </UButton>
      </VStack>

      <div class="text-center text-sm text-gray-600 dark:text-white mt-4">
        Нет аккаунта?
        <router-link to="/registration" class="text-blue-500 hover:underline">
          Зарегистрироваться
        </router-link>
      </div>
    </VStack>
  </VStack>
</template>
