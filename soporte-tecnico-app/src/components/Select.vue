<script>
import { ref, onMounted } from "vue"
import CategoriesService from "@/services/CategoriesService"
import { useHttpStore } from '@/stores/httpStore'
import { useRouter } from 'vue-router'

export default {
    setup() {
        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        // se hace uso de la api
        const categoriesService = CategoriesService(httpStore.apiClient)

        // se declara el router para redirifir
        const router = useRouter()

        // Generar el select de categorias
        const categories = ref([])

        // Datos del formulario
        const selectedCategory = ref("")

        onMounted(async () => {
            try {
                // LLamada a la api de categorias
                categories.value = await categoriesService.getCategories()
                console.log(categories.value)
            } catch (error) {
                console.error("Error al cargar categorías:", error)
            }
        })

        // evento al enviar el formulario
        const submitForm = async () => {
            // se decalra unnuevo caso con las propiedades sekleccionadas
            router.push({ name: "report-list-view", params: { id: selectedCategory.value } })
        }

        return {
            selectedCategory,
            submitForm,
            categories
        }
    },
}
</script>
<template>
    <form @submit.prevent="submitForm()">
        <select name="category" id="category" v-model="selectedCategory" required>
            <option disabled value="">Seleccione una categoría</option>
            <option v-for="cat in categories.data" :key="cat.id" :value="cat.id">
                {{ cat.name }}
            </option>
            <option value="0">Todas</option>
        </select>
        <button type="submit">Buscar</button>
    </form>
</template>

<style scoped>
form {
    margin: 2rem auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    border-radius: 12px;
    width: 35%;
}

select {
    padding: 1rem;
    border: 1px solid #ccc;
    border-radius: 6px;
    font-size: 1rem;
}

button {
    margin-top: 1rem;
    padding: 0.8rem;
    background-color: #1e293b;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    font-weight: 600;
    font-size: 1rem;
    height: 100%;
    transition: background-color 0.3s ease;
}

button:hover {
    background-color: #334155;
}
</style>