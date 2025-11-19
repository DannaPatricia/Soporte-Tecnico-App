import { defineStore } from 'pinia'
import axios from 'axios'

// se define el store
export const useHttpStore = defineStore('http', {
  // estado donde se guarda la instancia axios
  state: () => ({
    apiClient: axios.create({
      baseURL: 'https://soporte-tecnico-app.onrender.com',
      withCredentials: false,
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    })
  }),
})
