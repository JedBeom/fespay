<template>
  <div class="about">
    <h1>This is an about page</h1>
    <Scan v-on:on-detect="onDetect" />
  </div>
</template>

<script>
import Scan from "@/components/Scan.vue"
import Axios from "axios"
export default {
  components: {
    Scan
  },
  methods: {
    onDetect: function (code) {
      if (code === "") {
        return
      }
      let d = {
        type: 2, 
        amount: 1000,
        cardCode: code
      }
      Axios.post("https://fespay.aligo.space/api/v1/records", d, {
        headers: {
          "Content-Type": "application/json;charset=UTF-8",
          "Rabbit-Fur": "db3c6b39-fe75-4485-8dd9-c4ea454091eb"
        }
      }).then((response) => {
        alert("결제에 성공했습니다."+response.data.id)
      }).catch((error) => {
        alert("ERROR: "+error)
      })
    }
  } 
}
</script>