<template>
    <div>
        <el-row>
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6"></el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.employeeName" placeholder="业务员" />
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.status" placeholder="状态" clearable style="width: 100%;">
                            <el-option v-for="item in preResearchStatusItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="1">
                        <el-button type="primary" @click="query">查询</el-button>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
        <divTable :columnObj="columnObj" :tableData="tableData" :pageObj="pageObj" />
    </div>
</template>

<script setup>
import { reactive } from 'vue'
import { preResearchStatusItems } from '../../utils/magic'

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    employeeName: "",
    status: null,
})

const columnObj = {
    headers: [
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
            label: "业务员",
        },
        {
            prop: "auditor.name",
            label: "审批人",
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