<script setup lang="ts">
import { ref } from 'vue';
import { emailData } from '../models/email.model';
import EmailContent from './emailContent.vue';

interface Props {
  emails: emailData[];
}

defineProps<Props>();

// State for tracking selected email and content visibility
const selectedEmail = ref<emailData | null>(null);
const showEmailContent = ref(false);

const handleEmailClick = (email: emailData) => {
  selectedEmail.value = email;
  showEmailContent.value = true; // Show the email content panel
};
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
    <!-- Left Panel: Email List -->
    <div
      :class="{
        'md:col-span-4': !showEmailContent,
        'md:col-span-1': showEmailContent,
      }"
      class="bg-white rounded-lg shadow-md p-4"
    >
      <!-- Header Row -->
      <div class="grid grid-cols-4 gap-2 border-b border-gray-300 pb-2 font-bold text-gray-700">
        <div>Subject</div>
        <div>From</div>
        <div>To</div>
        <div>Date</div>
      </div>

      <!-- Email Rows -->
      <div
        v-for="email in emails"
        :key="email._source.ID"
        @click="handleEmailClick(email)"
        class="grid grid-cols-4 gap-2 border-b border-gray-300 py-2 cursor-pointer hover:bg-gray-100"
        :class="{ 'bg-gray-200': selectedEmail?._source.ID === email._source.ID }"
      >
        <div class="font-semibold truncate">{{ email._source.subject }}</div>
        <div class="text-sm text-gray-600 truncate">{{ email._source.from }}</div>
        <div class="text-sm text-gray-600 truncate">{{ email._source.to }}</div>
        <div class="text-sm text-gray-500 truncate">
          {{ new Date(email._source.Date).toLocaleDateString() }}
        </div>
      </div>
    </div>

    <!-- Right Panel: Email Details -->
    <div v-if="showEmailContent" class="md:col-span-3">
      <button
        class="mb-2 px-4 py-2 bg-blue-500 text-white rounded-lg shadow hover:bg-blue-600"
        @click="showEmailContent = false"
      >
        Back to Emails
      </button>
      <EmailContent v-if="selectedEmail" :email="selectedEmail" />
      <div v-else class="bg-white rounded-lg shadow-md p-4">
        <p class="text-gray-500">Select an email to view details.</p>
      </div>
    </div>
  </div>
</template>
