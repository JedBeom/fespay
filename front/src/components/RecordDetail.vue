<template>
<div>
<div v-if="user.name">
<p>유형: {{type}}</p>
<p>고객: {{user.name}} 님</p>
<p>금액: {{amount}}코인</p>
<p>결제 일시: {{hours}}시 {{minutes}}분 {{seconds}}초</p>
<p v-if="type === 1">잔액: {{user.coin}}</p>
</div>
<div v-else>
<h1 class="title">{{errMsg}}</h1>
</div>
</div>    
</template>

<script>
import api from '@/common/api.service'

export default {
    name: "RecordDetail",
    props: {
        recordID: String,
        displayCoin: Boolean
    },
    data: function(){
        return {errMsg: "", amount: "", user: {name: "", coin: 0}, hours: 0, minutes: 0, seconds: 0, type: ""}
    },
    created() {
        this.getRecord()
    },
    methods: {
        getRecord() {
            let loader = this.$loading.show()
            api.get("records", this.recordID).then((r) => {
                this.amount = r.data.amount
                this.user = r.data.user
                if (r.data.type === 1) {
                    this.type = "충전"
                } else {
                    this.type = "결제"
                }

                this.paidAt = new Date(r.data.paidAt);
                this.hours = this.paidAt.getHours();
                this.minutes = this.paidAt.getMinutes();
                this.seconds = this.paidAt.getSeconds();
            }).catch((err) => {
                if (err.response.status === 404) {
                    this.errMsg = "해당 결제 내역을 찾을 수 없습니다"
                } else {
                    this.errMsg = "결제 내역을 불러오는 중 문제가 발생했습니다"
                }
            }).finally(() => {
                loader.hide()
            })
        }
    },
}
</script>