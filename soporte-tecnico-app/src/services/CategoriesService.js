
// se exporta la funcion que obtienen como parametro la api pasada desde el componenten vue
export default function (apiClient) {
  return {
    getCategories() {
      return apiClient.get('/categories')
    },

    getCategoryById(categoryId) {
      return apiClient.get(`/category?id=${categoryId}`)
    }
  }
}