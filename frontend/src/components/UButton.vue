<script setup lang="ts">
import { computed } from "vue";

interface Props {
  isActive?: boolean;
  isDisabled?: boolean;
  type?: "button" | "submit" | "reset";
}

const props = withDefaults(defineProps<Props>(), {
  isActive: false,
  isDisabled: false,
  type: "button",
});

const buttonClasses = computed(() => [
  "px-2 py-1 rounded-md transition-colors",
  props.isActive
    ? "bg-blue-400 text-white hover:bg-blue-500 dark:bg-slate-600 dark:text-white dark:hover:bg-slate-700"
    : "bg-neutral-200 text-gray-700 hover:bg-neutral-300 dark:bg-slate-600 dark:hover:bg-slate-700 dark:text-white",
  props.isDisabled ? "opacity-50 cursor-not-allowed" : "cursor-pointer",
]);
</script>

<template>
  <button
    :type="type"
    :class="buttonClasses"
    :disabled="isDisabled"
    class="inline-flex items-center justify-center"
  >
    <slot />
  </button>
</template>
