<template>
    <el-header>
        <el-affix :offset="0">
            <el-menu :router="true" :default-active="activeMenu" mode="horizontal" background-color="#545c64"
                text-color="#fff" active-text-color="#ffd04b" @select="handleSelect">

                <div class="flex-grow" />
                <template v-for="value in menu">
                    <el-menu-item :index="value.node_link" :route="value.node_link" v-if="value.child == null">

                        <el-icon>
                            <component :is="value.node_menu_icon" />
                        </el-icon>

                        {{ value.node_name }}</el-menu-item>

                    <el-sub-menu :index="value.node_id.toString()" v-else>
                        <template #title><el-icon>
                                <component :is="value.node_menu_icon" />
                            </el-icon>{{ value.node_name }}</template>
                        <template v-for="item in value.child">
                            <el-menu-item :index="item.node_link" :route="item.node_link">
                                <el-icon>
                                    <component :is="item.node_menu_icon" />
                                </el-icon>
                                {{ item.node_name }}
                            </el-menu-item>
                        </template>
                    </el-sub-menu>
                </template>
            </el-menu>

            <div class="title-bar">
                <el-text @click="logout" type="warning" class="am-margin-left-xs">
                    退出登录
                </el-text>
            </div>

        </el-affix>
    </el-header>
</template>
<style @lang="less">
.el-header {
    width: 100%;
    padding: 0 !important;
    position: relative;

    .title-bar {
        position: absolute;
        right: 10px;
        top: 20px;

        span {

            margin-left: 10px;

            &::before {
                font-size: 20px;
                color: #fff;
            }

            &:hover {
                cursor: pointer;

                &::before {
                    color: #ffd04b;
                }
            }


        }
    }

}
</style>

<script>
import { ElMessageBox } from 'element-plus';
export default {
    data() {
        return {
            activeMenu: "/",
            menu: [],
        }
    },
    mounted() {
        this.$axios.get('/home/node/index').then(res => {
            let menu = res.data.data.menu
            for (let i in menu) {
                this.menu.push({
                    node_id: menu[i].node_id,
                    node_name: menu[i].node_name,
                    node_link: menu[i].node_link,
                    node_menu_icon: menu[i].node_icon,
                })
            }

            this.activeMenu = this.$route.path //设置菜单选中
        })

    },
    methods: {
        logout() {
            ElMessageBox.confirm('确定退出登录吗?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                localStorage.removeItem('token')
                this.$router.push('/login')
            }).catch(() => {

            });
        }
    }
}

</script>