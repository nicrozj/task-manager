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
  "px-4 py-2 rounded-md transition-colors focus:outline-none focus:ring-2 focus:ring-blue-300",
  props.isActive
    ? "bg-blue-500 text-white hover:bg-blue-600"
    : "bg-neutral-200 text-gray-700 hover:bg-neutral-300",
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
