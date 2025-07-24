<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

import UserService from '@/services/UserService.js'
// importar el store de pinia
import { useAuthStore } from '@/stores/auth'
import { useHttpStore } from '@/stores/httpStore'

export default {
    setup() {
        // router
        const router = useRouter()
        // session de pinia
        const auth = useAuthStore()
        if (auth.isAuthenticated) {
            //console.log("ya logeado")
            router.push('/')
        }

        // se declara el store que contiene la api
        const httpStore = useHttpStore()
        // se hace uso de la api
        const userService = UserService(httpStore.apiClient)

        // datos del form
        const username = ref('')
        const password = ref('')
        // Para validar si estan bien los datos y luego redirigir al user
        const error = ref('')
        const success = ref('')

        // Funcion del formulario
        const login = async () => {
            try {
                const res = await userService.login(username.value, password.value)
                success.value = res.data.message
                error.value = ''
                // Apartado para datos de pinia

                console.log(res.data)
                // guardar el role con pinia
                auth.setUser(res.data.user.username, res.data.user.role, res.data.user.id)
                router.push('/')
            } catch (err) {
                success.value = ''
                console.log(err)
                error.value = err.response.data || 'Error al iniciar sesión'
            }
        }

        return { username, password, login, error, success }
    }
}
</script>

<template>
    <main>
        <section id="login-container">
            <form @submit.prevent="login">
                <h1>Login</h1>
                <label for="username">Ingrese su usuario</label>
                <input id="username" name="username" type="text" v-model="username" required />

                <label for="password">Ingrese la clave</label>
                <input id="password" name="password" type="password" v-model="password" required />

                <button type="submit">Iniciar sesión</button>
                <!--De momento aqui para ver si va con go y si validacion-->
                <p v-if="error" style="color: red">{{ error }}</p>
                <p v-if="success" style="color: green">{{ success }}</p>
            </form>
            <img id="login-img" src="/login-back.jpg" alt="Login" />
        </section>
    </main>
</template>

<style scoped>
#login-container {
    margin: 5rem auto;
    background: white;
    display: flex;
    flex-wrap: wrap;
    padding: 2rem;
    border-radius: 12px;
    box-shadow: 0 4px 25px rgba(0, 0, 0, 0.15);
    max-width: 800px;
    width: 95%;
    gap: 2rem;
    align-items: center;
}

form {
    flex: 1;
    display: flex;
    flex-direction: column;
    min-width: 250px;
    gap: 1.2rem;
}

h1 {
    font-size: 2rem;
    color: #1e293b;
    text-align: center;
    margin-bottom: 1rem;
}

label {
    font-size: 1.1rem;
    font-weight: bold;
    color: #1e293b;
}

input {
    padding: 1rem;
    border: 1px solid #cbd5e1;
    border-radius: 6px;
    font-size: 1rem;
    background-color: #f9fafa;
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

#login-img {
    flex: 1;
    max-width: 400px;
    width: 100%;
    border-radius: 8px;
    object-fit: cover;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.08);
}

@media (max-width: 768px) {
    #login-container {
        flex-direction: column;
        text-align: center;
    }

    #login-img {
        max-width: 90%;
    }
}
</style>
