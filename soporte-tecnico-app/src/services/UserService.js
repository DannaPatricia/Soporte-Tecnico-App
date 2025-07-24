
// se exporta la funcion que obtienen como parametro la api pasada desde el componenten vue
export default function (apiClient) {
  return {
    login(username, password) {
      // usar post para darle los datos a go
      return apiClient.post("/user/login", {
        username: username,
        password: password,
      })
    },
  }
}
