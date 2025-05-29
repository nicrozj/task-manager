<script setup lang="ts">
import { reactive } from "vue";

interface Props {
  isActive?: boolean;
}

const props = withDefaults(defineProps<Props>(), { isActive: false });
</script>
<template>
  <button
    :class="
      props.isActive
        ? `bg-blue-400 text-white px-4 py-2 rounded-md cursor-pointer`
        : `bg-neutral-200 text-white px-4 py-2 rounded-md cursor-pointer`
    "
    class="transition-colors"
  >
    <slot></slot>
  </button>
</template>
