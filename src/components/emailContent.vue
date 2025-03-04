<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { emailData } from '../models/email.model';

interface Props {
  email: emailData;
}

onMounted(() => {
  console.log('EmailDetails mounted');
})

defineProps<Props>();
</script>

<template>
  <div class="container mx-auto p-4">
      <!-- Right Panel: Email Content (Placeholder for selected email) -->
      <div class="md:col-span-3 bg-white rounded-lg shadow-md p-4">
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
    </div>
</template>