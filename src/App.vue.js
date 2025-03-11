import { ref, onMounted } from "vue";
import PageHeader from './components/PageHeader.vue';
import emailsList from "./components/emailsList.vue";
import Pagination from "./components/pagination.vue";
import SearchBar from "./components/searchBar.vue";
import axiosClient from "./utils/axios";
const emails = ref([]);
const emailsNumber = ref(0);
const emailsHeaderString = ref("Initial value");
const currentEmailNumber = ref(0);
const search = ref("");
const fetchEmails = async (request) => {
    try {
        const response = await axiosClient.post("/emails", request);
        emails.value = response.data.hits.hits;
        emailsNumber.value = response.data.hits.total.value;
    }
    catch (error) {
        console.log(error);
    }
};
const filterEmails = async () => {
    var request = {
        term: search.value,
        from: currentEmailNumber.value
    };
    await fetchEmails(request);
    updateEmailsPageString();
};
const moveFromPage = (increaseValue) => {
    var modifiedValue = currentEmailNumber.value;
    let maxEmailsPerRequest = 20;
    if (increaseValue) {
        modifiedValue += maxEmailsPerRequest;
    }
    else {
        modifiedValue -= maxEmailsPerRequest;
    }
    if (modifiedValue < 0) {
        modifiedValue = 0;
    }
    else if (modifiedValue > emailsNumber.value) {
        modifiedValue = currentEmailNumber.value;
    }
    currentEmailNumber.value = modifiedValue;
    filterEmails();
};
const updateEmailsPageString = () => {
    let emailPageString = "";
    let maxEmails = 20;
    let lastEmailNumber = 20;
    lastEmailNumber = currentEmailNumber.value + maxEmails;
    if (lastEmailNumber > emailsNumber.value) {
        lastEmailNumber = emailsNumber.value;
    }
    emailPageString = (currentEmailNumber.value + 1) + "-" + (lastEmailNumber + 1)
        + " of " + (emailsNumber.value + 1);
    emailsHeaderString.value = emailPageString;
};
const handleSearch = (value) => {
    search.value = value;
    currentEmailNumber.value = 0;
    filterEmails();
};
onMounted(async () => {
    filterEmails();
}); /* PartiallyEnd: #3632/scriptSetup.vue */
const __VLS_ctx = {};
let __VLS_components;
let __VLS_directives;
/** @type {[typeof PageHeader, ]} */ ;
// @ts-ignore
const __VLS_0 = __VLS_asFunctionalComponent(PageHeader, new PageHeader({}));
const __VLS_1 = __VLS_0({}, ...__VLS_functionalComponentArgsRest(__VLS_0));
/** @type {[typeof Pagination, ]} */ ;
// @ts-ignore
const __VLS_3 = __VLS_asFunctionalComponent(Pagination, new Pagination({
    emailsHeaderString: (__VLS_ctx.emailsHeaderString),
    moveFromPage: (__VLS_ctx.moveFromPage),
}));
const __VLS_4 = __VLS_3({
    emailsHeaderString: (__VLS_ctx.emailsHeaderString),
    moveFromPage: (__VLS_ctx.moveFromPage),
}, ...__VLS_functionalComponentArgsRest(__VLS_3));
__VLS_asFunctionalElement(__VLS_intrinsicElements.div, __VLS_intrinsicElements.div)({
    ...{ class: "container mx-auto max-w-7xl p-4 bg-white rounded-lg shadow-md" },
});
/** @type {[typeof SearchBar, ]} */ ;
// @ts-ignore
const __VLS_6 = __VLS_asFunctionalComponent(SearchBar, new SearchBar({
    ...{ 'onSearch': {} },
}));
const __VLS_7 = __VLS_6({
    ...{ 'onSearch': {} },
}, ...__VLS_functionalComponentArgsRest(__VLS_6));
let __VLS_9;
let __VLS_10;
let __VLS_11;
const __VLS_12 = {
    onSearch: (__VLS_ctx.handleSearch)
};
var __VLS_8;
/** @type {[typeof emailsList, ]} */ ;
// @ts-ignore
const __VLS_13 = __VLS_asFunctionalComponent(emailsList, new emailsList({
    emails: (__VLS_ctx.emails),
}));
const __VLS_14 = __VLS_13({
    emails: (__VLS_ctx.emails),
}, ...__VLS_functionalComponentArgsRest(__VLS_13));
/** @type {__VLS_StyleScopedClasses['container']} */ ;
/** @type {__VLS_StyleScopedClasses['mx-auto']} */ ;
/** @type {__VLS_StyleScopedClasses['max-w-7xl']} */ ;
/** @type {__VLS_StyleScopedClasses['p-4']} */ ;
/** @type {__VLS_StyleScopedClasses['bg-white']} */ ;
/** @type {__VLS_StyleScopedClasses['rounded-lg']} */ ;
/** @type {__VLS_StyleScopedClasses['shadow-md']} */ ;
var __VLS_dollars;
const __VLS_self = (await import('vue')).defineComponent({
    setup() {
        return {
            PageHeader: PageHeader,
            emailsList: emailsList,
            Pagination: Pagination,
            SearchBar: SearchBar,
            emails: emails,
            emailsHeaderString: emailsHeaderString,
            moveFromPage: moveFromPage,
            handleSearch: handleSearch,
        };
    },
});
export default (await import('vue')).defineComponent({
    setup() {
        return {};
    },
});
; /* PartiallyEnd: #4569/main.vue */
//# sourceMappingURL=App.vue.js.map