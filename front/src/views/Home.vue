<template>
  <div class="content app">
    <figure class="image">
      <a href="/"><img src="@/assets/logo.png" class="logo" draggable="false"></a>
    </figure>
    <h1 class="title is-2">{{ me.name }} 님 </h1>
    <h2 class="title is-3">{{ me.coin }}코인 남았습니다.</h2>
    <router-link :to="{name:'BoothAmount'}">결제화면 바로가기</router-link>
    <list :records="records" v-if="!error"></list>
  </div>
</template>

<script>
import api from '@/common/api.service'
import list from '@/components/RecordList.vue'

export default {
  name: 'home',
  components: {
    list
  },
  beforeCreate() {
    api.get("https://fespay.aligo.space/api/v1/user").then((response) => {
      this.me = response.data
      api.get(`https://fespay.aligo.space/api/v1/booths/${this.me.boothID}/records`).then((r) => {
        this.records = r.data
        for (let i in this.records = r.data) {
          this.records[i].paidAt = new Date(this.records[i].paidAt);
          this.records[i].hours = this.records[i].paidAt.getHours();
          this.records[i].minutes = this.records[i].paidAt.getMinutes();
          this.records[i].seconds = this.records[i].paidAt.getSeconds();
          if (this.records[i].hours < 10 ) {
            this.records[i].hours = "0" + this.records[i].hours
          }
          if (this.records[i].minutes < 10 ) {
            this.records[i].minutes = "0" + this.records[i].minutes
          }
          if (this.records[i].seconds < 10 ) {
            this.records[i].seconds = "0" + this.records[i].seconds
          }
        }
      }).catch((err) => {
        this.error = err
      })
    }).catch((err) => {
      this.error = err
    })
  },
  data: function () {
    return {me: {}, records: {}, error: ""}
  },
}
</script>

<style scoped>
.app {
  margin: 8vw;
  padding-bottom: 8vw;
}
    
.logo {
  margin-bottom: 5vh;
  margin-left: auto;
  min-width: 30vw;
  max-width: 500px;
}

@media only screen and (min-width: 768px) {
  .app {
    margin: 8vw 30vw 8vw 30vw;
  }
}
</style>