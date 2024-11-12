<template>
    <Content_list :addButton="'addButton'">

        <template v-slot:addButton>
            <el-button size="small" @click="addModel">
                <el-icon>
                    <Plus />
                </el-icon>
                新增模型</el-button>
        </template>

        <template v-slot:default="slotProps">
            <el-button size="small" @click="operation(slotProps.scope)">字段管理</el-button>
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
        addModel() {
            this.$router.push({
                path: `/model/action`,
                query: {
                    model: "model",
                }
            })
        },
        operation(scope) {
            // const route = useRoute();
            let modelID = scope.row.model_id;
            this.$router.push({
                path: `/field/list`,
                query: {
                    model: "field",
                    modelID: modelID,
                }
            })
        }
    }
}
</script>