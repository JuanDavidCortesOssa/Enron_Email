<script setup lang="ts">
import { defineProps } from 'vue';
import { emailData } from '../models/email.model';

interface Props {
  email: emailData;
}

defineProps<Props>();
</script>

<template>
  <div class="email-content-container">
    <h2 class="text-lg font-bold mb-4">Email Details</h2>
    <div class="whitespace-pre-wrap">
      <p><strong>Subject:</strong> {{ email._source.subject }}</p>
      <p><strong>From:</strong> {{ email._source.from }}</p>
      <p><strong>To:</strong> {{ email._source.to }}</p>
      <p><strong>Date:</strong> {{ new Date(email._source.Date).toLocaleString() }}</p>
      <p><strong>Message:</strong></p>
      <p>{{ email._source.Message }}</p>
    </div>
  </div>
</template>

<style scoped>
.email-content-container {
  max-height: calc(130vh - 8rem); /* Adjust height based on layout */
  overflow-y: auto; /* Enable scrolling */
  padding: 1rem;
  background-color: #ffffff;
  border-radius: 0.5rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
