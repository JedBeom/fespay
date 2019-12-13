<template>
  <div class="content">
    <h1 class="title">{{ name }} 님 안녕하세요.</h1>
    <p class="subtitle">This is home</p>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'

export default {
  name: 'home',
  data: function () {
    return {name: ""}
  },
  components: {
    
  }
}
</script>
