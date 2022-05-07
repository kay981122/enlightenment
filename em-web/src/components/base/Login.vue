<!-- filename: Login.vue -->
<template>
    <div class="index">
        <div class="backgound-image"></div>
        <transition name="el-zoom-in-center">
            <div class="box-card" v-show="show">
                <div class="login-form">
                        <div class="input-instance">
                            <el-input type="username" v-model="userInput.username" placeholder="用户名"></el-input>
                            <el-input type="password" v-model="userInput.password" placeholder="密码"></el-input>
                            <el-input class="verify" v-model="userInput.verifyCode" placeholder="验证码" ></el-input>
                        </div>
                        <div>
                            <button class="btn-login" @click="submitForm('ruleForm')">登录</button>
                        </div>
                </div>
            </div>
        </transition>
    </div>
</template>

<script>
export default {
    name: 'LoginItem',
    data() {
         return{
             show:false,
             userInput:{
                 username:'',
                 password:'',
                 verifyCode:''
             }
         }
    },
    methods:{

    },
    created() {
        setInterval(()=>{
            this.show = true;
        },1000)
    }
}
</script>

<style lang="scss" scoped>
.index {

    .backgound-image{
        background-size: cover;
        width: 100%;
        height: 100%;
        position:absolute;
        z-index:-1;
        background-repeat: no-repeat;
        background-position:center 0 ;
        background-image: url('../../assets/base/login.jpg');
    }
    .box-card {
        margin: 15% 38%;
        position: absolute;
        width: 500px;
        height: 480px;
        opacity: 0.8;
        border-radius:5%;
        background: #fff;
    }
    .login-form{
        margin: 26% 10%;
    }

    .btn-login{
        width: 100%;
        height: 45px;
        font-size: 20px;
        color: #fff;
        background-color: #409eff;
        border-color: #409eff;
        padding: 0;
    }
    .input-instance{
        .el-input{
                margin-bottom: 20px;
                height: 40px;
                font-size: 18px;
        }
        .verify{
            width: 180px;
            float: left;
        }
    }

}
</style>