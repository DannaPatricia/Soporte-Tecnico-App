<template>
  <section class="error-page">
    <div class="error-box">
      <h1 class="error-code">404</h1>
      <p class="error-message">Página no encontrada</p>
      <RouterLink to="/" class="error-button">Volver al inicio</RouterLink>
    </div>
  </section>
</template>

<script setup>
// No necesitas lógica aquí a menos que quieras mostrar diferentes códigos/mensajes
</script>

<style scoped>
.error-page {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f8fafc;
}

.error-box {
  text-align: center;
  padding: 2rem;
  background-color: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 1rem;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

.error-code {
  font-size: 5rem;
  color: #1e293b;
  margin-bottom: 1rem;
}

.error-message {
  font-size: 1.5rem;
  color: #334155;
  margin-bottom: 2rem;
}

.error-button {
  padding: 0.8rem 1.5rem;
  background-color: #1e293b;
  color: #fff;
  text-decoration: none;
  border-radius: 0.5rem;
  font-weight: bold;
}

.error-button:hover {
  background-color: #334155;
}
</style>
