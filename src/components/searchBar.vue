<script setup lang="ts">
import { ref } from "vue";

const localSearch = ref("");

const emit = defineEmits<{
  (e: "search", value: string): void;
}>();

const handleEnter = () => {
  emit("search", localSearch.value);
};
</script>

<template>
  <div class="mb-4">
    <input
      type="text"
      class="bg-white border border-gray-300 rounded-lg shadow-md w-full px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
      placeholder="Search email"
      v-model="localSearch"
      @keyup.enter="handleEnter"
    />
  </div>
</template>

<style scoped>

</style>