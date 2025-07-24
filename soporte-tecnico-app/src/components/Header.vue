<script>
import CategoriesService from "@/services/CategoriesService"
import { useHttpStore } from '@/stores/httpStore'
import { useAuthStore } from '@/stores/auth'
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'

export default {
    setup() {
        // Stores
        const auth = useAuthStore()
        const httpStore = useHttpStore()
        const router = useRouter()

        // Services
        const categoriesService = CategoriesService(httpStore.apiClient)

        // Computeds
        const isLoggedIn = computed(() => auth.isAuthenticated)
        const nameUser = computed(() => auth.user || '')
        const rolUser = computed(() => auth.rol)

        // Refs
        const categories = ref([])

        // Métodos
        function logout() {
            auth.logout()
            router.push('/login')
        }

        // se indica que cada vez que el rol del usario cambie se ejecute
        // rolUser -> Valor que tiene que observar si cambia
        // newRol -> Valor nuevo del rol del usuario
        watch(rolUser, async (newRol) => {
            console.log('Rol actualizado:', newRol)
            if (newRol === 'support') {
                try {
                    const { data } = await categoriesService.getCategories()
                    categories.value = data
                    console.log('Categorías cargadas:', data)
                } catch (error) {
                    console.error("Error al cargar categorías:", error)
                }
            }
        }, { immediate: true }) // Ejecuta al cargar el componente, como sustitucion del onMounted


        return {
            isLoggedIn,
            nameUser,
            rolUser,
            categories,
            logout
        }
    }
}
</script>

<template>
    <header class="page-header">
        <div class="user-container">
            <div class="title">
                <RouterLink to="/">Soporte técnico App</RouterLink>
            </div>
            <nav class="user-nav">
                <ul class="user-nav-list">
                    <li v-if="isLoggedIn">
                        <a href="#">{{ nameUser }}</a>
                    </li>
                    <li v-if="isLoggedIn">
                        <a href="#" @click.prevent="logout">Cerrar sesión</a>
                    </li>
                    <li v-else>
                        <RouterLink to="/login">Iniciar sesión</RouterLink>
                    </li>
                </ul>
            </nav>
            <button class="menu-toggle" id="toggleMenu">☰</button>
        </div>
        <nav class="web-nav">
            <ul class="web-nav-list" v-if="rolUser">
                <li>
                    <RouterLink to="/form-support-case">Insertar incidencia</RouterLink>
                </li>
                <li>
                    <RouterLink to="/">Inicio</RouterLink>
                </li>
                <li>
                    <RouterLink :to="{name: 'report-list-view', params: {id:0}}" v-if="rolUser === 'user'">Ver mis incidencias</RouterLink>
                    <RouterLink :to="{name: 'report-list-view', params: {id:0}}" v-if="rolUser === 'support'">Ver incidencias</RouterLink>
                </li>
            </ul>
        </nav>
    </header>
</template>

<style scoped>
.page-header {
    position: fixed;
    z-index: 10;
    width: 100%;
    background-color: white;
}

.user-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    height: 4.5em;
    align-items: center;
    color: #1e293b;
}

.title {
    display: flex;
    font-size: 1.5em;
    font-weight: bold;
    margin: auto 2em;
}

.title a{
    font-size: 1.2em;
    text-decoration: none;
    color: #1e293b;
}

.title img {
    width: 9em;
}

.user-nav {
    display: flex;
    margin: auto 2em;
}

.support-container {
    display: flex;
    flex-direction: row;
    gap: 1.6em;
}

.web-nav {
    display: flex;
    justify-content: center;
    height: 2.7em;
    align-items: center;
    background-color: #1e293b;
}

.user-nav-list,
.web-nav-list {
    display: flex;
    list-style: none;
    gap: 2em;
    justify-content: flex-end;
}

.user-nav-list li a,
.web-nav-list li a {
    text-decoration: none;
    transition: color 0.3s ease;
}

.web-nav-list li a {
    color: white;
}

.user-nav-list li a {
    color: #1e293b;
}

.user-nav-list li a:hover,
.web-nav-list li a:hover {
    color: #28a1d5;
}

.menu-toggle {
    display: none;
    font-size: 2em;
    background: none;
    border: none;
    color: #1e293b;
    margin-right: 1em;
}

/* Responsive */
@media (max-width: 768px) {
    .user-nav {
        margin: 0px;
        max-height: 0;
        overflow: hidden;
        position: absolute;
        top: 4.5em;
        right: 0;
        background-color: white;
        flex-direction: column;
        width: 100%;
        padding: 0 1em;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
        transition: max-height 0.5s ease;
    }

    .user-nav.active {
        max-height: 300px;
        padding: 1em;
    }

    .web-nav.active {
        top: 14.4em;
        height: 4.5em;
    }

    .menu-toggle {
        display: block;
    }

    .user-nav-list {
        flex-direction: column;
        gap: 1em;
        align-items: center;
        justify-content: center;
    }
}
</style>