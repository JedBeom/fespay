<template>
<section>
<div class="notification is-danger navbar is-fixed-bottom" v-show="errMsg">
    {{ errMsg }}
</div>
<div class="app">
    <h1 class="title is-2">충전 완료</h1>
    <div class="result">
    <RecordDetail :recordID="recordID" displayCoin=true />
    </div>
    <button class="button is-outlined is-danger" :disabled="cancelDisabled" @click="cancel">{{ cancelMsg }}</button>
    <button class="button is-outlined is-success" @click="complete">결제 완료하기</button>
</div>
</section>
</template>

<script>
import RecordDetail from '@/components/RecordDetail.vue'
import api from '@/common/api.service'

export default {
    name: "AdminChargeDetail",
    components: {
        RecordDetail
    },
    computed: {
        recordID() {
            return this.$route.params.id
        },
        cancelDisabled() {
            if (this.cancelMsg !== "충전 취소하기") {
                return true
            }

            return false
        }
    },
    data: function() {
        return {cancelMsg: "충전 취소하기", errMsg: ""}
    },
    methods: {
        cancel() {
            let loader = this.$loading.show()
            api.delete("records", this.recordID).then(() => {
                this.cancelMsg = "충전 취소됨"
            }).catch((err) => {
                if (err.response.data.errorCode == -310) {
                    this.errMsg = "이미 취소된 충전입니다"
                    this.cancelMsg = "충전 취소됨"
                    return
                }
                this.errMsg = "충전 취소하는 도중 문제가 발생했습니다."
                this.cancelMsg = "문제 발생"
            }).finally(() => loader.hide())
        },
        complete() {
            this.$router.push({name: "AdminCharge"})
        }
    }
}
</script>

<style scoped>

.app {
    margin: 8vw;
    padding-bottom: 8vw;
}

.notification {
    width: 100%;
    border-radius: 0;
    margin-bottom: 0;
}

@media only screen and (min-width: 768px) {
    .app {
        margin: 8vw 30vw 8vw 30vw;
    }
}

.result {
    padding: 30px;
    border: 1px solid rgba(0, 0, 0, 0.658);
    box-sizing: border-box;
    margin-bottom: 30px;
}

svg {
    width: 2.5rem;
    height: auto;
}

span {
    text-align: center;
}

.is-success {
    float: right;
}
</style>