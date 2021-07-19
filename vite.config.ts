import { defineConfig } from 'vite';
import reactRefresh from '@vitejs/plugin-react-refresh';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [reactRefresh()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:7194',
        // target: 'https://analytics-production.up.railway.app',
        changeOrigin: true,
      },
      '/script.js': {
        target: 'http://localhost:7194',
      },
    },
  },
});
