<script setup lang="ts">
import { ref, watch } from "vue";
import VStack from "../stacks/VStack.vue";
import HStack from "../stacks/HStack.vue";

type status = "all" | "new" | "in-progress" | "completed";

const selectedStatus = ref<status>("all");
const selectedDate = ref<string>("");

const emit = defineEmits<{
  (e: "update:filters", filters: { status: status; date: string }): void;
}>();

function onFilterChange() {
  emit("update:filters", {
    status: selectedStatus.value,
    date: selectedDate.value,
  });
}

watch(selectedStatus, onFilterChange);
watch(selectedDate, onFilterChange);
</script>
<template>
  <HStack class="gap-2">
    <select
      v-model="selectedStatus"
      class="rounded-md px-3 py-2 bg-blue-400 text-white cursor-pointer"
    >
      <option value="all">Все</option>
      <option value="new">Новая</option>
      <option value="in-progress">В работе</option>
      <option value="completed">Завершена</option>
    </select>

    <input
      type="date"
      v-model="selectedDate"
      class="rounded-md px-3 py-2 bg-neutral-200 text-gray-800 cursor-pointer"
    />
  </HStack>
</template>
