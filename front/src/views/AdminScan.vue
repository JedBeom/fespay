<template>
<section>
    <div class="notification is-danger navbar is-fixed-bottom" v-show="errMsg">
        {{errMsg}}
    </div>
    <div class="app">
    <h2 class="title is-2">
        <span class="icon is-large"><i data-feather="dollar-sign"></i></span> 바코드 스캔
        </h2>
    <form autocomplete="off" @submit.prevent="onSubmit(code)" class="field">
        <label for="">
            <input class="input" type="text" placeholder="SXXXX" v-model="code"
            maxlength="5">
        </label>
    </form>
    <button class="button is-outlined" @click="back">돌아가기</button>
    <button class="button is-outlined is-success" @click="onSubmit(code)">충전하기</button>
    </div>
</section>
</template>

<script>
import api from '@/common/api.service'
import isNotValidAmount from '@/common/filterAmount'

export default {
    name: "AdminScan",
    data() {
        return {code: "", errMsg: ""}
    },
    methods: {
        onSubmit(code) {
            if (code === "") {
                return
            }

            let amount = this.$route.params.amount
            if (isNotValidAmount(amount)) {
                alert(isNotValidAmount(amount))
                this.$router.push({name:"AdminCharge"})
                return
            }
            amount = parseInt(amount)

            let d = {type: 1, cardCode: code, amount: amount}
            api.post("records", d).then((r) => {
                this.$router.push({name: "AdminChargeDetail", params: {id: r.data.id}})
            }).catch(() => {
                this.errMsg = "문제가 발생했습니다"
            })
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