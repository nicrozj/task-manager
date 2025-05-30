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
    enter-active-class="transition duration-150 ease-out"
    enter-from-class="opacity-0 scale-[1.03]"
    leave-active-class="transition duration-50 ease-in"
    leave-to-class="opacity-0 scale-[1.02]"
  >
    <div
      class="fixed inset-0 flex items-center justify-center bg-black/30 backdrop-blur-sm px-4"
      @click.self="close"
    >
      <div
        class="bg-white p-6 rounded-md w-full max-w-md shadow-lg dark:bg-slate-800"
      >
        <h2 class="text-xl font-semibold mb-4 dark:text-white">
          {{ isEdit ? "Редактировать" : "Создать" }} задачу
        </h2>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1 dark:text-white"
            >Заголовок</label
          >
          <input
            v-model="localTask.title"
            required
            class="w-full rounded-md border border-gray-300 px-3 py-2 dark:border-slate-700 dark:focus:outline-none"
          />
        </div>
        <div class="mb-4">
          <label class="block text-gray-700 mb-1 dark:text-white"
            >Описание</label
          >
          <textarea
            v-model="localTask.description"
            required
            class="w-full rounded-md border resize-none h-48 border-gray-300 px-3 py-2 dark:border-slate-700 dark:focus:outline-none"
          ></textarea>
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 mb-1 dark:text-white">Статус</label>
          <select
            v-model="localTask.status"
            class="w-full rounded-md border border-gray-300 px-3 py-2 cursor-pointer dark:border-slate-700 dark:focus:outline-none"
          >
            <option value="new">Новая</option>
            <option value="in-progress">В работе</option>
            <option value="completed">Завершена</option>
          </select>
        </div>
        <div class="flex justify-center gap-3">
          <button
            type="button"
            @click="onDelete()"
            class="p-2 bg-red-400 rounded-md text-white hover:bg-red-500 cursor-pointer flex items-center justify-center"
          >
            <span class="material-symbols-rounded">delete</span>
          </button>
          <button
            type="button"
            @click="close"
            class="p-2 bg-gray-200 rounded-md hover:bg-gray-300 text-gray-600 cursor-pointer flex items-center justify-center"
          >
            <span class="material-symbols-rounded">undo</span>
          </button>
          <button
            @click="onSave()"
            class="p-2 bg-indigo-400 rounded-md hover:bg-indigo-500 text-white cursor-pointer flex items-center justify-center"
          >
            <span class="material-symbols-rounded">check</span>
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>
