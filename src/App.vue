<script setup lang="ts">
import { onMounted } from "vue";
import PageHeader from './components/PageHeader.vue';
import axiosClient from "./utils/axios"

const fetchEmails = async() => {
  //const {data} = await axiosClient.get("/emails")
  console.log("asd asd JUAN DAVID cORTS OSSA")
}

onMounted(()=>{
  fetchEmails()
})

</script>

<template>
  <PageHeader/>
</template>
