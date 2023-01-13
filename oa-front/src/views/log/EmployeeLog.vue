<template>
    <el-row :gutter="20">
        <el-col :span="4" :offset="1">
            <el-select v-model="base.model.user.officeID" placeholder="办事处" clearable style="width: 100%;"
                :disabled="!user().my.pids.includes('88')">
                <el-option v-for="item in base.offices" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
        </el-col>
        <el-col :span="4">
            <el-input v-model="base.model.user.name" placeholder="员工" clearable maxlength="25" />
        </el-col>
        <el-col :span="4">
            <el-input v-model="base.model.remark" placeholder="备注" clearable maxlength="25" />
        </el-col>
        <el-col :span="4">
            <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
        </el-col>
        <el-col :span="4">
            <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
        </el-col>
        <el-col :span="1">
            <el-row>
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-row>
        </el-col>
    </el-row>

    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { reactive, onBeforeMount } from 'vue'
import { queryAllOffice } from "@/api/office"
import { queryHistoryEmployees } from "@/api/history"

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    offices: [],
    model: {
        remark: "",
        startDate: "",
        endDate: "",
        user: {
            officeID: "",
            name: "",
        },
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "时间",
                width: "10%",
            },
            {
                prop: "user.office.name",
                label: "办事处",
                width: "10%",
            },
            {
                prop: "user.name",
                label: "员工",
                width: "10%",
            },
            {
                prop: "oldMoney",
                label: "原补助额度",
                width: "10%",
            },
            {
                prop: "changeMoney",
                label: "改变值",
                width: "10%",
            },
            {
                prop: "newMoney",
                label: "新补助额度",
                width: "10%",
            },
            {
                prop: "employee.name",
                label: "操作人",
                width: "10%",
            },
            {
                prop: "remark",
                label: "备注",
                width: "30%",
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
        queryHistoryEmployees(base.model, base.pageData).then((res) => {
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
    queryAllOffice().then((res) => {
        if (res.status == 1) {
            base.offices = res.data
        }
    })
    if (!user().my.pids.includes('88')) {
        base.model.user.officeID = Number(localStorage.getItem("officeID"))
    }
    base.query()
})

</script>