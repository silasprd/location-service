
<template>
  <div>
    <h2>Conectar ao Benthos via WebSocket</h2>
    <button @click="getLocation">Conectar ao Benthos</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';


const socket = ref<WebSocket | null>(null)

const getLocation = () => {

  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    console.log('Já conectado ao Benthos via WebSocket.');
    return;
  }

  socket.value = new WebSocket('ws://localhost:8081/ws'); 

  socket.value.onopen = () => {
    console.log('Conectado ao Benthos via WebSocket.');
  };

  socket.value.onmessage = (event) => {
    console.log('Dados recebidos do Benthos:', event.data);
  };

  socket.value.onerror = (error) => {
    console.error('Erro na conexão WebSocket com o Benthos:', error);
  };

  socket.value.onclose = () => {
    console.log('Conexão WebSocket com o Benthos fechada.');
  };
};
</script>

<style scoped>

div{
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100vw;
  height: 100vh;
}
button{
  padding: 10px 20px;
}
</style>
