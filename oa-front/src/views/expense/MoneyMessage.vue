<template>
    <el-form :inline="true" :model="employee">
        <el-row>
            <el-col :span="6">
                <el-form-item label="个人补助金额（元）">
                    <el-input v-model="employee.money" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="业务费用金额（元）">
                    <el-input v-model="employee.office.businessMoney" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="办事处提成金额（元）">
                    <el-input v-model="employee.office.money" disabled />
                </el-form-item>
            </el-col>
            <el-col :span="6">
                <el-form-item label="办事处年底提成金额（元）">
                    <el-input v-model="employee.office.moneyCold" disabled />
                </el-form-item>
            </el-col>
        </el-row>
    </el-form>
</template>
  
<script setup>
const props = defineProps({
    employee: {
        type: Object,
        default: {
            money: -1,
            office: {
                businessMoney: -1,
                money: -1,
                moneyCold: -1,
            },
        }
    },
})
</script>