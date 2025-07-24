<script>
import { ref, onMounted } from "vue"
import CaseService from "@/services/CaseService"
import CategoriesService from "@/services/CategoriesService"
import { useAuthStore } from '@/stores/auth'
import { useHttpStore } from '@/stores/httpStore'
import { useRouter } from 'vue-router'

export default {
    setup() {
        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        // se hace uso de la api
        const caseService = CaseService(httpStore.apiClient)
        const categoriesService = CategoriesService(httpStore.apiClient)

        // se declara el router para redirifir
        const router = useRouter()

        // Generar el select de categorias
        const categories = ref([])

        // Datos del formulario
        const selectedCategory = ref("")
        const title = ref("")
        const description = ref("")
        const userId = ref("")

        onMounted(async () => {
            // Obtener el id del usuario logeado
            // si no esta logeado se redigira a la pagina de login
            const authStore = useAuthStore()
            if (!authStore.id) {
                router.push('/login')
            }
            // caso contrario se guardará el id del usuario
            userId.value = authStore.id
            console.log(userId)
            try {
                // LLamada a la api de categorias
                categories.value = await categoriesService.getCategories()
                console.log(categories.value)
            } catch (error) {
                console.error("Error al cargar categorías:", error)
            }
        })

        // variables de exito y de error del formulario
        const succesfull = ref(false)
        const errorForm = ref(false)

        // evento al enviar el formulario
        const submitForm = async () => {
            // se decalra unnuevo caso con las propiedades sekleccionadas
            const newCase = {
                title: title.value,
                description: description.value,
                categoryId: selectedCategory.value,
                userId: userId.value,
            }
            console.log("Datos enviados:", newCase)

            // se realiza el post a la api
            try {
                const response = await caseService.createCase(newCase)
                console.log("Caso creado:", response.data)
                // vaciar los campos del formulario
                title.value = ""
                description.value = ""
                selectedCategory.value = ""
                //alert("Incidencia registrada correctamente.")
                // Redirigir
                succesfull.value = true
                // redirigir con retardo para que se vea el mensaje
                setTimeout(() => {
                    router.push({ name: "report-list-view", params: { id: 0 } })
                }, 4000)
            } catch (error) {
                errorForm.value = true
                console.error("Error al enviar la incidencia:", error)
                //alert("Error al enviar la incidencia.")
                setTimeout(() => {
                    router.push('/form-support-case')
                }, 4000)
            }
        }

        return {
            categories,
            selectedCategory,
            title,
            description,
            submitForm,
            succesfull,
            errorForm
        }
    },
}
</script>

<template>
    <main>
        <section class="form">
            <form @submit.prevent="submitForm()" v-show="!succesfull && !errorForm">
                <h1>Cuéntenos su problema</h1>
                <label for="category">¿Qué tipo de problema tiene?</label>
                <!--Una llamada a la api de la tabla categoria-->
                <select name="category" id="category" v-model="selectedCategory" required>
                    <option disabled value="">Seleccione una categoría</option>
                    <option v-for="cat in categories.data" :key="cat.id" :value="cat.id">
                        {{ cat.name }}
                    </option>
                </select>
                <label for="title">¿Sobre qué es su problema?</label>
                <input type="text" name="title" id="title" v-model="title" required />
                <label for="description">Describa el problema detalladamente</label>
                <textarea name="description" id="description" v-model="description" required></textarea>
                <button type="submit">Enviar</button>
            </form>
            <!--Mensaje de exito-->
            <div class="succesfull" v-if="succesfull">
                Incidencia registrada correctamente.
            </div>
            <!--Mensaje de error-->
            <div class="error" v-if="errorForm">
                Hubo un error al registrar la incidencia. Inténtelo de nuevo más tarde.
            </div>
        </section>

        <section id="problems-and-solutions">
            <h3>Problemas frecuentes y soluciones</h3>
            <details>
                <summary>¿No puede iniciar sesión?</summary>
                <p>
                    Verifique que su contraseña sea correcta y que su cuenta esté activa.
                </p>
            </details>
            <details>
                <summary>Problemas con pagos</summary>
                <p>
                    Asegúrese de que su método de pago esté actualizado y tenga fondos
                    suficientes.
                </p>
            </details>
            <details>
                <summary>Otros problemas</summary>
                <p>
                    Revise nuestras preguntas frecuentes o contáctenos directamente para
                    más ayuda.
                </p>
            </details>
        </section>
    </main>
</template>

<style scoped>
.form {
    margin: 2rem auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    background: white;
    border-radius: 12px;
    padding: 2.5rem;
    gap: 1.2rem;
    box-shadow: 0 4px 25px rgba(0, 0, 0, 0.15);
    max-width: 800px;
    width: 95%;
}

form {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
}

h1 {
    margin-bottom: 1rem;
    font-size: 2rem;
    color: #1e293b;
    text-align: center;
}

label {
    font-size: 1.1rem;
    font-weight: bold;
    color: #1e293b;
}

input,
select,
textarea {
    padding: 1rem;
    border: 1px solid #ccc;
    border-radius: 6px;
    font-size: 1rem;
    background-color: #f9fafa;
}

textarea {
    height: 200px;
    resize: vertical;
}

form button {
    margin-top: 1rem;
    padding: 0.8rem;
    background-color: #1e293b;
    color: white;
    border: none;
    border-radius: 6px;
    cursor: pointer;
    font-weight: 600;
    font-size: 1rem;
    transition: background-color 0.3s ease;
}

form button:hover {
    background-color: #334155;
}

#problems-and-solutions {
    max-width: 800px;
    margin: 3rem auto;
    background: white;
    border-radius: 12px;
    padding: 2rem;
    box-shadow: 0 4px 25px rgba(0, 0, 0, 0.1);
}

#problems-and-solutions h3 {
    font-size: 1.5rem;
    margin-bottom: 1rem;
    color: #1e293b;
}

details {
    margin-bottom: 1rem;
    padding: 1rem;
    background-color: #f1f5f9;
    border-radius: 6px;
    border: 1px solid #cbd5e1;
    cursor: pointer;
}

summary {
    font-weight: bold;
    font-size: 1rem;
    color: #1e293b;
    cursor: pointer;
}

summary::marker {
    color: #1e293b;
}

details p {
    margin-top: 0.5rem;
    font-size: 0.95rem;
    color: #475569;
}

/*Mensajes de exito */
.succesfull {
    margin-top: 1rem;
    padding: 1rem;
    background-color: #d1fae5;
    color: #065f46;
    border: 1px solid #10b981;
    border-radius: 6px;
    font-weight: bold;
    text-align: center;
}

.error {
    margin-top: 1rem;
    padding: 1rem;
    background-color: #fee2e2;
    color: #991b1b;
    border: 1px solid #f87171;
    border-radius: 6px;
    font-weight: bold;
    text-align: center;
}
</style>
