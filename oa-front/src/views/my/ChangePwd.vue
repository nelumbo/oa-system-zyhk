<template>
    <el-form class="form" :model="model" :rules="rules" ref="baseForm" label-width="120px"
        style="margin-top: 10%;margin-left: 20%;">
        <el-row :gutter="20">
            <el-col :span="8" :offset="2">
                <el-form-item prop="password" label="密码">
                    <el-input class="input" v-model="model.password" show-password clearable prefix-icon="Lock">
                    </el-input>
                </el-form-item>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="8" :offset="2">
                <el-form-item prop="pwd" label="新密码">
                    <el-input class="input" v-model="model.pwd" show-password clearable prefix-icon="Lock">
                    </el-input>
                </el-form-item>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="8" :offset="2">
                <el-form-item prop="checkPwd" label="重复新密码">
                    <el-input class="input" v-model="model.checkPwd" show-password clearable prefix-icon="Lock">
                    </el-input>
                </el-form-item>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="8" :offset="4">
                <el-form-item>
                    <el-button class="btn" type="primary" @click="submit">登录</el-button>
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { updatePwd } from "@/api/my"
import { message } from '@/components/divMessage/index'

const baseForm = ref(null)

const model = reactive({
    password: '',
    pwd: '',
    checkPwd: '',
})

const ConfirmPassword = (rule, value, callback) => {
    if (!value) {
        callback(new Error('请重复输入新密码'))
    }
    if (value != model.pwd) {
        callback(new Error('两次输入密码不一致'))
    } else {
        callback()
    }
}

const rules = reactive({
    password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
    ],
    pwd: [
        { required: true, message: '请输入密码', trigger: 'blur' },
    ],
    checkPwd: [
        { required: true, validator: ConfirmPassword, trigger: 'blur' },
    ],
})

const submit = () => {
    baseForm.value.validate((valid) => {
        if (valid) {
            updatePwd(model).then((res) => {
                if (res.status == 1) {
                    localStorage.removeItem("token");
                    window.location.href = '/login'
                } else if (res.status == 2) {
                    message("修改失败，原密码输入错误", "error")
                }
            })
        } else {
            return false;
        }
    })
}
</script>