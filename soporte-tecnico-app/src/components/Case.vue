<script>
import { ref, onMounted } from 'vue'
// servicios
import CategoriesService from "@/services/CategoriesService"
import CaseService from '@/services/CaseService'
// Pinia
import { useHttpStore } from '@/stores/httpStore'
import { useAuthStore } from '@/stores/auth'

export default {
    props: {
        report: {
            type: Object,
            required: true
        }
    },
    setup(props) {
        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        const authStore = useAuthStore()
        // service
        const caseService = CaseService(httpStore.apiClient)
        // se hace uso de la api
        const categoriesService = CategoriesService(httpStore.apiClient)
        const category = ref()

        // console.log('report  ', props.report)
        onMounted(async () => {
            try {
                // Se realiza una llamada a la API para obtener las incidencias
                const { data } = await categoriesService.getCategoryById(props.report.categoryId)
                category.value = data
                // console.log('Categoria ', category.value)
            } catch (error) {
                console.log("Error al obtener la categoria: " + error)
            }
        })

        // funcion dinamica temporal
        function getStatusClass(status) {
            const statusLowerCase = status.toLowerCase()
            const style = {
                abierta: 'status-active',
                cerrado: 'status-inactive',
            }
            return style[statusLowerCase]
        }

        // Para los estados
        // Para controlar la opcion seleccionada en radios y que coincida con localReport.status
        const selectedStatus = ref(props.report.status.toLowerCase())

        function onStatusChange(id, event) {
            const newStatus = event.target.value
            const confir = confirm("¿Seguro que quieres cambiar el estado a " + newStatus + " ?")
            if (!confir) {
                // Si el usuario cancela no hace nada
                return;
            }
            // console.log("Nuevo estado seleccionado:", newStatus)

            if (authStore.rol === 'support') {
                caseService.setStatus(id, newStatus)
                    .then(() => {
                        // para cambiar el estado sin tener que recargar (segun documentacion es mala practica pero bueno :)
                        // no es lo mejor pero si recargo pagina se ronpe la session por que esta a medias
                        props.report.status = newStatus
                    })
                    .catch((err) => {
                        alert("Error al actualizar estado: " + err.message)
                    })
            }
        }

        return { getStatusClass, category, authStore, selectedStatus, onStatusChange }
    },
}
</script>

<template>
    <!-- routerlink hacia el chat correspondiente en la vista ResponseSupportCaseView.vue -->
    <article v-if="category" class="issue" :style="{ borderLeft: `6px solid ${category.color}` }">
        <RouterLink :to="{ name: 'response-support-case', params: { id: report.id } }">
            <!-- color dinamico por cada category -->
            <header class="article-header">
                <!-- color dinamico por cada category -->
                <p class="category" :style="{ color: category.color }">{{
                    category.name }}</p>
                <h3 class="article-title">
                    {{ report.title }}
                </h3>
                <time class="date">{{ new Date(report.createdAt).toLocaleString() }}</time>
            </header>
            <div class="content">
                <div class="top">
                    <!-- color dinamico por el status -->
                    <p :class="['status', getStatusClass(report.status)]">Estado: {{
                        report.status }}</p>
                </div>
                <p class="description">
                    Descripción: {{ report.description }}
                </p>
            </div>
        </RouterLink>
        <div id="sets-status" v-if="authStore.rol == 'support'" class="sets-status">
            <!--Tienen nombres unicos para que no se bugea al estar varios en el list-->
            <label class="status-option">
                <input type="radio" :name="`status-${report.id}`" value="abierta"
                    @change="onStatusChange(report.id, $event)" v-model="selectedStatus" />
                <span class="label" style="color: #528d2b;">Abierta</span>
            </label>

            <label class="status-option">
                <input type="radio" :name="`status-${report.id}`" value="en espera"
                    @change="onStatusChange(report.id, $event)" v-model="selectedStatus" />
                <span class="label" style="color: #c8a62b;">En proceso</span>
            </label>

            <label class="status-option">
                <input type="radio" :name="`status-${report.id}`" value="cerrado"
                    @change="onStatusChange(report.id, $event)" v-model="selectedStatus" />
                <span class="label" style="color: #99301d;">Cerrado</span>
            </label>
        </div>
    </article>
</template>

<style scoped>
a {
    text-decoration: none;
}

.issue {
    background-color: #ffffff;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
    padding: 1.2em;
    border-radius: 8px;
    transition: transform 0.2s ease;
}

.issue:hover {
    transform: translateY(-2px);
}

.article-header {
    margin-bottom: 1em;
}

.category {
    font-size: 0.9em;
    font-weight: bold;
    text-transform: uppercase;
    margin-bottom: 0.3em;
}

/* estilo del contenedor del content y del title de este */
.content .top {
    display: flex;
    justify-content: space-between;
    font-size: 0.9em;
    margin-bottom: 0.5em;
    color: #475569;
}

.description {
    font-size: 1em;
    color: #334155;
}

.article-title {
    font-size: 1.2em;
    margin: 0.2em 0;
    color: #1e293b;
}

.date {
    font-size: 0.85em;
    color: #64748b;
}

/* estilos de los statuss, dinammicos */
.status {
    font-weight: bold;
    color: #c8a62b;
}

.status-inactive {
    color: #99301d;
}

.status-active {
    color: #528d2b;
}

.sets-status {
    display: flex;
    gap: 1em;
    align-items: center;
    justify-content: flex-start;
    margin-top: 20px;
}

.status-option {
    display: flex;
    align-items: center;
    gap: 0.4em;
    background: #f1f5f9;
    padding: 0.5em 0.8em;
    border-radius: 8px;
    transition: background 0.3s ease;
    cursor: pointer;
}

.status-option:hover {
    background: #e2e8f0;
}

.status-option input[type="radio"] {
    accent-color: currentColor;
    cursor: pointer;
}
</style>