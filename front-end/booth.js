function viewKorean(num) {
    let hanA = ["", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구", "십"];
    let danA = ["", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천", "", "십", "백", "천"];
    let result = "";
    for (let i = 0; i < num.length; i++) {
        let str = "";
        let han = hanA[num.charAt(num.length - (i + 1))];
        if (han !== "") str += han + danA[i];
        if (i === 4) str += "만";
        result = str + result;
    }
    return result;
}

function filterAmount() {
    this.canContinue = false;
    this.amount = this.amount.replace(/\D/g, '').replace(/^0+/, '');
    let num = this.amount;
    if (num === "0" || num === "" || num === 0 || isNaN(num) || parseInt(num) <= 0) {
        return "금액을 입력하세요"
    }
    if (num > 50000) {
        return "50000 코인을 초과할 수 없습니다"
    }
    if (num % 100 !== 0) {
        return "100 코인 단위로 입력해야 합니다"
    }
    this.canContinue = true;
    return viewKorean(num) + " 코인"
}

Vue.component('type-amount', {
    template: `<div class="booth_box enter_amount app-margin">
    <h2>결제할 금액</h2>
    <form autocomplete="off" v-on:submit.prevent>
        <label>
            <input type="text" v-model="amount" placeholder="00000" pattern="\\d*" value="" autofocus
                    v-bind:class="[borderColor, textColor]" maxlength="5">
            코인
        </label>
    </form>
    <p class="hangulAmount" v-bind:class="textColor">{{ filterAmount }}</p>
    <div class="amount_next">
        <button v-if="canContinue" class="bottom-right" v-on:click="this.next">다음</button>
    </div>
</div>`,
    data: function () {
        return {amount: "", canContinue: false}
    },
    computed: {
        filterAmount: filterAmount,
        borderColor: function () {
            return this.canContinue ? "green_border" : "red_border"
        },
        textColor: function () {
            return this.canContinue ? "green" : "red"
        },
    },
    methods: {
        next: function () {
            this.$emit("set-amount", this.amount);
        },
    },
});

Vue.component('scan-barcode', {
    template: `<div>
    <h2 v-on:click="scanReady()" class="scan_h2">{{ message }}</h2>
    <div id="scan">
    </div>
</div>`,
    data: function() {
        return {code: "", message: "눌러서 스캔"}
    },
    methods: {
        scanReady: function () {
            if (this.message === "로딩 중...") {
                return
            }
            let that = this;
            Quagga.init({
                inputStream: {
                    name: "Live",
                    type: "LiveStream",
                    target: document.querySelector('#scan'),
                    constraints: {
                        width: {min: 720},
                        height: {min: 1280},
                        aspectRatio: {min: 1, max: 100},
                        facingMode: "environment"
                    }
                },
                numOfWorkers: 4,
                decoder: {
                    readers: ["code_39_reader"]
                },
            }, function (err) {
                if (err) {
                    that.message = "스캔 불가입니다";
                    alert("카메라를 실행 할 수 없습니다.\n카메라가 연결되어 있거나 권한을 허용했는지 확인해주세요.");
                    that.next("");
                    return
                }
                that.message = "스캔 중...";
                Quagga.onDetected(that.next);
                Quagga.start()
            })
        },
        next: _.debounce(
            function (data) {
                Quagga.stop();
                if (data !== "") {
                    this.code = data.codeResult.code;
                }
                this.message = "로딩 중...";
                this.$emit("set-code", this.code);
            }, 500
        )
    }
});

Vue.component('record-result', {
    template: `<div class="booth_box app-margin record-result">
    <h2>{{ getRecord }} 결제 완료</h2>
    <ul v-if="record.amount !== 0">
        <li>금액: {{ record.amount }} 코인</li>
        <li>고객: {{ record.user.name }} 님</li>
        <li>일시: {{ record.hours }}시 {{ record.minutes }}분</li>
    </ul>
    <ul v-else>
        <li>로드 중입니다...</li>
    </ul>
    <div class="amount_next">
        <button class="bottom-right" v-on:click="this.toAmount">새 결제</button>
        <button class="bottom-left" v-on:click="this.cancel">결제 취소</button>
    </div>
    </div>`,
    data: function () {
        return {record: {id: "", user: {name: ""}, amount: 0}}
    },
    props: ['recordId'],
    computed: {
        getRecord: function () {
            axios.get("/api/v1/records/" + this.recordId, {
                headers: {
                    "Content-Type": "application/json;charset=UTF-8",
                    "Rabbit-Fur": "db3c6b39-fe75-4485-8dd9-c4ea454091eb"
                }
            }).then((response => {
                this.record = response.data;
                this.record.paidAt = new Date(this.record.paidAt);
                this.record.hours = this.record.paidAt.getHours();
                this.record.minutes = this.record.paidAt.getMinutes();
                return "결제 완료"
            })).catch((error) => {
                alert("결제 내역 불러오기에 실패하였습니다.");
                console.log(error);
            })
        }
    },
    methods: {
        toAmount: function () {
            this.$emit("to", "amount")
        },
        cancel: function () {
            axios.delete("/api/v1/records/" + this.recordId, {
                headers: {
                    "Content-Type": "application/json;charset=UTF-8",
                    "Rabbit-Fur": "db3c6b39-fe75-4485-8dd9-c4ea454091eb"
                }
            }).then((response => {
                alert("결제 취소 완료되었습니다");
                this.toAmount();
            })).catch((error) => {
                alert("결제 취소에 실패하였습니다. 결제 내역에서 다시 시도해주세요.");
                console.log(error);
            })
        }
    }
});

let app = new Vue({
    el: '#app',
    data: {
        action: "amount",
        amount: 0,
        code: "",
        recordID: "",
    },
    methods: {
        to: function (action) {
            this.action = action
        },
        setAmount: function (amount) {
            this.amount = parseInt(amount);
            this.action = "scan";
        },
        setCode: function (code) {
            if (code === "") {
                this.action = "amount";
                return
            }
            this.code = code;
            this.processOrder();
        },
        processOrder: function () {
            let data = {
                type: 2,
                amount: this.amount,
                cardCode: this.code,
            };
            axios.post("/api/v1/records", data, {
                headers: {
                    "Content-Type": "application/json;charset=UTF-8",
                    "Rabbit-Fur": "db3c6b39-fe75-4485-8dd9-c4ea454091eb"
                }
            }).then((response) => {
                this.recordID = response.data.id;
                this.action = "result";
            }).catch((error) => {
                    let ec = error.response.data.errorCode;
                    if (ec === -103) {
                        alert("코인이 부족합니다")
                    } else if (ec === -201) {
                        alert("동결된 고객입니다. 자세한 내용은 페스페이 부스에 문의하세요.")
                    } else if (ec === -202) {
                        alert("부스가 동결 상태입니다. 자세한 내용은 페스페이 부스에 문의하세요.")
                    } else if (ec === -305) {
                        alert("잘못된 바코드입니다.")
                    }
                    this.action = "amount";
                }
            )
        }
    }
});