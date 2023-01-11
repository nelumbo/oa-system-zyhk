<template>
    <div>
        <el-row>
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6"></el-col>
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.startDate" type="date" placeholder="开始时间"
                            style="width: 100%;" />
                    </el-col>
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.endDate" type="date" placeholder="结束时间"
                            style="width: 100%;" />
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

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    employeeUID: "",
    startDate: "",
    endDate: "",
})

const columnObj = {
    headers: [
        {
            prop: "CreatedAt",
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
            prop: "remarks",
            label: "备注",
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
/* .el-row {
    margin-bottom: 20px;
} */
</style>