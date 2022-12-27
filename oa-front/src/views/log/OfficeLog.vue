<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-select v-model="queryObj.officeUID" placeholder="办事处" clearable style="width: 100%;">
                            <el-option v-for="item in officeItems" :key="item.UID" :label="item.name"
                                :value="item.UID" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.remarks" placeholder="备注" />
                    </el-col>
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

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    officeUID: "",
    startDate: "",
    endDate: "",
    remarks: "",
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
            prop: "CreatedAt",
            label: "时间",
        },
        {
            prop: "office.name",
            label: "办事处",
        },
        {
            prop: "old",
            label: "原值",
        },
        {
            prop: "change",
            label: "改变值",
        },
        {
            prop: "new",
            label: "新值",
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
.el-row {
    margin-bottom: 20px;
}
</style>