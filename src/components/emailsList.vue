<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { emailData } from '../models/email.model';
import emailContent from './emailContent.vue';
import EmailContent from './emailContent.vue';
// Define the props interface for the email list
interface Props {
  emails: emailData[];
}

// Use defineProps to receive emails
defineProps<Props>();

// State for tracking which email is selected
const selectedEmail = ref<emailData | null>(null);

// Handle click on an email to select it
const handleEmailClick = (email: emailData) => {
  selectedEmail.value = email;
};
</script>

<template>
  <div class="container mx-auto p-4">
    <!-- Main Grid Layout: Left Panel (Email List) and Right Panel (Email Details) -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <!-- Left Panel: Email List -->
      <div class="md:col-span-1 bg-white rounded-lg shadow-md p-4">
        <h2 class="text-lg font-bold mb-4">Emails</h2>
        <!-- Grid Headers for Columns -->
        <div class="grid grid-cols-4 gap-2 text-sm font-semibold text-gray-700 mb-2 border-b border-gray-300 pb-2">
          <div>Subject</div>
          <div>From</div>
          <div>To</div>
          <div>Date</div>
        </div>
        
        <!-- Email List with Grid Layout for Each Email -->
        <div 
          v-for="email in emails" 
          :key="email._source.ID" 
          class="grid grid-cols-4 gap-2 border-b border-gray-300 py-2 cursor-pointer hover:bg-gray-100"
          @click="handleEmailClick(email)"
          :class="{ 'bg-gray-200': selectedEmail?._source.ID === email._source.ID }"
        >
          <div class="font-semibold truncate">{{ email._source.subject }}</div>
          <div class="text-sm text-gray-600 truncate">{{ email._source.from }}</div>
          <div class="text-sm text-gray-600 truncate">{{ email._source.to }}</div>
          <div class="text-sm text-gray-500 truncate">{{ new Date(email._source.Date).toLocaleDateString() }}</div>
        </div>
      </div>

      <!-- Right Panel: Email Details -->
      <div class="md:col-span-3">
        <EmailContent v-if="selectedEmail" :email="selectedEmail" />
        <div v-else class="bg-white rounded-lg shadow-md p-4">
          <p class="text-gray-500">Select an email to view details.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* Optional: Additional custom styles for truncation and responsiveness */
.truncate {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  max-width: 100%;
}

.cursor-pointer:hover {
  transition: background-color 0.2s ease;
}

/* Ensure columns don’t overflow on smaller screens */
@media (max-width: 768px) {
  .grid-cols-4 {
    grid-template-columns: 1fr 1fr; /* Stack into 2 columns on mobile */
  }
  .font-semibold, .text-sm {
    font-size: 0.875rem; /* Adjust font size for mobile */
  }
}
</style>