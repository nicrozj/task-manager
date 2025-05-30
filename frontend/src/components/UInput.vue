<script setup lang="ts">
import { computed } from "vue";

const model = defineModel<string | number>();
const props = withDefaults(
  defineProps<{
    type?: string;
    placeholder?: string;
    hasError?: boolean;
    disabled?: boolean;
    inputClass?: string;
  }>(),
  {
    type: "text",
    hasError: false,
    disabled: false,
  }
);

const inputClasses = computed(() => [
  "w-full px-2 py-1 border rounded-md focus:outline-none transition-all duration-200",
  props.hasError
    ? "border-red-500 bg-red-50 focus:ring-2 focus:ring-red-200"
    : "border-gray-300 focus:border-blue-500 focus:ring-2 focus:ring-blue-200 dark:border-slate-900 dark:ring-slate-900 dark:bg-slate-700 dark:hover:bg-slate-800 dark:text-white",
  props.disabled ? "bg-gray-100 cursor-not-allowed opacity-75" : "bg-white",
  props.inputClass,
]);
</script>

<template>
  <input
    v-model="model"
    :type="type"
    :placeholder="placeholder"
    :class="inputClasses"
    :disabled="disabled"
    v-bind="$attrs"
  />
</template>
