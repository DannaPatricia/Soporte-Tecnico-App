
// se exporta la funcion que obtienen como parametro la api pasada desde el componenten vue
export default function (apiClient) {
  return {
    // se pasa por parametro el rol y el id del usuario para poder obtneerlo en el backend
    getCases(userRole, userId) {
      return apiClient.get('/cases', {
        params: {
          userId,
          userRole
        }
      })
    },


    getCasesByCategory(categoryId, userId, userRole) {
      return apiClient.get('/casesByCategory', {
        params: {
          categoryId,
          userId,
          userRole
        }
      })
    },

    // se obtiene la incidencia por id
    getCaseById(caseId) {
      return apiClient.get(`/case?id=${caseId}`)
    },

    createCase(newCase) {
      return apiClient.post("/case/create", {
        userId: newCase.userId,
        title: newCase.title,
        description: newCase.description,
        categoryId: newCase.categoryId
      })
    },

    setStatus(id, newStatus) {
      return apiClient.post("/case/set-status", {
        id: id,
        status: newStatus
      })
    },
  }
}
