<template>
    <div>
        <el-divider content-position="left">
            <h2>我的资金</h2>
        </el-divider>
        <moneyMessage />
    </div>
    <div>
        <el-divider content-position="left">
            <h2>我的报销</h2>
        </el-divider>
        <expenseManage />
    </div>
    <div>
        <el-divider content-position="left">
            <h2>补助额度变更记录</h2>
        </el-divider>
        <historyManage />
    </div>
</template>

<script setup>
import moneyMessage from './MoneyMessage.vue'
import expenseManage from './ExpenseManage.vue'
import historyManage from './HistoryManage.vue'
</script>