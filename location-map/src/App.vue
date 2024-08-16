
<template>
  <div>
    <button @click="watchLocation">Monitorar</button>
    <button @click="getLocations">GET</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const socket = ref<WebSocket | null>(null)

const getLocations = () => {
  axios.get("https://192.168.68.102:8082/locations").then((response: any) => {
    console.log(response)
  })
}

const watchLocation = () => {

  event?.preventDefault()

  socket.value = new WebSocket("wss://192.168.68.102:8082/ws");

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
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
button{
  padding: 10px 20px;
}
</style>
