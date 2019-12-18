<template>
  <div class="content app">
    <h1 class="title is-2">{{ myName }} 님 안녕하세요.</h1>
    <h2 class="title is-3">{{ myCoin }}코인 남았습니다.</h2>
    <router-link :to="{name:'BoothAmount'}">결제화면 바로가기</router-link>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import api from '@/common/api.service'

export default {
  name: 'home',
  beforeCreate() {
    if (!localStorage.getItem("token")) {
      this.$router.push({name: "login"})
    }
    api.get("https://fespay.aligo.space/api/v1/user").then((response) => {
      this.myName = response.data.name
      this.myCoin = response.data.coin
    }).catch((err) => {
      this.myName = err
    })
  },
  data: function () {
    return {myName: "", myCoin: 0}
  },
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
</style>