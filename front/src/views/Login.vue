<template>
   <div class="container">

       <p class="logo has-text-centered has-text-weight-bold">FESPAY</p>

       <div class="field">
        <div class="control">
            <p class="control has-icons-left">
            <input type="text" class="input" placeholder="ID를 입력하세요">
            <span class="icon is-small is-left">
                <i data-feather="user"></i>
            </span>
            </p>
        </div>
       </div>

       <div class="field">
        <div class="control">
            <p class="control has-icons-left">
            <input type="password" class="input" placeholder="암호를 입력하세요">
            <span class="icon is-small is-left">
                <i data-feather="key"></i>
            </span>
            </p>
        </div>
       </div>

        <button class="button is-link is-outlined">
        <span class="icon">
            <i data-feather="log-in"></i>
        </span>
        <span>로그인</span>
        </button>
        </div>

</template>

<script>
const feather = require('feather-icons')
export default {
    mounted() {
        this.$nextTick(() => {
            feather.replace()
        })
    }    
}
</script>

<style scoped>
    .logo {
        font-size: 50px;
    }

    .container {
        margin: 30px;
    }
</style>