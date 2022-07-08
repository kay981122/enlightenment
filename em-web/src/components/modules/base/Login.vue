<!-- filename: Login.vue -->
<template>
  <div class="index">
    <div class="backgound-image"></div>
    <transition name="el-zoom-in-center">
      <div class="box-card" v-show="show">
        <div class="login-form">
          <div class="input-instance">
            <el-input
              type="username"
              v-model="user.username"
              placeholder="用户名"
              :clearable="true"
            ></el-input>
            <el-input
              type="password"
              v-model="user.password"
              placeholder="密码"
              :clearable="true"
            ></el-input>
            <el-input
              class="verify"
              v-model="user.verifyCode"
              placeholder="验证码"
              :clearable="true"
            ></el-input>
          </div>
          <div>
            <button class="btn-login" @click="loginIn">登录</button>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  name: "LoginItem",
  data() {
    return {
      show: false,
      user: {
        username: "",
        password: "",
        verifyCode: "",
      },
    };
  },
  methods: {
    loginIn() {
       this.$http
        .post("/user/login",this.user)
        .then((res) => {
          if (res.code == 200) {
            localStorage.setItem("token",res.data)
            // 验证通过后，路由跳转到后台主页
            setTimeout(() => {
              this.$router.push({
                path: "/home/main",
              });
         }, 1500);
          } else {
            this.$ElMessage({ type: "error", message: res.msg });
          }
        })
        .catch((err) => {
          console.log(err);
        });

    },
  },
  created() {
    setInterval(() => {
      this.show = true;
    }, 1000);
  },
};
</script>

<style lang="scss" scoped>
.index {
  .backgound-image {
    background-size: cover;
    width: 100%;
    height: 100%;
    position: absolute;
    background-repeat: no-repeat;
    background-position: center 0;
    background-image: url("@/assets/base/login.jpg");
  }
  .box-card {
    margin: 15% 38%;
    position: absolute;
    width: 500px;
    height: 480px;
    opacity: 0.8;
    border-radius: 5%;
    background: #fff;
    z-index: 1;
  }
  .login-form {
    margin: 26% 10%;
  }

  .btn-login {
    width: 100%;
    height: 45px;
    font-size: 20px;
    color: #fff;
    background-color: #409eff;
    border-color: #409eff;
    padding: 0;
    cursor: pointer;
  }
  .input-instance {
    .el-input {
      margin-bottom: 20px;
      height: 40px;
      font-size: 18px;
    }
    .verify {
      width: 180px;
      float: left;
    }
  }
}
</style>