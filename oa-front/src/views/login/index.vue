<template>
    <el-container class="container" direction="vertical">
        <el-main>
            <el-form class="form" :model="model" :rules="rules">
                <h1 class="title">中研环科管理系统</h1>
                <el-form-item prop="phone">
                    <el-input class="input" v-model="model.phone" placeholder="手机号" clearable prefix-icon="User">
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input class="input" v-model="model.password" placeholder="密码" show-password clearable
                        prefix-icon="Lock">
                    </el-input>
                </el-form-item>
                <el-form-item>
                    <el-button class="btn" type="primary" @click="submit">登录</el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
</template>

<script setup>
import { login } from "@/api/login";
import { reactive } from 'vue'

const model = reactive({
    phone: '',
    password: '',
})
const rules = reactive({
    phone: [
        { required: true, message: '请输入手机号', trigger: 'blur' },
    ],
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
    ],
})
const submit = () => {
    login(model).then((res) => {
        if (res.status == 1) {
            localStorage.setItem("token", res.data.token);
            window.location.href = '/'
        } else if (res.status == 2) {

        }
    })
}
</script>

<style>
.container {
    background: #2d3a4b;

}

.form {
    width: 520px;
    max-width: 100%;
    margin: 200px auto 0;
    box-sizing: border-box;
    padding: 0 24px;
}

.title {
    color: #fff;
    text-align: center;
    font-weight: normal;
    font-size: 36px;
    margin: 0 0 40px;
}

.input {
    font-size: 18px;
}

.el-input__wrapper {
    background: #2d3a4b;
}

.el-input__inner {
    background: #2d3a4b;
    color: #fff;
    height: 60px;
    line-height: 60px;
}

.btn {
    height: 40px;
    width: 100%;
    font-size: 18px;
}
</style>