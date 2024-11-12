<template>
    <el-card v-if="!error">
        <el-space wrap>
            <el-text type="danger" size="large" @click="goback" style="cursor: pointer;"><el-icon>
                    <Back />
                </el-icon>返回</el-text>
            <h3>{{ title }}</h3>
        </el-space>

        <el-divider style="margin-top: 10px;" />

        <el-form label-position="top" label-width="auto" :model="form" :rules="rules" ref="formRef">

            <el-form-item v-for="item in formItems" :key="item.prop" :label="item.label" :prop="item.prop"
                v-show="item.show">

                <!-- 文本框 -->
                <el-input v-if="item.type == 'text'" v-model="form[item.prop]" @change="handleChange(item, form)" />

                <!-- 多行文本框 -->
                <el-input v-if="item.type == 'textarea' || item.type == 'editor'" type="textarea"
                    v-model="form[item.prop]" :rows="10" @change="handleChange(item, form)" />

                <!-- 单选框 -->
                <el-radio-group v-if="item.type == 'radio'" v-model="form[item.prop]"
                    @change="handleChange(item, form)">
                    <el-radio v-for="(value, name) in item.option" :key="value" :label="value" :value="value">
                        <template v-if="!item.render">
                            {{ name }}
                        </template>
                        <div v-else v-html="name"></div>
                    </el-radio>
                </el-radio-group>

                <!-- 多选框 -->
                <el-checkbox-group v-if="item.type == 'checkbox'" v-model="form[item.prop]"
                    @change="handleChange(item, form)">
                    <el-checkbox v-for="(value, name) in item.option" :key="value" :label="value" :value="value">
                        <template v-if="!item.render">
                            {{ name }}
                        </template>
                        <div v-else v-html="name"></div>
                    </el-checkbox>
                </el-checkbox-group>

                <!-- 单选下拉框 -->
                <el-select v-if="item.type == 'select' || item.type == 'category'" v-model="form[item.prop]"
                    placeholder="Select" @change="handleChange(item, form)">
                    <el-option v-for="(value, name) in item.option" :key="value" :label="cleanedHTML(name)"
                        :value="value">
                        <template v-if="!item.render">
                            {{ name }}
                        </template>
                        <div v-else v-html="name"></div>
                    </el-option>
                </el-select>

                <!-- 多选下拉框 -->
                <el-select v-if="item.type == 'multiple'" v-model="form[item.prop]" multiple placeholder="Select"
                    @change="handleChange(item, form)">
                    <el-option v-for="(value, name) in item.option" :key="value" :label="cleanedHTML(name)"
                        :value="value.toString()">
                        <template v-if="!item.render">
                            {{ name }}
                        </template>
                        <div v-else v-html="name"></div>
                    </el-option>
                </el-select>

                <!-- 选项框 -->
                <el-row :gutter="10" v-if="item.type == 'option'" v-for="(value, keys) in form[item.prop]"
                    style="width: 100%;">

                    <el-col :span="2" class="am-margin-bottom-xs">{{ item.op_title[0] ?? '显示名称' }}:</el-col>
                    <el-col :span="8">
                        <el-input v-model="form[item.prop][keys].key" @change="handleChange(item, form)" />
                    </el-col>

                    <el-col :span="2">{{ item.op_title[1] ?? '表单值' }}:</el-col>
                    <el-col :span="8">
                        <el-input v-model="form[item.prop][keys].value" @change="handleChange(item, form)" />
                    </el-col>

                    <el-col :span="4">
                        <el-button @click="addOp(form[item.prop])" type="primary"><el-icon>
                                <Plus />
                            </el-icon></el-button>
                        <el-button @click="removeOp(value, form[item.prop])" type="danger"><el-icon>
                                <Minus />
                            </el-icon></el-button>
                    </el-col>


                </el-row>

                <!-- 日期组件 -->
                <el-date-picker v-if="item.type == 'date'" v-model="form[item.prop]" type="date" format="YYYY-MM-DD"
                    value-format="YYYY-MM-DD" @change="handleChange(item, form)" />



                <el-alert v-if="item.explain" class="am-margin-top-xs" :title="item.explain" type="info"
                    :closable="false" />


            </el-form-item>

            <el-button type="primary" @click="submitForm">提交</el-button>

        </el-form>

    </el-card>

    <ErrorComponent :showError="error" :errorMsg="errorMsg" />

    <el-dialog v-model="imgPreview">
        <img w-full :src="previewScr" alt="Preview Image" />
    </el-dialog>

</template>

<style lang="less">
.pes-upload-actions {
    height: 20% !important;
    opacity: 1 !important;

    .el-upload-list__item-delete {
        display: block !important;
    }
}
</style>

<script>
import { inject } from 'vue';
import { useRoute } from 'vue-router';
import { genFileId } from 'element-plus'
import * as common from '/src/lib/common.js';
import ErrorComponent from '../Content/Content_error.vue';
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

    emits: ['axiosData', 'formChange'],
    data() {
        return {
            error: false,
            errorMsg: '',
            title: "loading...",
            route: useRoute(),
            form: [],
            formItems: [],
            rules: [
            ],
        }
    },

    computed: {
        //定义全局的model名称
        modelName() {
            return (this.route.params.model ?? this.route.meta.model).toLowerCase();
        }
    },

    mounted() {
        let modelName = this.modelName
        let id = this.route.query.id || ''


        this.$axios.get('/home/' + modelName + '/action', { params: this.route.query }).then((result) => {

            if (result.data.code != 0) {
                this.error = true
                this.errorMsg = result.data.msg
                this.$message.error(this.errorMsg)
                return
            }

            let res = result.data
            document.title = (id == '' ? '新增' : '编辑内容') + " - " + res.data.modelInfo.model_title
            this.title = document.title

            let content = res.data.content
            let fieldList = res.data.field

            for (let i = 0; i < fieldList.length; i++) {

                //表单组件不启用则跳开渲染
                if (fieldList[i]['field_status'] == 0) {
                    continue
                }

                let contentVal

                //只有在新增的时候才执行默认值
                if (this.route.query.id == undefined || this.route.query.id == '') {

                    this.form[fieldList[i]['field_name']] = fieldList[i]['field_default']
                } else {
                    try {
                        contentVal = content[`${modelName}_${fieldList[i]['field_name']}`];
                        this.form[fieldList[i]['field_name']] = contentVal == undefined ? '' : String(contentVal)
                    } catch (error) {

                        this.form[fieldList[i]['field_name']] = ""
                    }
                }



                switch (fieldList[i]['field_type']) {
                    case 'checkbox':
                    case 'multiple':
                    case 'imgs':
                    case 'files':
                    case 'videos':
                        this.form[fieldList[i]['field_name']] = content == null ? [] : common.trimSpecified(this.form[fieldList[i]['field_name']], ',').split(',')
                        break;
                    case 'date':
                        if (content != null) {
                            this.form[fieldList[i]['field_name']] = common.timestamp(this.form[fieldList[i]['field_name']], 'Y-m-d')
                        } else {
                            this.form[fieldList[i]['field_name']] = common.date('Y-m-d')
                        }
                        break;
                }

                let isRequire = fieldList[i]['field_required'] == 1 ? true : false
                let message = `${fieldList[i]['field_display_name']}为必填项`

                this.rules[fieldList[i]['field_name']] = [{ required: isRequire, message: message, trigger: 'blur' }]

                // this.rules.push({name:[ {required: true, message: '请输入模型名称', trigger: 'blur'} ]})



                let obj = {
                    label: fieldList[i]['field_display_name'],
                    prop: fieldList[i]['field_name'],
                    type: fieldList[i]['field_type'],
                    explain: fieldList[i]['field_explain'],
                    option: fieldList[i]['field_option'].length > 0 ? JSON.parse(fieldList[i]['field_option']) : '',
                    show: fieldList[i]['field_form'] == 1 ? true : false
                }

                if (fieldList[i]['field_type'] == 'option') {

                    const resultArray = Object.entries(obj.option);
                    if (resultArray.length > 0) {
                        obj['op_title'] = resultArray[0]
                    } else {
                        obj['op_title'] = ['显示名称', '表单值']
                    }

                    let newFieldValue = [{
                        key: '',
                        value: '',
                    }];

                    if (this.form[fieldList[i]['field_name']].length > 0) {

                        let splitJson = JSON.parse(common.decodeHtmlEntities(this.form[fieldList[i]['field_name']]))

                        if (Object.keys(splitJson).length > 0) {
                            newFieldValue = [];
                            Object.keys(splitJson).forEach(key => {
                                newFieldValue.push({
                                    key: key,
                                    value: splitJson[key]
                                })
                            });
                        }
                    }

                    this.form[fieldList[i]['field_name']] = newFieldValue

                }


                this.formItems.push(obj)
            }

            this.$emit('axiosData', {
                response: res,
            });

        })
    },
    methods: {
        submitForm(formRef) {
            this.$refs.formRef.validate((valid, fields) => {
                if (valid) {

                    let data = {}

                    for (let i in this.form) {
                        data[i] = this.form[i]
                    }

                    let id = this.route.query.id || ''

                    this.$axios({
                        method: id == '' ? 'post' : 'put',
                        url: '/home/' + this.modelName + '/action',
                        params: this.route.query,
                        data: data
                    }).then((result) => {
                        const res = result.data

                        if (res.code != 0) {
                            this.$message({
                                message: res.msg,
                                type: 'error'
                            });
                            return
                        }

                        this.$message({
                            message: res.msg,
                            type: 'success'
                        });
                        this.$router.go(-1)
                    }).catch((error) => {
                        this.$message({
                            message: error.response.data.msg,
                            type: 'error'
                        });
                    });

                } else {
                    console.log('error submit!', fields);
                }
            });
        },
        goback() {
            this.$router.go(-1)
        },
        /**
         * 表单更新
         * @param {*} currentForm 
         * @param {*} formData 
         */
        handleChange(currentForm, formData) {
            this.$emit('formChange', currentForm, formData, this.formItems);
        },
        addOp(form) {
            form.push({
                key: '',
                value: ''
            })
        },
        removeOp(item, form) {
            const index = form.indexOf(item);
            if (index !== -1) {
                form.splice(index, 1)
            }
        },
        updateData(func) {
            func(this.$data);
        },
        cleanedHTML(item) {
            return item.replace(/\&nbsp;/g, '')
        },

    }
}
</script>