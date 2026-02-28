import {DoLogin, GetLoginPageHTML, GetProductListHTML, SetToken} from '../wailsjs/go/main/App';

async function init() {
    const savedToken = localStorage.getItem("gf_token");
    const savedUser = localStorage.getItem("gf_user");

    if (savedToken && savedUser) {
        console.log("Token tapıldı, avtomatik giriş edilir...");

        await SetToken(savedToken);

        loadDashboard(savedUser);
    } else {
        document.querySelector('#app').innerHTML = await GetLoginPageHTML("");
    }
}

window.handleLogin = async () => {
    const u = document.getElementById("username").value;
    const p = document.getElementById("password").value;

    const res = await DoLogin(u, p);

    if (res.success) {
        localStorage.setItem("gf_token", res.token);
        localStorage.setItem("gf_user", res.user);

        loadDashboard(res.user);
    } else {
        document.querySelector('#app').innerHTML = await GetLoginPageHTML(res.message);
    }
};

function loadDashboard(userName) {
    document.querySelector('#app').innerHTML = `
        <div class="p-10 text-center animate-fade-in">
            <h1 class="text-2xl font-bold text-orange-600 tracking-tight">Xoş gəldin, ${userName}!</h1>
            <p class="text-slate-500 mt-2">Golden Fruit Dashboard hazırlanır...</p>
            <button onclick="logout()" class="mt-6 text-xs text-red-500 border border-red-200 px-3 py-1 rounded-lg hover:bg-red-50">Çıxış et</button>
        </div>
    `;
}

window.logout = () => {
    localStorage.removeItem("gf_token");
    localStorage.removeItem("gf_user");
    location.reload();
};

window.loadProducts = async () => {
    try {
        document.querySelector('#app').innerHTML = await GetProductListHTML();
    } catch (err) {
        console.error("Məhsullar yüklənərkən xəta:", err);
    }
};

init();