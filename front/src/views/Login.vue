<template>
<section>
    <div class="notification is-danger" v-show="errMsg">
        {{ errMsg }}
    </div>
    <div class="app container">


        <figure class="image">
        <a href="/"><img src="@/assets/logo.png" class="logo" draggable="false"></a>
        </figure>
        <h1 class="title">로그인</h1>

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
                <router-link to="/register">계정이 없습니다</router-link>
            </div>
            </footer>
            </form>
    </div>
</section>
</template>

<script>
import axios from 'axios'
const feather = require('feather-icons')
import api from '@/common/api.service'
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
                api.setHeader()
                this.$router.push({name: "home"})
            }).catch(() => {
                this.errMsg = "아이디 또는 암호가 올바르지 않습니다"
            }).finally(() => {
                this.isLoading = ""
            })
        }
    }   
}
</script>

<style scoped>
    section {
        position: relative;
    }

    .app {
        margin: 8vw;
        position: absolute;
        padding-bottom: 8vw;
    }
    
    .logo {
        margin-bottom: 5vh;
        display: block;
        min-width: 30vw;
        max-width: 500px;
        margin-right: auto;
    }

    input {
        /*
        max-width: 500px;
        min-width: 10vw;
        */
        width: 100%;
    }

    p .svg {
        margin: auto 0;
    }

@media only screen and (min-width: 768px) {
    .app {
        margin: 8vw 30vw 8vw 30vw;
    }
}
    
    .register {
        margin-top: 2vh;
        font-size: 20px;
        display: block;
    }
    
    button {
        float: right;
    }

    .login-button {
        display: block;
    }

    .notification {
        position: absolute;
        border-radius: 0;
        width: 100%;
    }

    footer {
        position: relative;
    }
</style>