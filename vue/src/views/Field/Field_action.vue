<template>
    <Content_action @formChange="formChange" @axiosData="axiosData" ref="childRef"></Content_action>
</template>

<script>

import Content_action from '../Content/Content_action.vue';

export default {
    components: {
        Content_action,
    },

    methods: {

        axiosData(res) {
            let modelID = this.$route.query.modelID;

            let id = this.$route.query.id

            this.$refs.childRef.updateData(function (data) {
                data.form.model_id = modelID
                if (id == undefined) {
                    data.form.sql_length = '11'
                }
            });
        },

        formChange(currentForm, formData) {
            if (currentForm.prop == 'sql_type') {
                switch (formData.sql_type) {
                    case "int":
                        formData.sql_length = "11"
                        break;
                    case "varchar":
                        formData.sql_length = "255"
                        break;
                    case "decimal":
                        formData.sql_length = "10,2"
                        break;
                    case "text":
                        formData.sql_length = "1000"
                        break;

                }
            }
        }
    }
}
</script>