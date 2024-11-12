<template>
    <el-card v-if="showError">
        <el-row justify="space-between">
            <el-col :span="6"></el-col>
            <el-col :span="6">
                <el-result icon="error" title="系统提示" :sub-title="errorMsg">
                    <template #extra>
                        <el-button type="primary" @click="goBack">
                            <el-icon>
                                <Back />
                            </el-icon>
                            返回上一页
                        </el-button>
                    </template>
                </el-result>
            </el-col>
            <el-col :span="6"></el-col>
        </el-row>
    </el-card>
</template>

<script>
export default {
    name: 'ErrorComponent',
    props: {
        showError: {
            type: Boolean,
            default: false
        },
        errorMsg: {
            type: String,
            default: '发生错误'
        }
    },
    methods: {
        goBack() {
            this.$router.go(-1);
        }
    }
}
</script>