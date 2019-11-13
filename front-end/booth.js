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
    this.amount = this.amount.replace(/,/g, '').replace(/\./g, '').replace(/^0+/, '');
    let num = this.amount
    if (num === "0" || num === "" || num === 0 || isNaN(num) || parseInt(num) <= 0) {
        return "금액을 입력하세요"
    }
    if (num > 50000) {
        return "50000 왕운 코인을 초과할 수 없습니다"
    }
    if (num % 100 !== 0) {
        return "100원 단위로 입력해야 합니다"
    }
    this.canContinue = true;
    return viewKorean(num) + " 왕운 코인"
}

let TypeAmount = {
    template: `<div class="enter_amount">
    <h2>결제할 금액</h2>
    <form autocomplete="off">
        <label>
            <input type="text" v-model="amount" placeholder="00000" pattern="\\d*" value="" autofocus
                    v-bind:class="[borderColor, textColor]" maxlength="5">
            왕운 코인
        </label>
    </form>
    <p class="hangulAmount" v-bind:class="textColor">{{ filterAmount }}</p>
    <div class="amount_next">
        <button v-if="canContinue" class="bottom-right" v-on:click="this.next">다음</button>
    </div>
</div>`,
    data: function () {
        return { amount: "", canContinue: false }
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
            this.$emit("set-amount", this.amount)
        }
    }
}

let ScanBarcode = {
    template: ``,
    data: function() {

    },
};

let routes = {
    '/type': TypeAmount
}

let app = new Vue({
    el: '#app',
    data: {
        currentRoute: "/type",
        amount: 0,
    },
    computed: {
        ViewComponent() {
            return routes[this.currentRoute]
        }
    },
    render(h) {
        return h(this.ViewComponent)
    }
});

app.$on("set-amount", function (amount) {
    this.amount = parseInt(amount)
    alert(amount)
});