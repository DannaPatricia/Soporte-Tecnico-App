<script>
import { ref, onMounted, onUnmounted } from 'vue'
import ChatMessage from './ChatMessage.vue'
import MessageService from '@/services/MessageService'
// Pinia stores
import { useAuthStore } from '@/stores/auth'
import { useHttpStore } from '@/stores/httpStore'

export default {
  name: 'ChatBox',
  components: {
    ChatMessage
  },
  props: {
    supportCaseId: {
      type: Number,
      required: true
    }
  },
  setup(props) {
    // se declara el store que contiene la api
    const httpStore = useHttpStore()

    // El usuario para los mensajes (talvez sea mejor que lo pase por prop desde el padre para no importar pero no se)
    const authStore = useAuthStore()
    // se hace uso de la api
    const messageService = MessageService(httpStore.apiClient)

    // Para el gift de carga y mostrar en caso de error
    const loading = ref(true)
    const loadingError = ref(false)

    // mensajes
    const messages = ref([])
    const newMessage = ref('')
    let interval = null
    // Cargar mensajes cada x tiempo
    const fetchMessages = async () => {
      try {
        const res = await messageService.getMessagesByCase(props.supportCaseId)
        messages.value = res.data
      } catch (err) {
        console.error('Error fetching messages:', err)
        loadingError.value = true
      } finally {
        // quitar gif de carga
        setTimeout(function () {
          loading.value = false
        }, 500)
      }
    }

    const sendMessage = async () => {
      const text = newMessage.value.trim()
      // Si no hay mensaje no hace nada
      if (!text) {
        return
      }

      // darle los datos que me pide la db
      const message = {
        content: text,
        sent_at: new Date().toISOString(),
        supportCaseId: props.supportCaseId,
        userId: authStore.id,
        sender: authStore.user
      }

      try {
        await messageService.sendMessage(message)
        newMessage.value = ''
        await fetchMessages() // recarga mensajes despues de enviar
      } catch (err) {
        console.error('Error sending message:', err)
      }
    }

    // cuando se monta ejecuta la primera llamada a la api y luego actualiza mensajes cada 3 segundos
    onMounted(() => {
      fetchMessages()
      interval = setInterval(fetchMessages, 3000)
    })

    // para eliminar el timer cuando el usaurio abandone el chat
    onUnmounted(() => {
      clearInterval(interval)
    })

    return {
      messages,
      newMessage,
      sendMessage,
      loading,
      loadingError
    }
  }
}

</script>

<template>
  <section class="chat-container">
    <div v-if="loading" class="loadingGift">
      <img src="/loading.gif" alt="Loading..." />
    </div>
    <!--Falta que te desplace al final del chat cuando esta lleno y escribes pero bueno-->
    <div v-show="!loading" class="chat-messages" v-if="messages && !loadingError">
      <ChatMessage v-for="message in messages" :message="message"/>
    </div>

    <div v-if="loadingError">
      <h2>Error al cargar el chat</h2>
      <p>Por favor intenta de nuevo más tarde.</p>
    </div>

    <div class="chat-input">
      <input v-model="newMessage" type="text" placeholder="Escribe tu mensaje..." @keyup.enter="sendMessage" />
      <button @click="sendMessage">Enviar</button>
    </div>
  </section>
</template>

<style scoped>
/*Parte del chat contenedor*/
.chat-container {
  display: flex;
  flex-direction: column;
  height: 60.9vh;
  max-height: 60.9vh;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  background-color: #f0f0f0;
}

.chat-input {
  display: flex;
  border-top: 1px solid #e5e7eb;
  padding: 0.8rem;
  background-color: #f9fafb;
}

.chat-input input {
  flex: 1;
  padding: 0.8rem;
  border: 1px solid #ccc;
  border-radius: 0.25rem;
  font-size: 1rem;
}

.chat-input button {
  margin-left: 0.5rem;
  padding: 0.8rem 1.2rem;
  background-color: #1e293b;
  color: white;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
}

.chat-input button:hover {
  background-color: #334155;
}

/* gif de carga*/
.loadingGift {
  margin: auto;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50vh;
  width: 50vh;
}
</style>