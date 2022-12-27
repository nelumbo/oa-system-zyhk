<template>
    <div>
        <el-row :gutter="20" align="middle">
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-select v-model="queryObj.type" placeholder="报销类型" clearable style="width: 100%;">
                            <el-option v-for="item in expenseTypeItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.status" placeholder="状态" clearable style="width: 100%;">
                            <el-option v-for="item in expenseStatusItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.officeUID" placeholder="办事处" clearable style="width: 100%;">
                            <el-option v-for="item in officeItems" :key="item.UID" :label="item.name"
                                :value="item.UID" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.employeeName" placeholder="员工" />
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.startDate" type="date" placeholder="开始时间"
                            style="width: 100%;" />
                    </el-col>
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.endDate" type="date" placeholder="结束时间"
                            style="width: 100%;" />
                    </el-col>
                </el-row>
            </el-col>
            <el-col :span="1">
                <el-row>
                    <el-button type="primary" @click="query">查询</el-button>
                </el-row>
            </el-col>
        </el-row>
        <divTable :columnObj="columnObj" :tableData="tableData" :pageObj="pageObj" />
    </div>
</template>

<script setup>
import { reactive } from 'vue'
import { expenseTypeItems, expenseStatusItems } from '../../utils/magic'

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    type: null,
    status: null,
    officeUID: "",
    employeeName: "",
    startDate: "",
    endDate: "",
})

let officeItems = [
    {
        UID: "bj1",
        name: '北京一区',
    },
    {
        UID: "bj2",
        name: '北京二区',
    },
]

const columnObj = {
    headers: [
        {
            prop: "expenseType.text",
            label: "类型",
        },
        {
            prop: "CreatedAt",
            label: "发起时间",
        },
        {
            prop: "employee.office.name",
            label: "办事处",
        },
        {
            prop: "employee.name",
            label: "员工",
        },
        {
            prop: "employee.phone",
            label: "员工电话",
        },
        {
            prop: "amount",
            label: "报销金额(元)",
        },
        {
            prop: "approver1.name",
            label: "办事处审批人",
        },
        {
            prop: "approver2.name",
            label: "财务",
        },
        {
            prop: "approver3.name",
            label: "出纳",
        },
        {
            prop: "status",
            label: "状态",
        },
        {
            type: "operation",
            label: "操作",
            operations: []
        },
    ],
}

const tableData = []

const pageObj = {
    total: 0,
    pageData: {
        page: 0,
        size: 10
    }
}

function query() {
    alert(JSON.stringify(queryObj))
}

</script>

<style>
.el-row {
    margin-bottom: 20px;
}
</style>