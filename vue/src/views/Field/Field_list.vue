<template>
    <Content_list :addButton="'addButton'" :editButton="'editButton'">

        <template v-slot:gobackButton>
            <el-text type="danger" size="large" @click="goback" style="cursor: pointer;"><el-icon>
                    <Back />
                </el-icon>返回模型管理</el-text>
        </template>

        <template v-slot:addButton>
            <el-button size="small" @click="addField">
                <el-icon>
                    <Plus />
                </el-icon>
                添加字段</el-button>
        </template>


        <template v-slot:default="scope">
            <el-button size="small" @click="operation(scope.scope, 'edit')">
                <el-icon>
                    <Edit />
                </el-icon>
                编辑</el-button>
        </template>

    </Content_list>
</template>

<script>

import Content_list from '../Content/Content_list.vue';

// import { useRoute } from 'vue-router';

export default {
    components: {
        Content_list,
    },

    methods: {
        addField() {
            this.$router.push({
                path: `/field/action`,
                query: this.$route.query
            })
        },
        operation(scope, type) {
            let modelName = this.modelName

            let param = Object.assign({}, this.$route.query, { id: scope.row[`field_id`] })

            if (type == 'edit' || type == 'add') {
                this.$router.push({
                    path: `/field/action`,
                    query: param
                })
            }
        },
        goback() {
            this.$router.push({
                path: `/model/list`,
            })
        }
    }
}
</script>