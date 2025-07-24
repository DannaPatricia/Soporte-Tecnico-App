<script setup>
import Select from '@/components/Select.vue'
</script>

<template>
  <main>
    <section class="hero">
      <h1>Tu plataforma de confiaza para resolver incidencias.</h1>
      <h2>¿Qué problema desea resolver?</h2>
    </section>
    <Select></Select>
    <section class="start-image">
      <img src="/comunicacion.png" alt="comunicacion" />
    </section>
    
  </main>
</template>

<style>
.hero h1,
.hero h2 {
  text-align: center;
}

.hero h1 {
  font-size: 2rem;
  margin: 1em auto 0.5em auto;
  color: #1e293b;
}

.hero h2 {
  font-size: 1.3rem;
  color: #64748b;
  margin-bottom: 1.5em;
}

.start-image {
  display: flex;
  margin: 2em auto;
  height: 50%;
  width: 30%;
}

.start-image img {
  max-width: auto;
  height: auto;
}
</style>