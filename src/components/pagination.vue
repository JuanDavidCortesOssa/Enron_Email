<!-- Pagination.vue -->
<script setup lang="ts">
defineProps<{
  emailsHeaderString: string;
  moveFromPage: (increaseValue: boolean) => void;
}>();
</script>

<template>
  <div class="flex items-center justify-center gap-4 py-4">
    <button
      class="flex items-center justify-center w-10 h-10 bg-blue-500 text-white rounded-full shadow-md hover:bg-blue-600 transition-colors duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
      @click="moveFromPage(false)"
    >
      <span class="text-lg font-bold">&lt;</span>
    </button>

    <div class="text-gray-700 font-medium text-lg bg-gray-100 px-4 py-2 rounded-lg shadow-sm">
      {{ emailsHeaderString }}
    </div>

    <button
      class="flex items-center justify-center w-10 h-10 bg-blue-500 text-white rounded-full shadow-md hover:bg-blue-600 transition-colors duration-200 disabled:bg-gray-400 disabled:cursor-not-allowed"
      @click="moveFromPage(true)"
    >
      <span class="text-lg font-bold">&gt;</span>
    </button>
  </div>
</template>

<style scoped>

</style>