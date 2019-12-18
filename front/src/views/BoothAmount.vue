<template>
<section>
<div class="notification is-danger navbar is-fixed-bottom" v-show="errMsg">
    {{ errMsg }}
</div>
<div class="app">
    <h2 class="title is-2">
        <span class="icon is-large"><i data-feather="dollar-sign"></i></span> 결제할 금액
        </h2>
    <form autocomplete="off" @submit.prevent="onSubmit(amount)" class="field">
        <label for="">
            <input class="input" type="text" placeholder="00000" v-model="amount"
            pattern="\d*" inputmode="numberic" maxlength="5">
        </label>
        <p>{{hangul}}</p>
    </form>
        <button class="button is-outlined" @click="back">돌아가기</button>
        <button class="button is-outlined is-success" @click="onSubmit(amount)">다음으로</button>
</div>
</section>
</template>

<script>
import numToKor from '@/common/numToKor'
import isAmountValid from '@/common/filterAmount'
const feather = require("feather-icons")

export default {
    name: "BoothAmount",
    data: function() {
        return {amount: "", errMsg: ""}
    },
    mounted() {
        this.$nextTick(() => feather.replace())
    },
    computed: {
        hangul() {
            let amount = this.amount.replace(/\D/g, '').replace(/^0+/, '')
            let isValid = isAmountValid(amount)
            if (isValid === "") {
                return numToKor(amount) + " 코인"
            }

            return isValid
        },
        amountValid() {
            return isAmountValid(this.amount)
        }
    },
    methods: {
        onSubmit(amount) {
            this.errMsg = ""
            let valid = isAmountValid(amount)
            if (valid) {
                this.errMsg = "금액 기준을 맞춰 주십시오"
                return
            }

            this.$router.push({name: "BoothScan", params: {amount: amount}})
        },
        back() {
            this.$router.push({name:"home"})
        }
    }
}
</script>

<style scoped>
.app {
    margin: 8vw;
    padding-bottom: 8vw;
}

@media only screen and (min-width: 768px) {
    .app {
        margin: 8vw 30vw 8vw 30vw;
    }
}

.notification {
    border-radius: 0;
    width: 100%;
    margin-bottom: 0;
}

form {
    box-sizing: border-box;
    display: block;
}

input {
    min-width: 30vw;
    max-width: 600px;
    width: 100%;
}

button {
    margin-top: 2vw;
}

.is-success {
    float: right;
}

p {
    text-align: center;
}
</style>