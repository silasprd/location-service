<template>
  <form action="">
    <input v-model="deviceId" placeholder="Informe o número identificador" type="text">
    <button @click="watchLocation">Monitorar Localização</button>
  </form>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const deviceId = ref()
const interval = ref()
const socket = ref<WebSocket | null>(null)
const location = ref({})

const watchLocation = () => {

  event?.preventDefault()

  if (!deviceId.value) {
    alert("Por favor, insira um ID de dispositivo.");
    return;
  }

  socket.value = new WebSocket("ws://localhost:8080/ws");

  socket.value.onopen = () => {
    console.log("Conectado ao WebSocket");
  };

  socket.value.onclose = () => {
    console.log("Desconectado do WebSocket");
  };

  if (navigator.geolocation) {
    interval.value = window.setInterval(() => {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const geolocationValue = {
            deviceId: deviceId.value,
            latitude: position.coords.latitude,
            longitude: position.coords.longitude,
            accuracy: position.coords.accuracy,
            speed: position.coords.speed,
            heading: position.coords.heading,
            altitude: position.coords.altitude,
            timestamp: position.timestamp
          }
          location.value = geolocationValue
          console.log("Location: ", geolocationValue.timestamp)
          if (socket.value) socket.value.send(JSON.stringify(location.value));
        },
        (error) => {
          console.error("Erro ao obter localização:", error);
        }
      )
    }, 10000);
  }

}

</script>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  width: 300px;
}

button {
  border: 1px solid;
}

input {
  border-radius: 10px;
  font-size: 16px;
  padding: 5px;
  margin-bottom: 10px
}
</style>
