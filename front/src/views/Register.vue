<template>
<section>
    <div class="notification is-danger" v-show="errMsg">
        {{ errMsg }}
    </div>
    <div class="app container">
        <figure class="image">
        <img src="@/assets/logo.png" class="logo" draggable="false">
        </figure>

        <h1 class="title is-3">회원가입</h1>
        <div v-show="!cardCode">
            <h2 class="title is-4">본인 인증</h2>
            <p class="">학생증의 바코드를 인식해 주세요</p>
            <Scan @on-detect="onDetect" />
            <div class="register">
                <router-link to="/login">계정이 있습니다</router-link>
            </div>
        </div>


        <div v-show="cardCode">
        <h2 class="title is-4">인적 작성</h2>
        <form class="forms" @submit.prevent="onSubmit(id, password, name, number)">
        <div class="field">
            <h3 class="title is-5">개인 정보</h3>
            <div class="control"> 
                <p class="control has-icons-left">
                    <input type="text" class="input" placeholder="이름을 입력하세요" v-model="name" required>
                    <span class="icon is-small is-left">
                        <i data-feather="at-sign"></i>
                    </span>
                </p>
            </div>
        </div>

        <div class="field">
            <div class="control">
                <p class="control has-icons-left">
                    <input type="text" class="input" placeholder="학생 번호를 입력하세요 ex) 23 (교사인 경우 입력x)" v-model="number" pattern="\d*" inputmode="numberic" required>
                    <span class="icon is-small is-left">
                        <i data-feather="hash"></i>
                    </span>
                </p>
            </div>
        </div>
        <h3 class="title is-5">가입 정보</h3>
        <div class="field">
            <div class="control">
                <p class="control has-icons-left">
                <input type="text" class="input" placeholder="아이디를 입력하세요" v-model="id" required>
                <span class="icon is-small is-left">
                    <i data-feather="user"></i>
                </span>
                </p>
            </div>
        </div>

        <div class="field">
            <div class="control">
                <p class="control has-icons-left">
                <input type="password" class="input" placeholder="암호를 입력하세요" v-model="password" required>
                <span class="icon is-small is-left">
                    <i data-feather="key"></i>
                </span>
                </p>
            </div>
        </div>

        <div class="field">
            <div class="control">
                <p class="control has-icons-left">
                <input type="password" class="input" placeholder="암호를 다시 입력하세요" v-model="passwordRetry" required>
                <span class="icon is-small is-left">
                    <i data-feather="key"></i>
                </span>
                </p>
            </div>
        </div>
        
        <p>{{isPasswordValid}}</p>

            <footer>
            <div class="login-button">
                <button class="button is-link is-outlined" v-bind:class="isLoading">
                <span class="icon">
                    <i data-feather="log-in"></i>
                </span>
                <span>회원가입</span>
                </button>
            </div>
            </footer>
            </form>
        </div>
    </div>
</section>
</template>

<script>
import axios from 'axios'
import Scan from '@/components/Scan.vue'
const feather = require('feather-icons')
export default {
    mounted() {
        this.$nextTick(() => {
            feather.replace()
        })
    },
    components: {
        Scan
    },
    data: function () {
        return {
            cardCode: "", name: "", number: "", id: "", password: "", passwordRetry: "", isLoading: "", errMsg: ""
        }
    },
    computed: {
        isPasswordValid() {
            if (this.password === "") {
                return ""
            }

            if (this.password.length < 8) {
                return "암호는 8자 이상이어야 합니다"
            }

            if (this.passwordRetry !== "" && this.password !== this.passwordRetry) {
                return "암호가 서로 일치하지 않습니다"
            }

            return ""
        }
    },
    methods: {
        onDetect: function(code) {
            this.checkAvailable(code).then((isAvailable) => {
                if (isAvailable) {
                    this.cardCode = code
                } else {
                    this.errMsg = "존재하지 않거나 이미 가입한 바코드입니다"
                }
            })
        },
        onSubmit(id, password, name, number) {
            this.isLoading = "is-loading"
            this.errMsg = ""
            if (number === "") {
                number = 0
            } else {
                number = parseInt(number)
            }

            let d = {loginID: id, password: password, name: name, number: number, cardCode: this.cardCode}
            alert(id, password, name, number, this.cardCode)
            axios.patch("https://fespay.aligo.space/api/v1/register", d).then(() => {
                this.$router.push({name: "login"})
            }).catch(() => {
                this.errMsg = "개인 정보가 일치하지 않거나 아이디가 고유하지 않습니다"
            }).finally(() => {
                this.isLoading = ""
            })
        },
        async checkAvailable(code) {
            return await axios.get(`https://fespay.aligo.space/api/v1/register/available?code=${code}`).then((response) => {
                return response.data.isAvailable
            }).catch(() => {
                return false
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
        padding-bottom: 8vw;
        position: absolute;
    }
    
    .logo {
        margin-bottom: 5vh;
        display: block;
        margin-left: auto;
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