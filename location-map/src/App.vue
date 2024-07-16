
<template>
  <div>
    <h2>Conectar ao Benthos via WebSocket</h2>
    <button @click="watchLocation">Monitorar</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';


const socket = ref<WebSocket | null>(null)

const watchLocation = () => {

  event?.preventDefault()

  socket.value = new WebSocket("ws://localhost:8081/ws");

  socket.value.onopen = () => {
    console.log("WebSocket conectado");
  };

  socket.value.onmessage = (event: MessageEvent) => {
    console.log("Mensagem recebida:", event.data);
  }

  socket.value.onclose = () => {
    console.log("WebSocket fechado");
  };

  socket.value.onerror = (error: Event) => {
    console.error("Erro WebSocket:", error);
  };
  
};
</script>

<style scoped>

div{
  display: flex;
  justify-content: center;
  align-items: center;
}
button{
  padding: 10px 20px;
}
</style>
