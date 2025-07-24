<script setup>
const props = defineProps({
    message: {
        type: Object,
        required: true,
    },
})
</script>

<template>
    <div class="message support">
        <b>{{ props.message.sender }}</b><br/>
        <p>{{ props.message.content }}</p>
        <p class="date">{{ new Date(props.message.sent_at).toLocaleString() }}</p>
    </div>
</template>
<style scoped>
.message {
    max-width: 100%;
    padding: 0.8rem;
    border-radius: 0.5rem;
    position: relative;
    word-wrap: break-word;
}

.message.support {
    background-color: #e0f2fe;
}

.message.user {
    background-color: #d1fae5;
}

.date{
    display: flex;
    justify-content: end;
    align-items: end;
    color: rgb(110, 110, 110);
}
</style>