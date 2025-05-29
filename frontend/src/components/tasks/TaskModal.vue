<script setup lang="ts">
import { createTask, deleteTask, updateTask } from "@/api/tasks";
import { useTaskModal } from "@/composables/useModals";
import { useTasks, type Task } from "@/composables/useTasks";
import { ref, watch } from "vue";
import HStack from "../stacks/HStack.vue";
const { updateTaskInList, deleteTaskInList } = useTasks();

const localTask = ref<Task>({
  title: "",
  description: "",
  status: "new",
});

const isEdit = ref<boolean>(false);

const { isOpen, editingTask, close } = useTaskModal();

watch(
  () => editingTask.value,
  (newTask) => {
    if (newTask) {
      isEdit.value = true;
      localTask.value = { ...newTask };
    } else {
      isEdit.value = false;
      localTask.value = {
        title: "",
        description: "",
        status: "new",
      };
    }
  },
  { immediate: true }
);

async function onSave() {
  if (!localTask.value.title.trim()) return;

  if (isEdit.value && localTask.value.id) {
    await updateTask(
      {
        title: localTask.value.title,
        description: localTask.value.description,
        status: localTask.value.status,
      },
      localTask.value.id
    );

    updateTaskInList(localTask.value);
  } else {
    const newTask: Task = {
      title: localTask.value.title,
      description: localTask.value.description,
      status: localTask.value.status,
      created_at: new Date().toISOString(),
    };
    await createTask(newTask);

    updateTaskInList(newTask);
  }

  close();
}

async function onDelete() {
  if (isEdit.value && localTask.value.id) {
    try {
      await deleteTask(localTask.value.id);
      deleteTaskInList(localTask.value.id);
      close();
    } catch (err) {
      alert("Не удалось удалить задачу");
      console.error(err);
    }
  }
  close();
}
</script>
<template>
  <Transition
    enter-from-class="opacity-0 scale-[1.03]"
    leave-to-class="opacity-0 scale-[1.02]"
  >
    <div
      class="fixed inset-0 flex items-center justify-center bg-black/30 backdrop-blur-sm"
      @click.self="close"
    >
      <div class="bg-white p-6 rounded-md w-full max-w-md shadow-lg">
        <h2 class="text-xl font-semibold mb-4">
          {{ isEdit ? "Редактировать" : "Создать" }} задачу
        </h2>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Заголовок</label>
          <input
            v-model="localTask.title"
            required
            class="w-full rounded-md border border-gray-300 px-3 py-2"
          />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1">Описание</label>
          <textarea
            v-model="localTask.description"
            required
            class="w-full rounded-md border border-gray-300 px-3 py-2"
          ></textarea>
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-1">Статус</label>
          <select
            v-model="localTask.status"
            class="w-full rounded-md border border-gray-300 px-3 py-2 cursor-pointer"
          >
            <option value="new">Новая</option>
            <option value="in-progress">В работе</option>
            <option value="completed">Завершена</option>
          </select>
        </div>
        <div class="flex justify-between gap-3">
          <button
            type="button"
            @click="onDelete()"
            class="px-4 py-2 bg-red-400 rounded-md text-white hover:bg-red-500 cursor-pointer"
          >
            Удалить
          </button>
          <HStack class="gap-2">
            <button
              type="button"
              @click="close"
              class="px-4 py-2 bg-gray-200 rounded-md hover:bg-gray-300 cursor-pointer"
            >
              Отмена
            </button>
            <button
              @click="onSave()"
              class="px-4 py-2 bg-indigo-400 rounded-md hover:bg-indigo-500 text-white cursor-pointer"
            >
              {{ isEdit ? "Сохранить" : "Создать" }}
            </button>
          </HStack>
        </div>
      </div>
    </div>
  </Transition>
</template>
