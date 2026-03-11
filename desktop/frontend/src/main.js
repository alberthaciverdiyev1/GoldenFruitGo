import {
    DoLogin,
    GetCustomerDetails,
    GetCustomerForm,
    GetCustomerList,
    GetLoginPageHTML,
    GetProductListHTML,
    ProductForm,
    GetSaleForm,
    GetDashboard,
    GetSaleList,
    SetToken, GetPurchaseForm, GetPurchaseList,
    CreateCustomer,
    UpdateCustomer
} from '../wailsjs/go/handler/App';

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

        await window.loadDashboard()
    } else {
        document.querySelector('#app').innerHTML = await GetLoginPageHTML(res.message);
    }
};



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

window.openProductForm = async (id = 0) => {
    try {
        document.querySelector('#app').innerHTML = await ProductForm(id);
    } catch (err) {
        console.error("Məhsul yaradarken xəta:", err);
    }
}

window.loadCustomers = async () => {
    try {
        document.querySelector('#app').innerHTML = await GetCustomerList();
    } catch (err) {
        console.error("Məhsullar yüklənərkən xəta:", err);
    }
};

window.openCustomerDetails = async (id) => {
    try {
        document.querySelector('#app').innerHTML = await GetCustomerDetails(id);
    } catch (err) {
        console.error("Musteri yadaraken xeta:", err);
    }
}

window.saveCustomer = async () => {
    try {
        const errorDiv = document.getElementById("error-msg");
        const form = document.getElementById("customer-form");
        const formData = new FormData(form);
        const data = Object.fromEntries(formData.entries());

        data.ID = data.id ? parseInt(data.id) : 0;
        delete data.id;

        if (!data.name || data.name.trim().length < 3) {
            errorDiv.innerText = "Ad 2 hərfdən kiçik ola bilməz";
            return;
        }
        if (!data.surname || data.surname.trim().length < 3) {
            errorDiv.innerText = "Soyad 2 hərfdən kiçik ola bilməz";
            return;
        }
        if (!data.phone || data.phone.trim().length < 3) {
            errorDiv.innerText = "Telefon 2 hərfdən kiçik ola bilməz";
            return;
        }
        if (!data.address || data.address.trim().length < 3) {
            errorDiv.innerText = "Ünvan 2 hərfdən kiçik ola bilməz";
            return;
        }

        errorDiv.innerText = "";

        let result;
        if (data.ID > 0) {
            result = await UpdateCustomer(data);
        } else {
            result = await CreateCustomer(data);
        }

        if (result && result.toLowerCase().includes("xəta")) {
            errorDiv.innerText = result;
        } else {
            window.loadCustomers();
        }
    } catch (err) {
        console.error("Müştəri yadda saxlanılarkən sistem xətası:", err);
        document.getElementById("error-msg").innerText = "Sistem xətası baş verdi.";
    }
}

window.openCustomerForm = async (id = 0) => {
    try {
        const app = document.querySelector('#app');
        app.innerHTML = '<div class="text-center p-10 text-slate-400">Yüklənir...</div>';

        const formHtml = await GetCustomerForm(id);
        app.innerHTML = formHtml;
    } catch (err) {
        console.error("Form açılarkən xəta:", err);
    }
}

window.openSaleForm = async (id = 0) => {
    try {
        document.querySelector('#app').innerHTML = await GetSaleForm(id);
    } catch (err) {
        console.error("Satis modalinda xeta:", err);
    }
}

window.loadSales = async () => {
    try {
        document.querySelector('#app').innerHTML = await GetSaleList();
    } catch (err) {
        console.error("Satis listi yuklenerken xeta:", err);
    }
}


window.openPurchaseForm = async (id = 0) => {
    try {
        document.querySelector('#app').innerHTML = await GetPurchaseForm(id);
    } catch (err) {
        console.error("Satis modalinda xeta:", err);
    }
}

window.loadPurchases = async () => {
    try {
        document.querySelector('#app').innerHTML = await GetPurchaseList();
    } catch (err) {
        console.error("Satis listi yuklenerken xeta:", err);
    }
}


window.loadDashboard = async () => {
    try {
        document.querySelector('#app').innerHTML = await GetDashboard();
    } catch (err) {
        console.error("Dashboard yuklenerken xeta:", err);
    }
}




window.switchTab = function (tabName) {
    document.querySelectorAll('.tab-panel').forEach(p => p.classList.add('hidden'));
    const panel = document.getElementById('panel-' + tabName);
    if (panel) panel.classList.remove('hidden');

    document.querySelectorAll('[id^="tab-"]').forEach(b => {
        b.classList.remove('bg-white', 'text-slate-900', 'shadow-sm');
        b.classList.add('text-slate-400');
    });
    const btn = document.getElementById('tab-' + tabName);
    if (btn) {
        btn.classList.add('bg-white', 'text-slate-900', 'shadow-sm');
        btn.classList.remove('text-slate-400');
    }
}

init();