import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import fs from 'fs';
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    https: {
      key: fs.readFileSync(path.resolve(__dirname, './src/ssl/locationSSL.key')),
      cert: fs.readFileSync(path.resolve(__dirname, './src/ssl/locationSSL.crt')),
    },
    host: '192.168.68.102',
    port: 4000
  }
})
