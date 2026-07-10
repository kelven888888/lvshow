
/**
 * 马来西亚地址数据与联动逻辑
 * 包含主要州属、地区和邮编的映射关系
 */

// 简化的马来西亚地址数据库 (State -> Areas -> Postcodes)
// 实际生产中建议从后端API获取完整数据
const malaysiaData = {
    "Johor": {
        "Johor Bahru": ["80000", "80050", "80100", "80150", "80200", "80250", "80300", "80350", "80400", "80500", "80550", "80600", "80650", "80700", "80750", "80800", "80850", "80900", "80950", "81000", "81100", "81200", "81300", "81400", "81500", "81600", "81700", "81750", "81800", "81900"],
        "Skudai": ["81300"],
        "Ulu Tiram": ["81800"],
        "Pasir Gudang": ["81700"],
        "Kulai": ["81000"],
        "Pontian": ["82000"],
        "Muar": ["84000"],
        "Batu Pahat": ["83000"],
        "Kluang": ["86000"],
        "Segamat": ["85000"],
        "Mersing": ["86800"],
        "Kota Tinggi": ["81900"]
    },
    "Kuala Lumpur": {
        "Kuala Lumpur City Centre": ["50000", "50050", "50088", "50100", "50150", "50200", "50250", "50300", "50350", "50400", "50450", "50460", "50470", "50480", "50490", "50500", "50550", "50560", "50570", "50580", "50590", "50600", "50650", "50660", "50670", "50680", "50690", "50700", "50750", "50800", "50900", "50950"],
        "Cheras": ["56000", "56100"],
        "Kepong": ["52000"],
        "Sentul": ["51000"],
        "Setapak": ["53000"],
        "Bangsar": ["59000"],
        "Brickfields": ["50470"]
    },
    "Selangor": {
        "Shah Alam": ["40000", "40100", "40150", "40170", "40200", "40300", "40400", "40450", "40460", "40470", "40500", "40550", "40560", "40570", "40580", "40590", "40600", "40650", "40660", "40670", "40680", "40690", "40700", "40750", "40780", "40790", "40800", "40900", "40950"],
        "Petaling Jaya": ["46000", "46050", "46100", "46150", "46200", "46300", "46400", "46500", "46600", "46700", "46800", "46900"],
        "Subang Jaya": ["47500", "47600"],
        "Puchong": ["47100", "47150", "47160", "47170", "47180", "47190"],
        "Klang": ["41000", "41050", "41100", "41200", "41300", "41400", "41500", "41600", "41700", "41800", "41900"],
        "Ampang": ["68000"],
        "Kajang": ["43000"],
        "Seri Kembangan": ["43300"],
        "Rawang": ["48000"],
        "Sepang": ["43900"]
    },
    "Penang": {
        "George Town": ["10000", "10050", "10100", "10150", "10200", "10250", "10300", "10350", "10400", "10450", "10460", "10470", "10500", "10550", "10600", "10650", "10700", "10750", "10800", "10850", "10900", "10950"],
        "Bayan Lepas": ["11900", "11950"],
        "Butterworth": ["12000", "12100", "12200", "12300"],
        "Bukit Mertajam": ["14000"],
        "Nibong Tebal": ["14300"]
    },
    "Perak": {
        "Ipoh": ["30000", "30050", "30100", "30150", "30200", "30250", "30300", "30350", "30400", "30450", "30500", "30550", "30600", "30650", "30660", "30670", "30680", "30690", "30700", "30750", "30800", "30900", "30950"],
        "Taiping": ["34000"],
        "Teluk Intan": ["36000"],
        "Sitiawan": ["32000"],
        "Kampar": ["31900"]
    },
    "Sabah": {
        "Kota Kinabalu": ["88000", "88100", "88200", "88300", "88400", "88450", "88500", "88600", "88700", "88800", "88900"],
        "Sandakan": ["90000"],
        "Tawau": ["91000"],
        "Lahad Datu": ["91100"],
        "Keningau": ["89000"]
    },
    "Sarawak": {
        "Kuching": ["93000", "93050", "93100", "93150", "93200", "93250", "93300", "93350", "93400", "93450", "93500", "93550", "93560", "93570", "93580", "93590", "93600", "93650", "93660", "93670", "93680", "93690", "93700", "93750", "93800", "93900", "93950"],
        "Miri": ["98000"],
        "Sibu": ["96000"],
        "Bintulu": ["97000"]
    },
    "Melaka": {
        "Melaka City": ["75000", "75050", "75100", "75150", "75200", "75250", "75300", "75350", "75400", "75450", "75460", "75500", "75550", "75600", "75650", "75660", "75670", "75680", "75690", "75700", "75750", "75800", "75850", "75900", "75950"],
        "Alor Gajah": ["78000"],
        "Jasin": ["77000"]
    },
    "Negeri Sembilan": {
        "Seremban": ["70000", "70100", "70200", "70300", "70400", "70450", "70500", "70550", "70600", "70650", "70700", "70750", "70800", "70900"],
        "Port Dickson": ["71000"],
        "Nilai": ["71800"]
    },
    "Pahang": {
        "Kuantan": ["25000", "25050", "25100", "25150", "25200", "25250", "25300", "25350", "25400", "25450", "25500", "25550", "25600", "25650", "25660", "25670", "25680", "25690", "25700", "25750", "25800", "25900"],
        "Temerloh": ["28000"],
        "Bentong": ["28700"],
        "Cameron Highlands": ["39000"]
    },
    "Terengganu": {
        "Kuala Terengganu": ["20000", "20050", "20100", "20150", "20200", "20250", "20300", "20350", "20400", "20450", "20500", "20550", "20600", "20650", "20660", "20670", "20680", "20690", "20700", "20750", "20800", "20900"],
        "Dungun": ["23000"],
        "Kemaman": ["24000"]
    },
    "Kelantan": {
        "Kota Bharu": ["15000", "15050", "15100", "15150", "15200", "15250", "15300", "15350", "15400", "15450", "15500", "15550", "15600", "15650", "15660", "15670", "15680", "15690", "15700", "15750", "15800", "15900"],
        "Tanah Merah": ["17500"],
        "Pasir Mas": ["17000"]
    },
    "Kedah": {
        "Alor Setar": ["05000", "05050", "05100", "05150", "05200", "05250", "05300", "05350", "05400", "05450", "05460", "05500", "05550", "05600", "05650", "05660", "05670", "05680", "05690", "05700", "05750", "05800", "05900"],
        "Sungai Petani": ["08000"],
        "Kulim": ["09000"],
        "Langkawi": ["07000"]
    },
    "Perlis": {
        "Kangar": ["01000"],
        "Arau": ["02600"]
    },
    "Putrajaya": {
        "Putrajaya": ["62000", "62100", "62200", "62300", "62500", "62600", "62700", "62800", "62900"]
    },
    "Labuan": {
        "Labuan": ["87000", "87010", "87020", "87030"]
    }
};

class AddressApp {
    constructor() {
        this.stateSelect = document.getElementById('stateSelect');
        this.areaSelect = document.getElementById('areaSelect');
        this.postcodeSelect = document.getElementById('postcodeSelect');
        this.savedList = document.getElementById('savedAddressesList');
        
        this.init();
    }

    init() {
        this.populateStates();
        this.loadSavedAddresses();
    }

    // 初始化州属下拉框
    populateStates() {
        const states = Object.keys(malaysiaData).sort();
        states.forEach(state => {
            const option = document.createElement('option');
            option.value = state;
            option.textContent = state;
            this.stateSelect.appendChild(option);
        });
    }

    // 州属改变事件
    onStateChange() {
        const selectedState = this.stateSelect.value;
        
        // 重置下级选项
        this.areaSelect.innerHTML = '<option value="">请选择地区...</option>';
        this.postcodeSelect.innerHTML = '<option value="">请先选择地区</option>';
        this.postcodeSelect.disabled = true;
        this.postcodeSelect.classList.add('bg-gray-50', 'text-gray-500', 'cursor-not-allowed');
        this.postcodeSelect.classList.remove('bg-white', 'text-gray-700', 'cursor-pointer');

        if (!selectedState) {
            this.areaSelect.disabled = true;
            this.areaSelect.classList.add('bg-gray-50', 'text-gray-500', 'cursor-not-allowed');
            this.areaSelect.classList.remove('bg-white', 'text-gray-700', 'cursor-pointer');
            return;
        }

        // 填充地区
        const areas = Object.keys(malaysiaData[selectedState]).sort();
        areas.forEach(area => {
            const option = document.createElement('option');
            option.value = area;
            option.textContent = area;
            this.areaSelect.appendChild(option);
        });

        // 启用地区选择
        this.areaSelect.disabled = false;
        this.areaSelect.classList.remove('bg-gray-50', 'text-gray-500', 'cursor-not-allowed');
        this.areaSelect.classList.add('bg-white', 'text-gray-700', 'cursor-pointer');
    }

    // 地区改变事件
    onAreaChange() {
        const selectedState = this.stateSelect.value;
        const selectedArea = this.areaSelect.value;

        // 重置邮编
        this.postcodeSelect.innerHTML = '<option value="">请选择邮编...</option>';

        if (!selectedArea) {
            this.postcodeSelect.disabled = true;
            this.postcodeSelect.classList.add('bg-gray-50', 'text-gray-500', 'cursor-not-allowed');
            this.postcodeSelect.classList.remove('bg-white', 'text-gray-700', 'cursor-pointer');
            return;
        }

        // 填充邮编
        const postcodes = malaysiaData[selectedState][selectedArea];
        postcodes.forEach(code => {
            const option = document.createElement('option');
            option.value = code;
            option.textContent = code;
            this.postcodeSelect.appendChild(option);
        });

        // 启用邮编选择
        this.postcodeSelect.disabled = false;
        this.postcodeSelect.classList.remove('bg-gray-50', 'text-gray-500', 'cursor-not-allowed');
        this.postcodeSelect.classList.add('bg-white', 'text-gray-700', 'cursor-pointer');
    }

    // 保存地址
    saveAddress() {
        const name = document.getElementById('receiverName').value.trim();
        const phone = document.getElementById('receiverPhone').value.trim();
        const state = this.stateSelect.value;
        const area = this.areaSelect.value;
        const postcode = this.postcodeSelect.value;
        const detail = document.getElementById('detailAddress').value.trim();

        if (!name || !phone || !state || !area || !postcode) {
            alert('请填写所有必填字段！');
            return;
        }

        const addressObj = {
            id: Date.now(),
            name,
            phone,
            state,
            area,
            postcode,
            detail,
            timestamp: new Date().toLocaleString('zh-CN')
        };

        // 保存到 LocalStorage
        let addresses = JSON.parse(localStorage.getItem('my_addresses') || '[]');
        addresses.unshift(addressObj);
        localStorage.setItem('my_addresses', JSON.stringify(addresses));

        this.renderSavedAddresses();
        this.resetForm(false); // 保留表单内容以便用户确认，或根据需求清空
        alert('地址保存成功！');
    }

    // 渲染已保存地址列表
    renderSavedAddresses() {
        const addresses = JSON.parse(localStorage.getItem('my_addresses') || '[]');
        const container = document.getElementById('savedAddressesList');

        if (addresses.length === 0) {
            container.innerHTML = `
                <div class="text-center py-8 text-gray-400 bg-gray-50 rounded-lg border border-dashed border-gray-300">
                    <i class="fa-regular fa-folder-open text-3xl mb-2"></i>
                    <p>暂无保存的地址</p>
                </div>
            `;
            return;
        }

        container.innerHTML = addresses.map(addr => `
            <div class="bg-white p-4 rounded-lg border border-gray-200 shadow-sm hover:shadow-md transition flex justify-between items-start group">
                <div>
                    <div class="flex items-center gap-2 mb-1">
                        <span class="font-bold text-gray-800">${addr.name}</span>
                        <span class="text-xs bg-blue-100 text-blue-700 px-2 py-0.5 rounded-full">${addr.phone}</span>
                    </div>
                    <p class="text-sm text-gray-600 mb-1">
                        ${addr.detail ? addr.detail + ', ' : ''}
                        ${addr.area}, ${addr.postcode} ${addr.state}
                    </p>
                    <p class="text-xs text-gray-400">保存于: ${addr.timestamp}</p>
                </div>
                <button onclick="app.deleteAddress(${addr.id})" class="text-gray-400 hover:text-red-500 transition p-2 opacity-0 group-hover:opacity-100">
                    <i class="fa-solid fa-trash"></i>
                </button>
            </div>
        `).join('');
    }

    // 删除地址
    deleteAddress(id) {
        if(!confirm('确定要删除这个地址吗？')) return;
        
        let addresses = JSON.parse(localStorage.getItem('my_addresses') || '[]');
        addresses = addresses.filter(a => a.id !== id);
        localStorage.setItem('my_addresses', JSON.stringify(addresses));
        this.renderSavedAddresses();
    }

    // 加载已保存地址
    loadSavedAddresses() {
        this.renderSavedAddresses();
    }

    // 重置表单
    resetForm(clearInputs = true) {
        if (clearInputs) {
            document.getElementById('receiverName').value = '';
            document.getElementById('receiverPhone').value = '';
            document.getElementById('detailAddress').value = '';
        }
        
        this.stateSelect.value = '';
        this.onStateChange(); // 这会重置后续选项
    }

    // 加载测试数据
    loadDemoData() {
        document.getElementById('receiverName').value = 'Ahmad Bin Ali';
        document.getElementById('receiverPhone').value = '+60 12-345 6789';
        document.getElementById('detailAddress').value = 'No. 123, Jalan Bukit Bintang, Taman Melati';
        
        // 模拟选择 KL
        this.stateSelect.value = 'Kuala Lumpur';
        this.onStateChange();
        
        setTimeout(() => {
            this.areaSelect.value = 'Kuala Lumpur City Centre';
            this.onAreaChange();
            
            setTimeout(() => {
                this.postcodeSelect.value = '50000';
            }, 100);
        }, 100);
    }

    // 调试信息
    debugInfo() {
        console.log('Current Data Structure:', malaysiaData);
        alert('数据结构已打印到控制台 (F12 -> Console)');
    }
}

// 初始化应用
const app = new AddressApp();
