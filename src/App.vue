<script setup lang="ts">
import { ref, onMounted } from "vue";
import PageHeader from './components/PageHeader.vue';
import emailsList from "./components/emailsList.vue";
import Pagination from "./components/pagination.vue";
import SearchBar from "./components/searchBar.vue";
import axiosClient from "./utils/axios"
import { emailData } from "./models/email.model";
import { emailRequest } from "./models/emailRequest.model";

const emails = ref<emailData[]>([]);
const emailsNumber = ref<number>(0);
const emailsHeaderString = ref<string>("Initial value");
const currentEmailNumber = ref<number>(0);
const initialRequest: emailRequest = {
  term: "example",
  from: 0
};
const search = ref("example");

const fetchEmails = async (request: emailRequest) => {
  try {
    const response = await axiosClient.post("/emails", request);
    emails.value = response.data.hits.hits;
    emailsNumber.value = response.data.hits.total.value; 

  } catch (error) {
    console.log(error);
  }
}

const filterEmails = async () =>
{
  var request: emailRequest = {
    term: search.value,
    from: currentEmailNumber.value
  }
  await fetchEmails(request)

  updateEmailsPageString()
}

const moveFromPage = (increaseValue: boolean) =>{
  var modifiedValue: number = currentEmailNumber.value;
  let maxEmailsPerRequest: number = 20;
  if(increaseValue){
    modifiedValue+= maxEmailsPerRequest;
  }else{
    modifiedValue-= maxEmailsPerRequest;
  }
  console.log(modifiedValue)
  if(modifiedValue<0){
    modifiedValue = 0;
  }else if(modifiedValue>emailsNumber.value){
    modifiedValue = currentEmailNumber.value
  }

  currentEmailNumber.value = modifiedValue;
  console.log(modifiedValue)
  filterEmails();
}

const updateEmailsPageString = () => {
    let emailPageString: string = ""; 
    let maxEmails: number = 20;
    let lastEmailNumber: number = 20;

    lastEmailNumber = currentEmailNumber.value + maxEmails
    if(lastEmailNumber > emailsNumber.value){
      lastEmailNumber = emailsNumber.value;
    }

    emailPageString = (currentEmailNumber.value + 1) + "-" + (lastEmailNumber + 1)
    + " of " + (emailsNumber.value + 1);

    emailsHeaderString.value = emailPageString; 
};

const handleSearch = (value: string) => {
  search.value = value;
  currentEmailNumber.value = 0;
  filterEmails();
};

onMounted(async ()=>{
  await fetchEmails(initialRequest)
  updateEmailsPageString()
})

</script>

<template>
  <PageHeader />
  <Pagination
    :emails-header-string="emailsHeaderString"
    :move-from-page="moveFromPage"
  />
  <div class="container mx-auto max-w-7xl p-4 bg-white rounded-lg shadow-md">
    <SearchBar @search="handleSearch" />
    <emailsList :emails="emails" />
  </div>
</template>
