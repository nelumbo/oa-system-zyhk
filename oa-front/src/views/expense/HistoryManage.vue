<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="6">
            <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
        </el-col>
        <el-col :span="6">
            <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />
</template>

<script setup>
import { reactive, onBeforeMount } from 'vue'
import { queryMyHistorys } from "@/api/my"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        startDate: "",
        endDate: "",
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "时间",
            },
            {
                prop: "oldMoney",
                label: "原补助额度",
            },
            {
                prop: "changeMoney",
                label: "改变值",
            },
            {
                prop: "newMoney",
                label: "新补助额度",
            },
            {
                prop: "employee.name",
                label: "操作人",
            },
            {
                prop: "remark",
                label: "备注",
            },
        ]
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        queryMyHistorys(base.model, base.pageData).then((res) => {
            if (res.status == 1) {
                base.tableData = res.data.data
                base.pageData.total = res.data.total
                base.pageData.pageSize = res.data.pageSize
                base.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    handleSizeChange: (e) => {
        base.pageData.pageSize = e
        base.pageData.pageNo = 1
        base.query()
    },
    handleCurrentChange: (e) => {
        base.pageData.pageNo = e
        base.query()
    },
})

onBeforeMount(() => {
    base.query()
})
</script>