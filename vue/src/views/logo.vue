<template>
    <div class="logo-container">
        <div style="display: flex;align-items: center;flex-direction: column;">
            <img src="https://www.pescms.com/Theme/assets/i/logo.png" alt="Logo" class="logo">

            <Transition>
                <el-text v-if="longtime" class="am-margin-top" type="warning"
                    style="display: flex;flex-direction: column;align-items: center;">
                    您好像长时间停留此页面了，页面可能已经失效，点击下方按钮返回首页
                    <div class="am-margin-top">
                        <el-button type="primary" @click="goHome"><el-icon>
                                <HomeFilled />
                            </el-icon> 返回房租管理</el-button>
                    </div>
                </el-text>
            </Transition>
        </div>

    </div>
</template>

<style>
.v-enter-active,
.v-leave-active {
    transition: opacity 0.8s ease;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}

.logo-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    /* 为了让logo居中显示，这里设置一个高度，可以根据实际情况调整 */
}

.logo {
    width: 150px;
    /* 根据你的需求设置合适的宽度 */
    height: 150px;
    /* 根据你的需求设置合适的高度 */
}

.logo {
    opacity: 0;
    /* 初始时透明 */
    animation: revealLogo 5s ease-in-out infinite;
    /* 这里是动画效果的设置 */
    clip-path: inset(0 0 100% 0);
    /* 初始时隐藏整个图标 */
}

@keyframes revealLogo {
    0% {
        clip-path: inset(0 0 100% 0);
        /* 隐藏整个图标 */
        opacity: 0;
        /* 透明度为0 */
    }

    25% {
        clip-path: inset(0 0 75% 0);
        /* 显示图标的四分之一 */
        opacity: 1;
        /* 透明度为1 */
    }

    50% {
        clip-path: inset(0 0 50% 0);
        /* 显示图标的一半 */
        opacity: 1;
        /* 透明度为1 */
    }

    75% {
        clip-path: inset(0 0 25% 0);
        /* 显示图标的四分之三 */
        opacity: 1;
        /* 透明度为1 */
    }

    100% {
        clip-path: inset(0 0 0 0);
        /* 完全显示图标 */
        opacity: 1;
        /* 透明度为1 */
    }
}
</style>

<script>
export default {
    data() {
        return {
            longtime: false
        }
    },
    mounted() {
        setTimeout(() => {
            this.longtime = true
        }, 10000)
    },
    methods: {
        goHome() {
            this.$router.push('/house/list');
        }
    }
}
</script>