<template>
  <div class="content app">
    <figure class="image">
      <a href="/"><img src="@/assets/logo.png" class="logo" draggable="false"></a>
    </figure>
    <h1 class="title is-2">{{ me.name }} 님 </h1>
    <template v-if="!me.boothID">
      <h2 class="title is-3">부스 직원이 아닙니다. 페스페이 부스(보건실)로 오셔서 부스 직원 등록을 진행해 주세요</h2>
    </template>
    <template v-else>
    <h2 class="title is-3">{{ me.booth.name }} 부스</h2>
    <router-link :to="{name:'BoothAmount'}"><button class="button is-outlined is-primary">결제화면 바로가기</button></router-link>
    <button @click="logout" class="button">로그아웃</button>
    <list :records="records" v-if="!error"></list>
    </template>
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
    api.get("user").then((response) => {
      this.me = response.data
      api.get(`booths/${this.me.boothID}/records`).then((r) => {
        this.records = r.data.reverse()
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

          if (this.records[i].canceledAt == undefined) {
            continue
          }

          this.records[i].canceledAt = new Date(this.records[i].canceledAt);
          this.records[i].canceledAt.hours = this.records[i].canceledAt.getHours();
          this.records[i].canceledAt.minutes = this.records[i].canceledAt.getMinutes();
          this.records[i].canceledAt.seconds = this.records[i].canceledAt.getSeconds();
          if (this.records[i].canceledAt.hours < 10 ) {
            this.records[i].canceledAt.hours = "0" + this.records[i].canceledAt.hours
          }
          if (this.records[i].canceledAt.minutes < 10 ) {
            this.records[i].canceledAt.minutes = "0" + this.records[i].canceledAt.minutes
          }
          if (this.records[i].canceledAt.seconds < 10 ) {
            this.records[i].canceledAt.seconds = "0" + this.records[i].canceledAt.seconds
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
  methods: {
    logout() {
      api.get(`logout`).then(() => {
        localStorage.setItem("token", "")
        window.location.href = "/login"
      })
    },
  }
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