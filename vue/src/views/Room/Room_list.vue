<template>
    <Content_list ref="contentListRef" @axiosData="axiosData" :addButton="'addButton'" :customTable="'customTable'">

        <template v-slot:gobackButton>
            <el-text type="danger" size="large" @click="goback" style="cursor: pointer;"><el-icon>
                    <Back />
                </el-icon>返回房屋管理</el-text>
        </template>

        <template v-slot:addButton>
            <el-button size="small" @click="addRoom">
                <el-icon>
                    <Plus />
                </el-icon>
                添加新房间</el-button>
        </template>


        <template v-slot:customTable>
            <template v-for="(room, floor) in table" v-if="Object.keys(table).length > 0">
                <el-divider>
                    {{ floor }}
                </el-divider>
                <div class="card-container">
                    <template v-for="(item, index) in room" :key="index">
                        <div :class="['card-item', item['room_status'] == '闲置' ? 'room-idle' : '']">
                            <el-card shadow="hover">

                                <template #header>
                                    <div class="card-header">
                                        <span>{{ item["room_number"] }}</span>
                                        <span>
                                            <el-text class="mx-1"
                                                :type="item['room_status'] == '闲置' ? 'danger' : 'success'"
                                                size="small">{{
                                                    item["room_status"]
                                                }}</el-text>
                                        </span>
                                    </div>
                                </template>

                                <div>

                                    <el-tag type="danger" class="am-margin-right-xs am-margin-bottom-xs">
                                        <el-icon>
                                            <User />
                                        </el-icon> 租金：{{ item["room_rent"] }}元/月</el-tag>

                                    <template v-for="(    fee, name    ) in item.room_cost    ">
                                        <el-tag type="info" class="am-margin-right-xs am-margin-bottom-xs">
                                            <el-icon>
                                                <Connection />
                                            </el-icon> {{ name }}：{{ fee }}元/{{ feeReulesOptions[name]["fee_rules_unit"]
                                                ==
                                                "" ? '月' : feeReulesOptions[name]["fee_rules_unit"]
                                            }}</el-tag>
                                    </template>

                                </div>

                                <template #footer>

                                    <el-button size="small" @click="operation({ row: item, $index: index }, 'edit')">
                                        <el-icon>
                                            <Edit />
                                        </el-icon>
                                        编辑信息</el-button>

                                    <el-button class="am-margin-left-xs" size="small" type="success"
                                        @click="roomManage(item)">
                                        <el-icon class="am-margin-right-xs">
                                            <Collection />
                                        </el-icon>
                                        收租记录
                                    </el-button>

                                    <el-popconfirm icon="warning" confirm-button-text="执行删除" confirm-button-type="info"
                                        cancel-button-text="取消" cancel-button-type="success" title="确认要删除此条记录吗?数据将无法恢复！"
                                        width="200px" @confirm="operation({ row: item, $index: index }, 'delete')">
                                        <template #reference>
                                            <el-button size="small" type="danger">
                                                <el-icon>
                                                    <Delete />
                                                </el-icon>
                                                移除
                                            </el-button>
                                        </template>
                                    </el-popconfirm>

                                </template>

                            </el-card>
                        </div>
                    </template>

                </div>
            </template>

            <template v-else>
                <el-empty description="房子还没添加房间"></el-empty>
            </template>


        </template>

    </Content_list>
</template>

<script>

import Content_list from '../Content/Content_list.vue';
import { printContract } from '/src/lib/common.js';
// import { useRoute } from 'vue-router';

export default {
    components: {
        Content_list,
    },
    data() {
        return {
            table: {},
            feeReulesOptions: {},
        }
    },

    methods: {

        axiosData(res) {
            let dom = this

            let feeReules = res.response.data.feeReules
            let feeReulesOptions = {}

            feeReules.forEach((item) => {
                this.feeReulesOptions[item.fee_rules_name] = item
            })

            this.$refs.contentListRef.updateData(function (data) {
                // dom.table = data.tableData

                data.tableData.forEach((item) => {
                    if (item.room_floor == '' || item.room_floor == undefined) {
                        item.room_floor = '所有楼层'

                    }
                    if (!dom.table[item.room_floor]) {
                        dom.table[item.room_floor] = [];
                    }

                    if (item.room_cost) {
                        item.room_cost = JSON.parse(item.room_cost)
                    }

                    dom.table[item.room_floor].push(item)
                })

            });
        },

        operation(scope, type) {
            let modelName = this.$refs.contentListRef.modelName
            let id = scope.row[`room_id`]

            let param = Object.assign({}, this.$route.query, { id })

            if (type == 'edit' || type == 'add') {
                this.$router.push({
                    path: `/room/action`,
                    query: param
                })
            } else {
                this.$axios.delete('/home/' + modelName + '/action?id=' + id, {}).then((res) => {
                    if (res.data.code != 0) {
                        this.$message({
                            message: res.msg,
                            type: 'error'
                        });
                        return
                    }

                    this.table[scope.row.room_floor].splice(scope.$index, 1);
                }).catch((err) => {
                    this.$message({
                        message: "执行删除失败，请重试",
                        type: 'error'
                    });
                })
            }
        },

        goback() {
            this.$router.push({
                path: `/house/list`,
            })
        },

        addRoom() {
            let houseID = this.$route.query.houseID || ''
            this.$router.push({
                path: `/room/action`,
                query: {
                    houseID: houseID,
                    model: 'room'
                }
            })
        },

        roomManage(item) {
            this.$router.push({
                path: `/room/fee`,
                query: {
                    roomID: item.room_id,
                }
            })
        },

        contract(item, type) {
            console.log(item)

            this.$axios.get('/home/setting/contract',).then(res => {
                let reponse = res.data
                if (reponse.code != 0) {
                    this.$message.error("读取合同目标失败")
                    return false
                }

                let contract = reponse.data.contract

                /**
                 * 房租费用构成
                 */
                let cost = item.room_cost
                let feeType = [];
                for (let key in cost) {
                    let unit = this.feeReulesOptions[key].fee_rules_price_type == 1 ? '元/' + this.feeReulesOptions[key].fee_rules_unit : '元'
                    feeType.push(`${key}<u>&nbsp;&nbsp;${cost[key]}&nbsp;&nbsp;</u>${unit}`)
                }

                contract = contract.replace(/{费用类目}/g, feeType.join('、'))


                /**
                 * 房间资产
                 */
                let room_fixed_assets = {}
                try {
                    room_fixed_assets = JSON.parse(item.room_fixed_assets)
                } catch (error) {
                    room_fixed_assets = {}
                }

                if (Object.keys(room_fixed_assets).length > 0) {

                    let device = []
                    let fee = '';

                    let deviceValue = contract.match(/{设备价值}/g)

                    for (let key in room_fixed_assets) {

                        if (deviceValue.length > 0) {
                            fee = `[${room_fixed_assets[key]}元]`
                        } else {
                            fee = ''
                        }

                        device.push(`（　）${key} ${fee}`)
                    }
                    contract = contract.replace(/{设备}/g, device.join('、'))

                } else {
                    contract = contract.replace(/{设备}/g, '')
                }
                contract = contract.replace(/{设备价值}/g, '')

                switch (type) {
                    case 'word':
                        asBlob(contract).then((data) => {
                            saveAs(data, `${item['room_number']}的房租合同.docx`);
                        });
                        break;
                    case 'print':
                        printContract(contract)
                        break;
                    default:

                        break;
                }


            })

        }

    }
}
</script>