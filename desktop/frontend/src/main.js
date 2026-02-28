import { DoLogin, GetLoginPageHTML, SetToken } from '../wailsjs/go/main/App';

async function init() {
    // 1. Yaddaşda token varmı yoxla?
    const savedToken = localStorage.getItem("gf_token");
    const savedUser = localStorage.getItem("gf_user");

    if (savedToken && savedUser) {
        console.log("Token tapıldı, avtomatik giriş edilir...");

        // Go tərəfindəki API Client-ə bu tokeni tanıtmalıyıq ki, sorğu ata bilsin
        await SetToken(savedToken);

        // Birbaşa Dashboard-u yükləyirik
        loadDashboard(savedUser);
    } else {
        // Token yoxdursa, Login səhifəsini göstər
        const html = await GetLoginPageHTML("");
        document.querySelector('#app').innerHTML = html;
    }
}

window.handleLogin = async () => {
    const u = document.getElementById("username").value;
    const p = document.getElementById("password").value;

    const res = await DoLogin(u, p);

    if (res.success) {
        // 2. Tokeni və İstifadəçi adını yaddaşa yazırıq
        localStorage.setItem("gf_token", res.token);
        localStorage.setItem("gf_user", res.user);

        loadDashboard(res.user);
    } else {
        const html = await GetLoginPageHTML(res.message);
        document.querySelector('#app').innerHTML = html;
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

// Çıxış funksiyası
window.logout = () => {
    localStorage.removeItem("gf_token");
    localStorage.removeItem("gf_user");
    location.reload(); // Proqramı yenilə ki, init() işə düşsün və login açılsın
};

init();