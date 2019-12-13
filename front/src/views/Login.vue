<template>
<div>
            <div class="notification is-danger container" v-show="errMsg">
                {{ errMsg }}
            </div>
   <div class="app container">


       <figure class="image">
        <img src="@/assets/logo.png" class="logo" draggable="false">
       </figure>

        <form class="forms" @submit.prevent="onSubmit(id, password)">
       <div class="field">
        <div class="control">
            <p class="control has-icons-left">
            <input type="text" class="input" placeholder="아이디를 입력하세요" v-model="id">
            <span class="icon is-small is-left">
                <i data-feather="user"></i>
            </span>
            </p>
        </div>
       </div>

       <div class="field">
        <div class="control">
            <p class="control has-icons-left">
            <input type="password" class="input" placeholder="암호를 입력하세요" v-model="password">
            <span class="icon is-small is-left">
                <i data-feather="key"></i>
            </span>
            </p>
        </div>
       </div>

        <footer>
        <div class="login-button">
            <button class="button is-link is-outlined" v-bind:class="isLoading">
            <span class="icon">
                <i data-feather="log-in"></i>
            </span>
            <span>로그인</span>
            </button>
        </div>
        <div class="register">
            <router-link to="/register">회원가입하기</router-link>
        </div>
        </footer>
        </form>
   </div>
</div>
</template>

<script>
import axios from 'axios'
const feather = require('feather-icons')
export default {
    mounted() {
        this.$nextTick(() => {
            feather.replace()
        })
    },
    data: function () {
        return {
            id: "", password: "", isLoading: "", errMsg: ""
        }
    },
    methods: {
        onSubmit(id, password) {
            this.isLoading = "is-loading"
            this.errMsg = ""
            let d = {loginID: id, password: password}
            axios.post("https://fespay.aligo.space/api/v1/login", d).then((response) => {
                localStorage.setItem("token", response.data.token)
                this.$router.push({path: "about"})
            }).catch((error) => {
                // if (error.response.data.errorCode === -100) {
                   this.errMsg = "아이디 또는 암호가 올바르지 않습니다" + error 
                // } else {
                //    this.errMsg = "지금 로그인을 할 수 없습니다."
                // }
            }).finally(() => {
                this.isLoading = ""
            })
        }
    }   
}
</script>

<style scoped>
    .app {
        margin: 8vw 10vw 10vw 10vw;
    }
    
    .logo {
        margin-top: 10vh;
        margin-bottom: 10vh;
        display: block;
        margin-left: auto;
        min-width: 30vw;
        max-width: 500px;
        margin-right: auto;
    }

    input {
        max-width: 500px;
        min-width: 10vw;
    }

    p .svg {
        margin: auto 0;
    }

@media only screen and (min-width: 768px) {
    .forms {
        display: block;
        margin-left: 20vw;
    }
}
    
    .register {
        margin-top: 2vh;
        font-size: 15px;
        text-align: center;
        display: block;
        position: absolute;
    }
    
    button {
        float: right;
    }

    .login-button {
        display: block;
        position: absolute;
    }

    @media only screen and (min-width: 768px) {
        button {
            margin-right: 40vh;
        }
    }

    .notification {
        border-radius: 0;
    }

    footer {
        position: relative;
    }
</style>