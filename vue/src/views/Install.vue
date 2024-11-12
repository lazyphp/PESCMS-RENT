<template>
    <div class="install-container">
        <el-card class="install-card" shadow="hover">
            <h2 class="install-title">欢迎使用PESCMS RENT 房租管理系统开源版</h2>


            <div class="agreement-section">
                <el-scrollbar ref="scrollbar" class="agreement-scroll" @scroll="handleScroll">
                    <div class="agreement-content">
                        <div v-html="agreement">

                        </div>
                    </div>
                </el-scrollbar>
            </div>


            <el-divider></el-divider>


            <el-form :model="form" :rules="rules" ref="formRef" label-width="120px">
                <el-form-item label="管理员账户" prop="account">
                    <el-input v-model="form.account" placeholder="请输入管理员账户"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input v-model="form.password" type="password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-form-item label="确认密码" prop="confirmPassword">
                    <el-input v-model="form.confirmPassword" type="password" placeholder="请确认密码"></el-input>
                </el-form-item>
            </el-form>


            <el-tooltip v-if="!isAgreementEnd" content="请先将用户协议滚动到最底部" placement="top">
                <el-checkbox v-model="form.agreementChecked" :disabled="!isAgreementEnd">
                    我已阅读并同意软件协议
                </el-checkbox>
            </el-tooltip>
            <el-checkbox v-else v-model="form.agreementChecked">
                我已阅读并同意软件协议
            </el-checkbox>


            <el-button class="am-margin-left" type="primary" :disabled="!form.agreementChecked" @click="handleSubmit">
                提交
            </el-button>
        </el-card>
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, inject } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';

const axios = inject("axios");

const router = useRouter();
const scrollbar = ref(null);
const formRef = ref(null);
const agreement = ref('用户协议正在加载中...')
const form = reactive({
    account: '',
    password: '',
    confirmPassword: '',
    agreementChecked: false,
});

const rules = {
    account: [{ required: true, message: '请输入管理员账户', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    confirmPassword: [
        { required: true, message: '请确认密码', trigger: 'blur' },
        { validator: (rule, value) => value === form.password, message: '两次输入的密码不一致', trigger: 'blur' },
    ],
};

const isAgreementEnd = ref(false);

const handleScroll = (e) => {
    const wrap = scrollbar.value?.wrapRef;

    if (wrap) {
        const { scrollTop, scrollHeight, clientHeight } = wrap;
        isAgreementEnd.value = scrollTop + clientHeight >= scrollHeight - 10;
    }
};

const handleSubmit = () => {
    formRef.value.validate((valid) => {
        if (valid) {
            axios.post('/home/setting/install', form).then(res => {
                if (res.data.code == 0) {
                    ElMessage.success('安装成功');
                    setTimeout(() => {
                        window.location.href = '/login';
                    }, 1000);
                } else {
                    ElMessage.error(res.data.msg);
                }
            }).catch(error => {
                ElMessage.error('安装失败，请稍后重试');
            });
        } else {
            ElMessage.error('请填写完整信息');
        }
    });
};


onMounted(() => {

    axios.get('/home/setting/install').then(response => {
        if (response.data.code == 0) {
            if (response.data.data.install != 0) {
                router.push("/");
            }
        }
    })

    axios.get('https://www.pescms.com/?g=Api&m=Article&a=rent').then(res => {
        if (res.data.status === 200) {
            agreement.value = res.data.data;
        }
    }).catch(error => {
        agreement.value = '加载用户协议时出错，请稍后重试。';
    });;
});

</script>

<style scoped>
.install-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f5f8fc;
}

.install-card {
    max-width: 1024px;
    width: 100%;
    padding: 20px;
}

.install-title {
    text-align: center;
    font-size: 24px;
    margin-bottom: 20px;
    color: #1d1d1f;
}

.agreement-section {
    height: 200px;
    border: 1px solid #ebeef5;
    overflow: hidden;
    margin-bottom: 20px;
}

.agreement-scroll {
    height: 100%;
}

.agreement-content {
    padding: 10px;
    line-height: 1.6;
}
</style>