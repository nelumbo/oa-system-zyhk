<template>
    <el-form :inline="true" :model="base.model">
        <el-row>
            <el-col :span="6">
                <el-form-item label="个人补助金额（元）">
                    <el-input v-model="base.model.money" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="业务费用金额（元）">
                    <el-input v-model="base.model.office.businessMoney" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="办事处提成金额（元）">
                    <el-input v-model="base.model.office.money" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="办事处年底提成金额（元）">
                    <el-input v-model="base.model.office.moneyCold" disabled />
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
</template>
  
<script setup>
import { reactive, onBeforeMount } from 'vue'
import { queryMe } from "@/api/my"
import { message } from '@/components/divMessage/index'

const base = reactive({
    model: {
        money: 0,
        office: {
            businessMoney: 0,
            money: 0,
            moneyCold: 0,
        },
    },
    query: () => {
        queryMe().then((res) => {
            if (res.status == 1) {
                base.model = res.data
            } else {
                message("查询失败", "error")
            }
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>