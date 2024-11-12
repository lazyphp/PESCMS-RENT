<template>
    <el-main class="pes-login-main">
        <div class="pes-login-wrapper">
            <div class="pes-login-panel">
                <div class="login-header">
                    <img src="https://www.pescms.com/Theme/assets/i/project/9.png" alt="PESCMS RENT Logo"
                        class="login-logo">
                    <h1>找回密码</h1>
                </div>
                <el-form :model="loginForm" :rules="loginRules" ref="loginForm" label-width="80px" class="demo-ruleForm"
                    @keydown.enter="login('loginForm')" label-position="top">
                    <el-form-item label="服务器密钥" prop="serverkey">
                        <el-input v-model="loginForm.serverkey" placeholder="请输入服务器密钥"></el-input>
                    </el-form-item>

                    <el-form-item label="账号" prop="account">
                        <el-input v-model="loginForm.account" placeholder="请输入账号"></el-input>
                    </el-form-item>

                    <el-form-item label="新密码" prop="password">
                        <el-input type="password" v-model="loginForm.password" placeholder="请输入新密码"
                            auto-complete="off"></el-input>
                    </el-form-item>

                    <el-form-item label="确认新密码" prop="confirmPassword">
                        <el-input v-model="loginForm.confirmPassword" type="password" placeholder="请确认新密码"></el-input>
                    </el-form-item>

                    <el-form-item>
                        <el-button type="primary" @click="login('loginForm')" class="login-button">登录</el-button>
                    </el-form-item>

                    <div style="display: flex;justify-content: space-between;">
                        <span>
                            <el-link href="/login" type="primary" :underline="false"
                                class="am-margin-left-xs">返回登录账户</el-link>
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
                serverkey: '',
                account: '',
                password: '',
                confirmPassword: '',
            },
            captchaImage: '',
            loginRules: {
                serverkey: [
                    { required: true, message: '请输入服务器密钥', trigger: 'blur' }
                ],
                account: [
                    { required: true, message: '请输入账号', trigger: 'blur' }
                ],
                password: [
                    { required: true, message: '请输入密码', trigger: 'blur' }
                ],
                confirmPassword: [
                    { required: true, message: '请确认密码', trigger: 'blur' },
                    { validator: (rule, value) => value === this.loginForm.password, message: '两次输入的新密码不一致', trigger: 'blur' },
                ]
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
                    this.$axios.put('/home/login/findpw', this.loginForm).then(function (result) {
                        if (result.data.code == 0) {
                            ElMessage({
                                showClose: true,
                                message: result.data.msg,
                                type: 'success',
                            });
                            route.push({ path: '/login' })
                        } else {
                            ElMessage({
                                showClose: true,
                                message: result.data.msg,
                                type: 'error',
                            });
                        }
                    }).catch(function (error) {
                        ElMessage({
                            showClose: true,
                            message: error,
                            type: 'error',
                        });
                    })
                } else {
                    return false;
                }
            })
        },
    }
}

</script>