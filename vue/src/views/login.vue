<template>
  <el-main class="pes-login-main">
    <div class="pes-login-wrapper">
      <div class="pes-login-panel">
        <div class="login-header">
          <img src="https://www.pescms.com/Theme/assets/i/project/9.png" alt="PESCMS RENT Logo" class="login-logo">
          <h1>PESCMS RENT 房租管理系统开源版</h1>
        </div>
        <el-form :model="loginForm" :rules="loginRules" ref="loginForm" label-width="80px" class="demo-ruleForm"
          @keydown.enter="login('loginForm')" label-position="top">
          <el-form-item label="账户" prop="username">
            <el-input v-model="loginForm.username" placeholder="请输入账户"></el-input>
          </el-form-item>


          <el-form-item label="密码" prop="password">
            <el-input type="password" v-model="loginForm.password" placeholder="请输入密码" auto-complete="off"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="login('loginForm')" class="login-button">登录</el-button>
          </el-form-item>

          <div style="display: flex;justify-content: space-between;">
            <span>
              <el-link href="/findpw" type="danger" :underline="false" class="am-margin-left-xs">忘记密码？</el-link>
            </span>
          </div>

        </el-form>
      </div>
    </div>
  </el-main>
</template>

<script>


import { ElMessage } from 'element-plus';


export default {
  data() {
    return {
      loginForm: {
        username: '',
        password: '',
      },
      captchaImage: '',
      loginRules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ],
      }
    }
  },

  mounted() {

  },

  methods: {
    login(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {

          let route = this.$router

          this.$axios.post('/home/login/login', this.loginForm).then(function (result) {
            let msg = result?.data?.msg || '登录异常，未知错误'
            ElMessage({
              showClose: true,
              message: msg,
              type: result.data.code == 0 ? 'success' : 'error',
            })
            if (result.data.code == 0) {
              // console.dir(result.data.data.token)
              localStorage.setItem('token', result.data.data.token)
              route.push({ path: '/house/list' })
            }
          }).catch(function (error) {
            let response = error.response
            let msg = error.response?.data?.msg || "请求出错了，可能服务器问题也可能本地网络不好"
            if (response.status == 400) {
              msg.forEach(element => {
                // console.dir(element)
                let dialogMsg = element.Field == '协议' ? '请勾选同意相关协议' : element.Field + ' ' + element.Msg
                ElMessage({
                  showClose: true,
                  message: dialogMsg,
                  type: 'error',
                })

              });
            } else {
              ElMessage({
                showClose: true,
                message: msg,
                type: 'error',
              })
            }
          })

          // this.$router.push({ path: '/' })
        } else {
          // ElMessage({
          //   showClose: true,
          //   message: '请确认所有字段填写完整',
          //   type: 'error',
          // });
          return false;
        }
      })
    },
  }
}

</script>
