import { defineConfig } from 'vite';
import reactRefresh from '@vitejs/plugin-react-refresh';

export default defineConfig({
  plugins: [reactRefresh()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:7194',
      },
      '/script.js': {
        target: 'http://localhost:7194',
      },
    },
  },
});
