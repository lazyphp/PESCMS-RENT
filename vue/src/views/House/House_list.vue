<template>
    <Content_list ref="contentListRef" @axiosData="axiosData" :customTable="'customTable'">

        <template v-slot:customTable>

            <div class="card-container">
                <div v-for="(item, index) in table" :key="index" class="card-item">
                    <el-card shadow="hover">

                        <template #header>
                            <div class="card-header">
                                <span>{{ item["house_name"] }}</span>

                                <el-button class="am-margin-left-xs" size="small" type="success"
                                    @click="roomManage(item)">
                                    <el-icon class="am-margin-right-xs">
                                        <Collection />
                                    </el-icon>
                                    管理 / 收租房间
                                </el-button>

                            </div>
                        </template>

                        <div>
                            <el-tag type="success" class="am-margin-right-xs">
                                <el-icon>
                                    <Connection />
                                </el-icon>在租：{{ roomStatus[item.house_id] ?
        roomStatus[item.house_id][1] : '0' }}
                            </el-tag>
                            <el-tag type="warning">
                                <el-icon>
                                    <WarnTriangleFilled />
                                </el-icon> 闲置：{{ roomStatus[item.house_id] ?
        roomStatus[item.house_id][0] : '0' }}
                            </el-tag>


                        </div>

                        <template #footer>

                            <el-button size="small" @click="operation({ row: item, $index: index }, 'edit')">
                                <el-icon>
                                    <Edit />
                                </el-icon>
                                编辑信息</el-button>

                            <el-popconfirm icon="warning" confirm-button-text="执行删除" confirm-button-type="info"
                                cancel-button-text="取消" cancel-button-type="success" title="确认要删除此条记录吗?数据将无法恢复！"
                                width="200px" @confirm="operation({ row: item, $index: index }, 'delete')">
                                <template #reference>
                                    <el-button size="small" type="danger">
                                        <el-icon>
                                            <Delete />
                                        </el-icon>
                                        删除房子
                                    </el-button>
                                </template>
                            </el-popconfirm>

                        </template>

                    </el-card>
                </div>
            </div>
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
    data() {
        return {
            table: [],
            roomStatus: []
        }
    },

    methods: {

        axiosData(res) {

            let dom = this

            this.roomStatus = res.response.data.roomStatus

            this.$refs.contentListRef.updateData(function (data) {
                dom.table = data.tableData
            });
        },

        operation(scope, type) {
            // 调用子组件的 operation 方法
            this.$refs.contentListRef.operation(scope, type)
        },

        roomManage(item) {
            this.$router.push({
                path: `/room/list`,
                query: {
                    houseID: item.house_id,
                }
            })
        },
    }
}
</script>