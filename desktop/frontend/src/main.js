import { DoLogin, GetLoginPageHTML } from '../wailsjs/go/main/App';

async function init() {
    const html = await GetLoginPageHTML("");
    document.querySelector('#app').innerHTML = html;
}

window.handleLogin = async () => {
    const u = document.getElementById("username").value;
    const p = document.getElementById("password").value;

    const res = await DoLogin(u, p);

    console.log("Sistem cavabı:", res);

    if (res.success) {
        console.log("Xoş gəldin:", res.user);
        console.log("Token:", res.token);

        // Uğurlu girişdən sonra Dashboard-u yükləyirik
        // loadDashboard(res.user);
        document.querySelector('#app').innerHTML = `<h1 class="p-10 text-orange-600 font-bold">Xoş gəldin, ${res.user}! Dashboard yüklənir...</h1>`;
    } else {
        const html = await GetLoginPageHTML(res.message);
        document.querySelector('#app').innerHTML = html;
    }
};

init();