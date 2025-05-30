<script setup lang="ts">
import { ref, watch } from "vue";
import VStack from "../stacks/VStack.vue";
import HStack from "../stacks/HStack.vue";
import UButton from "../UButton.vue";
import { useTaskModal } from "@/composables/useModals";

const { open } = useTaskModal();

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
  <div class="flex flex-col sm:flex-row justify-between gap-2">
    <UButton class="w-full" @click="open">Добавить задачу</UButton>
    <select
      v-model="selectedStatus"
      class="w-full rounded-md px-2 py-1 bg-neutral-200 text-gray-700 cursor-pointer hover:bg-neutral-300 transition-colors dark:bg-slate-600 dark:text-white dark:hover:bg-slate-700"
    >
      <option value="all">Все</option>
      <option value="new">Новая</option>
      <option value="in-progress">В работе</option>
      <option value="completed">Завершена</option>
    </select>

    <input
      type="date"
      v-model="selectedDate"
      class="w-full rounded-md px-2 py-1 bg-neutral-200 text-gray-800 cursor-pointer hover:bg-neutral-300 transition-colors dark:bg-slate-600 dark:text-white dark:hover:bg-slate-700"
    />
  </div>
</template>
