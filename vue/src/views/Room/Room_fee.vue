<template>
    <el-card v-if="!error">

        <el-space wrap>

            <el-text type="danger" size="large" @click="goback" style="cursor: pointer;"><el-icon>
                    <Back />
                </el-icon>返回房间管理</el-text>
            <h2>{{ title }}</h2>


        </el-space>

        <el-divider class="am-margin-top-xs" />

        <el-table :data="tableData" style="width: 100%">
            <!-- 第一行：标题行 -->
            <el-table-column v-for="(column, index) in columns " :key="index" :label="column.label" :prop="column.prop"
                align="center">

                <!-- 第二行：输入框 -->
                <template v-slot="scope">

                    <template v-if="scope.$index == 0">
                        <el-date-picker v-if="scope.cellIndex == 0" v-model="inputValues['date']" type="month"
                            placeholder="抄表日期" style="width: 120px;" />

                        <el-input v-else v-model="inputValues[column.prop]['fee']" :clearable="true"
                            :placeholder="inputValues[column.prop]['placeholder']" style="width: 100px;" />
                    </template>

                    <!-- 第三行开始显示数字 -->
                    <template v-else-if="scope.$index > 0">
                        <template v-if="scope.cellIndex == 0">
                            <span>
                                {{ scope.row[column.prop] }}
                            </span>
                            <el-popconfirm icon="warning" confirm-button-text="执行删除" confirm-button-type="info"
                                cancel-button-text="取消" cancel-button-type="success" title="确认要删除此条记录吗?数据将无法恢复！"
                                :hide-after="0" width="200px" @confirm="remove(scope)">
                                <template #reference>
                                    <el-icon style="cursor: pointer;">
                                        <Delete />
                                    </el-icon>
                                </template>
                            </el-popconfirm>
                        </template>
                        <template v-else>
                            <div class="am-text-primary">
                                {{ scope.row[column.prop]?.['fee_price'] || 0 }}元
                            </div>
                            <div v-if="feeReulesOptions[column.prop]['fee_rules_price_type'] == 1"
                                class="am-text-warning" style="font-size: 12px;">
                                记 {{ scope.row[column.prop]?.['fee_value'] || 0 }}
                                {{ feeReulesOptions[column.prop]['fee_rules_unit'] }}
                                <div>
                                    <el-text class="mx-1" type="info" size="small">(用量 {{
                                        scope.row[column.prop]?.['fee_used'] || 0 }}
                                        {{ feeReulesOptions[column.prop]['fee_rules_unit'] }})</el-text>
                                </div>
                            </div>
                        </template>
                    </template>
                </template>
            </el-table-column>

            <!-- 最后一列：操作按钮 -->
            <el-table-column label="合计" align="center" width="180">
                <template v-slot="scope">
                    <template v-if="scope.$index == 0">
                        <el-button size="small" @click="add">添加</el-button>
                    </template>

                    <template v-else>
                        <div>
                            总价:{{ scope.row['total'] }}元
                        </div>
                        <div class="am-text-warning" style="font-size: 12px;">
                            (费用:{{ scope.row['totalFee'] }}元)
                        </div>

                    </template>

                </template>
            </el-table-column>
        </el-table>
    </el-card>

    <ErrorComponent :showError="error" :errorMsg="errorMsg" />

</template>

<script>
import { inject } from 'vue';
import ErrorComponent from '../Content/Content_error.vue';
import { timestamp } from '/src/lib/common.js';
export default {
    components: {
        ErrorComponent
    },
    setup() {
        const apiUrl = inject('apiUrl');
        return {
            apiUrl,
        };
    },
    data() {
        return {
            title: '收租记录',
            invoiceStatus: {},
            error: false,
            errorMsg: '',
            room: [],
            feeReulesOptions: {},
            columns: [],
            tableData: [],
            inputValues: {},
        }
    },

    mounted() {
        this.list()
    },

    methods: {
        goback() {
            this.$router.go(-1)
        },
        add() {
            let data = {
                roomID: this.$route.query.roomID,
            };
            for (let key in this.inputValues) {
                if (key == 'date') {
                    let date = this.inputValues[key];
                    data[key] = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0');
                } else {
                    data[key] = this.inputValues[key]['fee']
                }
            }

            let params = this.$route.query

            this.$axios({
                method: 'post',
                url: '/home/room/fee',
                params: params,
                data: data
            }).then((res) => {
                let response = res.data
                this.$message({
                    message: response.msg,
                    type: response.code == 0 ? 'success' : 'error'
                });
                if (response.code == 0) {
                    this.list()
                }
            })

        },
        remove(scope) {
            if (this.tableData[scope.$index] == undefined) {
                this.$message({
                    message: "记录不存在，请刷新",
                    type: "error"
                });
                return
            }

            let date = this.tableData[scope.$index]['date']
            let params = this.$route.query

            this.$axios({
                method: 'DELETE',
                url: '/home/room/fee',
                params: params,
                data: { date: date }
            }).then((res) => {
                let response = res.data
                this.$message({
                    message: response.msg,
                    type: response.code == 0 ? 'success' : 'error'
                });

                if (response.code == 0) {
                    this.tableData.splice(scope.$index, 1);
                }

            })


        },
        list() {
            let dom = this

            // 先清空 columns、tableData 和 inputValues
            dom.columns = [
                { label: '收租日期', prop: 'date' },
            ];
            dom.tableData = [
                { 'date': '', '电费': '', '水费': '' },
            ];
            dom.inputValues = {
                date: null,
            };

            this.$axios.get('/home/room/fee', { params: this.$route.query }).then((res) => {
                let response = res.data

                if (response.code != 0) {
                    this.error = true
                    this.errorMsg = response.msg
                    this.$message.error(this.errorMsg)
                    return
                }

                this.room = response.data.room
                let feeReules = response.data.feeReules
                let feeList = {}
                let cost = JSON.parse(this.room.room_cost)

                this.title = this.room.room_number + '的收租记录'
                document.title = this.title

                feeReules.forEach((item) => {
                    this.feeReulesOptions[item.fee_rules_name] = item
                })

                // console.dir(this.feeReulesOptions)
                // console.dir(cost)

                //组装表格标题
                Object.keys(cost).forEach((key) => {
                    dom.columns.push({ label: key, prop: key })
                    dom.inputValues[key] = {
                        fee: this.feeReulesOptions[key]['fee_rules_price_type'] == '2' ? cost[key] : '',
                        placeholder: this.feeReulesOptions[key]['fee_rules_price_type'] == '2' ? key + '的费用' : key + '的用量',
                    }
                })

                //组装费用
                response.data.feeList.forEach((item) => {
                    let date = item.fee_date
                    if (!feeList[date]) {
                        feeList[date] = []
                    }
                    feeList[date].push(item)
                })

                for (let key in feeList) {
                    let data = {};
                    //费用合计
                    let totalFee = 0;
                    //费用 + 租金
                    let total = 0;
                    data["date"] = key;
                    feeList[key].forEach((item) => {
                        if (item["fee_name"] != '租金') {
                            totalFee += item.fee_price
                        }
                        data[item.fee_name] = item

                        total += item.fee_price
                    })

                    data["totalFee"] = totalFee.toFixed(2)
                    data["total"] = total.toFixed(2)

                    dom.tableData.push(data)
                }

            })
        },

    }

}
</script>
<style>
.el-input__inner::placeholder {
    font-size: 11px;
}
</style>