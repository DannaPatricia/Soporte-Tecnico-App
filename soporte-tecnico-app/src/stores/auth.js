import { defineStore } from 'pinia'

// se define el store
export const useAuthStore = defineStore('auth', {
    // se definen los datos que se vana almacenar en el store
    state: () => ({
        user: null,
        id: null,
        rol: null,
        isAuthenticated: false
    }),
    // funciones para modificar las propiedades del store
    actions: {
        setUser(user, rol, id) {
            this.user = user
            this.rol = rol
            this.id = id
            this.isAuthenticated = true
        },
        logout() {
            this.user = null
            this.rol = null
            this.id = null
            this.isAuthenticated = false
        }
    }
})
