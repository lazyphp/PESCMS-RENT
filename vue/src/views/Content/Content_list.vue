<template>
    <el-card v-if="!error">
        <el-skeleton :loading="skeletonLoading" :rows="2" animated>

            <el-space wrap>

                <slot :name="gobackButton"></slot>
                <h2>{{ title }}</h2>

                <el-button size="small" @click="operation(null, 'add')" v-show="addButton != 'addButton'">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    新增</el-button>

                <slot :name="addButton"></slot>

            </el-space>


            <el-divider class="am-margin-top-xs" />

            <el-table :data="tableData" border style="width: 100%" v-show="customTable != 'customTable'"
                :row-class-name="tableRowClassName">

                <el-table-column label="排序" width="70" v-if="hasSort">
                    <template v-slot:default="scope">
                        <el-form-item class="am-margin-bottom-0">
                            <el-input v-model="scope.row[modelName + '_listsort']" />
                        </el-form-item>
                    </template>
                </el-table-column>

                <template v-for="column in columns">
                    <el-table-column :prop="column.prop" :label="column.label" :width="column.width"
                        v-if="column.prop != modelName + '_listsort'">
                        <template v-if="column.render" v-slot="{ row }">
                            <div v-html="row[column.prop]"></div>
                        </template>
                    </el-table-column>
                </template>


                <el-table-column label="操作">

                    <template v-slot:default="scope">
                        <el-button size="small" @click="operation(scope, 'edit')" v-show="editButton != 'editButton'">
                            <el-icon>
                                <Edit />
                            </el-icon>
                            编辑</el-button>
                        <slot :name="editButton" :scope="scope"></slot>
                        <slot :scope="scope"></slot>

                        <slot :name="deleteButton"></slot>

                        <el-popconfirm icon="warning" confirm-button-text="执行删除" confirm-button-type="info"
                            cancel-button-text="取消" cancel-button-type="success" title="确认要删除此条记录吗?数据将无法恢复！"
                            :hide-after="0" width="200px" @confirm="operation(scope, 'delete')">
                            <template #reference>
                                <el-button size="small" type="danger" v-show="deleteButton != 'deleteButton'">
                                    <el-icon>
                                        <Delete />
                                    </el-icon>
                                    删除
                                </el-button>
                            </template>
                        </el-popconfirm>
                    </template>
                </el-table-column>

            </el-table>

            <slot :name="customTable"></slot>


            <el-row class="am-margin-vertical">
                <el-col :span="2" v-if="hasSort">
                    <el-button type="primary" @click="sort">
                        <el-icon>
                            <Edit />
                        </el-icon>
                        排序</el-button>
                </el-col>

                <el-col :span="19">
                    <el-pagination v-model:current-page="currentPage" background layout="prev, pager, next"
                        :total="pageTotal" :page-size="pageZise" />
                </el-col>
            </el-row>


        </el-skeleton>

    </el-card>

    <ErrorComponent :showError="error" :errorMsg="errorMsg" />


</template>

<script>

import { useRoute } from 'vue-router';
import { timestamp } from '/src/lib/common.js';
import ErrorComponent from '../Content/Content_error.vue';

export default {
    components: {
        ErrorComponent
    },
    emits: ['axiosData'],
    name: 'Content_list',
    data() {
        return {
            error: false,
            errorMsg: '',
            showTable: true,
            hasSort: false,
            columns: [],
            tableData: [],
            route: useRoute(),
            title: "Loading...",
            skeletonLoading: true,
            currentPage: 1,
            pageTotal: 10,
            pageZise: 10
        }
    },

    props: {
        gobackButton: {
            type: String,
            default: 'gobackButton'
        },
        addButton: {
            type: String,
            default: 'show'
        },
        customTable: {
            type: String,
            default: 'show'
        },
        editButton: {
            type: String,
            default: 'show'
        },
        deleteButton: {
            type: String,
            default: 'show'
        }
    },

    computed: {
        //定义全局的model名称
        modelName() {
            return (this.route.params.model ?? this.route.meta.model).toLowerCase();
        }
    },

    watch: {
        currentPage(val) {

            const query = Object.assign({}, this.$route.query, { page: val });

            this.$router.push({
                path: `/${this.modelName}/list`,
                query: query
            })

        },
        $route(to, from) {
            //@todo 存在一个BUG。当路由回滚，可能二次加载数据
            this.list()
        }
    },

    mounted() {
        this.list()
    },
    methods: {
        list() {

            let modelName = this.modelName
            this.route.query["model"] = modelName

            //初始化表头
            this.columns = [{ prop: `${modelName}_id`, label: 'ID', width: 60 }]
            //初始化表格数据
            this.tableData = [];

            if (this.route.query.page) {
                this.currentPage = parseInt(this.route.query.page)
            }

            this.$axios.get('/home/' + modelName + '/index', { params: this.route.query }).then((result) => {
                this.error = false
                if (result.data.code != 0) {
                    this.error = true
                    this.errorMsg = result.data.msg
                    this.$message.error(this.errorMsg)
                    return
                }
                let res = result.data
                let modelInfo = res.data.modelInfo

                //设置页面标题
                document.title = modelInfo.model_title
                this.title = modelInfo.model_title

                //设置分页数
                this.pageZise = modelInfo.model_page
                this.pageTotal = res.data.pageTotal

                let field = res.data.field
                let contentList = res.data.contentList
                let showColumns = [{ name: `${modelName}_id`, option: {}, type: 'text' }];
                // console.dir(field)
                // console.dir(contentList)

                //开始组装表格头部信息
                for (let i = 0; i < field.length; i++) {
                    if (field[i]['field_list'] == 0) {
                        continue
                    }

                    if (field[i]['field_status'] == 0) {
                        continue
                    }


                    //设置排序按钮显示
                    if (field[i]['field_name'] == 'listsort') {
                        this.hasSort = true
                    }

                    let sqlField = `${modelName}_${field[i]['field_name']}`

                    let fieldOption = {}
                    if (field[i]['field_option'].length > 0) {
                        fieldOption = JSON.parse(field[i]['field_option']);
                    }

                    //记录需要展示的列字段和对应的选项值（有得话）
                    showColumns.push({
                        name: sqlField,
                        option: fieldOption,
                        type: field[i]['field_type']
                    })
                    //组装表格头部信息
                    this.columns.push({
                        prop: sqlField,
                        label: field[i]['field_display_name']
                    })
                }

                //开始组装表格内容

                for (let i in contentList) {
                    let pushConetnt = {};
                    for (let j in showColumns) {
                        switch (showColumns[j]['type']) {
                            case 'text':
                            case 'textarea':
                            case 'editor':
                            case 'option':
                            case 'category':
                                pushConetnt[showColumns[j]['name']] = contentList[i][showColumns[j]['name']]
                                break;
                            case 'radio':
                            case 'select':
                                //拆分字段选项值，替换内容
                                pushConetnt[showColumns[j]['name']] = "未知选项值"
                                if (Object.keys(showColumns[j]['option']).length > 0) {
                                    for (let key in showColumns[j]['option']) {
                                        if (showColumns[j]['option'][key] == contentList[i][showColumns[j]['name']]) {
                                            pushConetnt[showColumns[j]['name']] = key
                                        }
                                    }
                                }
                                break;
                            case 'checkbox':
                            case 'multiple':
                                //多选组装数据
                                let multipleValue = [];
                                let newArray = contentList[i][showColumns[j]['name']].split(',')

                                if (Object.keys(showColumns[j]['option']).length > 0) {

                                    for (let key in showColumns[j]['option']) {

                                        if (newArray.includes(showColumns[j]['option'][key].toString())) {
                                            multipleValue.push(key)
                                        }
                                    }
                                }

                                pushConetnt[showColumns[j]['name']] = multipleValue.join(',')
                                break
                            case 'date':

                                pushConetnt[showColumns[j]['name']] = timestamp(contentList[i][showColumns[j]['name']], 'Y-m-d H:i:s')

                                break;
                            default:
                                break;
                        }


                    }
                    this.tableData.push(pushConetnt)
                }

                // console.dir(this.columns)
                // console.dir(this.tableData)

                this.$emit('axiosData', {
                    response: res,
                    // 其他需要传递的变量
                });

            }).finally(() => {

                this.skeletonLoading = false
            })
        },
        operation(scope, type) {

            let modelName = this.modelName

            let id = type == 'add' ? '' : scope.row[`${modelName}_id`]

            if (type == 'edit' || type == 'add') {
                this.$router.push({
                    path: `/${modelName}/action`,
                    query: {
                        id: id,
                        model: modelName
                    }
                })
            } else {
                this.$axios.delete('/home/' + modelName + '/action?id=' + id, {}).then((res) => {
                    this.tableData.splice(scope.$index, 1);
                }).catch((err) => {
                    console.dir(err)
                })
            }
        },

        sort() {
            let sort = []
            for (let i in this.tableData) {
                sort.push({ id: this.tableData[i][`${this.modelName}_id`], sort: this.tableData[i][`${this.modelName}_listsort`] })
            }
            this.$axios.put('/home/' + this.modelName + '/sort', { sort: sort }).then((res) => {
                this.$message({
                    message: '排序完成',
                    type: 'success'
                });
            }).finally(() => {
                this.list()
            })
        },

        //表格行样式设置
        tableRowClassName(param) {
            if (param.row.rowClass) {
                return param.row.rowClass
            } else {
                return ''
            }
        },

        goback() {
            this.$router.go(-1)
        },

        updateData(func) {
            func(this.$data);
        }
    }
}
</script>
