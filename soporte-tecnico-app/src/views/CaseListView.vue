<script>
import Case from '@/components/Case.vue'
import Select from '@/components/Select.vue'
import CaseService from '@/services/CaseService'
import { useHttpStore } from '@/stores/httpStore'
import { useAuthStore } from '@/stores/auth'
import { ref, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'

export default {
    components: {
        Select,
        Case,
    },
    setup() {
        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        // se hace uso de la api
        const caseService = CaseService(httpStore.apiClient)
        const authStore = useAuthStore()
        const router = useRouter()
        const route = useRoute()
        const casesList = ref([])

        // Funcion  para cargar casos dado  el id de cateogria
        const loadCases = async (categoryId) => {
            // Obtener el id del usuario logeado
            // si no esta logeado se redigira a la pagina de login
            if (!authStore.id) {
                router.push('/login')
                return
            }
            casesList.value = []

            try {
                console.log(authStore.rol, authStore)
                if (categoryId === '0') {
                    // Se realiza una llamada a la API para obtener las incidencias
                    const { data } = await caseService.getCases(authStore.rol, authStore.id)
                    casesList.value = data
                } else {
                    // se realiza unallamada a la papi para obtener las incidencias por categoria
                    const { data } = await caseService.getCasesByCategory(categoryId, authStore.id, authStore.rol)
                    casesList.value = data
                }
                console.log('Casos cargados:', casesList.value)
            } catch (error) {
                console.error('Error al obtener las incidencias:', error)
            }
        }

        // Se carga por primera vez
        onMounted(() => {
            loadCases(route.params.id)
        })

        // Se actualiza automaticamente cuando cambia el ID en la URL
        watch(
            () => route.params.id,
            (newId) => {
                loadCases(newId)
            }
        )
        return { casesList }
    }
}
</script>

<template>
    <main>
        <Select></Select>
        <div class="issues-container">
            <!-- se pasa por prop cada reporte o incidencia -->
            <Case v-for="report in casesList" :report="report"></Case>
        </div>
        <div class="message" v-if="casesList.length === 0">No tiene incidencias todavía</div>
    </main>
</template>

<style scoped>
.issues-container {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-template-rows: repeat(4, 1fr);
    grid-column-gap: 20px;
    grid-row-gap: 20px;
    margin: 20px auto;
    width: 80%;
}

.message {
    display: flex;
    justify-content: center;
    margin:  auto;
    width: 50%;
    padding: 2em;
    color: #b00020;
    background-color: #f8e8e8;
    border: 2px solid #f5c2c7;
    border-radius: 0.3em;
    font-size: 1.4em;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
}
</style>