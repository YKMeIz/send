import { createApp } from 'vue';
import VueQrcode from '@chenfengyuan/vue-qrcode';
import App from './App.vue';
import router from './router';
import installElementPlus from './plugins/element';

const app = createApp(App);
app.component(VueQrcode.name, VueQrcode);
installElementPlus(app);
app.use(router).mount('#app');
