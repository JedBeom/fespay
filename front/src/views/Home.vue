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
    <template v-if="me.booth.coin === 0">
      <h2 class="title is-4">우리 부스의 코인: 제로...</h2>
    </template>
    <template v-else>
      <h2 class="title is-4">{{ me.booth.coin }}코인 벌었다...!</h2>
    </template>
    <router-link :to="{name:'AdminCharge'}" v-if="me.boothID === 'fespay'"><button class="button is-link">
      <span class="icon">
          <i data-feather="zap"></i>
      </span>
      충전화면</button></router-link>
    <router-link :to="{name:'BoothAmount'}"><button class="button is-primary">
      <span class="icon">
          <i data-feather="dollar-sign"></i>
      </span>
      결제화면 바로가기</button></router-link>
    <button @click="logout" class="button is-warning">
      <span class="icon">
          <i data-feather="log-out"></i>
      </span>
      로그아웃</button>
    <article class="message is-info">
    <div class="message-header">
      <p>궁금한 게 있거나 문제가 발생했나요?</p>
    </div>
    <div class="message-body">
      1층 보건실의 페스페이 부스로 문의해주시기 바랍니다.  
    </div>
  </article>
    <list :records="records" v-if="!error"></list>
    <h2 class="title is-5" v-else>결제 내역이 없어요 ㅍ_ㅍ</h2>
    </template>
  </div>
</template>

<script>
import api from '@/common/api.service'
import list from '@/components/RecordList.vue'
import recordTime from '@/common/recordTime'
const feather = require('feather-icons')

export default {
  name: 'home',
  components: {
    list
  },
  mounted() {
      this.$nextTick(() => {
          feather.replace()
      })
  },
  updated() {
    feather.replace()
  },
  beforeCreate() {
    api.get("user").then((response) => {
      this.me = response.data
      if (this.me.boothID === 'fespay') {
        localStorage.setItem("is_admin", "true")
      }
      api.get(`booths/${this.me.boothID}/records`).then((r) => {
        this.records = r.data.reverse()
        for (let i in this.records = r.data) {
          this.records[i] = recordTime(this.records[i])
          if (this.records[i].type === 1) {
              this.records[i].type = "충전"
          } else {
              this.records[i].type = "결제"
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
        localStorage.setItem("is_admin", "")
        window.location.href = "/login"
      })
    },
  }
}
</script>

<style scoped>
.button {
  margin-right: 0.5vw;
}

.app {
  margin: 8vw;
  padding-bottom: 8vw;
}
    
.logo {
  margin-bottom: 1vh;
  margin-left: auto;
  min-width: 30vw;
  max-width: 500px;
}

@media only screen and (min-width: 768px) {
  .app {
    margin: 8vw 30vw 8vw 30vw;
  }
}

article {
  margin-top: 1vw;
}

.feather {
  margin-right: 0.3vw;
}
</style>