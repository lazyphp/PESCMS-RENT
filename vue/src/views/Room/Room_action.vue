<template>
    <Content_action @formChange="formChange" @axiosData="axiosData" ref="childRef"></Content_action>
</template>

<script>

import Content_action from '../Content/Content_action.vue';

export default {
    components: {
        Content_action,
    },

    data() {
        return {

        }
    },

    methods: {

        axiosData(res) {
            let houseID = this.$route.query.houseID;

            let feeReules = res.response.data.feeReules

            let feeReulesOptions = {}

            let newItem = []

            this.$refs.childRef.updateData(function (data) {

                if (data.form.cost.length == 0) {
                    data.form.cost = '{}'
                }
                let cost = JSON.parse(data.form.cost)

                feeReules.forEach((item) => {
                    feeReulesOptions[item.fee_rules_name] = item.fee_rules_name.toString()

                    //组装计费规则的表单
                    newItem.push({
                        prop: item.fee_rules_name,
                        label: item.fee_rules_name,
                        show: data.form.fee_rules.includes(item.fee_rules_name),
                        type: 'text',
                        option: {},
                    })

                    data.rules[item.fee_rules_name] = [{
                        required: true,
                        message: '请输入' + item.fee_rules_name,
                        trigger: 'blur'
                    }]

                    if (cost[item.fee_rules_name]) {
                        data.form[item.fee_rules_name] = cost[item.fee_rules_name].toString()
                    } else {
                        data.form[item.fee_rules_name] = item.fee_rules_unit_price.toString()
                    }

                })

                data.form.house_id = houseID

                data.formItems.forEach((item, key) => {
                    if (item.prop == 'fee_rules') {
                        item.option = feeReulesOptions

                        //追加计费规则的表单
                        data.formItems.splice(key + 1, 0, ...newItem);

                    }
                })

            });
        },

        formChange(currentForm, formData, formItem) {

            if (currentForm.prop == 'fee_rules') {
                // 获取 currentForm.option 的所有值
                const optionValues = Object.values(currentForm.option);

                // 找出 formData.fee_rules 中缺少的值
                const missingValues = optionValues.filter(value => !formData.fee_rules.includes(value));

                formItem.forEach(item => {
                    if (optionValues.includes(item.prop)) {
                        if (missingValues.includes(item.prop)) {
                            item.show = false;
                        } else {
                            item.show = true;
                        }
                    }

                });

            }
        }
    }
}
</script>