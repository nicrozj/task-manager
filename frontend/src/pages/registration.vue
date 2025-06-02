<script setup lang="ts">
import UButton from "@/components/UButton.vue";
import VStack from "@/components/stacks/VStack.vue";
import UInput from "@/components/UInput.vue";
import { computed, reactive, ref } from "vue";
import { register } from "@/api/login";
import { useRouter } from "vue-router";

const router = useRouter();

interface RegistrationData {
  username: string;
  password: string;
  repeatPassword: string;
}

const registrationData = reactive<RegistrationData>({
  username: "",
  password: "",
  repeatPassword: "",
});

const isLoginTouched = ref(false);
const isPasswordTouched = ref(false);
const isRepeatPasswordTouched = ref(false);

const isLoading = ref(false);
const serverError = ref<string | null>(null);
const isSubmitted = ref(false);

const isPasswordMatch = computed(
  () => registrationData.password === registrationData.repeatPassword
);

const usernameInputError = computed(() => {
  if (registrationData.username.length < 3) {
    return "Логин должен содержать минимум 3 символа";
  } else if (registrationData.username.length > 16) {
    return "Логин не должен быть длиннее 16 символов";
  }
});
const passwordInputError = computed(() => {
  if (registrationData.password.length < 6) {
    return "Пароль должен содержать минимум 6 символов";
  } else if (registrationData.password.length > 16) {
    return "Пароль не должен быть длиннее 16 символов";
  }
});

const shouldShowPasswordError = computed(
  () =>
    isRepeatPasswordTouched.value &&
    !isPasswordMatch.value &&
    registrationData.repeatPassword !== ""
);

const shouldShowValidationErrors = computed(
  () => isSubmitted.value || isRepeatPasswordTouched.value
);

const isFormValid = computed(
  () =>
    !usernameInputError.value &&
    !passwordInputError.value &&
    isPasswordMatch.value
);

const submitForm = async () => {
  isSubmitted.value = true;
  serverError.value = null;

  if (!isFormValid.value || isLoading.value) return;

  isLoading.value = true;

  try {
    await register({
      username: registrationData.username,
      password: registrationData.password,
    });
    router.push("/login");
  } catch (error: any) {
    serverError.value =
      error.message || "Ошибка регистрации. Попробуйте снова.";
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <VStack class="h-screen w-full justify-center items-center">
    <VStack
      class="gap-2 p-8 w-full max-w-md rounded-md dark:bg-slate-800 dark:text-white"
    >
      <VStack class="gap-1 text-center">
        <h1 class="text-3xl text-gray-800 dark:text-white">Регистрация</h1>
        <p class="text-gray-500">Создайте новый аккаунт</p>
      </VStack>
      <VStack class="gap-2">
        <VStack class="gap-2">
          <VStack class="gap-1">
            <UInput
              placeholder="Логин"
              v-model="registrationData.username"
              @blur="isLoginTouched = true"
              :hasError="isLoginTouched && Boolean(usernameInputError)"
            />
            <div
              v-if="isLoginTouched && Boolean(usernameInputError)"
              class="text-red-500 text-sm pl-2"
            >
              Логин должен содержать минимум 3 символа
            </div>
          </VStack>

          <VStack class="gap-1">
            <UInput
              placeholder="Пароль"
              type="password"
              v-model="registrationData.password"
              @blur="isPasswordTouched = true"
              :hasError="isPasswordTouched && Boolean(passwordInputError)"
            />
            <div
              v-if="isPasswordTouched && Boolean(passwordInputError)"
              class="text-red-500 text-sm pl-2"
            >
              Пароль должен содержать минимум 6 символов
            </div>
          </VStack>

          <VStack class="gap-1">
            <UInput
              placeholder="Повторите пароль"
              type="password"
              @blur="isRepeatPasswordTouched = true"
              v-model="registrationData.repeatPassword"
              :hasError="shouldShowPasswordError"
            />
            <div
              v-if="!isPasswordMatch && isRepeatPasswordTouched"
              class="text-red-500 text-sm pl-2"
            >
              Пароли не совпадают!
            </div>
          </VStack>
        </VStack>

        <div
          v-if="serverError"
          class="bg-red-50 border border-red-200 text-red-600 px-4 py-3 rounded-lg dark:bg-red-900 dark:text-white dark:border-red-600"
        >
          <div class="font-medium">Ошибка регистрации</div>
          <div class="text-sm mt-1">{{ serverError }}</div>
        </div>

        <UButton
          :is-active="isFormValid"
          :is-disabled="!isFormValid"
          :loading="isLoading"
          @click="submitForm"
          class="mt-2"
        >
          Зарегистрироваться
        </UButton>
      </VStack>

      <div class="text-center text-sm text-gray-600 mt-4 dark:text-white">
        Уже есть аккаунт?
        <router-link to="/login" class="text-blue-500 hover:underline">
          Войти
        </router-link>
      </div>
    </VStack>
  </VStack>
</template>
