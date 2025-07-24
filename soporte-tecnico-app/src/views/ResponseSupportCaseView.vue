<script>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import ChatBox from '@/components/ChatBox.vue'
import Report from '@/components/Case.vue'
import CaseService from '@/services/CaseService'
// stroes pinia
import { useAuthStore } from '@/stores/auth'
import { useHttpStore } from '@/stores/httpStore'

export default {
    components: {
        Report,
        ChatBox,
    },
    setup() {
        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        // se hace uso de la api
        const caseService = CaseService(httpStore.apiClient)
        // se declara el obtejo route
        const route = useRoute()
        const router = useRouter()

        const reportId = route.params.id
        const report = ref()


        // si no esta logeado se redigira a la pagina de login
        const authStore = useAuthStore()
        if (!authStore.id) {
            router.push('/login')
        }

        onMounted(async () => {
            try {
                // Se realiza una llamada a la API para obtener la incidencia correspondiente
                const { data } = await caseService.getCaseById(reportId)
                // Para verificar si el usuario id pertenece a su caso (Si hay tiempo realizarlo desde el backend)
                report.value = data
                if (authStore.rol !== 'support' && report.value.userId !== authStore.id) {
                    router.push('/error')
                }
                
            } catch (error) {
                console.log("Error al cargar la incidencia por id: " + error)
            }
        })
        return { report }
    },
}

</script>

<template>
    <main>
        <section>
            <!-- se muestra la incidencia o reporte seleccionada -->
            <Report :report="report" v-if="report"></Report>
            <!--Apartado del chat con el de soporte-->
            <ChatBox :supportCaseId="report.id" v-if="report" />
        </section>
    </main>
</template>

<style scoped>
main {
    width: 75%;
    margin: auto;
    margin-top: 8em;
}

section {
    margin: auto;
}
</style>