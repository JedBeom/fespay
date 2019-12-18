<template>
<div class="app">
    <div class="notification is-danger navbar is-fixed-bottom" v-show="errMsg">
        {{errMsg}}
    </div>
    <h1 class="title is-3">{{scanMsg}}</h1>
    <button @click="back" class="button is-outlined is-primary">뒤로 가기</button>
    <Scan v-on:on-detect="onDetect" :active="scanActive" />
</div>
</template>

<script>
import Api from '@/common/api.service'
import Scan from '@/components/Scan.vue'
import isNotValidAmount from '@/common/filterAmount'
export default {
    name: "BoothScan",
    components: {
        Scan
    },
    data: function () {
        return {errMsg: "", scanMsg: "바코드 스캔 중", scanActive: true}
    },
    methods: {
        onDetect(code) {
            let amount = this.$route.params.amount
            if (isNotValidAmount(amount)) {
                alert(isNotValidAmount(amount))
                this.$router.push({name:"BoothAmount"})
                return
            }
            amount = parseInt(amount)
            let loader = this.$loading.show();
            this.scanMsg = "처리 중"
            Api.post("records", {
                type: 2,
                amount: amount,
                cardCode: code
            }).then((r) => {
                this.$router.push({name: "BoothRecordDetail", params: {id: r.data.id}})
            }).catch((err) => {
                if (err.response.data.errorCode === -103) {
                    this.errMsg = "돈이 충분하지 않습니다"
                } else {
                    this.errMsg = "알 수 없는 문제가 발생했습니다"
                }
                this.scanMsg = "결제에 실패했습니다"
            }).finally(() => {

                loader.hide()
            })            
        },
        back() {
            this.$router.push({name:"BoothAmount"})
        }
    },
}
</script>

<style scoped>
.app {
    margin: 8vw;
}

.title {
    font-weight: 900;
    text-shadow: 0px 0px 10px white;
}

.notification {
    margin-bottom: 0;
    width: 100%;
    border-radius: 0;
}
 
@media only screen and (min-width: 768px) {
    video {
        width: 100%;
    }

    .app {
        margin: 8vw 30vw 8vw 30vw;
    }
}
</style>

<style>
#video {
    top: 0;
    left: 0;
    z-index: -100;
    position: fixed;
    max-width: none;
}
</style>
