<script setup lang="ts">
import { ref, onMounted } from "vue";
import PageHeader from './components/PageHeader.vue';
import emailsList from "./components/emailsList.vue";
import axiosClient from "./utils/axios"
import { emailData } from "./models/email.model";

const emails = ref<emailData[]>([]);
const postData = {
  term: "example",
  from: 0
};

const fetchEmails = async () => {
  try {
    const response = await axiosClient.post("/emails", postData);
    emails.value = response.data.hits.hits;
  } catch (error) {
    console.log(error);
  }
}

onMounted(()=>{
  fetchEmails()
})

</script>

<template>
  <PageHeader/>
  <div> 
    <emailsList :emails="emails"/>
  </div>
</template>
